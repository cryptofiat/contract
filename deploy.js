// npm install -g ethereumjs-testrpc
// testrpc
//

const fs = require("fs");
const Web3 = require("web3");
const TestRPC = require("ethereumjs-testrpc");

var web3 = new Web3();
web3.setProvider(TestRPC.provider());
web3.eth.getAccounts(function(err, accounts) {
	var GAS = 2000000;
	console.log(accounts);
	var MasterAccount = accounts[0];

	var SETUP = {
		MasterAccount: MasterAccount,

		AccountApprover: MasterAccount,
		ReserveBank: MasterAccount,

		LawEnforcer: MasterAccount,
		EnforcementAccountDesignator: MasterAccount,
		EnforcementAccount: MasterAccount
	};

	function Work() {
		this.schedule = [];
	}
	Work.prototype = {
		add: function(fn) {
			this.schedule.push(fn);
		},
		run: function() {
			var self = this;
			var fn = self.schedule.shift();
			if (fn) {
				fn(function() {
					self.run();
				});
			} else {
				console.log("<<< DONE >>>");
			}
		}
	};

	function load(filename, success) {
		fs.readFile(filename, "utf8", function(err, data) {
			if (err) {
				console.log(err);
				throw err;
			}
			success(data);
		})
	}

	function countdown(number, done) {
		if (number <= 0) {
			done();
			return function() {};
		}
		return function() {
			number--;
			if (number == 0) {
				done();
			}
		}
	}

	var work = new Work();
	var $Contract = {};

	([
		"CryptoFiat",
		"Data",
		"AccountRecovery",
		"Accounts",
		"Approving",
		"Delegation",
		"Enforcement",
		"Reserve"
	]).map(function(name) {
		$Contract[name] = {
			contract: null,
			address: null,
			abi: null,
			bin: null,
			create: function(args, success) {
				var self = $Contract[name];

				var result = self.abi.new({
					from: SETUP.MasterAccount,
					data: self.bin,
					gas: GAS
				}, function(err, contract) {
					if (err) {
						console.log(err);
						throw err;
					}
					if (!contract.address) {

					} else {
						console.log(contract);
						console.log("Contract " + name + " deployed to " + contract.address);
						self.contract = contract;
						self.address = contract.address;
						success(contract);
					}
				});
			}
		};

		work.add(function(next) {
			load(".bin/" + name + ".abi", function(data) {
				$Contract[name].abi = web3.eth.contract(JSON.parse(data));
				load(".bin/" + name + ".bin", function(data) {
					$Contract[name].bin = data;
					next();
				})
			});
		})
	});

	work.add(function(next) {
		$Contract.CryptoFiat.create([], next);
	});

	work.add(function(next) {
		var sync = countdown(7, next);

		$Contract.Data.create([
			$Contract.CryptoFiat.address
		], sync);

		$Contract.Accounts.create([
			$Contract.CryptoFiat.address
		], sync);

		$Contract.Approving.create([
			$Contract.CryptoFiat.address,
			SETUP.AccountApprover
		], sync);

		$Contract.Reserve.create([
			$Contract.CryptoFiat.address,
			SETUP.ReserveBank
		], sync);

		$Contract.Enforcement.create([
			$Contract.CryptoFiat.address,
			SETUP.LawEnforcer,
			SETUP.EnforcementAccount,
			SETUP.EnforcementAccountDesignator
		], sync);

		$Contract.AccountRecovery.create([
			$Contract.CryptoFiat.address
		], sync);

		$Contract.Delegation.create([
			$Contract.CryptoFiat.address
		], sync);
	});

	work.add(function(next) {
		var sync = countdown(7, next);

		$Contract.CryptoFiat.upgrade(0, $Contract.Data.address, sync);
		$Contract.CryptoFiat.upgrade(1, $Contract.Accounts.address, sync);
		$Contract.CryptoFiat.upgrade(2, $Contract.Approving.address, sync);
		$Contract.CryptoFiat.upgrade(3, $Contract.Reserve.address, sync);
		$Contract.CryptoFiat.upgrade(4, $Contract.Enforcement.address, sync);
		$Contract.CryptoFiat.upgrade(5, $Contract.AccountRecovery.address, sync);
		$Contract.CryptoFiat.upgrade(6, $Contract.Delegation.address, sync);
	});

	work.run();
});