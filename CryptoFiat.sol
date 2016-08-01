// Events pre-defines all events in the contract
contract Events {
	event Transfer(address source, address destination, uint256 amount);
	event SupplyChanged(address source, int256 amount);
}

// Appointed defines all appointed roles and specifies how they can be changed
contract Appointed {
	address public reserveBank;
	address public lawEnforcer;
	address public accountApprover;
	address public enforcementAccountDesignator;

	modifier onlyReserveBank                   { if(msg.sender != reserveBank) throw; _ }
	modifier onlyLawEnforcer                   { if(msg.sender != lawEnforcer) throw; _ }
	modifier onlyAccountApprover               { if(msg.sender != accountApprover) throw; _ }
	modifier onlyEnforcementAccountDesignator  { if(msg.sender != enforcementAccountDesignator) throw; _ }

	// appoint a new reserve bank
	// only previous reserve bank
	function appointReserveBank(address account) onlyReserveBank {
		if (account == 0) throw;
		reserveBank = account;
	}

	// appoint a new account approver
	// only previous accountApprover or reserveBank
	function appointAccountApprover(address account) {
		if (account == 0) throw;
		if (msg.sender != accountApprover && msg.sender != reserveBank) throw;

		accountApprover = account;
	}

	// appoint a new law enforcement account
	// only previous lawEnforcer or reserveBank
	function appointLawEnforcer(address account) {
		if (account == 0) throw;
		if (msg.sender != lawEnforcer && msg.sender != reserveBank) throw;

		lawEnforcer = account;
	}

	// appoint a new law enforcement account designator
	// only previous enforcementAccountDesignator or reserveBank
	function appointEnforcementAccountDesignator(address account) {
		if (account == 0) throw;
		if (msg.sender != enforcementAccountDesignator && msg.sender != reserveBank) throw;

		accountApprover = account;
	}
}

// Accounts defines basic account and capabilities
contract Accounts is Events, Appointed {
	// TODO: convert status to an enum instead of having two mappings

	// accounts that can send money
	mapping (address => bool) public approved;
	// accounts that cannot receive money
    mapping (address => bool) public closed;

	event AccountApproved(address source);
	event AccountFrozen(address source);
	event AccountUnfrozen(address source);
	event AccountClosed(address source);

	function approveAccount(address account) onlyAccountApprover {
		approved[account] = true;
		AccountApproved(account);
	}

	// closeAccount closes the account for receiving money
    function closeAccount(address target) onlyAccountApprover {
        closed[target] = true;
        AccountClosed(target);
    }

    modifier canSend(address account)    {
    	if(!approved[account]) throw;
    	if(account == 0) throw;
    	_
    }
    modifier canReceive(address account) {
    	if(closed[account]) throw;
    	if(account == 0) throw;
    	_
    }
}

// Balance defines the balance for Accounts
contract Balance is Events, Accounts {
	mapping (address => uint256) public balanceOf;

	modifier hasFunds(address account, uint256 amount) {
		if(balanceOf[account] < amount) throw;
		_
	}

	function transfer(address destination, uint256 amount)
		canSend(msg.sender)
		hasFunds(msg.sender, amount)
		canReceive(destination)
	{
		// check for potential overflow
		if(balanceOf[destination] + amount < balanceOf[destination]) throw;

		address source = msg.sender;
		balanceOf[source] -= amount;
		balanceOf[destination] += amount;

		Transfer(source, destination, amount);
	}
}

// Supply defines how tokens are increased/decreased in/from circulation
contract Supply is Appointed, Balance {
	// totalSupply is the total amount of tokens in circulation
	uint256 public totalSupply;

	// increaseSupply increases the tokens in circulation
    function increaseSupply(uint256 amount)
    	onlyReserveBank
    	canReceive(reserveBank)
    {
        // check for potential overflow
        if(balanceOf[reserveBank] + amount < balanceOf[reserveBank]) throw;
        if(totalSupply + amount < totalSupply) throw;

        balanceOf[reserveBank] += amount;
        totalSupply += amount;

        Transfer(0, reserveBank, amount);
    }

    // decreaseSupply decreases the amount of tokens in circulation
    function decreaseSupply(uint256 amount)
    	onlyReserveBank
    	canSend(reserveBank)
    	hasFunds(reserveBank, amount)
    {
    	if(totalSupply < amount) throw; // invalid state

        balanceOf[reserveBank] -= amount;
        totalSupply -= amount;
        Transfer(reserveBank, 0, amount);
    }
}

// DelegatedTransfer allows a third-party to transfer money on behalf of an account
contract DelegatedTransfer is Balance {
   	// delegatedTransferNonce protects against using same signed transfer
   	// multiple times
    mapping (address => uint256) public delegatedTransferNonce;

    function delegatedTransfer(
    	// transfer request
    	address source, address destination, uint256 amount, uint256 fee,
    	uint256 nonce,
    	// transfer request signed by source
    	uint256 signature,
    	// whom to pay for fulfilling transfer
    	address delegate
    )
    	canSend(source)
    	hasFunds(source, amount + fee)
    	canReceive(destination)
		canReceive(delegate)
    {
    	// protect against replayed transactions
    	if(delegatedTransferNonce[source] >= nonce) throw;

    	// check for overflow
    	if(balanceOf[destination] + amount < balanceOf[destination]) throw;
		if(balanceOf[delegate] + fee < balanceOf[delegate]) throw;

    	// verify signature
    	address signer = ecrecover(
    		sha3(source, destination, amount, fee, nonce),
    		0, 0, 0
    	);
    	if(signer != source) throw;

    	balanceOf[source] -= amount + fee;
    	balanceOf[destination] += amount;
    	Transfer(source, destination, amount);

    	if(fee > 0){
	    	balanceOf[delegate] += fee;
    		Transfer(source, delegate, fee);
    	}
    }
}

// AccountRecovery specifies mechhanisms for recovering tokens from an account
contract AccountRecovery is Accounts, Balance {
	// each account can, optionally, specify a recovery account in case
	// the original address cannot be accessed due to lost keys
	mapping (address => address) public recoveryAccountOf;

	// designateRecoveryAccount allows msg.sender to specify a trusted account
	// that can recover the tokens
	function designateRecoveryAccount(address recoveryAccount){
		if(recoveryAccount == 0){
			delete recoveryAccountOf[msg.sender];
		} else {
			recoveryAccountOf[msg.sender] = recoveryAccount;
		}
	}

	// recoverBalance allows to recover tokens on a particular account
	// recovering a balance will automatically close the account
	function recoverBalance(address from, address into)
		canSend(from)
		canReceive(into)
	{
		if(msg.sender != recoveryAccountOf[from]) throw;

		// close the account
		closed[from] = true;
		AccountClosed(from);

		uint256 amount = balanceOf[from];

		// check for overflow
		if(balanceOf[into] + amount < balanceOf[into]) throw;

		balanceOf[from] = 0;
		balanceOf[into] += amount;
		Transfer(from, into, amount);
	}
}

// Enforcement allows to freeze/unfreeze accounts
contract Enforcement is Appointed, Balance {
	// enforcementAccount defines where all enforced withdraws will be made
   	address public enforcementAccount;

    // designateEnforcementAccount allows changing the enforcementAccount
    function designateEnforcementAccount(address account)
        onlyEnforcementAccountDesignator
        canReceive(account)
    {
        enforcementAccount = account;
    }

    // enforcedWithdraw allows law enforcerer to withdraw to a dedicated account
    function enforcedWithdraw(address from, uint256 amount)
    	onlyLawEnforcer
    	hasFunds(from, amount)
    	canReceive(enforcementAccount)
    {
        // check for overflow
        if(balanceOf[enforcementAccount] + amount < balanceOf[enforcementAccount]) throw;

        balanceOf[from] -= amount;
        balanceOf[enforcementAccount] += amount;
        Transfer(from, enforcementAccount, amount);
    }

    // freezeAccount disallows account to send money
    function freezeAccount(address target) onlyLawEnforcer {
        approved[target] = false;
        AccountFrozen(target);
    }

    // unfreezeAccount re-allows account to send money
    function unfreezeAccount(address target) onlyLawEnforcer {
        approved[target] = true;
        AccountUnfrozen(target);
    }
}

// CryptoFiat defines crypto currency with government regulations
contract CryptoFiat is
	Events,
	Appointed,
	Accounts,
	Balance,
	DelegatedTransfer,
	Supply,
	AccountRecovery,
	Enforcement
{
	string public standard = 'CryptoFiat 0.5';
    string public name;
    string public symbol;
    uint8  public decimals = 2;

    function CryptoFiat(
    	string tokenName,
    	string tokenSymbol,
    	uint8  decimalUnits,

    	address _reserveBank,
    	address _accountApprover,
    	address _lawEnforcer,
    	address _enforcementAccountDesignator,
    	address _enforcementAccount
    ) {
    	name = tokenName;
    	symbol = tokenSymbol;
    	decimals = decimalUnits;

    	// ensure that appointed accounts are not empty
    	if(_reserveBank == 0) _reserveBank = msg.sender;
		if(_accountApprover == 0) _accountApprover = msg.sender;
		if(_lawEnforcer == 0) _lawEnforcer = msg.sender;
		if(_enforcementAccountDesignator == 0) _enforcementAccountDesignator = msg.sender;
		if(_enforcementAccount == 0) _enforcementAccount = msg.sender;

		// initialize appointed accounts
		reserveBank = _reserveBank;
		accountApprover = _accountApprover;
		lawEnforcer = _lawEnforcer;
		enforcementAccountDesignator = _enforcementAccountDesignator;
		enforcementAccount = _enforcementAccount;
	}

	// don't allow sending ether to the contract
	function() { throw; }
}