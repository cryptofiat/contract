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
		fs.mkdirSync(".bin", 0o755);
	} catch (ignore) {};

	console.log("Writing:");
	for (var name in compiled.contracts) {
		console.log("\t" + name);
		var contract = compiled.contracts[name];
		fs.writeFileSync(".bin/" + name + ".bin", contract.bin);
		fs.writeFileSync(".bin/" + name + ".abi", contract.abi);
	}
}

console.log("Done");