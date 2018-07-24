pragma solidity ^0.4.24;

contract CryptoFiat {
    address public masterAccount;
    modifier onlyMasterAccount {
        require(msg.sender == masterAccount);
        _ ;
    } 

    // lookup table for finding a particular contract by ID
    mapping(uint256 => address) public contractAddress;
    mapping(address => uint256) public contractId;

    function contractActive(address addr) public view returns (bool) { return contractId[addr] > 0; }

    // list of contracts involved in this CryptoFiat instance
    // use this list of contracts for filtering for events
    address[] public contracts;
    function contractsLength() public view returns (uint256) { return contracts.length; }

    constructor() public {
        masterAccount = msg.sender;
        contracts.push(address(this));
    }

    function appointMasterAccount(address next) public onlyMasterAccount { masterAccount = next; }

    event ContractUpgraded(uint256 indexed id, address previous, address next);
    function upgrade(uint256 id, address next) public {
        require(id != 0);
        address prev = contractAddress[id];

        require(prev != next);

        // message sender or the previous contract
        bool canUpgrade = (msg.sender == masterAccount) || (msg.sender == prev);
        require(canUpgrade);
        // check double use of contract
        require(!contractActive(next));

        // disable previous contract
        contractId[prev] = 0;

        // activate next contract
        contractAddress[id] = next;
        if (next != 0x0000000000000000000000000000000000000000)
            contractId[next] = id;

        // finalize
        emit ContractUpgraded(id, prev, next);
        contracts.push(next);
    }
}

contract Data {
    CryptoFiat public cryptoFiat;

    constructor(CryptoFiat _cryptoFiat) public {
        cryptoFiat = _cryptoFiat;
    }

    modifier onlyContracts {
        require(cryptoFiat.contractActive(msg.sender));
        _;
    }
    
    mapping(bytes32 => bytes32) private _data;

    function set(uint256 bucket, bytes32 key, bytes32 value)
        public
        onlyContracts
    {
        _data[keccak256(abi.encodePacked(bucket, key))] = value;
    }

    function get(uint256 bucket, bytes32 key)
        public view
        returns (bytes32)
    {
        return _data[keccak256(abi.encodePacked(bucket, key))];
    }
}

contract Constants {
    address constant WORLD = 0x0000000000000000000000000000000000000000;

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

contract InternalData is Constants {
    CryptoFiat public cryptoFiat;
    Data       public data;

    modifier onlyMasterAccount {
        require(cryptoFiat.masterAccount() == msg.sender);
        _;
    }
    function switchCryptoFiat(CryptoFiat next) public onlyMasterAccount { cryptoFiat = next; }
    function switchData(Data next) public onlyMasterAccount { data = next; }
    function cacheData() internal {
        data = Data(contractAddress(DATA));
        require(address(data) != WORLD);
    }

    function contractAddress(uint256 id) internal view returns (address) {
        return cryptoFiat.contractAddress(id);
    }

    function accounts() internal view returns (Accounts) { return Accounts(contractAddress(ACCOUNTS)); }
    function approving() internal view returns (Approving) { return Approving(contractAddress(APPROVING)); }
    function reserve() internal view returns (Reserve) { return Reserve(contractAddress(RESERVE)); }
    function enforcement() internal view returns (Enforcement) { return Enforcement(contractAddress(ENFORCEMENT)); }
    function accountRecovery() internal view returns (AccountRecovery) { return AccountRecovery(contractAddress(ACCOUNT_RECOVERY)); }
    function delegation() internal view returns (Delegation) { return Delegation(contractAddress(DELEGATION)); }

    // balance contains the balance of an account
    function _balanceOf(address addr) internal view returns (uint256) { return uint256(data.get(BALANCE, bytes20(addr))); }
    function _setBalanceOf(address addr, uint256 value) internal { data.set(BALANCE, bytes20(addr), bytes32(value)); }

    // state contains the current state of an account
    function _statusOf(address addr) internal view returns (uint256) { return uint256(data.get(STATUS, bytes20(addr))); }
    function _setStatusOf(address addr, uint256 value) internal { data.set(STATUS, bytes20(addr), bytes32(value)); }

    // delegated trancfer nonce contains the last nonce used in delegatedTransfer
    function _delegatedTransferNonceOf(address addr) internal view returns (uint256) { return uint256(data.get(DELEGATED_TRANSFER_NONCE, bytes20(addr))); }
    function _setDelegatedTransferNonceOf(address addr, uint256 value) internal { data.set(DELEGATED_TRANSFER_NONCE, bytes20(addr), bytes32(value)); }

    // recovery account contains a fallback account that can be used to recover funds
    function _recoveryAccountOf(address addr) internal view returns (address) { return address(bytes20(data.get(RECOVERY_ACCOUNT, bytes20(addr)))); }
    function _setRecoveryAccountOf(address addr, address value) internal { data.set(RECOVERY_ACCOUNT, bytes20(addr), bytes20(value)); }

    // totalSupply is the total amount of tokens in circulation
    function _totalSupply() internal view returns (uint256) { return uint256(data.get(TOTAL_SUPPLY, bytes32(0))); }
    function _setTotalSupply(uint256 value) internal { data.set(TOTAL_SUPPLY, bytes32(0), bytes32(value)); }

    // for checking account status
    function _isApproved(address account) internal view returns (bool) { return _statusOf(account) & APPROVED == APPROVED; }
    function _isClosed(address account)   internal view returns (bool) { return _statusOf(account) & CLOSED == CLOSED; }
    function _isFrozen(address account)   internal view returns (bool) { return _statusOf(account) & FROZEN == FROZEN; }

    modifier canSend(address account) {
        uint256 status = _statusOf(account);
        require(status & APPROVED == APPROVED);
        require(status & FROZEN   != FROZEN);
        require(account != WORLD);
        _;
    }
    function assertSend(address account) internal view canSend(account) {}

    modifier canReceive(address account) {
        require(!_isClosed(account));
        require(account != WORLD);
        _;
    }
    function assertReceive(address account) internal view canReceive(account) {}

    // internal modification of balance
    function _withdraw(address account, uint256 amount) internal {
        // check for underflow
        uint256 balance = _balanceOf(account);
        require(balance >= amount);
        _setBalanceOf(account, balance - amount);
    }

    function _deposit(address account, uint256 amount) internal {
        // check for overflow
        uint256 balance = _balanceOf(account);
        require(balance + amount >= balance);
        _setBalanceOf(account, balance + amount);
    }
}

contract Accounts is InternalData {
    constructor(CryptoFiat _cryptoFiat) public {
        cryptoFiat = _cryptoFiat;
        cacheData();
    }

    // balance contains the balance of an account
    function balanceOf(address addr) public view returns (uint256) { return _balanceOf(addr); }
    function statusOf(address addr)  public view returns (uint256) { return _statusOf(addr); }

    function isApproved(address account) public view returns (bool) { return _isApproved(account); }
    function isClosed(address account)   public view returns (bool) { return _isClosed(account); }
    function isFrozen(address account)   public view returns (bool) { return _isFrozen(account); }

    function transfer(address destination, uint256 amount)
        public
        canSend(msg.sender)
        canReceive(destination)
    {
        address source = msg.sender;

        _withdraw(source, amount);
        _deposit(destination, amount);
        emit Transfer(source, destination, amount);
    }
}

contract Approving is InternalData {
    address public accountApprover;
    constructor(CryptoFiat _cryptoFiat, address _accountApprover) public {
        cryptoFiat = _cryptoFiat;
        cacheData();
        accountApprover = _accountApprover;
    }
    modifier onlyAccountApprover {
        require(msg.sender == accountApprover);
        _;
    }
    function appointAccountApprover(address next) public onlyAccountApprover {
        accountApprover = next;
    }

    // approveAccount makes it possible for an account to send money
    function approveAccount(address account) public onlyAccountApprover {
        _setStatusOf(account, _statusOf(account) | APPROVED);
        emit AccountApproved(account);
    }

    // approveAccounts approves multiple accounts
    function approveAccounts(address[] accounts) public {
        for (uint i = 0; i < accounts.length; i += 1) {
            approveAccount(accounts[i]);
        }
    }

    // closeAccount closes the account for receiving money
    function closeAccount(address account) public onlyAccountApprover {
        _setStatusOf(account, _statusOf(account) | CLOSED);
        emit AccountClosed(account);
    }
}


contract Reserve is InternalData {
    address public reserveBank;
    constructor(CryptoFiat _cryptoFiat, address _reserveBank) public {
        cryptoFiat = _cryptoFiat;
        cacheData();
        reserveBank = _reserveBank;
    }
    modifier onlyReserveBank {
        require(msg.sender == reserveBank);
        _;
    }
    function appointReserveBank(address next) public onlyReserveBank { reserveBank = next; }

    function totalSupply() public view returns (uint256) { return _totalSupply(); }

    // increaseSupply increases the tokens in circulation
    function increaseSupply(uint256 amount)
        public
        onlyReserveBank
        canReceive(reserveBank)
    {
        uint256 supply = totalSupply();
        require(supply + amount >= supply);
        supply += amount;
        _setTotalSupply(supply);

        _deposit(reserveBank, amount);
        emit Transfer(WORLD, reserveBank, amount);

        emit SupplyChanged(supply);
    }

    // decreaseSupply decreases the amount of tokens in circulation
    function decreaseSupply(uint256 amount)
        public
        onlyReserveBank
        canSend(reserveBank)
    {
        uint256 supply = _totalSupply();
        require(supply >= amount);
        supply -= amount;
        _setTotalSupply(supply);

        _withdraw(reserveBank, amount);
        emit Transfer(reserveBank, WORLD, amount);

        emit SupplyChanged(supply);
    }
}


contract Enforcement is InternalData {
    address public lawEnforcer;

    address public accountDesignator;
    address public account;

    constructor(CryptoFiat _cryptoFiat, address _lawEnforcer, address _enforcementAccountDesignator, address _enforcementAccount) public {
        cryptoFiat = _cryptoFiat;
        cacheData();
        lawEnforcer = _lawEnforcer;

        accountDesignator = _enforcementAccountDesignator;
        account = _enforcementAccount;
    }
    modifier onlyLawEnforcer {
        require(msg.sender == lawEnforcer);
        _;
    }
    modifier onlyAccountDesignator {
        require(msg.sender == accountDesignator);
        _;
    }

    function appointLawEnforcer(address next) public onlyLawEnforcer { lawEnforcer = next; }
    function appointAccountDesignator(address next) public onlyAccountDesignator { accountDesignator = next; }

    // withdraw allows law enforcerer to withdraw to a dedicated account
    function withdraw(address from, uint256 amount)
        public
        onlyLawEnforcer
        canReceive(account)
    {
        _withdraw(from, amount);
        _deposit(account, amount);
        emit Transfer(from, account, amount);
    }

    // freezeAccount disallows account to send money
    function freezeAccount(address target) public onlyLawEnforcer {
        _setStatusOf(target, _statusOf(target) | FROZEN);
        emit AccountFreeze(target, true);
    }

    // unfreezeAccount re-allows account to send money
    function unfreezeAccount(address target) public onlyLawEnforcer {
        _setStatusOf(target, _statusOf(target) & ~FROZEN);
        emit AccountFreeze(target, false);
    }

    // designateAccount allows changing the account
    function designateAccount(address _account)
        public
        onlyAccountDesignator
        canReceive(_account)
    {
        account = _account;
    }
}

contract AccountRecovery is InternalData {
    constructor(CryptoFiat _cryptoFiat) public {
        cryptoFiat = _cryptoFiat;
        cacheData();
    }

    // designateRecoveryAccount allows msg.sender to specify a trusted account
    // that can recover the tokens
    function designateRecoveryAccount(address recoveryAccount) public {
        _setRecoveryAccountOf(msg.sender, recoveryAccount);
    }

    // recoverBalance allows to recover tokens on a particular account
    // recovering a balance will automatically close the account
    function recoverBalance(address from, address into)
        public
        canSend(from)
        canReceive(into)
    {
        require(msg.sender == _recoveryAccountOf(from));

        // close the account
        _setStatusOf(from, _statusOf(from) | CLOSED);
        emit AccountClosed(from);

        uint256 amount = _balanceOf(from);

        _withdraw(from, amount);
        _deposit(into, amount);

        emit Transfer(from, into, amount);
    }
}

contract Delegation is InternalData {
    constructor(CryptoFiat _cryptoFiat) public {
        cryptoFiat = _cryptoFiat;
        cacheData();
    }

    function nonceOf(address account) public view returns (uint256) {
        return _delegatedTransferNonceOf(account);
    }

    function recoverSigner(bytes32 hash, bytes signature)
        internal pure
        returns (address)
    {
        bytes32 r;
        bytes32 s;
        uint8 v;

        require(signature.length == 65);

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
        public
        canReceive(destination)
        canReceive(delegate)
    {
        // extract source from signature
        address source = recoverSigner(
            keccak256(abi.encodePacked(nonce, destination, amount, fee)),
            signature
        );

        // check whether source can send
        assertSend(source);

        // protect against replayed transactions

        require(_delegatedTransferNonceOf(source) < nonce);
        _setDelegatedTransferNonceOf(source, nonce);

        _withdraw(source, amount + fee);
        _deposit(destination, amount);

        emit Transfer(source, destination, amount);
        if (fee > 0) {
            _deposit(delegate, fee);
            emit Transfer(source, delegate, fee);
        }
    }

    uint constant XFER_SIZE = 32+20+32+32+(32+32+1);
    // expected format
    // struct XferEncoded {
    //     uint256 nonce;       // 32 =  32
    //     address destination; // 20 =  52
    //     uint256 amount;      // 32 =  84
    //     uint256 fee;         // 32 = 116
    //     bytes32 r;           // 32 = 148
    //     bytes32 s;           // 32 = 180
    //     uint8   v;           //  1 = 181
    // }

    struct Xfer {
        uint256 nonce;
        address source;
        address destination;
        uint256 amount;
        uint256 fee;
    }

     function recoverXfer(bytes data, uint offset)
        internal pure
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
            data := add(data, base)
            nonce := mload(add(data, 32))
            destination := mload(add(data, 52))
            amount := mload(add(data, 84))
            fee := mload(add(data, 116))
            r := mload(add(data, 148))
            s := mload(add(data, 180))
            v := and(mload(add(data, 181)), 255)
        }
        if (v < 27) {
            v += 27;
        }

        bytes32 hash = keccak256(abi.encodePacked(nonce, destination, amount, fee));

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
    )
        public
        canReceive(delegate)
    {
        for (uint i = 0; i < count; i++) {
            Xfer memory xfer = recoverXfer(transfers, i);

            // check whether source can send
            assertSend(xfer.source);
            assertReceive(xfer.destination);

            // protect against replayed transactions
            require(_delegatedTransferNonceOf(xfer.source) < xfer.nonce);
            _setDelegatedTransferNonceOf(xfer.source, xfer.nonce);

            _withdraw(xfer.source, xfer.amount + xfer.fee);
            _deposit(xfer.destination, xfer.amount);

            emit Transfer(xfer.source, xfer.destination, xfer.amount);
            if (xfer.fee > 0) {
                _deposit(delegate, xfer.fee);
                emit Transfer(xfer.source, delegate, xfer.fee);
            }
        }
    }
}
