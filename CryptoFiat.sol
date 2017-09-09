pragma solidity ^0.4.16;

contract CryptoFiat {
    address public masterAccount;
    modifier onlyMasterAccount {
        if (msg.sender != masterAccount)
            revert();
        _ ;
    }

    // lookup table for finding a particular contract by ID
    mapping(uint256 => address) public contractAddress;
    mapping(address => uint256) public contractId;

    function contractActive(address addr) constant returns (bool) { return contractId[addr] > 0; }

    // list of contracts involved in this CryptoFiat instance
    // use this list of contracts for filtering for events
    address[] public contracts;
    function contractsLength() constant returns (uint256) { return contracts.length; }

    function CryptoFiat() {
        masterAccount = msg.sender;
        contracts.push(address(this));
    }

    function appointMasterAccount(address next) onlyMasterAccount { masterAccount = next; }

    event ContractUpgraded(uint256 indexed id, address previous, address next);
    function upgrade(uint256 id, address next) {
        if (id == 0)
            revert();
        address prev = contractAddress[id];

        if (prev == next)
            revert();

        // message sender or the previous contract
        bool canUpgrade = (msg.sender == masterAccount) || (msg.sender == prev);
        if (!canUpgrade)
            revert();

        // check double use of contract
        if (contractActive(next))
            revert();

        // disable previous contract
        contractId[prev] = 0;

        // activate next contract
        contractAddress[id] = next;
        if (next != 0)
            contractId[next] = id;

        // finalize
        ContractUpgraded(id, prev, next);
        contracts.push(next);
    }
}

contract Constants {
    // contract id
    uint256 constant DATA             = 1;
    uint256 constant ACCOUNTS         = 2;
    uint256 constant APPROVING        = 3;
    uint256 constant RESERVE          = 4;
    uint256 constant ENFORCEMENT      = 5;
    uint256 constant ACCOUNT_RECOVERY = 6;
    uint256 constant DELEGATION       = 7;

    // bucket identifiers
    uint256 constant STATUS                   = 1;

    uint256 constant BALANCE                  = 2;
    uint256 constant DELEGATED_TRANSFER_NONCE = 3;
    uint256 constant RECOVERY_ACCOUNT         = 4;
    uint256 constant TOTAL_SUPPLY             = 5;

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
    address public cryptoFiat;

    modifier onlyMasterAccount {
        if (CryptoFiat(cryptoFiat).masterAccount() != msg.sender)
            revert();
        _;
    }
    modifier onlyContracts {
        if (!CryptoFiat(cryptoFiat).contractActive(msg.sender))
            revert();
        _;
    }
    function switchCryptoFiat(address next) onlyMasterAccount { cryptoFiat = next; }

    function contractAddress(uint256 id) constant internal returns (address) { return CryptoFiat(cryptoFiat).contractAddress(id); }

    function data() constant internal returns (Data) { return Data(contractAddress(DATA)); }
    function accounts() constant internal returns (Accounts) { return Accounts(contractAddress(ACCOUNTS)); }
    function approving() constant internal returns (Approving) { return Approving(contractAddress(APPROVING)); }
    function reserve() constant internal returns (Reserve) { return Reserve(contractAddress(RESERVE)); }
    function enforcement() constant internal returns (Enforcement) { return Enforcement(contractAddress(ENFORCEMENT)); }
    function accountRecovery() constant internal returns (AccountRecovery) { return AccountRecovery(contractAddress(ACCOUNT_RECOVERY)); }
    function delegation() constant internal returns (Delegation) { return Delegation(contractAddress(DELEGATION)); }
}

contract Data is Relay {
    function Data(address _cryptoFiat) {
        cryptoFiat = _cryptoFiat;
    }

    mapping(bytes32 => bytes32) private _data;

    function set(uint256 bucket, bytes32 key, bytes32 value)
        onlyContracts
    {
        _data[sha3(bucket, key)] = value;
    }

    function get(uint256 bucket, bytes32 key)
        constant
        returns (bytes32)
    {
        return _data[sha3(bucket, key)];
    }
}


contract InternalData is Constants, Relay {
    // balance contains the balance of an account
    function _balanceOf(address addr) constant internal returns (uint256) { return uint256(data().get(BALANCE, bytes32(addr))); }
    function _setBalanceOf(address addr, uint256 value) internal { data().set(BALANCE, bytes32(addr), bytes32(value)); }

    // state contains the current state of an account
    function _statusOf(address addr) constant internal returns (uint256) { return uint256(data().get(STATUS, bytes32(addr))); }
    function _setStatusOf(address addr, uint256 value) internal { data().set(STATUS, bytes32(addr), bytes32(value)); }

    // delegated trancfer nonce contains the last nonce used in delegatedTransfer
    function _delegatedTransferNonceOf(address addr) constant internal returns (uint256) { return uint256(data().get(DELEGATED_TRANSFER_NONCE, bytes32(addr))); }
    function _setDelegatedTransferNonceOf(address addr, uint256 value) internal { data().set(DELEGATED_TRANSFER_NONCE, bytes32(addr), bytes32(value)); }

    // recovery account contains a fallback account that can be used to recover funds
    function _recoveryAccountOf(address addr) constant internal returns (address) { return address(data().get(RECOVERY_ACCOUNT, bytes32(addr))); }
    function _setRecoveryAccountOf(address addr, address value) internal { data().set(RECOVERY_ACCOUNT, bytes32(addr), bytes32(value)); }

    // totalSupply is the total amount of tokens in circulation
    function _totalSupply() constant internal returns (uint256) { return uint256(data().get(TOTAL_SUPPLY, bytes32(0))); }
    function _setTotalSupply(uint256 value) internal { data().set(TOTAL_SUPPLY, bytes32(0), bytes32(value)); }

    // for checking account status
    function _isApproved(address account) constant internal returns (bool) { return _statusOf(account) & APPROVED == APPROVED; }
    function _isClosed(address account)   constant internal returns (bool) { return _statusOf(account) & CLOSED == CLOSED; }
    function _isFrozen(address account)   constant internal returns (bool) { return _statusOf(account) & FROZEN == FROZEN; }

    modifier canSend(address account) {
        if (!_isApproved(account))
            revert();
        if (_isFrozen(account))
            revert();

        if (account == 0)
            revert();
        _;
    }
    function assertSend(address account) constant internal canSend(account) {}

    modifier canReceive(address account) {
        if (_isClosed(account))
            revert();
        if (account == 0)
            revert();
        _;
    }
    function assertReceive(address account) constant internal canReceive(account) {}

    // internal modification of balance
    function _withdraw(address account, uint256 amount) internal {
        // check for underflow
        uint256 balance = _balanceOf(account);
        if (balance < amount)
            revert();
        _setBalanceOf(account, balance - amount);
    }

    function _deposit(address account, uint256 amount) internal {
        // check for overflow
        uint256 balance = _balanceOf(account);
        if (balance + amount < balance)
            revert();
        _setBalanceOf(account, balance + amount);
    }
}

contract Accounts is InternalData {
    function Accounts(address _cryptoFiat) {
        cryptoFiat = _cryptoFiat;
    }

    // balance contains the balance of an account
    function balanceOf(address addr) constant returns (uint256) { return _balanceOf(addr); }
    function statusOf(address addr) constant returns (uint256) { return _statusOf(addr); }

    function isApproved(address account) constant returns (bool) { return _isApproved(account); }
    function isClosed(address account)   constant returns (bool) { return _isClosed(account); }
    function isFrozen(address account)   constant returns (bool) { return _isFrozen(account); }

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
    function Approving(address _cryptoFiat, address _accountApprover) {
        cryptoFiat = _cryptoFiat;
        accountApprover = _accountApprover;
    }
    modifier onlyAccountApprover {
        if (msg.sender != accountApprover)
            revert();
        _;
    }
    function appointAccountApprover(address next) onlyAccountApprover {
        accountApprover = next;
    }

    // approveAccount makes it possible for an account to send money
    function approveAccount(address account) onlyAccountApprover {
        _setStatusOf(account, _statusOf(account) | APPROVED);
        AccountApproved(account);
    }

    // approveAccounts approves multiple accounts
    function approveAccounts(address[] accounts) {
        for (uint i = 0; i < accounts.length; i += 1) {
            approveAccount(accounts[i]);
        }
    }

    // closeAccount closes the account for receiving money
    function closeAccount(address account) onlyAccountApprover {
        _setStatusOf(account, _statusOf(account) | CLOSED);
        AccountClosed(account);
    }
}


contract Reserve is InternalData {
    address public reserveBank;
    function Reserve(address _cryptoFiat, address _reserveBank) {
        cryptoFiat = _cryptoFiat;
        reserveBank = _reserveBank;
    }
    modifier onlyReserveBank {
        if (msg.sender != reserveBank)
            revert();
        _;
    }
    function appointReserveBank(address next) onlyReserveBank { reserveBank = next; }

    function totalSupply() constant returns (uint256) { return _totalSupply(); }

    // increaseSupply increases the tokens in circulation
    function increaseSupply(uint256 amount)
        onlyReserveBank
        canReceive(reserveBank)
    {
        uint256 supply = totalSupply();
        if (supply + amount < supply)
            revert();
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
        if (supply < amount)
            revert(); // invalid state
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

    function Enforcement(address _cryptoFiat, address _lawEnforcer, address _enforcementAccountDesignator, address _enforcementAccount) {
        cryptoFiat = _cryptoFiat;
        lawEnforcer = _lawEnforcer;

        accountDesignator = _enforcementAccountDesignator;
        account = _enforcementAccount;
    }
    modifier onlyLawEnforcer {
        if (msg.sender != lawEnforcer)
            revert();
        _;
    }
    modifier onlyAccountDesignator {
        if (msg.sender != accountDesignator)
            revert();
        _;
    }

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
        _setStatusOf(target, _statusOf(target) | FROZEN);
        AccountFreeze(target, true);
    }

    // unfreezeAccount re-allows account to send money
    function unfreezeAccount(address target) onlyLawEnforcer {
        _setStatusOf(target, _statusOf(target) & ~FROZEN);
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
    function AccountRecovery(address _cryptoFiat) {
        cryptoFiat = _cryptoFiat;
    }

    // designateRecoveryAccount allows msg.sender to specify a trusted account
    // that can recover the tokens
    function designateRecoveryAccount(address recoveryAccount) {
        _setRecoveryAccountOf(msg.sender, recoveryAccount);
    }

    // recoverBalance allows to recover tokens on a particular account
    // recovering a balance will automatically close the account
    function recoverBalance(address from, address into)
        canSend(from)
        canReceive(into)
    {
        if (msg.sender != _recoveryAccountOf(from))
            revert();

        // close the account
        _setStatusOf(from, _statusOf(from) | CLOSED);
        AccountClosed(from);

        uint256 amount = _balanceOf(from);

        _withdraw(from, amount);
        _deposit(into, amount);

        Transfer(from, into, amount);
    }
}

contract Delegation is InternalData {
    function Delegation(address _cryptoFiat) {
        cryptoFiat = _cryptoFiat;
    }

    function nonceOf(address account) constant returns (uint256) {
        return _delegatedTransferNonceOf(account);
    }

    function recoverSigner(bytes32 hash, bytes signature)
        internal
        returns (address)
    {
        bytes32 r;
        bytes32 s;
        uint8 v;

        if (signature.length != 65)
            revert();

        assembly {
            r := mload(add(signature, 32))
            s := mload(add(signature, 64))
            v := and(mload(add(signature, 65)), 255)
        }

        if (v < 27) {
            v += 27;
        }

        return ecrecover(hash, v, r, s);
    }

    function transfer(
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

        if (_delegatedTransferNonceOf(source) >= nonce)
            revert();
        _setDelegatedTransferNonceOf(source, nonce);

        _withdraw(source, amount + fee);
        _deposit(destination, amount);

        Transfer(source, destination, amount);
        if (fee > 0) {
            _deposit(delegate, fee);
            Transfer(source, delegate, fee);
        }
    }

    uint constant XFER_SIZE = 32+32+32+32+32+32+1;
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
        uint base = XFER_SIZE * offset;

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
        if (v < 27) {
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

    function multitransfer(
        uint256 count,
        bytes   transfers,
        address delegate
    ) {
        for (uint i = 0; i < count; i++) {
            Xfer memory xfer = recoverXfer(transfers, i);

            // check whether source can send
            assertSend(xfer.source);

            // protect against replayed transactions
            if (_delegatedTransferNonceOf(xfer.source) >= xfer.nonce)
                revert();
            _setDelegatedTransferNonceOf(xfer.source, xfer.nonce);

            _withdraw(xfer.source, xfer.amount + xfer.fee);
            _deposit(xfer.destination, xfer.amount);

            Transfer(xfer.source, xfer.destination, xfer.amount);
            if (xfer.fee > 0) {
                _deposit(delegate, xfer.fee);
                Transfer(xfer.source, delegate, xfer.fee);
            }
        }
    }
}
