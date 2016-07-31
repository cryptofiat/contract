// Mintable allows an account to be responsible for minting.
contract Mintable {
    address public centralMinter;

    function Mintable() {
        centralMinter = msg.sender;
    }

    modifier onlyMinter {
        if (msg.sender != centralMinter) throw;
        _
    }

    function transferMinter(address newMinter)
        onlyMinter
    {
        centralMinter = newMinter;
    }
}
