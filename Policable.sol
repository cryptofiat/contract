contract Policable {
    address public lawEnforcement;

    function Policable() {
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