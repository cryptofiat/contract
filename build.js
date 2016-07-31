const fs = require("fs");
const solc = require("solc");

var contracts = {};

{
	console.log("");
	console.log("====================");
	console.log("Loading contract files: ")

	fs.readdirSync(".").forEach(function(file) {
		if (!file.endsWith(".sol")) {
			return;
		}

		contracts[file] = fs.readFileSync(file, "utf8");
		console.log("\t" + file);
	});
}

{
	console.log("Compiling:");

	var compiled = solc.compile({
		sources: contracts
	}, 1);

	if (compiled.errors) {
		console.log("ERRORS:");
		compiled.errors.forEach(function(error) {
			console.error(error);
		})
	}
}

{
	try {
		// fs.rmdirSync(".bin");
		fs.mkdirSync(".bin", 0o755);
	} catch (ignore) {};

	console.log("Writing:");
	for (var name in compiled.contracts) {
		console.log("\t" + name);
		var contract = compiled.contracts[name];

		// var buffer = Buffer.from(contract.bytecode, "hex");
		// fs.writeFileSync(".bin/" + name + ".bin", buffer);
		fs.writeFileSync(".bin/" + name + ".abi", contract.interface);

		var log = "Gas estimates:\n";

		var estimates = contract.gasEstimates;
		log += "\tcreation: " + estimates.creation + "\n";
		log += "\tinternal:\n";
		for (var fn in estimates.internal) {
			var est = estimates.internal[fn];
			log += "\t\t" + fn + ": " + (est || "delegated") + "\n";
		};

		log += "\texternal:\n";
		for (var fn in estimates.external) {
			var est = estimates.external[fn];
			log += "\t\t" + fn + ": " + (est || "delegated") + "\n";
		};

		fs.writeFileSync(".bin/" + name + ".log", log);
	}
}

console.log("Done");