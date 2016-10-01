// loadScript("abi.js");

const ROOT = "0x7f99ACA8A0C23f107729aa9929D9Dc1481B087F6";

const fs = require("fs");

var abijs = "";

function emit(line) {
	abijs += line + "\n";
}

emit("// GENERATED CODE");
emit("// DO NOT MODIFY");
emit("var Contract = {};")

fs.readdirSync(".bin").forEach(function(file) {
	if (!file.endsWith(".abi")) {
		return;
	}
	var name = file.substr(0, file.length - 4);

	var data = fs.readFileSync(".bin\\" + file, "utf8");
	emit("Contract['" + name + "'] = web3.eth.contract(" + data.trim() + ");");
});

emit("");
emit("var CryptoFiat = Contract.CryptoFiat.at(ROOT || '" + ROOT + "');");
emit("var Accounts = Contract.Accounts.at(CryptoFiat.contracts(1));");
emit("var Approving = Contract.Approving.at(CryptoFiat.contracts(2));");
emit("var Reserve = Contract.Reserve.at(CryptoFiat.contracts(3));");
emit("var Enforcement = Contract.Enforcement.at(CryptoFiat.contracts(4));");
emit("var AccountRecovery = Contract.AccountRecovery.at(CryptoFiat.contracts(5));");
emit("var Delegation = Contract.Delegation.at(CryptoFiat.contracts(6));");
emit("var MultiDelegation = Contract.MultiDelegation.at(CryptoFiat.contracts(7));");

fs.writeFileSync("cryptofiat.abi.js", abijs);