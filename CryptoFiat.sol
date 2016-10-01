pragma solidity ^0.4.0;

contract CryptoFiat {
    address public masterAccount;
    modifier onlyMasterAccount {
        if(msg.sender != masterAccount) throw;
        _ ;
    }

    address public data;
    address public logic;

    function CryptoFiat(){
        masterAccount = msg.sender;
    }

    function appointMasterAccount(address next) onlyMasterAccount { masterAccount = next; }
    function upgradeData(address next) onlyMasterAccount { data = next; }
    function upgradeLogic(address next) onlyMasterAccount { logic = next; }

    function () {
        if(!logic.delegatecall(msg.data)) throw;
    }
}

contract Relay {
    address relay;

    modifier onlyLogic {
        address logic = CryptoFiat(relay).logic();
        if(msg.sender != logic) throw;
        _;
    }

    function data()  internal returns (Storage)         { return Storage(CryptoFiat(relay).data()); }
    function logic() internal returns (CryptoFiatLogic) { return CryptoFiatLogic(CryptoFiat(relay).logic()); }
}

contract Storage is Relay {
    mapping(bytes32 => bytes32) private data;

    function set(bytes32 context, bytes32 key, bytes32 value)
        onlyLogic
    {
        data[sha3(context, key)] = value;
    }

    function get(bytes32 context, bytes32 key)
        onlyLogic
        returns (bytes32)
    {
        return data[sha3(context, key)];
    }
}

contract ExternalStorage is Relay {
    bytes32 constant BALANCE = 1;
    bytes32 constant STATE = 2;
    bytes32 constant DELEGATED_TRANSFER_NONCE = 3;
    bytes32 constant RECOVERY_ACCOUNT = 4;
    bytes32 constant TOTAL_SUPPLY = 5;

    // balance contains the balance of an account
    function balanceOf(address addr) returns (uint256) { return uint256(data().get(BALANCE, bytes32(addr))); }
    function setBalanceOf(address addr, uint256 value) internal { data().set(BALANCE, bytes32(addr), bytes32(value)); }

    // state contains the current state of an account
    function stateOf(address addr) returns (uint256) { return uint256(data().get(STATE, bytes32(addr))); }
    function setStateOf(address addr, uint256 value) internal { data().set(STATE, bytes32(addr), bytes32(value)); }

    // delegated trancfer nonce contains the last nonce used in delegatedTransfer
    function delegatedTransferNonceOf(address addr) returns (uint256) { return uint256(data().get(DELEGATED_TRANSFER_NONCE, bytes32(addr))); }
    function setDelegatedTransferNonceOf(address addr, uint256 value) internal { data().set(DELEGATED_TRANSFER_NONCE, bytes32(addr), bytes32(value)); }

    // recovery account contains a fallback account that can be used to recover funds
    function recoveryAccountOf(address addr) returns (address) { return address(data().get(RECOVERY_ACCOUNT, bytes32(addr))); }
    function setRecoveryAccountOf(address addr, address value) internal { data().set(RECOVERY_ACCOUNT, bytes32(addr), bytes32(value)); }

    // totalSupply is the total amount of tokens in circulation
    function totalSupply() returns (uint256) { return uint256(data().get(TOTAL_SUPPLY, bytes32(0))); }
    function setTotalSupply(uint256 value) internal { data().set(TOTAL_SUPPLY, bytes32(0), bytes32(value)); }
}

// Appointed defines all appointed roles and specifies how they can be changed
contract Appointed {
    address public reserveBank;
    address public lawEnforcer;
    address public accountApprover;
    address public enforcementAccountDesignator;

    modifier onlyReserveBank                   { if(msg.sender != reserveBank) throw; _ ; }
    modifier onlyLawEnforcer                   { if(msg.sender != lawEnforcer) throw; _ ; }
    modifier onlyAccountApprover               { if(msg.sender != accountApprover) throw; _ ; }
    modifier onlyEnforcementAccountDesignator  { if(msg.sender != enforcementAccountDesignator) throw; _ ; }

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
contract Accounts is ExternalStorage, Appointed {
    uint8 constant APPROVED = 1;
    uint8 constant CLOSED   = 2;
    uint8 constant FROZEN   = 4;

    function isApproved(address account) internal returns (bool) { return stateOf(account) & APPROVED == APPROVED; }
    function isClosed(address account)   internal returns (bool) { return stateOf(account) & CLOSED   == CLOSED;   }
    function isFrozen(address account)   internal returns (bool) { return stateOf(account) & FROZEN   == FROZEN;   }

    event AccountApproved(address indexed source);
    event AccountFreeze(address indexed source, bool frozen);
    event AccountClosed(address indexed source);

    function approveAccount(address account) onlyAccountApprover {
        setStateOf(account, stateOf(account) | APPROVED);
        AccountApproved(account);
    }

    // closeAccount closes the account for receiving money
    function closeAccount(address account) onlyAccountApprover {
        setStateOf(account, stateOf(account) | CLOSED);
        AccountClosed(account);
    }

    modifier canSend(address account) {
        if(!isApproved(account)) throw;
        if(isFrozen(account)) throw;

        if(account == 0) throw;
        _ ;
    }
    function assertSend(address account) internal canSend(account) {}

    modifier canReceive(address account) {
        if(isClosed(account)) throw;
        if(account == 0) throw;
        _ ;
    }
    function assertReceive(address account) internal canReceive(account) {}
}

// Balance defines the balance for Accounts
contract Balance is ExternalStorage, Accounts {
    event Transfer(address indexed source, address indexed destination, uint256 amount);

    function transfer(address destination, uint256 amount)
        canSend(msg.sender)
        canReceive(destination)
    {
        address source = msg.sender;

        withdraw(source, amount);
        deposit(destination, amount);
        Transfer(source, destination, amount);
    }

    function withdraw(address account, uint256 amount) internal {
        // check for underflow
        uint256 balance = balanceOf(account);
        if(balance < amount) throw;
        setBalanceOf(account, balance - amount);
    }

    function deposit(address account, uint256 amount) internal {
        // check for overflow
        uint256 balance = balanceOf(account);
        if(balance + amount < balance) throw;
        setBalanceOf(account, balance + amount);
    }
}

// Supply defines how tokens are increased/decreased in/from circulation
contract Supply is ExternalStorage, Appointed, Balance {
    event SupplyChanged(uint256 totalSupply);

    // increaseSupply increases the tokens in circulation
    function increaseSupply(uint256 amount)
        onlyReserveBank
        canReceive(reserveBank)
    {
        uint256 supply = totalSupply();
        if(supply + amount < supply) throw;
        supply += amount;
        setTotalSupply(supply);

        deposit(reserveBank, amount);
        Transfer(0, reserveBank, amount);

        SupplyChanged(supply);
    }

    // decreaseSupply decreases the amount of tokens in circulation
    function decreaseSupply(uint256 amount)
        onlyReserveBank
        canSend(reserveBank)
    {
        uint256 supply = totalSupply();
        if(supply < amount) throw; // invalid state
        supply -= amount;
        setTotalSupply(supply);

        withdraw(reserveBank, amount);
        Transfer(reserveBank, 0, amount);

        SupplyChanged(supply);
    }
}

// DelegatedTransfer allows a third-party to transfer money on behalf of an account
contract DelegatedTransfer is ExternalStorage, Balance {
    function recoverSigner(bytes32 hash, bytes signature)
        internal
        returns (address)
    {
        bytes32 r;
        bytes32 s;
        uint8 v;

        if(signature.length != 65) throw;

        assembly {
            r := mload(add(signature, 32))
            s := mload(add(signature, 64))
            v := and(mload(add(signature, 65)), 255)
        }

        if(v < 27){
            v += 27;
        }

        return ecrecover(hash, v, r, s);
    }

    function delegatedTransfer(
        // transfer request
        uint256 nonce, address destination, uint256 amount, uint256 fee,
        // transfer request signed by source
        bytes signature,
        // whom to pay for fulfilling transfer
        address delegate
    )
        canReceive(destination)
        canReceive(delegate)
    {
        // extract source from signature
        address source = recoverSigner(
            sha3(nonce, destination, amount, fee),
            signature
        );

        // check whether source can send
        assertSend(source);

        // protect against replayed transactions

        if(delegatedTransferNonceOf(source) >= nonce) throw;
        setDelegatedTransferNonceOf(source, nonce);

        withdraw(source, amount + fee);
        deposit(destination, amount);

        Transfer(source, destination, amount);
        if(fee > 0){
            deposit(delegate, fee);
            Transfer(source, delegate, fee);
        }
    }
}

// AccountRecovery specifies mechhanisms for recovering tokens from an account
contract AccountRecovery is ExternalStorage, Accounts, Balance {
    // designateRecoveryAccount allows msg.sender to specify a trusted account
    // that can recover the tokens
    function designateRecoveryAccount(address recoveryAccount){
        setRecoveryAccountOf(msg.sender, recoveryAccount);
    }

    // recoverBalance allows to recover tokens on a particular account
    // recovering a balance will automatically close the account
    function recoverBalance(address from, address into)
        canSend(from)
        canReceive(into)
    {
        if(msg.sender != recoveryAccountOf(from)) throw;

        // close the account
        setStateOf(from, stateOf(from) | CLOSED);
        AccountClosed(from);

        uint256 amount = balanceOf(from);

        withdraw(from, amount);
        deposit(into, amount);

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
        canReceive(enforcementAccount)
    {
        withdraw(from, amount);
        deposit(enforcementAccount, amount);
        Transfer(from, enforcementAccount, amount);
    }

    // freezeAccount disallows account to send money
    function freezeAccount(address target) onlyLawEnforcer {
        setStateOf(target, stateOf(target) | FROZEN);
        AccountFreeze(target, true);
    }

    // unfreezeAccount re-allows account to send money
    function unfreezeAccount(address target) onlyLawEnforcer {
        setStateOf(target, stateOf(target) & ~FROZEN);
        AccountFreeze(target, false);
    }
}

// CryptoFiat defines crypto currency with government regulations
contract CryptoFiatLogic is
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

    function CryptoFiatLogic(
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
