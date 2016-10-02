pragma solidity ^0.4.0;

contract CryptoFiat {
    address public masterAccount;
    modifier onlyMasterAccount {
        if(msg.sender != masterAccount) throw;
        _ ;
    }

    mapping(bytes32 => address) public contracts;
    mapping(address => uint256) public approved;

    function CryptoFiat(){
        masterAccount = msg.sender;
    }

    function appointMasterAccount(address next) onlyMasterAccount { masterAccount = next; }

    function upgrade(bytes32 id, address next) {
        address prev = contracts[id];
        // message sender or the previous contract
        bool canUpgrade = (msg.sender == masterAccount) || (msg.sender == prev);
        if(!canUpgrade) throw;

        if(prev != 0) approved[prev] -= 1;
        contracts[id] = next;
        if(next != 0) approved[next] += 1;
    }
}

contract Constants {
    // contracts
    bytes32 constant DATA             = 0;
    bytes32 constant ACCOUNTS         = 1;
    bytes32 constant APPROVING        = 2;
    bytes32 constant RESERVE          = 3;
    bytes32 constant ENFORCEMENT      = 4;
    bytes32 constant ACCOUNT_RECOVERY = 5;
    bytes32 constant DELEGATION       = 6;
    bytes32 constant MULTI_DELEGATION = 7;

    // data
    bytes32 constant BALANCE                  = 1;
    bytes32 constant STATE                    = 2;
    bytes32 constant DELEGATED_TRANSFER_NONCE = 3;
    bytes32 constant RECOVERY_ACCOUNT         = 4;
    bytes32 constant TOTAL_SUPPLY             = 5;

    // account states
    uint256 constant APPROVED = 1;
    uint256 constant CLOSED   = 2;
    uint256 constant FROZEN   = 4;

    // events
    event Transfer(address indexed source, address indexed destination, uint256 amount);

    event AccountApproved(address indexed source);
    event AccountClosed(address indexed source);
    event AccountFreeze(address indexed source, bool frozen);

    event SupplyChanged(uint256 totalSupply);
}

contract Relay is Constants {
    address cryptoFiat;

    modifier onlyContracts {
        if(CryptoFiat(cryptoFiat).approved(msg.sender) <= 0) throw;
        _;
    }

    function contractFor(bytes32 id) internal returns (address) { return CryptoFiat(cryptoFiat).contracts(id); }

    function data() internal returns (Data) { return Data(contractFor(DATA)); }
    function accounts() internal returns (Accounts) { return Accounts(contractFor(ACCOUNTS)); }
    function approving() internal returns (Approving) { return Approving(contractFor(APPROVING)); }
    function reserve() internal returns (Reserve) { return Reserve(contractFor(RESERVE)); }
    function enforcement() internal returns (Enforcement) { return Enforcement(contractFor(ENFORCEMENT)); }
    function accountRecovery() internal returns (AccountRecovery) { return AccountRecovery(contractFor(ACCOUNT_RECOVERY)); }
    function delegation() internal returns (Delegation) { return Delegation(contractFor(DELEGATION)); }
    function multiDelegation() internal returns (MultiDelegation) { return MultiDelegation(contractFor(MULTI_DELEGATION)); }
}

contract Data is Relay {
    function Data(address _cryptoFiat) {
        cryptoFiat = _cryptoFiat;
    }

    mapping(bytes32 => bytes32) private data;

    function set(bytes32 context, bytes32 key, bytes32 value)
        onlyContracts
    {
        data[sha3(context, key)] = value;
    }

    function get(bytes32 context, bytes32 key)
        onlyContracts
        returns (bytes32)
    {
        return data[sha3(context, key)];
    }
}


contract InternalData is Constants, Relay {
    // balance contains the balance of an account
    function _balanceOf(address addr) internal returns (uint256) { return uint256(data().get(BALANCE, bytes32(addr))); }
    function _setBalanceOf(address addr, uint256 value) internal { data().set(BALANCE, bytes32(addr), bytes32(value)); }

    // state contains the current state of an account
    function _stateOf(address addr) internal returns (uint256) { return uint256(data().get(STATE, bytes32(addr))); }
    function _setStateOf(address addr, uint256 value) internal { data().set(STATE, bytes32(addr), bytes32(value)); }

    // delegated trancfer nonce contains the last nonce used in delegatedTransfer
    function _delegatedTransferNonceOf(address addr) internal returns (uint256) { return uint256(data().get(DELEGATED_TRANSFER_NONCE, bytes32(addr))); }
    function _setDelegatedTransferNonceOf(address addr, uint256 value) internal { data().set(DELEGATED_TRANSFER_NONCE, bytes32(addr), bytes32(value)); }

    // recovery account contains a fallback account that can be used to recover funds
    function _recoveryAccountOf(address addr) internal returns (address) { return address(data().get(RECOVERY_ACCOUNT, bytes32(addr))); }
    function _setRecoveryAccountOf(address addr, address value) internal { data().set(RECOVERY_ACCOUNT, bytes32(addr), bytes32(value)); }

    // totalSupply is the total amount of tokens in circulation
    function _totalSupply() internal returns (uint256) { return uint256(data().get(TOTAL_SUPPLY, bytes32(0))); }
    function _setTotalSupply(uint256 value) internal { data().set(TOTAL_SUPPLY, bytes32(0), bytes32(value)); }

    // for checking account status
    function isApproved(address account) internal returns (bool) { return _stateOf(account) & APPROVED == APPROVED; }
    function isClosed(address account)   internal returns (bool) { return _stateOf(account) & CLOSED   == CLOSED;   }
    function isFrozen(address account)   internal returns (bool) { return _stateOf(account) & FROZEN   == FROZEN;   }

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

    // internal modification of balance
    function _withdraw(address account, uint256 amount) internal {
        // check for underflow
        uint256 balance = _balanceOf(account);
        if(balance < amount) throw;
        _setBalanceOf(account, balance - amount);
    }

    function _deposit(address account, uint256 amount) internal {
        // check for overflow
        uint256 balance = _balanceOf(account);
        if(balance + amount < balance) throw;
        _setBalanceOf(account, balance + amount);
    }
}

contract Accounts is InternalData {
    function Accounts(address _cryptoFiat){
        cryptoFiat = _cryptoFiat;
    }

    // balance contains the balance of an account
    function balanceOf(address addr) returns (uint256) { return _balanceOf(addr); }
    function stateOf(address addr) returns (uint256) { return _stateOf(addr); }

    function transfer(address destination, uint256 amount)
        canSend(msg.sender)
        canReceive(destination)
    {
        address source = msg.sender;

        _withdraw(source, amount);
        _deposit(destination, amount);
        Transfer(source, destination, amount);
    }
}

contract Approving is InternalData {
    address public accountApprover;
    function Approving(address _cryptoFiat, address _accountApprover){
        cryptoFiat = _cryptoFiat;
        accountApprover = _accountApprover;
    }
    modifier onlyAccountApprover { if(msg.sender != accountApprover) throw; _ ; }
    function appointAccountApprover(address next) onlyAccountApprover {
        accountApprover = next;
    }

    // approveAccount makes it possible for an account to send money
    function approveAccount(address account) onlyAccountApprover {
        _setStateOf(account, _stateOf(account) | APPROVED);
        AccountApproved(account);
    }

    // approveAccounts approves multiple accounts
    function approveAccounts(address[] accounts) {
        for(uint i = 0; i < accounts.length; i += 1){
            approveAccount(accounts[i]);
        }
    }

    // closeAccount closes the account for receiving money
    function closeAccount(address account) onlyAccountApprover {
        _setStateOf(account, _stateOf(account) | CLOSED);
        AccountClosed(account);
    }
}


contract Reserve is InternalData {
    address public reserveBank;
    function Reserve(address _cryptoFiat, address _reserveBank){
        cryptoFiat = _cryptoFiat;
        reserveBank = _reserveBank;
    }
    modifier onlyReserveBank { if(msg.sender != reserveBank) throw; _ ; }
    function appointReserveBank(address next) onlyReserveBank { reserveBank = next; }

    function totalSupply() returns (uint256) { return _totalSupply(); }

    // increaseSupply increases the tokens in circulation
    function increaseSupply(uint256 amount)
        onlyReserveBank
        canReceive(reserveBank)
    {
        uint256 supply = totalSupply();
        if(supply + amount < supply) throw;
        supply += amount;
        _setTotalSupply(supply);

        _deposit(reserveBank, amount);
        Transfer(0, reserveBank, amount);

        SupplyChanged(supply);
    }

    // decreaseSupply decreases the amount of tokens in circulation
    function decreaseSupply(uint256 amount)
        onlyReserveBank
        canSend(reserveBank)
    {
        uint256 supply = _totalSupply();
        if(supply < amount) throw; // invalid state
        supply -= amount;
        _setTotalSupply(supply);

        _withdraw(reserveBank, amount);
        Transfer(reserveBank, 0, amount);

        SupplyChanged(supply);
    }
}


contract Enforcement is InternalData {
    address public lawEnforcer;

    address public accountDesignator;
    address public account;

    function Enforcement(address _cryptoFiat, address _lawEnforcer, address _enforcementAccountDesignator, address _enforcementAccount){
        cryptoFiat = _cryptoFiat;
        lawEnforcer = _lawEnforcer;

        accountDesignator = _enforcementAccountDesignator;
        account = _enforcementAccount;
    }
    modifier onlyLawEnforcer { if(msg.sender != lawEnforcer) throw; _ ; }
    modifier onlyAccountDesignator { if(msg.sender != accountDesignator) throw; _ ; }

    function appointLawEnforcer(address next) onlyLawEnforcer { lawEnforcer = next; }
    function appointAccountDesignator(address next) onlyAccountDesignator { accountDesignator = next; }

    // withdraw allows law enforcerer to withdraw to a dedicated account
    function withdraw(address from, uint256 amount)
        onlyLawEnforcer
        canReceive(account)
    {
        _withdraw(from, amount);
        _deposit(account, amount);
        Transfer(from, account, amount);
    }

    // freezeAccount disallows account to send money
    function freezeAccount(address target) onlyLawEnforcer {
        _setStateOf(target, _stateOf(target) | FROZEN);
        AccountFreeze(target, true);
    }

    // unfreezeAccount re-allows account to send money
    function unfreezeAccount(address target) onlyLawEnforcer {
        _setStateOf(target, _stateOf(target) & ~FROZEN);
        AccountFreeze(target, false);
    }

    // designateAccount allows changing the account
    function designateAccount(address account)
        onlyAccountDesignator
        canReceive(account)
    {
        account = account;
    }
}

contract AccountRecovery is InternalData {
    function AccountRecovery(address _cryptoFiat){
        cryptoFiat = _cryptoFiat;
    }

    // designateRecoveryAccount allows msg.sender to specify a trusted account
    // that can recover the tokens
    function designateRecoveryAccount(address recoveryAccount){
        _setRecoveryAccountOf(msg.sender, recoveryAccount);
    }

    // recoverBalance allows to recover tokens on a particular account
    // recovering a balance will automatically close the account
    function recoverBalance(address from, address into)
        canSend(from)
        canReceive(into)
    {
        if(msg.sender != _recoveryAccountOf(from)) throw;

        // close the account
        _setStateOf(from, _stateOf(from) | CLOSED);
        AccountClosed(from);

        uint256 amount = _balanceOf(from);

        _withdraw(from, amount);
        _deposit(into, amount);

        Transfer(from, into, amount);
    }
}

contract Delegation is InternalData {
    function Delegation(address _cryptoFiat){
        cryptoFiat = _cryptoFiat;
    }

    function delegatedTransferNonceOf(address account) returns (uint256) { return _delegatedTransferNonceOf(account); }

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

        if(_delegatedTransferNonceOf(source) >= nonce) throw;
        _setDelegatedTransferNonceOf(source, nonce);

        _withdraw(source, amount + fee);
        _deposit(destination, amount);

        Transfer(source, destination, amount);
        if(fee > 0){
            _deposit(delegate, fee);
            Transfer(source, delegate, fee);
        }
    }
}

contract MultiDelegation is InternalData {
    function MultiDelegation(address _cryptoFiat){
        cryptoFiat = _cryptoFiat;
    }

    function delegatedTransferNonceOf(address account) returns (uint256) { return _delegatedTransferNonceOf(account); }

    uint constant xfersize = 32+32+32+32+32+32+1;
    // expected format
    // struct XferEncoded {
    //     uint256 nonce;
    //     address destination;
    //     uint256 amount;
    //     uint256 fee;
    //     bytes32 r;
    //     bytes32 s;
    //     uint8   v;
    // }

    struct Xfer {
        uint256 nonce;
        address source;
        address destination;
        uint256 amount;
        uint256 fee;
    }

     function recoverXfer(bytes data, uint offset)
        internal
        returns (Xfer)
    {
        uint base = xfersize * offset;

        uint256 nonce;
        address destination;
        uint256 amount;
        uint256 fee;
        bytes32 r;
        bytes32 s;
        uint8 v;

        assembly {
            nonce := mload(add(data, 32))
            destination := mload(add(data, 64))
            amount := mload(add(data, 96))
            fee := mload(add(data, 128))
            r := mload(add(data, 160))
            s := mload(add(data, 192))
            v := and(mload(add(data, 193)), 255)
        }
        if(v < 27){
            v += 27;
        }

        bytes32 hash = sha3(nonce, destination, amount, fee);

        Xfer memory xfer;
        xfer.nonce = nonce;
        xfer.source = ecrecover(hash, v, r, s);
        xfer.destination = destination;
        xfer.amount = amount;
        xfer.fee = fee;

        return xfer;
    }

    function delegatedTransfers(
        uint256 count,
        bytes   transfers,
        address delegate
    ) {
        for(uint i = 0; i < count; i++){
            Xfer memory xfer = recoverXfer(transfers, i);

            // check whether source can send
            assertSend(xfer.source);

            // protect against replayed transactions
            if(_delegatedTransferNonceOf(xfer.source) >= xfer.nonce) throw;
            _setDelegatedTransferNonceOf(xfer.source, xfer.nonce);

            _withdraw(xfer.source, xfer.amount + xfer.fee);
            _deposit(xfer.destination, xfer.amount);

            Transfer(xfer.source, xfer.destination, xfer.amount);
            if(xfer.fee > 0){
                _deposit(delegate, xfer.fee);
                Transfer(xfer.source, delegate, xfer.fee);
            }
        }
    }
}
