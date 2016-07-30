// Mintable allows an account to be responsible for minting.
contract Mintable {
    address public centralBank;

    function Mintable() {
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