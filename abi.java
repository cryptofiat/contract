
// This file is an automatically generated Java binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package eu.cryptoeuro.contract;

import org.ethereum.geth.*;
import org.ethereum.geth.internal.*;


	public class AccountRecovery {
		// ABI is the input ABI used to generate the binding from.
		public final static String ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"into\",\"type\":\"address\"}],\"name\":\"recoverBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recoveryAccount\",\"type\":\"address\"}],\"name\":\"designateRecoveryAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x6060604052341561000f57600080fd5b604051602080610a3b833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b6109e8806100536000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632d1c5ff8811461005e578063516c4b8414610092578063f1375f38146100ce578063fb55a055146100fc575b600080fd5b341561006957600080fd5b61009073ffffffffffffffffffffffffffffffffffffffff6004358116906024351661012a565b005b341561009d57600080fd5b6100a56102d0565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100d957600080fd5b61009073ffffffffffffffffffffffffffffffffffffffff600435166102ec565b005b341561010757600080fd5b61009073ffffffffffffffffffffffffffffffffffffffff600435166102fa565b005b600082610136816103f0565b151561014157600080fd5b61014a81610408565b1561015457600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561017657600080fd5b8261018081610420565b1561018a57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156101ac57600080fd5b6101b585610438565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156101ee57600080fd5b6102038560026101fd886104f8565b176105b8565b8473ffffffffffffffffffffffffffffffffffffffff167fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de60405160405180910390a261024f85610670565b925061025b8584610730565b610265848461075c565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405190815260200160405180910390a35b5b505b50505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6102f6338261078b565b5b50565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561036e57600080fd5b6102c65a03f1151561037f57600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff161415156103ab57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60006001806103fe846104f8565b161490505b919050565b60006004806103fe846104f8565b161490505b919050565b60006002806103fe846104f8565b161490505b919050565b6000610442610859565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760048473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156104d657600080fd5b6102c65a03f115156104e757600080fd5b50505060405180519150505b919050565b6000610502610859565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760018473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156104d657600080fd5b6102c65a03f115156104e757600080fd5b50505060405180519150505b919050565b6105c0610859565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060018473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561065757600080fd5b6102c65a03f115156102c757600080fd5b5050505b5050565b600061067a610859565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760028473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156104d657600080fd5b6102c65a03f115156104e757600080fd5b50505060405180519150505b919050565b600061073b83610670565b90508181101561074a57600080fd5b6107568383830361086b565b5b505050565b600061076783610670565b90508181018190101561077957600080fd5b6107568383830161086b565b5b505050565b610793610859565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060048473ffffffffffffffffffffffffffffffffffffffff166001028473ffffffffffffffffffffffffffffffffffffffff166001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561065757600080fd5b6102c65a03f115156102c757600080fd5b5050505b5050565b60006108656001610923565b90505b90565b610873610859565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060028473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561065757600080fd5b6102c65a03f115156102c757600080fd5b5050505b5050565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b15156104d657600080fd5b6102c65a03f115156104e757600080fd5b50505060405180519150505b9190505600a165627a7a7230582088333de5751e56115e8f777fcf472970f2858ef191fdf67782336925141ded010029".getBytes();

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
		
	}

	public class Accounts {
		// ABI is the input ABI used to generate the binding from.
		public final static String ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"statusOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFrozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x6060604052341561000f57600080fd5b6040516020806108e1833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b61088e806100536000396000f3006060604052361561008b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663516c4b848114610090578063673448dd146100cc5780636943b0171461010c57806370a082311461014c57806397a5d5b51461018a578063a9059cbb146101c8578063e5839836146101f9578063fb55a05514610239575b600080fd5b341561009b57600080fd5b6100a3610267565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100d757600080fd5b6100f873ffffffffffffffffffffffffffffffffffffffff60043516610283565b604051901515815260200160405180910390f35b341561011757600080fd5b6100f873ffffffffffffffffffffffffffffffffffffffff60043516610296565b604051901515815260200160405180910390f35b341561015757600080fd5b61017873ffffffffffffffffffffffffffffffffffffffff600435166102a9565b60405190815260200160405180910390f35b341561019557600080fd5b61017873ffffffffffffffffffffffffffffffffffffffff600435166102bc565b60405190815260200160405180910390f35b34156101d357600080fd5b6101f773ffffffffffffffffffffffffffffffffffffffff600435166024356102cf565b005b341561020457600080fd5b6100f873ffffffffffffffffffffffffffffffffffffffff600435166103d3565b604051901515815260200160405180910390f35b341561024457600080fd5b6101f773ffffffffffffffffffffffffffffffffffffffff600435166103e6565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600061028e826104dc565b90505b919050565b600061028e826104f4565b90505b919050565b600061028e8261050c565b90505b919050565b600061028e826105cc565b90505b919050565b6000336102db816104dc565b15156102e657600080fd5b6102ef8161068c565b156102f957600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561031b57600080fd5b83610325816104f4565b1561032f57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561035157600080fd5b33925061035e83856106a4565b61036885856106d0565b8473ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8660405190815260200160405180910390a35b5b505b50505050565b600061028e8261068c565b90505b919050565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561045a57600080fd5b6102c65a03f1151561046b57600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561049757600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60006001806104ea846105cc565b161490505b919050565b60006002806104ea846105cc565b161490505b919050565b60006105166106ff565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760028473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156105aa57600080fd5b6102c65a03f115156105bb57600080fd5b50505060405180519150505b919050565b60006105d66106ff565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760018473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156105aa57600080fd5b6102c65a03f115156105bb57600080fd5b50505060405180519150505b919050565b60006004806104ea846105cc565b161490505b919050565b60006106af8361050c565b9050818110156106be57600080fd5b6106ca83838303610711565b5b505050565b60006106db8361050c565b9050818101819010156106ed57600080fd5b6106ca83838301610711565b5b505050565b600061070b60016107c9565b90505b90565b6107196106ff565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060028473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156107b057600080fd5b6102c65a03f115156103ca57600080fd5b5050505b5050565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b15156105aa57600080fd5b6102c65a03f115156105bb57600080fd5b50505060405180519150505b9190505600a165627a7a723058202676d7c634dcf467185f6cdf407e31a02e8f1aff953815a772f713a416b118960029".getBytes();

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
		public final static String ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"approveAccounts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountApprover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointAccountApprover\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"closeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"approveAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_accountApprover\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x6060604052341561000f57600080fd5b6040516040806107a283398101604052808051919060200180519150505b60008054600160a060020a03808516600160a060020a03199283161790925560018054928416929091169190911790555b50505b610732806100706000396000f300606060405236156100805763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663071a8b53811461008557806307a385e6146100d6578063516c4b8414610112578063c8b091091461014e578063dd336b941461017c578063f89f4e77146101aa578063fb55a055146101d8575b600080fd5b341561009057600080fd5b6100d4600460248135818101908301358060208181020160405190810160405280939291908181526020018383602002808284375094965061020695505050505050565b005b34156100e157600080fd5b6100e961023e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561011d57600080fd5b6100e961025a565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561015957600080fd5b6100d473ffffffffffffffffffffffffffffffffffffffff60043516610276565b005b341561018757600080fd5b6100d473ffffffffffffffffffffffffffffffffffffffff600435166102e3565b005b34156101b557600080fd5b6100d473ffffffffffffffffffffffffffffffffffffffff60043516610368565b005b34156101e357600080fd5b6100d473ffffffffffffffffffffffffffffffffffffffff600435166103ed565b005b60005b81518110156102395761023082828151811061022157fe5b90602001906020020151610368565b5b600101610209565b5b5050565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6001543373ffffffffffffffffffffffffffffffffffffffff90811691161461029e57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6001543373ffffffffffffffffffffffffffffffffffffffff90811691161461030b57600080fd5b61032081600261031a846104e3565b176105a3565b8073ffffffffffffffffffffffffffffffffffffffff167fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de60405160405180910390a25b5b50565b6001543373ffffffffffffffffffffffffffffffffffffffff90811691161461039057600080fd5b6103a581600161031a846104e3565b176105a3565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7abdf8533487db28f8c616affbb4e122d90c5ab8deb258fd21b09cee59573060405160405180910390a25b5b50565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561046157600080fd5b6102c65a03f1151561047257600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561049e57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60006104ed61065b565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760018473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561058157600080fd5b6102c65a03f1151561059257600080fd5b50505060405180519150505b919050565b6105ab61065b565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060018473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561064257600080fd5b6102c65a03f1151561065357600080fd5b5050505b5050565b6000610667600161066d565b90505b90565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561058157600080fd5b6102c65a03f1151561059257600080fd5b50505060405180519150505b9190505600a165627a7a72305820cf9948277410e5e7ba8c93ba7a00aa82c896d65cd0a1e3d18e2c359e1dfbdad20029".getBytes();

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
		
	}

	public class CryptoFiat {
		// ABI is the input ABI used to generate the binding from.
		public final static String ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointMasterAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"masterAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"contractActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"next\",\"type\":\"address\"}],\"name\":\"ContractUpgraded\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a0316179055600380546001810161003d8382610070565b916000526020600020900160005b8154600160a060020a033081166101009390930a92830292021916179055505b6100bb565b8154818355818115116100945760008381526020902061009491810190830161009a565b5b505050565b6100b891905b808211156100b457600081556001016100a0565b5090565b90565b6105ec806100ca6000396000f3006060604052361561008b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663028f4e47811461009057806313c01368146100c15780633fad74ad14610100578063474da79a146101255780635db4380d14610164578063874c3473146101925780639afd453c146101d0578063e814861e1461020c575b600080fd5b341561009b57600080fd5b6100bf60043573ffffffffffffffffffffffffffffffffffffffff6024351661024c565b005b34156100cc57600080fd5b6100d760043561043f565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561010b57600080fd5b610113610467565b60405190815260200160405180910390f35b341561013057600080fd5b6100d760043561046e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561016f57600080fd5b6100bf73ffffffffffffffffffffffffffffffffffffffff600435166104ad565b005b341561019d57600080fd5b61011373ffffffffffffffffffffffffffffffffffffffff6004351661051a565b60405190815260200160405180910390f35b34156101db57600080fd5b6100d761052c565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561021757600080fd5b61023873ffffffffffffffffffffffffffffffffffffffff60043516610548565b604051901515815260200160405180910390f35b60008083151561025b57600080fd5b60008481526001602052604090205473ffffffffffffffffffffffffffffffffffffffff9081169250831682141561029257600080fd5b6000543373ffffffffffffffffffffffffffffffffffffffff908116911614806102e757508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b90508015156102f557600080fd5b6102fe83610548565b1561030857600080fd5b73ffffffffffffffffffffffffffffffffffffffff82811660009081526002602090815260408083208390558783526001909152902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169185169182179055156103985773ffffffffffffffffffffffffffffffffffffffff831660009081526002602052604090208490555b837fdc69b57038334451ee12fd1742228917cea7f40dbd33cda5162e7e5754acee1c838560405173ffffffffffffffffffffffffffffffffffffffff9283168152911660208201526040908101905180910390a260038054600181016103fe8382610575565b916000526020600020900160005b815473ffffffffffffffffffffffffffffffffffffffff8088166101009390930a92830292021916179055505b50505050565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003545b90565b600380548290811061047c57fe5b906000526020600020900160005b915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146104d557600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60026020526000908152604090205481565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260026020526040812054115b919050565b8154818355818115116105995760008381526020902061059991810190830161059f565b5b505050565b61046b91905b808211156105b957600081556001016105a5565b5090565b905600a165627a7a72305820c9c2946e5b59e4c1b7e2931686d0ca86c872e1fd7b43cbf8c3f84f124b34cb6f0029".getBytes();

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
		public final static String ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"bucket\",\"type\":\"uint256\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bucket\",\"type\":\"uint256\"},{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x6060604052341561000f57600080fd5b6040516020806103c7833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b610374806100536000396000f300606060405263ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663295f36d7811461005e578063461b09c014610089578063516c4b84146100a7578063fb55a055146100e3575b600080fd5b341561006957600080fd5b610077600435602435610111565b60405190815260200160405180910390f35b341561009457600080fd5b6100a560043560243560443561014b565b005b34156100b257600080fd5b6100ba610236565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100ee57600080fd5b6100a573ffffffffffffffffffffffffffffffffffffffff60043516610252565b005b6000600160008484604051918252602082015260409081019051908190039020815260208101919091526040016000205490505b92915050565b6000805473ffffffffffffffffffffffffffffffffffffffff169063e814861e903390604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff841602815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381600087803b15156101d957600080fd5b6102c65a03f115156101ea57600080fd5b5050506040518051905015156101ff57600080fd5b8060016000858560405191825260208201526040908101905190819003902081526020810191909152604001600020555b5b505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156102c657600080fd5b6102c65a03f115156102d757600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561030357600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b505600a165627a7a723058202cf6dc14646b90fe4f1a372c4186dbe1f838bb093662ef88225bd8fab3c4bce20029".getBytes();

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
		
			// switchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
			//
			// Solidity: function switchCryptoFiat(next address) returns()
			public Transaction switchCryptoFiat(TransactOpts opts, Address next) throws Exception {
				Interfaces args = Geth.newInterfaces(1);
				args.set(0, Geth.newInterface()); args.get(0).setAddress(next);
				

				return this.Contract.transact(opts, "switchCryptoFiat"	, args);
			}
		
	}

	public class Delegation {
		// ABI is the input ABI used to generate the binding from.
		public final static String ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\"},{\"name\":\"transfers\",\"type\":\"bytes\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"multitransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"nonceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x6060604052341561000f57600080fd5b604051602080610ece833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b610e7b806100536000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305bafaa48114610069578063516c4b84146100d9578063e218e6d214610115578063ed2a2d64146101a9578063fb55a055146101e7575b600080fd5b341561007457600080fd5b6100d7600480359060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505050923573ffffffffffffffffffffffffffffffffffffffff169250610215915050565b005b34156100e457600080fd5b6100ec61039d565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561012057600080fd5b6100d760048035906024803573ffffffffffffffffffffffffffffffffffffffff169160443591606435919060a49060843590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505050923573ffffffffffffffffffffffffffffffffffffffff1692506103b9915050565b005b34156101b457600080fd5b6101d573ffffffffffffffffffffffffffffffffffffffff6004351661059f565b60405190815260200160405180910390f35b34156101f257600080fd5b6100d773ffffffffffffffffffffffffffffffffffffffff600435166105b2565b005b600061021f610e21565b600091505b848210156103955761023684836106a8565b9050610245816020015161081d565b8051610254826020015161086d565b1061025e57600080fd5b61026d8160200151825161092d565b610285816020015182608001518360600151016109e5565b61029781604001518260600151610a11565b806040015173ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836060015160405190815260200160405180910390a36000816080015111156103895761031f838260800151610a11565b8273ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836080015160405190815260200160405180910390a35b5b600190910190610224565b5b5050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000856103c581610a40565b156103cf57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156103f157600080fd5b826103fb81610a40565b1561040557600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561042757600080fd5b61047c8989898960405193845273ffffffffffffffffffffffffffffffffffffffff929092166c0100000000000000000000000002602084015260348301526054820152607401604051809103902086610a58565b92506104878361081d565b886104918461086d565b1061049b57600080fd5b6104a5838a61092d565b6104b1838789016109e5565b6104bb8888610a11565b8773ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8960405190815260200160405180910390a360008611156105915761052f8487610a11565b8373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8860405190815260200160405180910390a35b5b5b505b5050505050505050565b60006105aa8261086d565b90505b919050565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561062657600080fd5b6102c65a03f1151561063757600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561066357600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6106b0610e21565b60008060008060008060008060006106c6610e21565b60b58c029c8d019c995060208d0151985060348d0151975060548d0151965060748d0151955060948d0151945060b48d0151935060ff60b58e0151169250601b8360ff16101561071757601b830192505b8888888860405193845273ffffffffffffffffffffffffffffffffffffffff929092166c010000000000000000000000000260208401526034830152605482015260740160405190819003902089825291506001828487876040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f115156107ce57600080fd5b50506020604051035173ffffffffffffffffffffffffffffffffffffffff90811660208301528816604082015260608101879052608081018690529950895b5050505050505050505092915050565b8061082781610b0e565b151561083257600080fd5b61083b81610b26565b1561084557600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561086757600080fd5b5b5b5050565b6000610877610b3e565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760038473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561090b57600080fd5b6102c65a03f1151561091c57600080fd5b50505060405180519150505b919050565b610935610b3e565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060038473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156109cc57600080fd5b6102c65a03f1151561039557600080fd5b5050505b5050565b60006109f083610b50565b9050818110156109ff57600080fd5b610a0b83838303610c10565b5b505050565b6000610a1c83610b50565b905081810181901015610a2e57600080fd5b610a0b83838301610c10565b5b505050565b6000600280610a4e84610cc8565b161490505b919050565b6000806000808451604114610a6c57600080fd5b602085015192506040850151915060ff6041860151169050601b8160ff161015610a9457601b015b6001868285856040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f11515610af957600080fd5b50506020604051035193505b50505092915050565b6000600180610a4e84610cc8565b161490505b919050565b6000600480610a4e84610cc8565b161490505b919050565b6000610b4a6001610d88565b90505b90565b6000610b5a610b3e565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760028473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561090b57600080fd5b6102c65a03f1151561091c57600080fd5b50505060405180519150505b919050565b610c18610b3e565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060028473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156109cc57600080fd5b6102c65a03f1151561039557600080fd5b5050505b5050565b6000610cd2610b3e565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760018473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561090b57600080fd5b6102c65a03f1151561091c57600080fd5b50505060405180519150505b919050565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561090b57600080fd5b6102c65a03f1151561091c57600080fd5b50505060405180519150505b919050565b60a06040519081016040908152600080835260208301819052908201819052606082018190526080820152905600a165627a7a72305820d3441e3e0e0c5387ea204412f83f13443c59111c719da89715875a159332ef8d0029".getBytes();

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
		public final static String ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"account\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lawEnforcer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"unfreezeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountDesignator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointAccountDesignator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointLawEnforcer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"designateAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"freezeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_lawEnforcer\",\"type\":\"address\"},{\"name\":\"_enforcementAccountDesignator\",\"type\":\"address\"},{\"name\":\"_enforcementAccount\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x6060604052341561000f57600080fd5b604051608080610cd8833981016040528080519190602001805191906020018051919060200180519150505b60008054600160a060020a03808716600160a060020a031992831617909255600180548684169083161790556002805485841690831617905560038054928416929091169190911790555b505050505b610c3e8061009a6000396000f300606060405236156100ac5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663516c4b8481146100b15780635dab2420146100ed57806372cfc9dc14610129578063788649ea1461016557806385a0f2821461019357806390f28e74146101cf578063b10725e8146101fd578063b9b0330f1461022b578063f26c159f14610259578063f3fef3a314610287578063fb55a055146102b8575b600080fd5b34156100bc57600080fd5b6100c46102e6565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100f857600080fd5b6100c4610302565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561013457600080fd5b6100c461031e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561017057600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff6004351661033a565b005b341561019e57600080fd5b6100c46103e8565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156101da57600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff60043516610404565b005b341561020857600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff60043516610471565b005b341561023657600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff600435166104de565b005b341561026457600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff60043516610583565b005b341561029257600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff60043516602435610612565b005b34156102c357600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff6004351661070e565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6001543373ffffffffffffffffffffffffffffffffffffffff90811691161461036257600080fd5b610396817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb61039082610804565b166108c4565b8073ffffffffffffffffffffffffffffffffffffffff167fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee6000604051901515815260200160405180910390a25b5b50565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b6002543373ffffffffffffffffffffffffffffffffffffffff90811691161461042c57600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6001543373ffffffffffffffffffffffffffffffffffffffff90811691161461049957600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6002543373ffffffffffffffffffffffffffffffffffffffff90811691161461050657600080fd5b806105108161097c565b1561051a57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561053c57600080fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b5b505b50565b6001543373ffffffffffffffffffffffffffffffffffffffff9081169116146105ab57600080fd5b6105c08160046105ba84610804565b176108c4565b8073ffffffffffffffffffffffffffffffffffffffff167fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee6001604051901515815260200160405180910390a25b5b50565b6001543373ffffffffffffffffffffffffffffffffffffffff90811691161461063a57600080fd5b60035473ffffffffffffffffffffffffffffffffffffffff1661065c8161097c565b1561066657600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561068857600080fd5b6106928383610994565b6003546106b59073ffffffffffffffffffffffffffffffffffffffff16836109c0565b60035473ffffffffffffffffffffffffffffffffffffffff9081169084167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35b5b505b5050565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561078257600080fd5b6102c65a03f1151561079357600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff161415156107bf57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b600061080e6109ef565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760018473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156108a257600080fd5b6102c65a03f115156108b357600080fd5b50505060405180519150505b919050565b6108cc6109ef565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060018473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561096357600080fd5b6102c65a03f1151561097457600080fd5b5050505b5050565b600060028061098a84610804565b161490505b919050565b600061099f83610a01565b9050818110156109ae57600080fd5b61070783838303610ac1565b5b505050565b60006109cb83610a01565b9050818101819010156109dd57600080fd5b61070783838301610ac1565b5b505050565b60006109fb6001610b79565b90505b90565b6000610a0b6109ef565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760028473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156108a257600080fd5b6102c65a03f115156108b357600080fd5b50505060405180519150505b919050565b610ac96109ef565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060028473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561096357600080fd5b6102c65a03f1151561097457600080fd5b5050505b5050565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b15156108a257600080fd5b6102c65a03f115156108b357600080fd5b50505060405180519150505b9190505600a165627a7a72305820292a12740e45c08d841e9c2fe6027549605d4be7f81f45f522f03401c93cd4d20029".getBytes();

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
		public final static String ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"reserveBank\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointReserveBank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_reserveBank\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]";

		
			// BYTECODE is the compiled bytecode used for deploying new contracts.
			public final static byte[] BYTECODE = "0x6060604052341561000f57600080fd5b604051604080610b6783398101604052808051919060200180519150505b60008054600160a060020a03808516600160a060020a03199283161790925560018054928416929091169190911790555b50505b610af7806100706000396000f300606060405236156100805763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302946804811461008557806318160ddd146100c1578063516c4b84146100e657806398e52f9a14610122578063b921e1631461013a578063ddf05f5914610152578063fb55a05514610180575b600080fd5b341561009057600080fd5b6100986101ae565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100cc57600080fd5b6100d46101ca565b60405190815260200160405180910390f35b34156100f157600080fd5b6100986101da565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561012d57600080fd5b6101386004356101f6565b005b341561014557600080fd5b610138600435610355565b005b341561015d57600080fd5b61013873ffffffffffffffffffffffffffffffffffffffff600435166104a1565b005b341561018b57600080fd5b61013873ffffffffffffffffffffffffffffffffffffffff6004351661050e565b005b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60006101d4610604565b90505b90565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6001546000903373ffffffffffffffffffffffffffffffffffffffff90811691161461022157600080fd5b60015473ffffffffffffffffffffffffffffffffffffffff16610243816106a9565b151561024e57600080fd5b610257816106c1565b1561026157600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561028357600080fd5b61028b610604565b91508282101561029a57600080fd5b82820391506102a8826106d9565b6001546102cb9073ffffffffffffffffffffffffffffffffffffffff1684610775565b60015460009073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405190815260200160405180910390a37ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e98260405190815260200160405180910390a15b5b505b5050565b6001546000903373ffffffffffffffffffffffffffffffffffffffff90811691161461038057600080fd5b60015473ffffffffffffffffffffffffffffffffffffffff166103a2816107a1565b156103ac57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156103ce57600080fd5b6103d66101ca565b9150828201829010156103e857600080fd5b908201906103f5826106d9565b6001546104189073ffffffffffffffffffffffffffffffffffffffff16846107b9565b60015473ffffffffffffffffffffffffffffffffffffffff1660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405190815260200160405180910390a37ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e98260405190815260200160405180910390a15b5b505b5050565b6001543373ffffffffffffffffffffffffffffffffffffffff9081169116146104c957600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561058257600080fd5b6102c65a03f1151561059357600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff161415156105bf57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b600061060e6107e8565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d76005600080604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561068957600080fd5b6102c65a03f1151561069a57600080fd5b50505060405180519150505b90565b60006001806106b7846107fa565b161490505b919050565b60006004806106b7846107fa565b161490505b919050565b6106e16107e8565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060056000846040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561075d57600080fd5b6102c65a03f1151561076e57600080fd5b5050505b50565b6000610780836108ba565b90508181101561078f57600080fd5b61034e8383830361097a565b5b505050565b60006002806106b7846107fa565b161490505b919050565b60006107c4836108ba565b9050818101819010156107d657600080fd5b61034e8383830161097a565b5b505050565b60006101d46001610a32565b90505b90565b60006108046107e8565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760018473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561089857600080fd5b6102c65a03f115156108a957600080fd5b50505060405180519150505b919050565b60006108c46107e8565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760028473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561089857600080fd5b6102c65a03f115156108a957600080fd5b50505060405180519150505b919050565b6109826107e8565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060028473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b1515610a1957600080fd5b6102c65a03f11515610a2a57600080fd5b5050505b5050565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561089857600080fd5b6102c65a03f115156108a957600080fd5b50505060405180519150505b9190505600a165627a7a723058202a365205e7fed71551eab7f444ee7914391b0f49e96aa0cc4519a82bc67b0b770029".getBytes();

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
		
	}

