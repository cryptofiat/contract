
// This file is an automatically generated Java binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package eu.cryptoeuro.contract;

import org.ethereum.geth.*;
import org.ethereum.geth.internal.*;


	public class AccountRecovery {
		// ABI is the input ABI used to generate the binding from.
		public final static String ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"into\",\"type\":\"address\"}],\"name\":\"recoverBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recoveryAccount\",\"type\":\"address\"}],\"name\":\"designateRecoveryAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x608060405234801561001057600080fd5b50604051602080610ae7833981016040525160008054600160a060020a031916600160a060020a03831617905561004e640100000000610054810204565b5061012c565b6100676001640100000000610093810204565b60018054600160a060020a031916600160a060020a03928316179081905516151561009157600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b1580156100fa57600080fd5b505af115801561010e573d6000803e3d6000fd5b505050506040513d602081101561012457600080fd5b505192915050565b6109ac8061013b6000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632d1c5ff8811461007c5780633363375c146100b2578063516c4b84146100e057806373d4a13a1461011e578063f1375f3814610133578063fb55a05514610161575b600080fd5b34801561008857600080fd5b506100b073ffffffffffffffffffffffffffffffffffffffff6004358116906024351661018f565b005b3480156100be57600080fd5b506100b073ffffffffffffffffffffffffffffffffffffffff60043516610322565b3480156100ec57600080fd5b506100f5610451565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561012a57600080fd5b506100f561046d565b34801561013f57600080fd5b506100b073ffffffffffffffffffffffffffffffffffffffff60043516610489565b34801561016d57600080fd5b506100b073ffffffffffffffffffffffffffffffffffffffff60043516610496565b600082600061019d826105c5565b90506001808216146101ae57600080fd5b600481811614156101be57600080fd5b73ffffffffffffffffffffffffffffffffffffffff821615156101e057600080fd5b836101ea81610678565b156101f457600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561021657600080fd5b61021f8661068e565b73ffffffffffffffffffffffffffffffffffffffff16331461024057600080fd5b61025586600261024f896105c5565b1761070c565b60405173ffffffffffffffffffffffffffffffffffffffff8716907fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de90600090a261029f866107aa565b93506102ab8685610829565b6102b58585610854565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef866040518082815260200191505060405180910390a3505050505050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156103be57600080fd5b505af11580156103d2573d6000803e3d6000fd5b505050506040513d60208110156103e857600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461040a57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b610493338261087c565b50565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561053257600080fd5b505af1158015610546573d6000803e3d6000fd5b505050506040513d602081101561055c57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461057e57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60018054604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff84811660248501529051600093919092169163295f36d79160448082019260209290919082900301818787803b15801561064657600080fd5b505af115801561065a573d6000803e3d6000fd5b505050506040513d602081101561067057600080fd5b505192915050565b6000600280610686846105c5565b161492915050565b600154604080517f295f36d700000000000000000000000000000000000000000000000000000000815260048181015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b15801561064657600080fd5b60018054604080517f461b09c0000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff858116602485015260448401859052905191169163461b09c091606480830192600092919082900301818387803b15801561078e57600080fd5b505af11580156107a2573d6000803e3d6000fd5b505050505050565b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b15801561064657600080fd5b6000610834836107aa565b90508181101561084357600080fd5b61084f838383036108fe565b505050565b600061085f836107aa565b905081810181111561087057600080fd5b61084f838383016108fe565b600154604080517f461b09c000000000000000000000000000000000000000000000000000000000815260048181015273ffffffffffffffffffffffffffffffffffffffff858116602483015284811660448301529151919092169163461b09c091606480830192600092919082900301818387803b15801561078e57600080fd5b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff8581166024830152604482018590529151919092169163461b09c091606480830192600092919082900301818387803b15801561078e57600080fd00a165627a7a723058200188215b9ff1ee8d796476ad18cd5b823e46677f13aaeb488c532983cc0d65c00029".getBytes();

			// deploy deploys a new Ethereum contract, binding an instance of AccountRecovery to it.
			public static AccountRecovery deploy(TransactOpts auth, EthereumClient client, Address _cryptoFiat) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				
				  args.set(0, Geth.newInterface()); args.get(0).setAddress(_cryptoFiat);
				
				return new AccountRecovery(Geth.deployContract(auth, ABI, BYTECODE, client, args));
			}

			// Internal constructor used by contract deployment.
			private AccountRecovery(BoundContract deployment) {
				this.Address  = deployment.getAddress();
				this.Deployer = deployment.getDeployer();
				this.Contract = deployment;
			}
		

		// Ethereum address where this contract is located at.
		public final Address Address;

		// Ethereum transaction in which this contract was deployed (if known!).
		public final Transaction Deployer;

		// Contract instance bound to a blockchain address.
		private final BoundContract Contract;

		// Creates a new instance of AccountRecovery, bound to a specific deployed contract.
		public AccountRecovery(Address address, EthereumClient client) throws Exception {
			this(Geth.bindContract(address, ABI, client));
		}

		
			

			// cryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
			//
			// Solidity: function cryptoFiat() constant returns(address)
			public Address cryptoFiat(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "cryptoFiat", args);
				return results.get(0).getAddress();
				
			}
		
			

			// data is a free data retrieval call binding the contract method 0x73d4a13a.
			//
			// Solidity: function data() constant returns(address)
			public Address data(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "data", args);
				return results.get(0).getAddress();
				
			}
		

		
			// designateRecoveryAccount is a paid mutator transaction binding the contract method 0xf1375f38.
			//
			// Solidity: function designateRecoveryAccount(recoveryAccount address) returns()
			public Transaction designateRecoveryAccount(TransactOpts opts, Address recoveryAccount) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(recoveryAccount);
				

				return this.Contract.transact(opts, "designateRecoveryAccount"	, args);
			}
		
			// recoverBalance is a paid mutator transaction binding the contract method 0x2d1c5ff8.
			//
			// Solidity: function recoverBalance(from address, into address) returns()
			public Transaction recoverBalance(TransactOpts opts, Address from, Address into) throws Exception {
				Interfaces args = Geth.newInterfaces(2);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(from);
				args.set(1, Geth.newInterface()); args.get(1).setAddress(into);
				

				return this.Contract.transact(opts, "recoverBalance"	, args);
			}
		
			// switchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
			//
			// Solidity: function switchCryptoFiat(next address) returns()
			public Transaction switchCryptoFiat(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchCryptoFiat"	, args);
			}
		
			// switchData is a paid mutator transaction binding the contract method 0x3363375c.
			//
			// Solidity: function switchData(next address) returns()
			public Transaction switchData(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchData"	, args);
			}
		
	}

	public class Accounts {
		// ABI is the input ABI used to generate the binding from.
		public final static String ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"statusOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFrozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x608060405234801561001057600080fd5b50604051602080610a30833981016040525160008054600160a060020a031916600160a060020a03831617905561004e640100000000610054810204565b5061012c565b6100676001640100000000610093810204565b60018054600160a060020a031916600160a060020a03928316179081905516151561009157600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b1580156100fa57600080fd5b505af115801561010e573d6000803e3d6000fd5b505050506040513d602081101561012457600080fd5b505192915050565b6108f58061013b6000396000f3006080604052600436106100a35763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633363375c81146100a8578063516c4b84146100d8578063673448dd146101165780636943b0171461015857806370a082311461018657806373d4a13a146101c657806397a5d5b5146101db578063a9059cbb14610209578063e58398361461023a578063fb55a05514610268575b600080fd5b3480156100b457600080fd5b506100d673ffffffffffffffffffffffffffffffffffffffff60043516610296565b005b3480156100e457600080fd5b506100ed6103c5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561012257600080fd5b5061014473ffffffffffffffffffffffffffffffffffffffff600435166103e1565b604080519115158252519081900360200190f35b34801561016457600080fd5b5061014473ffffffffffffffffffffffffffffffffffffffff600435166103f2565b34801561019257600080fd5b506101b473ffffffffffffffffffffffffffffffffffffffff600435166103fd565b60408051918252519081900360200190f35b3480156101d257600080fd5b506100ed610408565b3480156101e757600080fd5b506101b473ffffffffffffffffffffffffffffffffffffffff60043516610424565b34801561021557600080fd5b506100d673ffffffffffffffffffffffffffffffffffffffff6004351660243561042f565b34801561024657600080fd5b5061014473ffffffffffffffffffffffffffffffffffffffff6004351661053a565b34801561027457600080fd5b506100d673ffffffffffffffffffffffffffffffffffffffff60043516610545565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561033257600080fd5b505af1158015610346573d6000803e3d6000fd5b505050506040513d602081101561035c57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461037e57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60006103ec82610674565b92915050565b60006103ec8261068a565b60006103ec82610698565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60006103ec82610749565b600033600061043d82610749565b905060018082161461044e57600080fd5b6004818116141561045e57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8216151561048057600080fd5b8461048a8161068a565b1561049457600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156104b657600080fd5b3393506104c384866107ca565b6104cd86866107f5565b8573ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef876040518082815260200191505060405180910390a3505050505050565b60006103ec8261081d565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156105e157600080fd5b505af11580156105f5573d6000803e3d6000fd5b505050506040513d602081101561060b57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461062d57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600060018061068284610749565b161492915050565b600060028061068284610749565b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b15801561071757600080fd5b505af115801561072b573d6000803e3d6000fd5b505050506040513d602081101561074157600080fd5b505192915050565b60018054604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff84811660248501529051600093919092169163295f36d79160448082019260209290919082900301818787803b15801561071757600080fd5b60006107d583610698565b9050818110156107e457600080fd5b6107f08383830361082b565b505050565b600061080083610698565b905081810181111561081157600080fd5b6107f08383830161082b565b600060048061068284610749565b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff8581166024830152604482018590529151919092169163461b09c091606480830192600092919082900301818387803b1580156108ad57600080fd5b505af11580156108c1573d6000803e3d6000fd5b5050505050505600a165627a7a72305820a8a2953760b304b29f6a6fa8db54be1e13c3305a8662e705f94396e3a4c1e0cf0029".getBytes();

			// deploy deploys a new Ethereum contract, binding an instance of Accounts to it.
			public static Accounts deploy(TransactOpts auth, EthereumClient client, Address _cryptoFiat) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				
				  args.set(0, Geth.newInterface()); args.get(0).setAddress(_cryptoFiat);
				
				return new Accounts(Geth.deployContract(auth, ABI, BYTECODE, client, args));
			}

			// Internal constructor used by contract deployment.
			private Accounts(BoundContract deployment) {
				this.Address  = deployment.getAddress();
				this.Deployer = deployment.getDeployer();
				this.Contract = deployment;
			}
		

		// Ethereum address where this contract is located at.
		public final Address Address;

		// Ethereum transaction in which this contract was deployed (if known!).
		public final Transaction Deployer;

		// Contract instance bound to a blockchain address.
		private final BoundContract Contract;

		// Creates a new instance of Accounts, bound to a specific deployed contract.
		public Accounts(Address address, EthereumClient client) throws Exception {
			this(Geth.bindContract(address, ABI, client));
		}

		
			

			// balanceOf is a free data retrieval call binding the contract method 0x70a08231.
			//
			// Solidity: function balanceOf(addr address) constant returns(uint256)
			public BigInt balanceOf(CallOpts opts, Address addr) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(addr);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultBigInt(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "balanceOf", args);
				return results.get(0).getBigInt();
				
			}
		
			

			// cryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
			//
			// Solidity: function cryptoFiat() constant returns(address)
			public Address cryptoFiat(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "cryptoFiat", args);
				return results.get(0).getAddress();
				
			}
		
			

			// data is a free data retrieval call binding the contract method 0x73d4a13a.
			//
			// Solidity: function data() constant returns(address)
			public Address data(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "data", args);
				return results.get(0).getAddress();
				
			}
		
			

			// isApproved is a free data retrieval call binding the contract method 0x673448dd.
			//
			// Solidity: function isApproved(account address) constant returns(bool)
			public bool isApproved(CallOpts opts, Address account) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(account);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultBool(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "isApproved", args);
				return results.get(0).getBool();
				
			}
		
			

			// isClosed is a free data retrieval call binding the contract method 0x6943b017.
			//
			// Solidity: function isClosed(account address) constant returns(bool)
			public bool isClosed(CallOpts opts, Address account) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(account);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultBool(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "isClosed", args);
				return results.get(0).getBool();
				
			}
		
			

			// isFrozen is a free data retrieval call binding the contract method 0xe5839836.
			//
			// Solidity: function isFrozen(account address) constant returns(bool)
			public bool isFrozen(CallOpts opts, Address account) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(account);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultBool(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "isFrozen", args);
				return results.get(0).getBool();
				
			}
		
			

			// statusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
			//
			// Solidity: function statusOf(addr address) constant returns(uint256)
			public BigInt statusOf(CallOpts opts, Address addr) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(addr);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultBigInt(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "statusOf", args);
				return results.get(0).getBigInt();
				
			}
		

		
			// switchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
			//
			// Solidity: function switchCryptoFiat(next address) returns()
			public Transaction switchCryptoFiat(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchCryptoFiat"	, args);
			}
		
			// switchData is a paid mutator transaction binding the contract method 0x3363375c.
			//
			// Solidity: function switchData(next address) returns()
			public Transaction switchData(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchData"	, args);
			}
		
			// transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
			//
			// Solidity: function transfer(destination address, amount uint256) returns()
			public Transaction transfer(TransactOpts opts, Address destination, BigInt amount) throws Exception {
				Interfaces args = Geth.newInterfaces(2);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(destination);
				args.set(1, Geth.newInterface()); args.get(1).setBigInt(amount);
				

				return this.Contract.transact(opts, "transfer"	, args);
			}
		
	}

	public class Approving {
		// ABI is the input ABI used to generate the binding from.
		public final static String ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"approveAccounts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountApprover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointAccountApprover\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"closeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"approveAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_accountApprover\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x608060405234801561001057600080fd5b5060405160408061096a83398101604052805160209091015160008054600160a060020a031916600160a060020a03841617905561005564010000000061007b810204565b60028054600160a060020a031916600160a060020a039290921691909117905550610153565b61008e60016401000000006100ba810204565b60018054600160a060020a031916600160a060020a0392831617908190551615156100b857600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b15801561012157600080fd5b505af1158015610135573d6000803e3d6000fd5b505050506040513d602081101561014b57600080fd5b505192915050565b610808806101626000396000f3006080604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663071a8b53811461009d57806307a385e6146100f45780633363375c14610132578063516c4b841461016057806373d4a13a14610175578063c8b091091461018a578063dd336b94146101b8578063f89f4e77146101e6578063fb55a05514610214575b600080fd5b3480156100a957600080fd5b50604080516020600480358082013583810280860185019096528085526100f2953695939460249493850192918291850190849080828437509497506102429650505050505050565b005b34801561010057600080fd5b5061010961027a565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561013e57600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff60043516610296565b34801561016c57600080fd5b506101096103c5565b34801561018157600080fd5b506101096103e1565b34801561019657600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff600435166103fd565b3480156101c457600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff60043516610468565b3480156101f257600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff600435166104e5565b34801561022057600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff6004351661055c565b60005b81518110156102765761026e828281518110151561025f57fe5b906020019060200201516104e5565b600101610245565b5050565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561033257600080fd5b505af1158015610346573d6000803e3d6000fd5b505050506040513d602081101561035c57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461037e57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60025473ffffffffffffffffffffffffffffffffffffffff16331461042157600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60025473ffffffffffffffffffffffffffffffffffffffff16331461048c57600080fd5b6104a181600261049b8461068b565b1761073e565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de90600090a250565b60025473ffffffffffffffffffffffffffffffffffffffff16331461050957600080fd5b61051881600161049b8461068b565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7abdf8533487db28f8c616affbb4e122d90c5ab8deb258fd21b09cee59573090600090a250565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156105f857600080fd5b505af115801561060c573d6000803e3d6000fd5b505050506040513d602081101561062257600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461064457600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60018054604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff84811660248501529051600093919092169163295f36d79160448082019260209290919082900301818787803b15801561070c57600080fd5b505af1158015610720573d6000803e3d6000fd5b505050506040513d602081101561073657600080fd5b505192915050565b60018054604080517f461b09c0000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff858116602485015260448401859052905191169163461b09c091606480830192600092919082900301818387803b1580156107c057600080fd5b505af11580156107d4573d6000803e3d6000fd5b5050505050505600a165627a7a7230582085224779143f6669d52695d45d40cbcebb9f30cc878327219164147116ffcee00029".getBytes();

			// deploy deploys a new Ethereum contract, binding an instance of Approving to it.
			public static Approving deploy(TransactOpts auth, EthereumClient client, Address _cryptoFiat, Address _accountApprover) throws Exception {
				Interfaces args = Geth.newInterfaces(2);
				
				  args.set(0, Geth.newInterface()); args.get(0).setAddress(_cryptoFiat);
				
				  args.set(1, Geth.newInterface()); args.get(1).setAddress(_accountApprover);
				
				return new Approving(Geth.deployContract(auth, ABI, BYTECODE, client, args));
			}

			// Internal constructor used by contract deployment.
			private Approving(BoundContract deployment) {
				this.Address  = deployment.getAddress();
				this.Deployer = deployment.getDeployer();
				this.Contract = deployment;
			}
		

		// Ethereum address where this contract is located at.
		public final Address Address;

		// Ethereum transaction in which this contract was deployed (if known!).
		public final Transaction Deployer;

		// Contract instance bound to a blockchain address.
		private final BoundContract Contract;

		// Creates a new instance of Approving, bound to a specific deployed contract.
		public Approving(Address address, EthereumClient client) throws Exception {
			this(Geth.bindContract(address, ABI, client));
		}

		
			

			// accountApprover is a free data retrieval call binding the contract method 0x07a385e6.
			//
			// Solidity: function accountApprover() constant returns(address)
			public Address accountApprover(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "accountApprover", args);
				return results.get(0).getAddress();
				
			}
		
			

			// cryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
			//
			// Solidity: function cryptoFiat() constant returns(address)
			public Address cryptoFiat(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "cryptoFiat", args);
				return results.get(0).getAddress();
				
			}
		
			

			// data is a free data retrieval call binding the contract method 0x73d4a13a.
			//
			// Solidity: function data() constant returns(address)
			public Address data(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "data", args);
				return results.get(0).getAddress();
				
			}
		

		
			// appointAccountApprover is a paid mutator transaction binding the contract method 0xc8b09109.
			//
			// Solidity: function appointAccountApprover(next address) returns()
			public Transaction appointAccountApprover(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "appointAccountApprover"	, args);
			}
		
			// approveAccount is a paid mutator transaction binding the contract method 0xf89f4e77.
			//
			// Solidity: function approveAccount(account address) returns()
			public Transaction approveAccount(TransactOpts opts, Address account) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(account);
				

				return this.Contract.transact(opts, "approveAccount"	, args);
			}
		
			// approveAccounts is a paid mutator transaction binding the contract method 0x071a8b53.
			//
			// Solidity: function approveAccounts(accounts address[]) returns()
			public Transaction approveAccounts(TransactOpts opts, Addresses accounts) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddresses(accounts);
				

				return this.Contract.transact(opts, "approveAccounts"	, args);
			}
		
			// closeAccount is a paid mutator transaction binding the contract method 0xdd336b94.
			//
			// Solidity: function closeAccount(account address) returns()
			public Transaction closeAccount(TransactOpts opts, Address account) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(account);
				

				return this.Contract.transact(opts, "closeAccount"	, args);
			}
		
			// switchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
			//
			// Solidity: function switchCryptoFiat(next address) returns()
			public Transaction switchCryptoFiat(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchCryptoFiat"	, args);
			}
		
			// switchData is a paid mutator transaction binding the contract method 0x3363375c.
			//
			// Solidity: function switchData(next address) returns()
			public Transaction switchData(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchData"	, args);
			}
		
	}

	public class CryptoFiat {
		// ABI is the input ABI used to generate the binding from.
		public final static String ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointMasterAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"masterAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"contractActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"next\",\"type\":\"address\"}],\"name\":\"ContractUpgraded\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x608060405234801561001057600080fd5b5060008054600160a060020a0319908116331782556003805460018101825592527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9091018054909116301790556105468061006d6000396000f30060806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663028f4e47811461009257806313c01368146100c55780633fad74ad14610106578063474da79a1461012d5780635db4380d14610145578063874c3473146101735780639afd453c146101a1578063e814861e146101b6575b600080fd5b34801561009e57600080fd5b506100c360043573ffffffffffffffffffffffffffffffffffffffff602435166101f8565b005b3480156100d157600080fd5b506100dd6004356103f5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561011257600080fd5b5061011b61041d565b60408051918252519081900360200190f35b34801561013957600080fd5b506100dd600435610423565b34801561015157600080fd5b506100c373ffffffffffffffffffffffffffffffffffffffff60043516610458565b34801561017f57600080fd5b5061011b73ffffffffffffffffffffffffffffffffffffffff600435166104c3565b3480156101ad57600080fd5b506100dd6104d5565b3480156101c257600080fd5b506101e473ffffffffffffffffffffffffffffffffffffffff600435166104f1565b604080519115158252519081900360200190f35b60008083151561020757600080fd5b60008481526001602052604090205473ffffffffffffffffffffffffffffffffffffffff9081169250831682141561023e57600080fd5b60005473ffffffffffffffffffffffffffffffffffffffff1633148061027957503373ffffffffffffffffffffffffffffffffffffffff8316145b905080151561028757600080fd5b610290836104f1565b1561029a57600080fd5b73ffffffffffffffffffffffffffffffffffffffff82811660009081526002602090815260408083208390558783526001909152902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001691851691821790551561032a5773ffffffffffffffffffffffffffffffffffffffff831660009081526002602052604090208490555b6040805173ffffffffffffffffffffffffffffffffffffffff808516825285166020820152815186927fdc69b57038334451ee12fd1742228917cea7f40dbd33cda5162e7e5754acee1c928290030190a25050600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60035490565b600380548290811061043157fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60005473ffffffffffffffffffffffffffffffffffffffff16331461047c57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60026020526000908152604090205481565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604081205411905600a165627a7a7230582038bd3cc8ec8ee41f7b2e7468c55293150e27e0c7ab678e5e4162233dbef3f4360029".getBytes();

			// deploy deploys a new Ethereum contract, binding an instance of CryptoFiat to it.
			public static CryptoFiat deploy(TransactOpts auth, EthereumClient client) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				
				return new CryptoFiat(Geth.deployContract(auth, ABI, BYTECODE, client, args));
			}

			// Internal constructor used by contract deployment.
			private CryptoFiat(BoundContract deployment) {
				this.Address  = deployment.getAddress();
				this.Deployer = deployment.getDeployer();
				this.Contract = deployment;
			}
		

		// Ethereum address where this contract is located at.
		public final Address Address;

		// Ethereum transaction in which this contract was deployed (if known!).
		public final Transaction Deployer;

		// Contract instance bound to a blockchain address.
		private final BoundContract Contract;

		// Creates a new instance of CryptoFiat, bound to a specific deployed contract.
		public CryptoFiat(Address address, EthereumClient client) throws Exception {
			this(Geth.bindContract(address, ABI, client));
		}

		
			

			// contractActive is a free data retrieval call binding the contract method 0xe814861e.
			//
			// Solidity: function contractActive(addr address) constant returns(bool)
			public bool contractActive(CallOpts opts, Address addr) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(addr);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultBool(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "contractActive", args);
				return results.get(0).getBool();
				
			}
		
			

			// contractAddress is a free data retrieval call binding the contract method 0x13c01368.
			//
			// Solidity: function contractAddress( uint256) constant returns(address)
			public Address contractAddress(CallOpts opts, BigInt arg0) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setBigInt(arg0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "contractAddress", args);
				return results.get(0).getAddress();
				
			}
		
			

			// contractId is a free data retrieval call binding the contract method 0x874c3473.
			//
			// Solidity: function contractId( address) constant returns(uint256)
			public BigInt contractId(CallOpts opts, Address arg0) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(arg0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultBigInt(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "contractId", args);
				return results.get(0).getBigInt();
				
			}
		
			

			// contracts is a free data retrieval call binding the contract method 0x474da79a.
			//
			// Solidity: function contracts( uint256) constant returns(address)
			public Address contracts(CallOpts opts, BigInt arg0) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setBigInt(arg0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "contracts", args);
				return results.get(0).getAddress();
				
			}
		
			

			// contractsLength is a free data retrieval call binding the contract method 0x3fad74ad.
			//
			// Solidity: function contractsLength() constant returns(uint256)
			public BigInt contractsLength(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultBigInt(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "contractsLength", args);
				return results.get(0).getBigInt();
				
			}
		
			

			// masterAccount is a free data retrieval call binding the contract method 0x9afd453c.
			//
			// Solidity: function masterAccount() constant returns(address)
			public Address masterAccount(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "masterAccount", args);
				return results.get(0).getAddress();
				
			}
		

		
			// appointMasterAccount is a paid mutator transaction binding the contract method 0x5db4380d.
			//
			// Solidity: function appointMasterAccount(next address) returns()
			public Transaction appointMasterAccount(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "appointMasterAccount"	, args);
			}
		
			// upgrade is a paid mutator transaction binding the contract method 0x028f4e47.
			//
			// Solidity: function upgrade(id uint256, next address) returns()
			public Transaction upgrade(TransactOpts opts, BigInt id, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(2);
				args.set(0, Geth.newInterface()); args.get(0).setBigInt(id);
				args.set(1, Geth.newInterface()); args.get(1).setAddress(next);
				

				return this.Contract.transact(opts, "upgrade"	, args);
			}
		
	}

	public class Data {
		// ABI is the input ABI used to generate the binding from.
		public final static String ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"bucket\",\"type\":\"uint256\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bucket\",\"type\":\"uint256\"},{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x608060405234801561001057600080fd5b506040516020806103d8833981016040525160008054600160a060020a03909216600160a060020a0319909216919091179055610386806100526000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663295f36d7811461005b578063461b09c014610088578063516c4b84146100a8575b600080fd5b34801561006757600080fd5b506100766004356024356100e6565b60408051918252519081900360200190f35b34801561009457600080fd5b506100a66004356024356044356101bc565b005b3480156100b457600080fd5b506100bd61033e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60408051602080820185905281830184905282518083038401815260609092019283905281516000936001938593909282918401908083835b6020831061015c57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0909201916020918201910161011f565b51815160209384036101000a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01801990921691161790526040805192909401829003909120865285019590955292909201600020549695505050505050565b60008054604080517fe814861e000000000000000000000000000000000000000000000000000000008152336004820152905173ffffffffffffffffffffffffffffffffffffffff9092169263e814861e926024808401936020939083900390910190829087803b15801561023057600080fd5b505af1158015610244573d6000803e3d6000fd5b505050506040513d602081101561025a57600080fd5b5051151561026757600080fd5b60408051602080820186905281830185905282518083038401815260609092019283905281518493600193600093909282918401908083835b602083106102dd57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe090920191602091820191016102a0565b51815160209384036101000a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0180199092169116179052604080519290940182900390912086528501959095529290920160002093909355505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820831ad6bd7960b490d17b231075877e71ec8d74a792e3731e6ea67035ff8586900029".getBytes();

			// deploy deploys a new Ethereum contract, binding an instance of Data to it.
			public static Data deploy(TransactOpts auth, EthereumClient client, Address _cryptoFiat) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				
				  args.set(0, Geth.newInterface()); args.get(0).setAddress(_cryptoFiat);
				
				return new Data(Geth.deployContract(auth, ABI, BYTECODE, client, args));
			}

			// Internal constructor used by contract deployment.
			private Data(BoundContract deployment) {
				this.Address  = deployment.getAddress();
				this.Deployer = deployment.getDeployer();
				this.Contract = deployment;
			}
		

		// Ethereum address where this contract is located at.
		public final Address Address;

		// Ethereum transaction in which this contract was deployed (if known!).
		public final Transaction Deployer;

		// Contract instance bound to a blockchain address.
		private final BoundContract Contract;

		// Creates a new instance of Data, bound to a specific deployed contract.
		public Data(Address address, EthereumClient client) throws Exception {
			this(Geth.bindContract(address, ABI, client));
		}

		
			

			// cryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
			//
			// Solidity: function cryptoFiat() constant returns(address)
			public Address cryptoFiat(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "cryptoFiat", args);
				return results.get(0).getAddress();
				
			}
		
			

			// get is a free data retrieval call binding the contract method 0x295f36d7.
			//
			// Solidity: function get(bucket uint256, key bytes32) constant returns(bytes32)
			public byte[] get(CallOpts opts, BigInt bucket, byte[] key) throws Exception {
				Interfaces args = Geth.newInterfaces(2);
				args.set(0, Geth.newInterface()); args.get(0).setBigInt(bucket);
				args.set(1, Geth.newInterface()); args.get(1).setBinary(key);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultBinary(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "get", args);
				return results.get(0).getBinary();
				
			}
		

		
			// set is a paid mutator transaction binding the contract method 0x461b09c0.
			//
			// Solidity: function set(bucket uint256, key bytes32, value bytes32) returns()
			public Transaction set(TransactOpts opts, BigInt bucket, byte[] key, byte[] value) throws Exception {
				Interfaces args = Geth.newInterfaces(3);
				args.set(0, Geth.newInterface()); args.get(0).setBigInt(bucket);
				args.set(1, Geth.newInterface()); args.get(1).setBinary(key);
				args.set(2, Geth.newInterface()); args.get(2).setBinary(value);
				

				return this.Contract.transact(opts, "set"	, args);
			}
		
	}

	public class Delegation {
		// ABI is the input ABI used to generate the binding from.
		public final static String ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\"},{\"name\":\"transfers\",\"type\":\"bytes\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"multitransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"nonceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x608060405234801561001057600080fd5b50604051602080611213833981016040525160008054600160a060020a031916600160a060020a03831617905561004e640100000000610054810204565b5061012c565b6100676001640100000000610093810204565b60018054600160a060020a031916600160a060020a03928316179081905516151561009157600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b1580156100fa57600080fd5b505af115801561010e573d6000803e3d6000fd5b505050506040513d602081101561012457600080fd5b505192915050565b6110d88061013b6000396000f3006080604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305bafaa481146100875780633363375c146100ff578063516c4b841461012d57806373d4a13a1461016b578063e218e6d214610180578063ed2a2d6414610219578063fb55a05514610259575b600080fd5b34801561009357600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100fd9583359536956044949193909101919081908401838280828437509497505050923573ffffffffffffffffffffffffffffffffffffffff16935061028792505050565b005b34801561010b57600080fd5b506100fd73ffffffffffffffffffffffffffffffffffffffff6004351661045d565b34801561013957600080fd5b5061014261058c565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561017757600080fd5b506101426105a8565b34801561018c57600080fd5b50604080516020600460843581810135601f81018490048402850184019095528484526100fd948235946024803573ffffffffffffffffffffffffffffffffffffffff169560443595606435953695919460a49490939101919081908401838280828437509497505050923573ffffffffffffffffffffffffffffffffffffffff1693506105c492505050565b34801561022557600080fd5b5061024773ffffffffffffffffffffffffffffffffffffffff60043516610833565b60408051918252519081900360200190f35b34801561026557600080fd5b506100fd73ffffffffffffffffffffffffffffffffffffffff60043516610844565b6000610291611050565b8261029b81610973565b156102a557600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156102c757600080fd5b600092505b85831015610455576102de8584610989565b91506102ed8260200151610bc8565b6102fa8260400151610c1c565b8151602083015161030a90610c56565b1061031457600080fd5b61032682602001518360000151610d07565b61033e82602001518360800151846060015101610d9d565b61035082604001518360600151610dc3565b816040015173ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84606001516040518082815260200191505060405180910390a360008260800151111561044a576103dc848360800151610dc3565b8373ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84608001516040518082815260200191505060405180910390a35b6001909201916102cc565b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156104f957600080fd5b505af115801561050d573d6000803e3d6000fd5b505050506040513d602081101561052357600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461054557600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6000856105d081610973565b156105da57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156105fc57600080fd5b8261060681610973565b1561061057600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561063257600080fd5b6040805160208082018c90526c0100000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8c160282840152605482018a905260748083018a905283518084039091018152609490920192839052815161070b93918291908401908083835b602083106106d857805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0909201916020918201910161069b565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902086610deb565b925061071683610bc8565b8861072084610c56565b1061072a57600080fd5b610734838a610d07565b61074083878901610d9d565b61074a8888610dc3565b8773ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef896040518082815260200191505060405180910390a36000861115610828576107c28487610dc3565b8373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef886040518082815260200191505060405180910390a35b505050505050505050565b600061083e82610c56565b92915050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156108e057600080fd5b505af11580156108f4573d6000803e3d6000fd5b505050506040513d602081101561090a57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461092c57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600060028061098184610ece565b161492915050565b610991611050565b60008060008060008060008060006109a7611050565b8b60b5029950898d019c5060208d0151985060348d0151975060548d0151965060748d0151955060948d0151945060b48d0151935060ff60b58e0151169250601b8360ff1610156109f957601b830192505b6040805160208082018c90526c0100000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8c160282840152605482018a905260748083018a905283518084039091018152609490920192839052815191929182918401908083835b60208310610a9c57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101610a5f565b5181517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209485036101000a0190811690199190911617905260408051949092018490039093208e87528151600080825281860180855283905260ff8b1682850152606082018d9052608082018c905292519198506001965060a080820196507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083019450908290030191865af1158015610b5c573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015173ffffffffffffffffffffffffffffffffffffffff90811660208501528a1690830152506060810187905260808101869052995050505050505050505092915050565b806000610bd482610ece565b9050600180821614610be557600080fd5b60048181161415610bf557600080fd5b73ffffffffffffffffffffffffffffffffffffffff82161515610c1757600080fd5b505050565b80610c2681610973565b15610c3057600080fd5b73ffffffffffffffffffffffffffffffffffffffff81161515610c5257600080fd5b5050565b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526003600482015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b158015610cd557600080fd5b505af1158015610ce9573d6000803e3d6000fd5b505050506040513d6020811015610cff57600080fd5b505192915050565b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526003600482015273ffffffffffffffffffffffffffffffffffffffff8581166024830152604482018590529151919092169163461b09c091606480830192600092919082900301818387803b158015610d8957600080fd5b505af1158015610455573d6000803e3d6000fd5b6000610da883610f4f565b905081811015610db757600080fd5b610c1783838303610fce565b6000610dce83610f4f565b9050818101811115610ddf57600080fd5b610c1783838301610fce565b60008060008084516041141515610e0157600080fd5b50505060208201516040830151604184015160ff16601b811015610e2357601b015b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a08085019491937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0840193928390039091019190865af1158015610e9b573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00151979650505050505050565b60018054604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff84811660248501529051600093919092169163295f36d79160448082019260209290919082900301818787803b158015610cd557600080fd5b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b158015610cd557600080fd5b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff8581166024830152604482018590529151919092169163461b09c091606480830192600092919082900301818387803b158015610d8957600080fd5b60a06040519081016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815250905600a165627a7a723058208df1c8c23e4a032647a53cb610af0d9c7ba3e2433e6499858f341d399d6efca90029".getBytes();

			// deploy deploys a new Ethereum contract, binding an instance of Delegation to it.
			public static Delegation deploy(TransactOpts auth, EthereumClient client, Address _cryptoFiat) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				
				  args.set(0, Geth.newInterface()); args.get(0).setAddress(_cryptoFiat);
				
				return new Delegation(Geth.deployContract(auth, ABI, BYTECODE, client, args));
			}

			// Internal constructor used by contract deployment.
			private Delegation(BoundContract deployment) {
				this.Address  = deployment.getAddress();
				this.Deployer = deployment.getDeployer();
				this.Contract = deployment;
			}
		

		// Ethereum address where this contract is located at.
		public final Address Address;

		// Ethereum transaction in which this contract was deployed (if known!).
		public final Transaction Deployer;

		// Contract instance bound to a blockchain address.
		private final BoundContract Contract;

		// Creates a new instance of Delegation, bound to a specific deployed contract.
		public Delegation(Address address, EthereumClient client) throws Exception {
			this(Geth.bindContract(address, ABI, client));
		}

		
			

			// cryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
			//
			// Solidity: function cryptoFiat() constant returns(address)
			public Address cryptoFiat(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "cryptoFiat", args);
				return results.get(0).getAddress();
				
			}
		
			

			// data is a free data retrieval call binding the contract method 0x73d4a13a.
			//
			// Solidity: function data() constant returns(address)
			public Address data(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "data", args);
				return results.get(0).getAddress();
				
			}
		
			

			// nonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
			//
			// Solidity: function nonceOf(account address) constant returns(uint256)
			public BigInt nonceOf(CallOpts opts, Address account) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(account);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultBigInt(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "nonceOf", args);
				return results.get(0).getBigInt();
				
			}
		

		
			// multitransfer is a paid mutator transaction binding the contract method 0x05bafaa4.
			//
			// Solidity: function multitransfer(count uint256, transfers bytes, delegate address) returns()
			public Transaction multitransfer(TransactOpts opts, BigInt count, byte[] transfers, Address delegate) throws Exception {
				Interfaces args = Geth.newInterfaces(3);
				args.set(0, Geth.newInterface()); args.get(0).setBigInt(count);
				args.set(1, Geth.newInterface()); args.get(1).setBinary(transfers);
				args.set(2, Geth.newInterface()); args.get(2).setAddress(delegate);
				

				return this.Contract.transact(opts, "multitransfer"	, args);
			}
		
			// switchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
			//
			// Solidity: function switchCryptoFiat(next address) returns()
			public Transaction switchCryptoFiat(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchCryptoFiat"	, args);
			}
		
			// switchData is a paid mutator transaction binding the contract method 0x3363375c.
			//
			// Solidity: function switchData(next address) returns()
			public Transaction switchData(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchData"	, args);
			}
		
			// transfer is a paid mutator transaction binding the contract method 0xe218e6d2.
			//
			// Solidity: function transfer(nonce uint256, destination address, amount uint256, fee uint256, signature bytes, delegate address) returns()
			public Transaction transfer(TransactOpts opts, BigInt nonce, Address destination, BigInt amount, BigInt fee, byte[] signature, Address delegate) throws Exception {
				Interfaces args = Geth.newInterfaces(6);
				args.set(0, Geth.newInterface()); args.get(0).setBigInt(nonce);
				args.set(1, Geth.newInterface()); args.get(1).setAddress(destination);
				args.set(2, Geth.newInterface()); args.get(2).setBigInt(amount);
				args.set(3, Geth.newInterface()); args.get(3).setBigInt(fee);
				args.set(4, Geth.newInterface()); args.get(4).setBinary(signature);
				args.set(5, Geth.newInterface()); args.get(5).setAddress(delegate);
				

				return this.Contract.transact(opts, "transfer"	, args);
			}
		
	}

	public class Enforcement {
		// ABI is the input ABI used to generate the binding from.
		public final static String ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"account\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lawEnforcer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"unfreezeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountDesignator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointAccountDesignator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointLawEnforcer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"designateAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"freezeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_lawEnforcer\",\"type\":\"address\"},{\"name\":\"_enforcementAccountDesignator\",\"type\":\"address\"},{\"name\":\"_enforcementAccount\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x608060405234801561001057600080fd5b50604051608080610dce833981016040908152815160208301519183015160609093015160008054600160a060020a031916600160a060020a0384161790559092906100636401000000006100a6810204565b60028054600160a060020a03948516600160a060020a0319918216179091556003805493851693821693909317909255600480549190931691161790555061017e565b6100b960016401000000006100e5810204565b60018054600160a060020a031916600160a060020a0392831617908190551615156100e357600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b15801561014c57600080fd5b505af1158015610160573d6000803e3d6000fd5b505050506040513d602081101561017657600080fd5b505192915050565b610c418061018d6000396000f3006080604052600436106100c45763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633363375c81146100c9578063516c4b84146100f95780635dab24201461013757806372cfc9dc1461014c57806373d4a13a14610161578063788649ea1461017657806385a0f282146101a457806390f28e74146101b9578063b10725e8146101e7578063b9b0330f14610215578063f26c159f14610243578063f3fef3a314610271578063fb55a055146102a2575b600080fd5b3480156100d557600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff600435166102d0565b005b34801561010557600080fd5b5061010e6103ff565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561014357600080fd5b5061010e61041b565b34801561015857600080fd5b5061010e610437565b34801561016d57600080fd5b5061010e610453565b34801561018257600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff6004351661046f565b3480156101b057600080fd5b5061010e610517565b3480156101c557600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff60043516610533565b3480156101f357600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff6004351661059e565b34801561022157600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff60043516610609565b34801561024f57600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff600435166106ab565b34801561027d57600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff60043516602435610734565b3480156102ae57600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff6004351661082b565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561036c57600080fd5b505af1158015610380573d6000803e3d6000fd5b505050506040513d602081101561039657600080fd5b505173ffffffffffffffffffffffffffffffffffffffff16146103b857600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60025473ffffffffffffffffffffffffffffffffffffffff16331461049357600080fd5b6104c7817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb6104c18261095a565b16610a0d565b6040805160008152905173ffffffffffffffffffffffffffffffffffffffff8316917fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee919081900360200190a250565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff16331461055757600080fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60025473ffffffffffffffffffffffffffffffffffffffff1633146105c257600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60035473ffffffffffffffffffffffffffffffffffffffff16331461062d57600080fd5b8061063781610aab565b1561064157600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561066357600080fd5b50600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60025473ffffffffffffffffffffffffffffffffffffffff1633146106cf57600080fd5b6106e48160046106de8461095a565b17610a0d565b6040805160018152905173ffffffffffffffffffffffffffffffffffffffff8316917fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee919081900360200190a250565b60025473ffffffffffffffffffffffffffffffffffffffff16331461075857600080fd5b60045473ffffffffffffffffffffffffffffffffffffffff1661077a81610aab565b1561078457600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156107a657600080fd5b6107b08383610ac1565b6004546107d39073ffffffffffffffffffffffffffffffffffffffff1683610aec565b60045460408051848152905173ffffffffffffffffffffffffffffffffffffffff928316928616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a3505050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156108c757600080fd5b505af11580156108db573d6000803e3d6000fd5b505050506040513d60208110156108f157600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461091357600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60018054604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff84811660248501529051600093919092169163295f36d79160448082019260209290919082900301818787803b1580156109db57600080fd5b505af11580156109ef573d6000803e3d6000fd5b505050506040513d6020811015610a0557600080fd5b505192915050565b60018054604080517f461b09c0000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff858116602485015260448401859052905191169163461b09c091606480830192600092919082900301818387803b158015610a8f57600080fd5b505af1158015610aa3573d6000803e3d6000fd5b505050505050565b6000600280610ab98461095a565b161492915050565b6000610acc83610b14565b905081811015610adb57600080fd5b610ae783838303610b93565b505050565b6000610af783610b14565b9050818101811115610b0857600080fd5b610ae783838301610b93565b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b1580156109db57600080fd5b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff8581166024830152604482018590529151919092169163461b09c091606480830192600092919082900301818387803b158015610a8f57600080fd00a165627a7a72305820a17e0428ce3a8ae44793fc733a66c6af15d663b98a45b15cb96119b2dcac6fba0029".getBytes();

			// deploy deploys a new Ethereum contract, binding an instance of Enforcement to it.
			public static Enforcement deploy(TransactOpts auth, EthereumClient client, Address _cryptoFiat, Address _lawEnforcer, Address _enforcementAccountDesignator, Address _enforcementAccount) throws Exception {
				Interfaces args = Geth.newInterfaces(4);
				
				  args.set(0, Geth.newInterface()); args.get(0).setAddress(_cryptoFiat);
				
				  args.set(1, Geth.newInterface()); args.get(1).setAddress(_lawEnforcer);
				
				  args.set(2, Geth.newInterface()); args.get(2).setAddress(_enforcementAccountDesignator);
				
				  args.set(3, Geth.newInterface()); args.get(3).setAddress(_enforcementAccount);
				
				return new Enforcement(Geth.deployContract(auth, ABI, BYTECODE, client, args));
			}

			// Internal constructor used by contract deployment.
			private Enforcement(BoundContract deployment) {
				this.Address  = deployment.getAddress();
				this.Deployer = deployment.getDeployer();
				this.Contract = deployment;
			}
		

		// Ethereum address where this contract is located at.
		public final Address Address;

		// Ethereum transaction in which this contract was deployed (if known!).
		public final Transaction Deployer;

		// Contract instance bound to a blockchain address.
		private final BoundContract Contract;

		// Creates a new instance of Enforcement, bound to a specific deployed contract.
		public Enforcement(Address address, EthereumClient client) throws Exception {
			this(Geth.bindContract(address, ABI, client));
		}

		
			

			// account is a free data retrieval call binding the contract method 0x5dab2420.
			//
			// Solidity: function account() constant returns(address)
			public Address account(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "account", args);
				return results.get(0).getAddress();
				
			}
		
			

			// accountDesignator is a free data retrieval call binding the contract method 0x85a0f282.
			//
			// Solidity: function accountDesignator() constant returns(address)
			public Address accountDesignator(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "accountDesignator", args);
				return results.get(0).getAddress();
				
			}
		
			

			// cryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
			//
			// Solidity: function cryptoFiat() constant returns(address)
			public Address cryptoFiat(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "cryptoFiat", args);
				return results.get(0).getAddress();
				
			}
		
			

			// data is a free data retrieval call binding the contract method 0x73d4a13a.
			//
			// Solidity: function data() constant returns(address)
			public Address data(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "data", args);
				return results.get(0).getAddress();
				
			}
		
			

			// lawEnforcer is a free data retrieval call binding the contract method 0x72cfc9dc.
			//
			// Solidity: function lawEnforcer() constant returns(address)
			public Address lawEnforcer(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "lawEnforcer", args);
				return results.get(0).getAddress();
				
			}
		

		
			// appointAccountDesignator is a paid mutator transaction binding the contract method 0x90f28e74.
			//
			// Solidity: function appointAccountDesignator(next address) returns()
			public Transaction appointAccountDesignator(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "appointAccountDesignator"	, args);
			}
		
			// appointLawEnforcer is a paid mutator transaction binding the contract method 0xb10725e8.
			//
			// Solidity: function appointLawEnforcer(next address) returns()
			public Transaction appointLawEnforcer(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "appointLawEnforcer"	, args);
			}
		
			// designateAccount is a paid mutator transaction binding the contract method 0xb9b0330f.
			//
			// Solidity: function designateAccount(_account address) returns()
			public Transaction designateAccount(TransactOpts opts, Address _account) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(_account);
				

				return this.Contract.transact(opts, "designateAccount"	, args);
			}
		
			// freezeAccount is a paid mutator transaction binding the contract method 0xf26c159f.
			//
			// Solidity: function freezeAccount(target address) returns()
			public Transaction freezeAccount(TransactOpts opts, Address target) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(target);
				

				return this.Contract.transact(opts, "freezeAccount"	, args);
			}
		
			// switchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
			//
			// Solidity: function switchCryptoFiat(next address) returns()
			public Transaction switchCryptoFiat(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchCryptoFiat"	, args);
			}
		
			// switchData is a paid mutator transaction binding the contract method 0x3363375c.
			//
			// Solidity: function switchData(next address) returns()
			public Transaction switchData(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchData"	, args);
			}
		
			// unfreezeAccount is a paid mutator transaction binding the contract method 0x788649ea.
			//
			// Solidity: function unfreezeAccount(target address) returns()
			public Transaction unfreezeAccount(TransactOpts opts, Address target) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(target);
				

				return this.Contract.transact(opts, "unfreezeAccount"	, args);
			}
		
			// withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
			//
			// Solidity: function withdraw(from address, amount uint256) returns()
			public Transaction withdraw(TransactOpts opts, Address from, BigInt amount) throws Exception {
				Interfaces args = Geth.newInterfaces(2);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(from);
				args.set(1, Geth.newInterface()); args.get(1).setBigInt(amount);
				

				return this.Contract.transact(opts, "withdraw"	, args);
			}
		
	}

	public class Reserve {
		// ABI is the input ABI used to generate the binding from.
		public final static String ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"reserveBank\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointReserveBank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_reserveBank\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x608060405234801561001057600080fd5b50604051604080610ccc83398101604052805160209091015160008054600160a060020a031916600160a060020a03841617905561005564010000000061007b810204565b60028054600160a060020a031916600160a060020a039290921691909117905550610153565b61008e60016401000000006100ba810204565b60018054600160a060020a031916600160a060020a0392831617908190551615156100b857600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b15801561012157600080fd5b505af1158015610135573d6000803e3d6000fd5b505050506040513d602081101561014b57600080fd5b505192915050565b610b6a806101626000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302946804811461009d57806318160ddd146100db5780633363375c14610102578063516c4b841461013257806373d4a13a1461014757806398e52f9a1461015c578063b921e16314610174578063ddf05f591461018c578063fb55a055146101ba575b600080fd5b3480156100a957600080fd5b506100b26101e8565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156100e757600080fd5b506100f0610204565b60408051918252519081900360200190f35b34801561010e57600080fd5b5061013073ffffffffffffffffffffffffffffffffffffffff60043516610213565b005b34801561013e57600080fd5b506100b2610342565b34801561015357600080fd5b506100b261035e565b34801561016857600080fd5b5061013060043561037a565b34801561018057600080fd5b506101306004356104da565b34801561019857600080fd5b5061013073ffffffffffffffffffffffffffffffffffffffff60043516610623565b3480156101c657600080fd5b5061013073ffffffffffffffffffffffffffffffffffffffff6004351661068e565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b600061020e6107bd565b905090565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156102af57600080fd5b505af11580156102c3573d6000803e3d6000fd5b505050506040513d60208110156102d957600080fd5b505173ffffffffffffffffffffffffffffffffffffffff16146102fb57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60025460009073ffffffffffffffffffffffffffffffffffffffff1633146103a157600080fd5b60025473ffffffffffffffffffffffffffffffffffffffff1660006103c58261086a565b90506001808216146103d657600080fd5b600481811614156103e657600080fd5b73ffffffffffffffffffffffffffffffffffffffff8216151561040857600080fd5b6104106107bd565b92508383101561041f57600080fd5b838303925061042d8361091d565b6002546104509073ffffffffffffffffffffffffffffffffffffffff16856109b8565b60025460408051868152905160009273ffffffffffffffffffffffffffffffffffffffff16917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a36040805184815290517ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e99181900360200190a150505050565b60025460009073ffffffffffffffffffffffffffffffffffffffff16331461050157600080fd5b60025473ffffffffffffffffffffffffffffffffffffffff16610523816109e3565b1561052d57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561054f57600080fd5b610557610204565b915082820182111561056857600080fd5b908201906105758261091d565b6002546105989073ffffffffffffffffffffffffffffffffffffffff16846109f9565b60025460408051858152905173ffffffffffffffffffffffffffffffffffffffff909216916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a36040805183815290517ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e99181900360200190a1505050565b60025473ffffffffffffffffffffffffffffffffffffffff16331461064757600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561072a57600080fd5b505af115801561073e573d6000803e3d6000fd5b505050506040513d602081101561075457600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461077657600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600154604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600560048201526000602482018190529151919273ffffffffffffffffffffffffffffffffffffffff169163295f36d79160448082019260209290919082900301818787803b15801561083957600080fd5b505af115801561084d573d6000803e3d6000fd5b505050506040513d602081101561086357600080fd5b5051905090565b60018054604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff84811660248501529051600093919092169163295f36d79160448082019260209290919082900301818787803b1580156108eb57600080fd5b505af11580156108ff573d6000803e3d6000fd5b505050506040513d602081101561091557600080fd5b505192915050565b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526005600482015260006024820181905260448201859052915173ffffffffffffffffffffffffffffffffffffffff9093169263461b09c09260648084019391929182900301818387803b15801561099d57600080fd5b505af11580156109b1573d6000803e3d6000fd5b5050505050565b60006109c383610a21565b9050818110156109d257600080fd5b6109de83838303610aa0565b505050565b60006002806109f18461086a565b161492915050565b6000610a0483610a21565b9050818101811115610a1557600080fd5b6109de83838301610aa0565b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b1580156108eb57600080fd5b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff8581166024830152604482018590529151919092169163461b09c091606480830192600092919082900301818387803b158015610b2257600080fd5b505af1158015610b36573d6000803e3d6000fd5b5050505050505600a165627a7a72305820c31c813b4344ce6b585169a838e95f32e023d0b734a4d6f9c10daaba58d10e6b0029".getBytes();

			// deploy deploys a new Ethereum contract, binding an instance of Reserve to it.
			public static Reserve deploy(TransactOpts auth, EthereumClient client, Address _cryptoFiat, Address _reserveBank) throws Exception {
				Interfaces args = Geth.newInterfaces(2);
				
				  args.set(0, Geth.newInterface()); args.get(0).setAddress(_cryptoFiat);
				
				  args.set(1, Geth.newInterface()); args.get(1).setAddress(_reserveBank);
				
				return new Reserve(Geth.deployContract(auth, ABI, BYTECODE, client, args));
			}

			// Internal constructor used by contract deployment.
			private Reserve(BoundContract deployment) {
				this.Address  = deployment.getAddress();
				this.Deployer = deployment.getDeployer();
				this.Contract = deployment;
			}
		

		// Ethereum address where this contract is located at.
		public final Address Address;

		// Ethereum transaction in which this contract was deployed (if known!).
		public final Transaction Deployer;

		// Contract instance bound to a blockchain address.
		private final BoundContract Contract;

		// Creates a new instance of Reserve, bound to a specific deployed contract.
		public Reserve(Address address, EthereumClient client) throws Exception {
			this(Geth.bindContract(address, ABI, client));
		}

		
			

			// cryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
			//
			// Solidity: function cryptoFiat() constant returns(address)
			public Address cryptoFiat(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "cryptoFiat", args);
				return results.get(0).getAddress();
				
			}
		
			

			// data is a free data retrieval call binding the contract method 0x73d4a13a.
			//
			// Solidity: function data() constant returns(address)
			public Address data(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "data", args);
				return results.get(0).getAddress();
				
			}
		
			

			// reserveBank is a free data retrieval call binding the contract method 0x02946804.
			//
			// Solidity: function reserveBank() constant returns(address)
			public Address reserveBank(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultAddress(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "reserveBank", args);
				return results.get(0).getAddress();
				
			}
		
			

			// totalSupply is a free data retrieval call binding the contract method 0x18160ddd.
			//
			// Solidity: function totalSupply() constant returns(uint256)
			public BigInt totalSupply(CallOpts opts) throws Exception {
				Interfaces args = Geth.newInterfaces(0);
				

				Interfaces results = Geth.newInterfaces(1);
				Interface result0 = Geth.newInterface(); result0.setDefaultBigInt(); results.set(0, result0);
				

				if (opts == null) {
					opts = Geth.newCallOpts();
				}
				this.Contract.call(opts, results, "totalSupply", args);
				return results.get(0).getBigInt();
				
			}
		

		
			// appointReserveBank is a paid mutator transaction binding the contract method 0xddf05f59.
			//
			// Solidity: function appointReserveBank(next address) returns()
			public Transaction appointReserveBank(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "appointReserveBank"	, args);
			}
		
			// decreaseSupply is a paid mutator transaction binding the contract method 0x98e52f9a.
			//
			// Solidity: function decreaseSupply(amount uint256) returns()
			public Transaction decreaseSupply(TransactOpts opts, BigInt amount) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setBigInt(amount);
				

				return this.Contract.transact(opts, "decreaseSupply"	, args);
			}
		
			// increaseSupply is a paid mutator transaction binding the contract method 0xb921e163.
			//
			// Solidity: function increaseSupply(amount uint256) returns()
			public Transaction increaseSupply(TransactOpts opts, BigInt amount) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setBigInt(amount);
				

				return this.Contract.transact(opts, "increaseSupply"	, args);
			}
		
			// switchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
			//
			// Solidity: function switchCryptoFiat(next address) returns()
			public Transaction switchCryptoFiat(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchCryptoFiat"	, args);
			}
		
			// switchData is a paid mutator transaction binding the contract method 0x3363375c.
			//
			// Solidity: function switchData(next address) returns()
			public Transaction switchData(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchData"	, args);
			}
		
	}

