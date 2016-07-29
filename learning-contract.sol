contract mintable {
    address public centralBank;

    function mintable() {
        centralBank = msg.sender;
    }

    modifier onlyMinter {
        if (msg.sender != centralBank) throw;
        _
    }

    function transferMinter(address newMinter) 
        onlyMinter
    {
        centralBank = newMinter;
    }
}

contract policable {
    address public lawEnforcement;

    function policable() {
        lawEnforcement = msg.sender;
    }

    modifier onlyLawEnforcement {
        if (msg.sender != lawEnforcement) throw;
        _
    }

    function transferLawEnforcement(address newLawEnforcement) 
        onlyLawEnforcement
    {
        lawEnforcement = newLawEnforcement;
    }
}

contract CryptoEur02 is mintable, policable {

    /* Public variables of the contract */
	string public standard = 'Token 0.2';
	string public name;
	string public symbol;
	uint8 public decimals;

	/* This creates an array with all balances */
	mapping (address => uint256) public balanceOf;
	mapping (address => mapping (address => uint256)) public allowance;

	/* This generates a public event on the blockchain that will notify clients */
	event Transfer(address indexed from, address indexed to, uint256 value);

    uint256 public totalSupply;

    mapping (address => bool) public approvedAccount;

    /* This generates a public event on the blockchain that will notify clients */
    event ApprovedAccount(address target, bool approved);

    /* Initializes contract with initial supply tokens to the creator of the contract */
    function CryptoEur02(
        uint256 initialSupply,
        string tokenName,
        uint8 decimalUnits,
        string tokenSymbol,
        address _centralMinter,
        address _lawEnforcement
    ) {
        if(_centralMinter != 0 ) centralBank = _centralMinter;         // Sets the minter
        if(_lawEnforcement != 0 ) lawEnforcement = _lawEnforcement;  // Sets the law enforcer if set in arg
        balanceOf[msg.sender] = initialSupply;              // Give the creator all initial tokens
        name = tokenName;                                   // Set the name for display purposes
        symbol = tokenSymbol;                               // Set the symbol for display purposes
        decimals = decimalUnits;                            // Amount of decimals for display purposes
        totalSupply = initialSupply;
    }

    /* Send crypto Euros */
    function transfer(address _to, uint256 _value) {
        if (balanceOf[msg.sender] < _value) throw;           // Check if the sender has enough
        if (balanceOf[_to] + _value < balanceOf[_to]) throw; // Check for overflows
        if (!approvedAccount[msg.sender]) throw;                // Check if frozen
        balanceOf[msg.sender] -= _value;                     // Subtract from the sender
        balanceOf[_to] += _value;                            // Add the same to the recipient
        Transfer(msg.sender, _to, _value);                   // Notify anyone listening that this transfer took place
    }


    /* A contract attempts to get the coins */
    /*
    function transferFrom(address _from, address _to, uint256 _value) returns (bool success) {
        if (frozenAccount[_from]) throw;                        // Check if frozen            
        if (balanceOf[_from] < _value) throw;                 // Check if the sender has enough
        if (balanceOf[_to] + _value < balanceOf[_to]) throw;  // Check for overflows
        if (_value > allowance[_from][msg.sender]) throw;   // Check allowance
        balanceOf[_from] -= _value;                          // Subtract from the sender
        balanceOf[_to] += _value;                            // Add the same to the recipient
        allowance[_from][msg.sender] -= _value;
        Transfer(_from, _to, _value);
        return true;
    }
    */

    function mintToken(address target, uint256 mintedAmount)
        onlyMinter
    {
        balanceOf[target] += mintedAmount;
        totalSupply += mintedAmount;
        Transfer(0, centralBank, mintedAmount);
        Transfer(centralBank, target, mintedAmount);
    }

    function approveAccount(address target, bool approve)
        onlyLawEnforcement
    {
        approvedAccount[target] = approve;
        ApprovedAccount(target, approve);
    }


	/* Allow another contract to spend some tokens in your behalf */
	/*
	function approveAndCall(address _spender, uint256 _value, bytes _extraData)
		returns (bool success) {
		allowance[msg.sender][_spender] = _value;
		tokenRecipient spender = tokenRecipient(_spender);
		spender.receiveApproval(msg.sender, _value, this, _extraData);
		return true;
	}
	*/

	/* A contract attempts to get the coins */
	/*
	function transferFrom(address _from, address _to, uint256 _value) returns (bool success) {
		if (balanceOf[_from] < _value) throw;                 // Check if the sender has enough
		if (balanceOf[_to] + _value < balanceOf[_to]) throw;  // Check for overflows
		if (_value > allowance[_from][msg.sender]) throw;   // Check allowance
		balanceOf[_from] -= _value;                          // Subtract from the sender
		balanceOf[_to] += _value;                            // Add the same to the recipient
		allowance[_from][msg.sender] -= _value;
		Transfer(_from, _to, _value);
		return true;
	}
	*/

	/* This unnamed function is called whenever someone tries to send ether to it */
	function () {
		throw;     // Prevents accidental sending of ether
	}
}
