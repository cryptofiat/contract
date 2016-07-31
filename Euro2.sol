//
//
//



// Helper Contract: Reserve Bank is an appointed account to be responsible for increasing and decreasing the supply of crypto tokens.
contract ReserveBank { 
    address public reserveBank;

    //Set in constructor, if not overridden
    function ReserveBank() {
        reserveBank = msg.sender;
    }

    modifier onlyReserveBank {
        if (msg.sender != reserveBank) throw;
        _
    }

    function appointReserveBank(address newReserveBank)
        onlyReserveBank
    {
        reserveBank = newReserveBank;
    }
}

// Helper Contract: Policable implements law enforcement institutions
contract Policable is ReserveBank {
    address public lawEnforcer; // account or multisig contract
    address public accountApprover; // account or multisig contract
    address public enforcementAccountDesignator; // people, who can reconfigure the destination account
    address public enforcementAccount; // Account where lawEnforcers can pull money to - e.g. some

    function Policable() {
        lawEnforcer = msg.sender;
    }

// Modifiers only to allow actions by appointed accounts

    modifier onlyLawEnforcer {
        if (msg.sender != lawEnforcer) throw;
        _
    }

    modifier onlyAccountApprover {
        if (msg.sender != accountApprover) throw;
        _
    }

    modifier onlyEnforcementAccountDesignator {
        if (msg.sender != enforcementAccountDesignator) throw;
        _
    }


// Functions to change the appointed accounts/committees

    function appointLawEnforcer(address newLawEnforcer)
    {
	// Allow the reserveBank to set new law enforcement if keys are lost
	if (msg.sender != lawEnforcer && msg.sender != reserveBank) throw;
        lawEnforcer = newLawEnforcer;
    }

    function appointAccountApprover(address newAccountApprover)
    {
	// Allow the reserveBank to set a new account approver if keys are lost
	if (msg.sender != accountApprover && msg.sender != reserveBank) throw;
        accountApprover = newAccountApprover;
    }

    function appointEnforcementAccountDesignator(address newEnforcementAccountDesignator)
    {
	// Allow the reserveBank to set a new enforcement account designator if keys are lost
	if (msg.sender != enforcementAccountDesignator && msg.sender != reserveBank) throw;
        enforcementAccountDesignator = newEnforcementAccountDesignator;
    }

// The appointed committee can set the destination account for enforcements by lawEnforcer

    function designateEnforcementAccount(address target)
        onlyEnforcementAccountDesignator
    {
	//TODO: should also check it is approvedAccount and not suspendedAccount
	if (target != 0)  {
        	enforcementAccount = target;
	}
    }

}



// Euro2 implements crypto-currency that can be recovered and requires
// central approval.
contract CryptoFiat is ReserveBank, Policable {
    /* Data of the contract */
    string public standard = 'CryptoEUR 0.4';
    string public name;
    string public symbol;
    uint8  public decimals;

    // tracking total balance
    uint256 public totalSupply;
    // tracks balance of a particular account
    mapping (address => uint256) public balanceOf;
    // set of approved accounts that can transfer out Euro2 tokens. 
    //Unapproved accounts can still receive tokens, the account holders  can't use them  until they get approved.
    mapping (address => bool) public approvedAccount;
    // set of accounts that can not receive tokens
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

    // Initializes crypto currency with the appropriate initial institutions
    function CryptoFiat(
        uint256 initialSupply,
        string tokenName,
        uint8 decimalUnits,
        string tokenSymbol,
        address _reserveBank,
        address _accountApprover,
        address _enforcementAccount,
        address _enforcementAccountDesignator,
        address _lawEnforcer
    ) {
        // reassign appointed accounts, if specified
        if(_reserveBank != 0) reserveBank = _reserveBank;
        if(_lawEnforcer != 0 ) lawEnforcer = _lawEnforcer;
        if(_accountApprover != 0 ) accountApprover = _accountApprover;
        if(_enforcementAccount != 0 ) enforcementAccount = _enforcementAccount;
        if(_enforcementAccountDesignator != 0 ) enforcementAccountDesignator = _enforcementAccountDesignator;

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
    // (v,r,s) is signature of (to, amount, fee, nonce)
    // signed by _from account.
    //
    // from must have amount + fee available for the transfer
    // nonce must be larger than previous nonce used by that account
    function signedTransfer(
        // requested transfer
        address _from, address _to, uint256 _amount, uint256 _fee, uint256 _nonce,
        // signature of message above signed by _from
        uint8 sig_v, bytes32 sig_r, bytes32 sig_s,
        // where to transfer the transaction fee
        address _sponsor
    ){
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

        // verify that the requested transfer is properly signed
        // TODO: research whether _from field can be removed from the arguments
        address verify = ecrecover(sha3(_from, _to, _amount, _fee, _nonce), sig_v, sig_r, sig_s);
        if(verify != _from) throw;

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
        if(balanceOf[enforcementAccount] + _amount < balanceOf[enforcementAccount]) throw;
	// check if we can send money to destination
        if(suspendedAccount[enforcementAccount]) throw;

        balanceOf[_from] -= _amount;
        balanceOf[enforcementAccount] += _amount;
        Transfer(_from, enforcementAccount, _amount);
    }

    // allows Reserve Bank to issue new tokens to the total supply
    function increaseSupply(uint256 _amount) onlyReserveBank {
        if(suspendedAccount[reserveBank]) throw;

        // check for overflow
        if(balanceOf[reserveBank] + _amount < balanceOf[reserveBank]) throw;

        balanceOf[reserveBank] += _amount;
        totalSupply += _amount;
        Transfer(0, reserveBank, _amount);
    }

    // allows reserve bank to remove tokens from the total supply
    function decreaseSupply(uint256 _amount) onlyReserveBank {
        if(!approvedAccount[reserveBank]) throw;

        // check if have enough tokens to burn
        if(balanceOf[reserveBank] - _amount < 0) throw;

        balanceOf[reserveBank] -= _amount;
        totalSupply -= _amount;
        Transfer(reserveBank, 0, _amount);
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
