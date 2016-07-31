import "./Mintable.sol";
import "./Policable.sol";

// Euro2 implements crypto-currency that can be recovered and requires
// central approval.
contract Euro2 is Mintable, Policable {
    /* Data of the contract */
    string public standard = 'CryptoEUR 0.2';
    string public name;
    string public symbol;
    uint8  public decimals;

    // tracking total balance
    uint256 public totalSupply;
    // tracks balance of a particular account
    mapping (address => uint256) public balanceOf;
    // set of approved accounts that can transfer Euro2 tokens
    mapping (address => bool) public approvedAccount;
    // set of accounts that should not receive tokens
    mapping (address => bool) public suspendedAccount;

    // each address can, optionally, specify a recovery account in case the
    // original address keys are lost
    mapping (address => address) public recoveryAccount;

    // when using signed transfer it is required to protect against replays of
    // transfer requests. the nonces used, must be increasing
    // e.g. use lastSignatureNonce[sender]++ as the nonce
    mapping (address => uint256) public lastSignatureNonce;


    // EVENTS

    // Transfer notifies about transfer of tokens
    event Transfer(address indexed from, address indexed to, uint256 value);

    // AccountApproved notifies that account has been granted or revoked the right
    // to transfer tokens
    event AccountApproved(address account, bool approved);

    // Initializes Euro2 crypto currency with the appropriate initial institutions
    function Euro2(
        uint256 initialSupply,
        string tokenName,
        uint8 decimalUnits,
        string tokenSymbol,
        address _centralMinter,
        address _accountApprover,
        address _enforcementDestination,
        address _enforcementDestinationSetter,
        address _lawEnforcer
    ) {
        // reassign appointed accounts, if specified
        if(_centralMinter != 0) centralBank = _centralMinter;
        if(_lawEnforcer != 0 ) lawEnforcer = _lawEnforcer;
        if(_accountApprover != 0 ) accountApprover = _accountApprover;
        if(_enforcementDestination != 0 ) enforcementDestination = _enforcementDestination;
        if(_enforcementDestinationSetter != 0 ) enforcementDestinationSetter = _enforcementDestinationSetter;

        // starting balance
        balanceOf[msg.sender] = initialSupply;
        totalSupply = initialSupply;

        // display information for currency
        name = tokenName;
        symbol = tokenSymbol;
        decimals = decimalUnits;
    }

    // transfer transfers tokens from msg.sender to _to
    // msg.sender must be approved and have sufficient funds
    function transfer(address _to, uint256 _amount){
        // check whether we can transfer from sender
        if(!approvedAccount[msg.sender]) throw;
        if(balanceOf[msg.sender] < _amount) throw;

        // check whether we can send to the account
        if(suspendedAccount[_to]) throw;
        // check for overflow
        if(balanceOf[_to] + _amount < balanceOf[_to]) throw;

        balanceOf[msg.sender] -= _amount;
        balanceOf[_to] += _amount;
        Transfer(msg.sender, _to, _amount);
    }

    // signedTransfer transfers tokens on behalf of _from,
    // where _sponser is paid _fee for services
    //
    // (v,r,s) is signature of (from, to, amount, fee, nonce)
    // signed by _from account.
    //
    // from must have amount + fee available for the transfer
    // nonce must be larger than previous nonce used by that account
    //
    // TODO: _from is not needed it can be derived from (v,r,s) and message
    function signedTransfer(
        // requested transfer
        address _to, uint256 _amount, uint256 _fee, uint256 _nonce,
        // where to transfer the transaction fee
        address _sponsor,
        // signature of _from
        uint8 v, bytes32 r, bytes32 s
    ){
	// @EGON  TODO: replace removal of  _from
	address _from = 0;

        // check whether we can transfer from sender
        if(!approvedAccount[_from]) throw;
        if(balanceOf[_from] < _amount + _fee) throw;
        // protect against duplicate transfers
        if(lastSignatureNonce[_from] >= _nonce) throw;

        // check whether we can send to the accounts
        if(suspendedAccount[_to]) throw;
        if((_fee > 0) && suspendedAccount[_sponsor]) throw;

        // check for overflow
        if(balanceOf[_to] + _amount < balanceOf[_to]) throw;
        if(balanceOf[_sponsor] + _fee < balanceOf[_sponsor]) throw;

        // verify that the transfer request is properly signed by "from"
        if (ecrecover(sha3(_to, _amount, _fee, _nonce), v, r, s) != _from)
            throw;

        balanceOf[_from] -= _amount;
        balanceOf[_to] += _amount;
        Transfer(_from, _to, _amount);

        if (_fee > 0) {
            balanceOf[_from] -= _fee;
            balanceOf[_sponsor] += _fee;
            Transfer(_from, _sponsor, _fee);
        }
    }

    // enforcedWithdraw allows law enforcerer to withdraw to a dedicated account
    function enforcedWithdraw(address _from, uint256 _amount, uint256 _reference) onlyLawEnforcer {
        if(balanceOf[_from] < _amount) throw;

        // check for overflow
        if(balanceOf[enforcementDestination] + _amount < balanceOf[enforcementDestination]) throw;
	// check if we can send money to destination
        if(suspendedAccount[enforcementDestination]) throw;

        balanceOf[_from] -= _amount;
        balanceOf[enforcementDestination] += _amount;
        Transfer(_from, enforcementDestination, _amount);
    }

    // mintToken allows minter to issue new tokens to the total supply
    function mintToken(address target, uint256 mintedAmount) onlyMinter {
        if(suspendedAccount[target]) throw;

        balanceOf[target] += mintedAmount;
        totalSupply += mintedAmount;
        Transfer(0, centralBank, mintedAmount);
        Transfer(centralBank, target, mintedAmount);
    }

    // approveAccount changes target approval status
    function approveAccount(address target, bool approve) onlyAccountApprover {
        approvedAccount[target] = approve;
        AccountApproved(target, approve);
    }

    // assignRecoveryAccount allows an account owner to specify a
    // trusted party that can retransfer all money to a backup/new account
    //
    // if the specified account is 0, the recovery account will be removed
    function assignRecoveryAccount(address _trusted){
        if(_trusted == 0) {
            delete recoveryAccount[msg.sender];
        } else {
            recoveryAccount[msg.sender] = _trusted;
        }
    }

    // recoverAccount allows a trusted party to retransfer all money from
    // a lost account to a new or another account
    function recoverAccount(address _recover, address _to){
        // check whether there is the correct recoverer
        if(msg.sender != recoveryAccount[_recover]) throw;

        // check whether we can recover into _to address
        if(suspendedAccount[_to]) throw;

        // check whether we can transfer from sender
        if(!approvedAccount[_recover]) throw;
        suspendedAccount[_recover] = true;

        // get the current balance of the recover account
        uint256 amount = balanceOf[_recover];

        // check for overflow
        if(balanceOf[_to] + amount < balanceOf[_to]) throw;

        // move the tokens around
        balanceOf[_recover] = 0;
        balanceOf[_to] += amount;
        Transfer(_recover, _to, amount);
    }

    /* This unnamed function is called whenever someone tries to send ether to it */
    function () {
        throw;     // Prevents accidental sending of ether
    }
}


// TBC: DOESN'T WORK AT THE MOMENT

contract Relay {
    address public currentVersion;
    address public owner;

    function Relay(address initialAddress){
        currentVersion = initialAddress;
        owner = msg.sender;
    }

    function update(address newAddress){
        if(msg.sender != owner) throw;
        currentVersion = newAddress;
    }

    function(){
        if(!currentVersion.delegatecall(msg.data)) throw;
    }
}
