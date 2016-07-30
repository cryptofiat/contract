import "./Mintable.sol";
import "./Policable.sol";

contract Euro2 is Mintable, Policable {
    /* Data of the contract */
    string public standard = 'CryptoEUR 0.2';
    string public name;
    string public symbol;
    uint8  public decimals;

    /* This creates an array with all balances */
    mapping (address => bool) public approvedAccount;
    mapping (address => address) public recoveryAccount;
    mapping (address => uint256) public balanceOf;

    uint256 public totalSupply;

    /* This generates a public event on the blockchain that will notify clients */
    event Transfer(address indexed from, address indexed to, uint256 value, uint256 reference);

    /* This generates a public event on the blockchain that will notify clients */
    event ApprovedAccount(address target, bool approved);

    /* Initializes contract with initial supply tokens to the creator of the contract */
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
        if(_centralMinter != 0 ) centralBank = _centralMinter;         // Sets the minter
        if(_lawEnforcer != 0 ) lawEnforcer = _lawEnforcer;  // Sets the law enforcer if set in arg
        if(_accountApprover != 0 ) accountApprover = _accountApprover;  // Sets the law enforcer if set in arg
        if(_enforcementDestination != 0 ) enforcementDestination = _enforcementDestination;  // Sets the law enforcer if set in arg
        if(_enforcementDestinationSetter != 0 ) enforcementDestinationSetter = _enforcementDestinationSetter;  // Sets the law enforcer if set in arg
        balanceOf[msg.sender] = initialSupply;              // Give the creator all initial tokens
        name = tokenName;                                   // Set the name for display purposes
        symbol = tokenSymbol;                               // Set the symbol for display purposes
        decimals = decimalUnits;                            // Amount of decimals for display purposes
        totalSupply = initialSupply;
    }

    function transfer(address _to, uint256 _amount, uint256 _reference){
        // check whether we can transfer from sender
        if(!approvedAccount[msg.sender]) throw;
        if(balanceOf[msg.sender] < _amount) throw;

        // check for overflow
        if(balanceOf[_to] + _amount < balanceOf[_to]) throw;

        balanceOf[msg.sender] -= _amount;
        balanceOf[_to] += _amount;
        Transfer(msg.sender, _to, _amount, _reference);
    }

    /* Execute transfers signed by sender */
    function signedTransfer(
        // requested transfer
        address _from, address _to, uint256 _amount, uint256 _reference, uint256 _fee,
        // where to transfer the transaction fee
        address _sponsor,
        // signature of _from
        uint8 v, bytes32 r, bytes32 s
    ){
        // check whether we can transfer from sender
        if(!approvedAccount[_from]) throw;
        if(balanceOf[_from] < _amount + _fee) throw;

        // check for overflow
        if(balanceOf[_to] + _amount < balanceOf[_to]) throw;
        if(balanceOf[_sponsor] + _fee < balanceOf[_sponsor]) throw;

        // Verify that _from requested the transfer
        if (ecrecover(sha3(_from, _to, _amount, _fee, _reference), v, r, s) != _from)
            throw;

        balanceOf[_from] -= _amount;
        balanceOf[_to] += _amount;
        Transfer(_from, _to, _amount, _reference);

        if (_fee > 0) {
            balanceOf[_from] -= _fee;
            balanceOf[_sponsor] += _fee;
            Transfer(_from, _sponsor, _fee, 0);
        }
    }


    function mintToken(address target, uint256 mintedAmount)
        onlyMinter
    {
        balanceOf[target] += mintedAmount;
        totalSupply += mintedAmount;
        Transfer(0, centralBank, mintedAmount, 0);
        Transfer(centralBank, target, mintedAmount, 0);
    }

    function approveAccount(address target, bool approve)
        onlyAccountApprover
    {
        approvedAccount[target] = approve;
        ApprovedAccount(target, approve);
    }

    function assignRecoveryAccount(address _trusted){
        if(_trusted == 0) {
            delete recoveryAccount[msg.sender];
        } else {
            recoveryAccount[msg.sender] = _trusted;
        }
    }

    function recoverAccount(address _recover, address _to){
        // check whether there is the correct recoverer
        if(msg.sender != recoveryAccount[_recover]) throw;

        // check whether we can transfer from sender
        if(!approvedAccount[_recover]) throw;

        uint256 amount = balanceOf[_recover];

        // check for overflow
        if(balanceOf[_to] + amount < balanceOf[_to]) throw;

        balanceOf[_to] += amount;
        Transfer(_recover, _to, amount, 0);

        // tear down the unusable account
        delete balanceOf[_recover];
        delete recoveryAccount[_recover];
        delete approvedAccount[_recover];
    }

    /* This unnamed function is called whenever someone tries to send ether to it */
    function () {
        throw;     // Prevents accidental sending of ether
    }
}

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
