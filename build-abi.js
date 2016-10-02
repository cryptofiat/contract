// loadScript("abi.js");

const ROOT = "0xB8E7245B42529B905a8909B8FD5fC690097e71f7";

const fs = require("fs");

var abijs = "";

function emit(line) {
	abijs += line + "\n";
}

var Constants = {
	Data: 0,
	Accounts: 1,
	Approving: 2,
	Reserve: 3,
	Enforcement: 4,
	AccountRecovery: 5,
	Delegation: 6,
	MultiDelegation: 7
};

emit("// GENERATED CODE");
emit("// DO NOT MODIFY");
emit("var Contract = {};")

fs.readdirSync(".bin").forEach(function(file) {
	if (!file.endsWith(".abi")) {
		return;
	}
	var name = file.substr(0, file.length - 4);

	var data = fs.readFileSync(".bin\\" + file, "utf8");

	if (Constants.hasOwnProperty(name) || (name == "CryptoFiat")) {
		fs.writeFileSync("abi\\" + file, data);
	}
	emit("Contract['" + name + "'] = web3.eth.contract(" + data.trim() + ");");
});

emit("");
emit("var CryptoFiat = Contract.CryptoFiat.at('" + ROOT + "');");

for (var name in Constants) {
	emit("var " + name + " = Contract.Accounts.at(CryptoFiat.contracts(" + Constants[name] + "));");
}

fs.writeFileSync("cryptofiat.abi.js", abijs);