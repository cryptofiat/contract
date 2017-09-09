// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AccountRecoveryABI is the input ABI used to generate the binding from.
const AccountRecoveryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"into\",\"type\":\"address\"}],\"name\":\"recoverBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recoveryAccount\",\"type\":\"address\"}],\"name\":\"designateRecoveryAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// AccountRecoveryBin is the compiled bytecode used for deploying new contracts.
const AccountRecoveryBin = `0x6060604052341561000f57600080fd5b604051602080610a3b833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b6109e8806100536000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632d1c5ff8811461005e578063516c4b8414610092578063f1375f38146100ce578063fb55a055146100fc575b600080fd5b341561006957600080fd5b61009073ffffffffffffffffffffffffffffffffffffffff6004358116906024351661012a565b005b341561009d57600080fd5b6100a56102d0565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100d957600080fd5b61009073ffffffffffffffffffffffffffffffffffffffff600435166102ec565b005b341561010757600080fd5b61009073ffffffffffffffffffffffffffffffffffffffff600435166102fa565b005b600082610136816103f0565b151561014157600080fd5b61014a81610408565b1561015457600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561017657600080fd5b8261018081610420565b1561018a57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156101ac57600080fd5b6101b585610438565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156101ee57600080fd5b6102038560026101fd886104f8565b176105b8565b8473ffffffffffffffffffffffffffffffffffffffff167fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de60405160405180910390a261024f85610670565b925061025b8584610730565b610265848461075c565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405190815260200160405180910390a35b5b505b50505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6102f6338261078b565b5b50565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561036e57600080fd5b6102c65a03f1151561037f57600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff161415156103ab57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60006001806103fe846104f8565b161490505b919050565b60006004806103fe846104f8565b161490505b919050565b60006002806103fe846104f8565b161490505b919050565b6000610442610859565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760048473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156104d657600080fd5b6102c65a03f115156104e757600080fd5b50505060405180519150505b919050565b6000610502610859565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760018473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156104d657600080fd5b6102c65a03f115156104e757600080fd5b50505060405180519150505b919050565b6105c0610859565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060018473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561065757600080fd5b6102c65a03f115156102c757600080fd5b5050505b5050565b600061067a610859565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760028473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156104d657600080fd5b6102c65a03f115156104e757600080fd5b50505060405180519150505b919050565b600061073b83610670565b90508181101561074a57600080fd5b6107568383830361086b565b5b505050565b600061076783610670565b90508181018190101561077957600080fd5b6107568383830161086b565b5b505050565b610793610859565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060048473ffffffffffffffffffffffffffffffffffffffff166001028473ffffffffffffffffffffffffffffffffffffffff166001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561065757600080fd5b6102c65a03f115156102c757600080fd5b5050505b5050565b60006108656001610923565b90505b90565b610873610859565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060028473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561065757600080fd5b6102c65a03f115156102c757600080fd5b5050505b5050565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b15156104d657600080fd5b6102c65a03f115156104e757600080fd5b50505060405180519150505b9190505600a165627a7a7230582088333de5751e56115e8f777fcf472970f2858ef191fdf67782336925141ded010029`

// DeployAccountRecovery deploys a new Ethereum contract, binding an instance of AccountRecovery to it.
func DeployAccountRecovery(auth *bind.TransactOpts, backend bind.ContractBackend, _cryptoFiat common.Address) (common.Address, *types.Transaction, *AccountRecovery, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountRecoveryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountRecoveryBin), backend, _cryptoFiat)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountRecovery{AccountRecoveryCaller: AccountRecoveryCaller{contract: contract}, AccountRecoveryTransactor: AccountRecoveryTransactor{contract: contract}}, nil
}

// AccountRecovery is an auto generated Go binding around an Ethereum contract.
type AccountRecovery struct {
	AccountRecoveryCaller     // Read-only binding to the contract
	AccountRecoveryTransactor // Write-only binding to the contract
}

// AccountRecoveryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountRecoveryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountRecoveryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountRecoveryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountRecoverySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountRecoverySession struct {
	Contract     *AccountRecovery  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountRecoveryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountRecoveryCallerSession struct {
	Contract *AccountRecoveryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AccountRecoveryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountRecoveryTransactorSession struct {
	Contract     *AccountRecoveryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AccountRecoveryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountRecoveryRaw struct {
	Contract *AccountRecovery // Generic contract binding to access the raw methods on
}

// AccountRecoveryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountRecoveryCallerRaw struct {
	Contract *AccountRecoveryCaller // Generic read-only contract binding to access the raw methods on
}

// AccountRecoveryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountRecoveryTransactorRaw struct {
	Contract *AccountRecoveryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountRecovery creates a new instance of AccountRecovery, bound to a specific deployed contract.
func NewAccountRecovery(address common.Address, backend bind.ContractBackend) (*AccountRecovery, error) {
	contract, err := bindAccountRecovery(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountRecovery{AccountRecoveryCaller: AccountRecoveryCaller{contract: contract}, AccountRecoveryTransactor: AccountRecoveryTransactor{contract: contract}}, nil
}

// NewAccountRecoveryCaller creates a new read-only instance of AccountRecovery, bound to a specific deployed contract.
func NewAccountRecoveryCaller(address common.Address, caller bind.ContractCaller) (*AccountRecoveryCaller, error) {
	contract, err := bindAccountRecovery(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AccountRecoveryCaller{contract: contract}, nil
}

// NewAccountRecoveryTransactor creates a new write-only instance of AccountRecovery, bound to a specific deployed contract.
func NewAccountRecoveryTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountRecoveryTransactor, error) {
	contract, err := bindAccountRecovery(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AccountRecoveryTransactor{contract: contract}, nil
}

// bindAccountRecovery binds a generic wrapper to an already deployed contract.
func bindAccountRecovery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountRecoveryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountRecovery *AccountRecoveryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountRecovery.Contract.AccountRecoveryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountRecovery *AccountRecoveryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountRecovery.Contract.AccountRecoveryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountRecovery *AccountRecoveryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountRecovery.Contract.AccountRecoveryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountRecovery *AccountRecoveryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountRecovery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountRecovery *AccountRecoveryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountRecovery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountRecovery *AccountRecoveryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountRecovery.Contract.contract.Transact(opts, method, params...)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_AccountRecovery *AccountRecoveryCaller) CryptoFiat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountRecovery.contract.Call(opts, out, "cryptoFiat")
	return *ret0, err
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_AccountRecovery *AccountRecoverySession) CryptoFiat() (common.Address, error) {
	return _AccountRecovery.Contract.CryptoFiat(&_AccountRecovery.CallOpts)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_AccountRecovery *AccountRecoveryCallerSession) CryptoFiat() (common.Address, error) {
	return _AccountRecovery.Contract.CryptoFiat(&_AccountRecovery.CallOpts)
}

// DesignateRecoveryAccount is a paid mutator transaction binding the contract method 0xf1375f38.
//
// Solidity: function designateRecoveryAccount(recoveryAccount address) returns()
func (_AccountRecovery *AccountRecoveryTransactor) DesignateRecoveryAccount(opts *bind.TransactOpts, recoveryAccount common.Address) (*types.Transaction, error) {
	return _AccountRecovery.contract.Transact(opts, "designateRecoveryAccount", recoveryAccount)
}

// DesignateRecoveryAccount is a paid mutator transaction binding the contract method 0xf1375f38.
//
// Solidity: function designateRecoveryAccount(recoveryAccount address) returns()
func (_AccountRecovery *AccountRecoverySession) DesignateRecoveryAccount(recoveryAccount common.Address) (*types.Transaction, error) {
	return _AccountRecovery.Contract.DesignateRecoveryAccount(&_AccountRecovery.TransactOpts, recoveryAccount)
}

// DesignateRecoveryAccount is a paid mutator transaction binding the contract method 0xf1375f38.
//
// Solidity: function designateRecoveryAccount(recoveryAccount address) returns()
func (_AccountRecovery *AccountRecoveryTransactorSession) DesignateRecoveryAccount(recoveryAccount common.Address) (*types.Transaction, error) {
	return _AccountRecovery.Contract.DesignateRecoveryAccount(&_AccountRecovery.TransactOpts, recoveryAccount)
}

// RecoverBalance is a paid mutator transaction binding the contract method 0x2d1c5ff8.
//
// Solidity: function recoverBalance(from address, into address) returns()
func (_AccountRecovery *AccountRecoveryTransactor) RecoverBalance(opts *bind.TransactOpts, from common.Address, into common.Address) (*types.Transaction, error) {
	return _AccountRecovery.contract.Transact(opts, "recoverBalance", from, into)
}

// RecoverBalance is a paid mutator transaction binding the contract method 0x2d1c5ff8.
//
// Solidity: function recoverBalance(from address, into address) returns()
func (_AccountRecovery *AccountRecoverySession) RecoverBalance(from common.Address, into common.Address) (*types.Transaction, error) {
	return _AccountRecovery.Contract.RecoverBalance(&_AccountRecovery.TransactOpts, from, into)
}

// RecoverBalance is a paid mutator transaction binding the contract method 0x2d1c5ff8.
//
// Solidity: function recoverBalance(from address, into address) returns()
func (_AccountRecovery *AccountRecoveryTransactorSession) RecoverBalance(from common.Address, into common.Address) (*types.Transaction, error) {
	return _AccountRecovery.Contract.RecoverBalance(&_AccountRecovery.TransactOpts, from, into)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_AccountRecovery *AccountRecoveryTransactor) SwitchCryptoFiat(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _AccountRecovery.contract.Transact(opts, "switchCryptoFiat", next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_AccountRecovery *AccountRecoverySession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _AccountRecovery.Contract.SwitchCryptoFiat(&_AccountRecovery.TransactOpts, next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_AccountRecovery *AccountRecoveryTransactorSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _AccountRecovery.Contract.SwitchCryptoFiat(&_AccountRecovery.TransactOpts, next)
}

// AccountsABI is the input ABI used to generate the binding from.
const AccountsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"statusOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFrozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// AccountsBin is the compiled bytecode used for deploying new contracts.
const AccountsBin = `0x6060604052341561000f57600080fd5b6040516020806108e1833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b61088e806100536000396000f3006060604052361561008b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663516c4b848114610090578063673448dd146100cc5780636943b0171461010c57806370a082311461014c57806397a5d5b51461018a578063a9059cbb146101c8578063e5839836146101f9578063fb55a05514610239575b600080fd5b341561009b57600080fd5b6100a3610267565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100d757600080fd5b6100f873ffffffffffffffffffffffffffffffffffffffff60043516610283565b604051901515815260200160405180910390f35b341561011757600080fd5b6100f873ffffffffffffffffffffffffffffffffffffffff60043516610296565b604051901515815260200160405180910390f35b341561015757600080fd5b61017873ffffffffffffffffffffffffffffffffffffffff600435166102a9565b60405190815260200160405180910390f35b341561019557600080fd5b61017873ffffffffffffffffffffffffffffffffffffffff600435166102bc565b60405190815260200160405180910390f35b34156101d357600080fd5b6101f773ffffffffffffffffffffffffffffffffffffffff600435166024356102cf565b005b341561020457600080fd5b6100f873ffffffffffffffffffffffffffffffffffffffff600435166103d3565b604051901515815260200160405180910390f35b341561024457600080fd5b6101f773ffffffffffffffffffffffffffffffffffffffff600435166103e6565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600061028e826104dc565b90505b919050565b600061028e826104f4565b90505b919050565b600061028e8261050c565b90505b919050565b600061028e826105cc565b90505b919050565b6000336102db816104dc565b15156102e657600080fd5b6102ef8161068c565b156102f957600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561031b57600080fd5b83610325816104f4565b1561032f57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561035157600080fd5b33925061035e83856106a4565b61036885856106d0565b8473ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8660405190815260200160405180910390a35b5b505b50505050565b600061028e8261068c565b90505b919050565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561045a57600080fd5b6102c65a03f1151561046b57600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561049757600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60006001806104ea846105cc565b161490505b919050565b60006002806104ea846105cc565b161490505b919050565b60006105166106ff565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760028473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156105aa57600080fd5b6102c65a03f115156105bb57600080fd5b50505060405180519150505b919050565b60006105d66106ff565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760018473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156105aa57600080fd5b6102c65a03f115156105bb57600080fd5b50505060405180519150505b919050565b60006004806104ea846105cc565b161490505b919050565b60006106af8361050c565b9050818110156106be57600080fd5b6106ca83838303610711565b5b505050565b60006106db8361050c565b9050818101819010156106ed57600080fd5b6106ca83838301610711565b5b505050565b600061070b60016107c9565b90505b90565b6107196106ff565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060028473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156107b057600080fd5b6102c65a03f115156103ca57600080fd5b5050505b5050565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b15156105aa57600080fd5b6102c65a03f115156105bb57600080fd5b50505060405180519150505b9190505600a165627a7a723058202676d7c634dcf467185f6cdf407e31a02e8f1aff953815a772f713a416b118960029`

// DeployAccounts deploys a new Ethereum contract, binding an instance of Accounts to it.
func DeployAccounts(auth *bind.TransactOpts, backend bind.ContractBackend, _cryptoFiat common.Address) (common.Address, *types.Transaction, *Accounts, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountsBin), backend, _cryptoFiat)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Accounts{AccountsCaller: AccountsCaller{contract: contract}, AccountsTransactor: AccountsTransactor{contract: contract}}, nil
}

// Accounts is an auto generated Go binding around an Ethereum contract.
type Accounts struct {
	AccountsCaller     // Read-only binding to the contract
	AccountsTransactor // Write-only binding to the contract
}

// AccountsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountsSession struct {
	Contract     *Accounts         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountsCallerSession struct {
	Contract *AccountsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AccountsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountsTransactorSession struct {
	Contract     *AccountsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AccountsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountsRaw struct {
	Contract *Accounts // Generic contract binding to access the raw methods on
}

// AccountsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountsCallerRaw struct {
	Contract *AccountsCaller // Generic read-only contract binding to access the raw methods on
}

// AccountsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountsTransactorRaw struct {
	Contract *AccountsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccounts creates a new instance of Accounts, bound to a specific deployed contract.
func NewAccounts(address common.Address, backend bind.ContractBackend) (*Accounts, error) {
	contract, err := bindAccounts(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accounts{AccountsCaller: AccountsCaller{contract: contract}, AccountsTransactor: AccountsTransactor{contract: contract}}, nil
}

// NewAccountsCaller creates a new read-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsCaller(address common.Address, caller bind.ContractCaller) (*AccountsCaller, error) {
	contract, err := bindAccounts(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsCaller{contract: contract}, nil
}

// NewAccountsTransactor creates a new write-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountsTransactor, error) {
	contract, err := bindAccounts(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AccountsTransactor{contract: contract}, nil
}

// bindAccounts binds a generic wrapper to an already deployed contract.
func bindAccounts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.AccountsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_Accounts *AccountsCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "balanceOf", addr)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_Accounts *AccountsSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Accounts.Contract.BalanceOf(&_Accounts.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_Accounts *AccountsCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Accounts.Contract.BalanceOf(&_Accounts.CallOpts, addr)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Accounts *AccountsCaller) CryptoFiat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "cryptoFiat")
	return *ret0, err
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Accounts *AccountsSession) CryptoFiat() (common.Address, error) {
	return _Accounts.Contract.CryptoFiat(&_Accounts.CallOpts)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Accounts *AccountsCallerSession) CryptoFiat() (common.Address, error) {
	return _Accounts.Contract.CryptoFiat(&_Accounts.CallOpts)
}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(account address) constant returns(bool)
func (_Accounts *AccountsCaller) IsApproved(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "isApproved", account)
	return *ret0, err
}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(account address) constant returns(bool)
func (_Accounts *AccountsSession) IsApproved(account common.Address) (bool, error) {
	return _Accounts.Contract.IsApproved(&_Accounts.CallOpts, account)
}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(account address) constant returns(bool)
func (_Accounts *AccountsCallerSession) IsApproved(account common.Address) (bool, error) {
	return _Accounts.Contract.IsApproved(&_Accounts.CallOpts, account)
}

// IsClosed is a free data retrieval call binding the contract method 0x6943b017.
//
// Solidity: function isClosed(account address) constant returns(bool)
func (_Accounts *AccountsCaller) IsClosed(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "isClosed", account)
	return *ret0, err
}

// IsClosed is a free data retrieval call binding the contract method 0x6943b017.
//
// Solidity: function isClosed(account address) constant returns(bool)
func (_Accounts *AccountsSession) IsClosed(account common.Address) (bool, error) {
	return _Accounts.Contract.IsClosed(&_Accounts.CallOpts, account)
}

// IsClosed is a free data retrieval call binding the contract method 0x6943b017.
//
// Solidity: function isClosed(account address) constant returns(bool)
func (_Accounts *AccountsCallerSession) IsClosed(account common.Address) (bool, error) {
	return _Accounts.Contract.IsClosed(&_Accounts.CallOpts, account)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(account address) constant returns(bool)
func (_Accounts *AccountsCaller) IsFrozen(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "isFrozen", account)
	return *ret0, err
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(account address) constant returns(bool)
func (_Accounts *AccountsSession) IsFrozen(account common.Address) (bool, error) {
	return _Accounts.Contract.IsFrozen(&_Accounts.CallOpts, account)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(account address) constant returns(bool)
func (_Accounts *AccountsCallerSession) IsFrozen(account common.Address) (bool, error) {
	return _Accounts.Contract.IsFrozen(&_Accounts.CallOpts, account)
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(addr address) constant returns(uint256)
func (_Accounts *AccountsCaller) StatusOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "statusOf", addr)
	return *ret0, err
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(addr address) constant returns(uint256)
func (_Accounts *AccountsSession) StatusOf(addr common.Address) (*big.Int, error) {
	return _Accounts.Contract.StatusOf(&_Accounts.CallOpts, addr)
}

// StatusOf is a free data retrieval call binding the contract method 0x97a5d5b5.
//
// Solidity: function statusOf(addr address) constant returns(uint256)
func (_Accounts *AccountsCallerSession) StatusOf(addr common.Address) (*big.Int, error) {
	return _Accounts.Contract.StatusOf(&_Accounts.CallOpts, addr)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Accounts *AccountsTransactor) SwitchCryptoFiat(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "switchCryptoFiat", next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Accounts *AccountsSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SwitchCryptoFiat(&_Accounts.TransactOpts, next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Accounts *AccountsTransactorSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SwitchCryptoFiat(&_Accounts.TransactOpts, next)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(destination address, amount uint256) returns()
func (_Accounts *AccountsTransactor) Transfer(opts *bind.TransactOpts, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "transfer", destination, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(destination address, amount uint256) returns()
func (_Accounts *AccountsSession) Transfer(destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Transfer(&_Accounts.TransactOpts, destination, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(destination address, amount uint256) returns()
func (_Accounts *AccountsTransactorSession) Transfer(destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Accounts.Contract.Transfer(&_Accounts.TransactOpts, destination, amount)
}

// ApprovingABI is the input ABI used to generate the binding from.
const ApprovingABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"approveAccounts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountApprover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointAccountApprover\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"closeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"approveAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_accountApprover\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// ApprovingBin is the compiled bytecode used for deploying new contracts.
const ApprovingBin = `0x6060604052341561000f57600080fd5b6040516040806107a283398101604052808051919060200180519150505b60008054600160a060020a03808516600160a060020a03199283161790925560018054928416929091169190911790555b50505b610732806100706000396000f300606060405236156100805763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663071a8b53811461008557806307a385e6146100d6578063516c4b8414610112578063c8b091091461014e578063dd336b941461017c578063f89f4e77146101aa578063fb55a055146101d8575b600080fd5b341561009057600080fd5b6100d4600460248135818101908301358060208181020160405190810160405280939291908181526020018383602002808284375094965061020695505050505050565b005b34156100e157600080fd5b6100e961023e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561011d57600080fd5b6100e961025a565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561015957600080fd5b6100d473ffffffffffffffffffffffffffffffffffffffff60043516610276565b005b341561018757600080fd5b6100d473ffffffffffffffffffffffffffffffffffffffff600435166102e3565b005b34156101b557600080fd5b6100d473ffffffffffffffffffffffffffffffffffffffff60043516610368565b005b34156101e357600080fd5b6100d473ffffffffffffffffffffffffffffffffffffffff600435166103ed565b005b60005b81518110156102395761023082828151811061022157fe5b90602001906020020151610368565b5b600101610209565b5b5050565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6001543373ffffffffffffffffffffffffffffffffffffffff90811691161461029e57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6001543373ffffffffffffffffffffffffffffffffffffffff90811691161461030b57600080fd5b61032081600261031a846104e3565b176105a3565b8073ffffffffffffffffffffffffffffffffffffffff167fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de60405160405180910390a25b5b50565b6001543373ffffffffffffffffffffffffffffffffffffffff90811691161461039057600080fd5b6103a581600161031a846104e3565b176105a3565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7abdf8533487db28f8c616affbb4e122d90c5ab8deb258fd21b09cee59573060405160405180910390a25b5b50565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561046157600080fd5b6102c65a03f1151561047257600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561049e57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60006104ed61065b565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760018473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561058157600080fd5b6102c65a03f1151561059257600080fd5b50505060405180519150505b919050565b6105ab61065b565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060018473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561064257600080fd5b6102c65a03f1151561065357600080fd5b5050505b5050565b6000610667600161066d565b90505b90565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561058157600080fd5b6102c65a03f1151561059257600080fd5b50505060405180519150505b9190505600a165627a7a72305820cf9948277410e5e7ba8c93ba7a00aa82c896d65cd0a1e3d18e2c359e1dfbdad20029`

// DeployApproving deploys a new Ethereum contract, binding an instance of Approving to it.
func DeployApproving(auth *bind.TransactOpts, backend bind.ContractBackend, _cryptoFiat common.Address, _accountApprover common.Address) (common.Address, *types.Transaction, *Approving, error) {
	parsed, err := abi.JSON(strings.NewReader(ApprovingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ApprovingBin), backend, _cryptoFiat, _accountApprover)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Approving{ApprovingCaller: ApprovingCaller{contract: contract}, ApprovingTransactor: ApprovingTransactor{contract: contract}}, nil
}

// Approving is an auto generated Go binding around an Ethereum contract.
type Approving struct {
	ApprovingCaller     // Read-only binding to the contract
	ApprovingTransactor // Write-only binding to the contract
}

// ApprovingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApprovingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApprovingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApprovingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApprovingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApprovingSession struct {
	Contract     *Approving        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApprovingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApprovingCallerSession struct {
	Contract *ApprovingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ApprovingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApprovingTransactorSession struct {
	Contract     *ApprovingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ApprovingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApprovingRaw struct {
	Contract *Approving // Generic contract binding to access the raw methods on
}

// ApprovingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApprovingCallerRaw struct {
	Contract *ApprovingCaller // Generic read-only contract binding to access the raw methods on
}

// ApprovingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApprovingTransactorRaw struct {
	Contract *ApprovingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApproving creates a new instance of Approving, bound to a specific deployed contract.
func NewApproving(address common.Address, backend bind.ContractBackend) (*Approving, error) {
	contract, err := bindApproving(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Approving{ApprovingCaller: ApprovingCaller{contract: contract}, ApprovingTransactor: ApprovingTransactor{contract: contract}}, nil
}

// NewApprovingCaller creates a new read-only instance of Approving, bound to a specific deployed contract.
func NewApprovingCaller(address common.Address, caller bind.ContractCaller) (*ApprovingCaller, error) {
	contract, err := bindApproving(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ApprovingCaller{contract: contract}, nil
}

// NewApprovingTransactor creates a new write-only instance of Approving, bound to a specific deployed contract.
func NewApprovingTransactor(address common.Address, transactor bind.ContractTransactor) (*ApprovingTransactor, error) {
	contract, err := bindApproving(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ApprovingTransactor{contract: contract}, nil
}

// bindApproving binds a generic wrapper to an already deployed contract.
func bindApproving(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApprovingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Approving *ApprovingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Approving.Contract.ApprovingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Approving *ApprovingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Approving.Contract.ApprovingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Approving *ApprovingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Approving.Contract.ApprovingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Approving *ApprovingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Approving.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Approving *ApprovingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Approving.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Approving *ApprovingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Approving.Contract.contract.Transact(opts, method, params...)
}

// AccountApprover is a free data retrieval call binding the contract method 0x07a385e6.
//
// Solidity: function accountApprover() constant returns(address)
func (_Approving *ApprovingCaller) AccountApprover(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Approving.contract.Call(opts, out, "accountApprover")
	return *ret0, err
}

// AccountApprover is a free data retrieval call binding the contract method 0x07a385e6.
//
// Solidity: function accountApprover() constant returns(address)
func (_Approving *ApprovingSession) AccountApprover() (common.Address, error) {
	return _Approving.Contract.AccountApprover(&_Approving.CallOpts)
}

// AccountApprover is a free data retrieval call binding the contract method 0x07a385e6.
//
// Solidity: function accountApprover() constant returns(address)
func (_Approving *ApprovingCallerSession) AccountApprover() (common.Address, error) {
	return _Approving.Contract.AccountApprover(&_Approving.CallOpts)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Approving *ApprovingCaller) CryptoFiat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Approving.contract.Call(opts, out, "cryptoFiat")
	return *ret0, err
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Approving *ApprovingSession) CryptoFiat() (common.Address, error) {
	return _Approving.Contract.CryptoFiat(&_Approving.CallOpts)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Approving *ApprovingCallerSession) CryptoFiat() (common.Address, error) {
	return _Approving.Contract.CryptoFiat(&_Approving.CallOpts)
}

// AppointAccountApprover is a paid mutator transaction binding the contract method 0xc8b09109.
//
// Solidity: function appointAccountApprover(next address) returns()
func (_Approving *ApprovingTransactor) AppointAccountApprover(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Approving.contract.Transact(opts, "appointAccountApprover", next)
}

// AppointAccountApprover is a paid mutator transaction binding the contract method 0xc8b09109.
//
// Solidity: function appointAccountApprover(next address) returns()
func (_Approving *ApprovingSession) AppointAccountApprover(next common.Address) (*types.Transaction, error) {
	return _Approving.Contract.AppointAccountApprover(&_Approving.TransactOpts, next)
}

// AppointAccountApprover is a paid mutator transaction binding the contract method 0xc8b09109.
//
// Solidity: function appointAccountApprover(next address) returns()
func (_Approving *ApprovingTransactorSession) AppointAccountApprover(next common.Address) (*types.Transaction, error) {
	return _Approving.Contract.AppointAccountApprover(&_Approving.TransactOpts, next)
}

// ApproveAccount is a paid mutator transaction binding the contract method 0xf89f4e77.
//
// Solidity: function approveAccount(account address) returns()
func (_Approving *ApprovingTransactor) ApproveAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Approving.contract.Transact(opts, "approveAccount", account)
}

// ApproveAccount is a paid mutator transaction binding the contract method 0xf89f4e77.
//
// Solidity: function approveAccount(account address) returns()
func (_Approving *ApprovingSession) ApproveAccount(account common.Address) (*types.Transaction, error) {
	return _Approving.Contract.ApproveAccount(&_Approving.TransactOpts, account)
}

// ApproveAccount is a paid mutator transaction binding the contract method 0xf89f4e77.
//
// Solidity: function approveAccount(account address) returns()
func (_Approving *ApprovingTransactorSession) ApproveAccount(account common.Address) (*types.Transaction, error) {
	return _Approving.Contract.ApproveAccount(&_Approving.TransactOpts, account)
}

// ApproveAccounts is a paid mutator transaction binding the contract method 0x071a8b53.
//
// Solidity: function approveAccounts(accounts address[]) returns()
func (_Approving *ApprovingTransactor) ApproveAccounts(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Approving.contract.Transact(opts, "approveAccounts", accounts)
}

// ApproveAccounts is a paid mutator transaction binding the contract method 0x071a8b53.
//
// Solidity: function approveAccounts(accounts address[]) returns()
func (_Approving *ApprovingSession) ApproveAccounts(accounts []common.Address) (*types.Transaction, error) {
	return _Approving.Contract.ApproveAccounts(&_Approving.TransactOpts, accounts)
}

// ApproveAccounts is a paid mutator transaction binding the contract method 0x071a8b53.
//
// Solidity: function approveAccounts(accounts address[]) returns()
func (_Approving *ApprovingTransactorSession) ApproveAccounts(accounts []common.Address) (*types.Transaction, error) {
	return _Approving.Contract.ApproveAccounts(&_Approving.TransactOpts, accounts)
}

// CloseAccount is a paid mutator transaction binding the contract method 0xdd336b94.
//
// Solidity: function closeAccount(account address) returns()
func (_Approving *ApprovingTransactor) CloseAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Approving.contract.Transact(opts, "closeAccount", account)
}

// CloseAccount is a paid mutator transaction binding the contract method 0xdd336b94.
//
// Solidity: function closeAccount(account address) returns()
func (_Approving *ApprovingSession) CloseAccount(account common.Address) (*types.Transaction, error) {
	return _Approving.Contract.CloseAccount(&_Approving.TransactOpts, account)
}

// CloseAccount is a paid mutator transaction binding the contract method 0xdd336b94.
//
// Solidity: function closeAccount(account address) returns()
func (_Approving *ApprovingTransactorSession) CloseAccount(account common.Address) (*types.Transaction, error) {
	return _Approving.Contract.CloseAccount(&_Approving.TransactOpts, account)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Approving *ApprovingTransactor) SwitchCryptoFiat(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Approving.contract.Transact(opts, "switchCryptoFiat", next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Approving *ApprovingSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Approving.Contract.SwitchCryptoFiat(&_Approving.TransactOpts, next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Approving *ApprovingTransactorSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Approving.Contract.SwitchCryptoFiat(&_Approving.TransactOpts, next)
}

// CryptoFiatABI is the input ABI used to generate the binding from.
const CryptoFiatABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointMasterAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"masterAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"contractActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"next\",\"type\":\"address\"}],\"name\":\"ContractUpgraded\",\"type\":\"event\"}]"

// CryptoFiatBin is the compiled bytecode used for deploying new contracts.
const CryptoFiatBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a0316179055600380546001810161003d8382610070565b916000526020600020900160005b8154600160a060020a033081166101009390930a92830292021916179055505b6100bb565b8154818355818115116100945760008381526020902061009491810190830161009a565b5b505050565b6100b891905b808211156100b457600081556001016100a0565b5090565b90565b6105ec806100ca6000396000f3006060604052361561008b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663028f4e47811461009057806313c01368146100c15780633fad74ad14610100578063474da79a146101255780635db4380d14610164578063874c3473146101925780639afd453c146101d0578063e814861e1461020c575b600080fd5b341561009b57600080fd5b6100bf60043573ffffffffffffffffffffffffffffffffffffffff6024351661024c565b005b34156100cc57600080fd5b6100d760043561043f565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561010b57600080fd5b610113610467565b60405190815260200160405180910390f35b341561013057600080fd5b6100d760043561046e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561016f57600080fd5b6100bf73ffffffffffffffffffffffffffffffffffffffff600435166104ad565b005b341561019d57600080fd5b61011373ffffffffffffffffffffffffffffffffffffffff6004351661051a565b60405190815260200160405180910390f35b34156101db57600080fd5b6100d761052c565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561021757600080fd5b61023873ffffffffffffffffffffffffffffffffffffffff60043516610548565b604051901515815260200160405180910390f35b60008083151561025b57600080fd5b60008481526001602052604090205473ffffffffffffffffffffffffffffffffffffffff9081169250831682141561029257600080fd5b6000543373ffffffffffffffffffffffffffffffffffffffff908116911614806102e757508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b90508015156102f557600080fd5b6102fe83610548565b1561030857600080fd5b73ffffffffffffffffffffffffffffffffffffffff82811660009081526002602090815260408083208390558783526001909152902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169185169182179055156103985773ffffffffffffffffffffffffffffffffffffffff831660009081526002602052604090208490555b837fdc69b57038334451ee12fd1742228917cea7f40dbd33cda5162e7e5754acee1c838560405173ffffffffffffffffffffffffffffffffffffffff9283168152911660208201526040908101905180910390a260038054600181016103fe8382610575565b916000526020600020900160005b815473ffffffffffffffffffffffffffffffffffffffff8088166101009390930a92830292021916179055505b50505050565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003545b90565b600380548290811061047c57fe5b906000526020600020900160005b915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146104d557600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60026020526000908152604090205481565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260026020526040812054115b919050565b8154818355818115116105995760008381526020902061059991810190830161059f565b5b505050565b61046b91905b808211156105b957600081556001016105a5565b5090565b905600a165627a7a72305820c9c2946e5b59e4c1b7e2931686d0ca86c872e1fd7b43cbf8c3f84f124b34cb6f0029`

// DeployCryptoFiat deploys a new Ethereum contract, binding an instance of CryptoFiat to it.
func DeployCryptoFiat(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CryptoFiat, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoFiatABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CryptoFiatBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CryptoFiat{CryptoFiatCaller: CryptoFiatCaller{contract: contract}, CryptoFiatTransactor: CryptoFiatTransactor{contract: contract}}, nil
}

// CryptoFiat is an auto generated Go binding around an Ethereum contract.
type CryptoFiat struct {
	CryptoFiatCaller     // Read-only binding to the contract
	CryptoFiatTransactor // Write-only binding to the contract
}

// CryptoFiatCaller is an auto generated read-only Go binding around an Ethereum contract.
type CryptoFiatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoFiatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CryptoFiatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoFiatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CryptoFiatSession struct {
	Contract     *CryptoFiat       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CryptoFiatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CryptoFiatCallerSession struct {
	Contract *CryptoFiatCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CryptoFiatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CryptoFiatTransactorSession struct {
	Contract     *CryptoFiatTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CryptoFiatRaw is an auto generated low-level Go binding around an Ethereum contract.
type CryptoFiatRaw struct {
	Contract *CryptoFiat // Generic contract binding to access the raw methods on
}

// CryptoFiatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CryptoFiatCallerRaw struct {
	Contract *CryptoFiatCaller // Generic read-only contract binding to access the raw methods on
}

// CryptoFiatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CryptoFiatTransactorRaw struct {
	Contract *CryptoFiatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCryptoFiat creates a new instance of CryptoFiat, bound to a specific deployed contract.
func NewCryptoFiat(address common.Address, backend bind.ContractBackend) (*CryptoFiat, error) {
	contract, err := bindCryptoFiat(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CryptoFiat{CryptoFiatCaller: CryptoFiatCaller{contract: contract}, CryptoFiatTransactor: CryptoFiatTransactor{contract: contract}}, nil
}

// NewCryptoFiatCaller creates a new read-only instance of CryptoFiat, bound to a specific deployed contract.
func NewCryptoFiatCaller(address common.Address, caller bind.ContractCaller) (*CryptoFiatCaller, error) {
	contract, err := bindCryptoFiat(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoFiatCaller{contract: contract}, nil
}

// NewCryptoFiatTransactor creates a new write-only instance of CryptoFiat, bound to a specific deployed contract.
func NewCryptoFiatTransactor(address common.Address, transactor bind.ContractTransactor) (*CryptoFiatTransactor, error) {
	contract, err := bindCryptoFiat(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &CryptoFiatTransactor{contract: contract}, nil
}

// bindCryptoFiat binds a generic wrapper to an already deployed contract.
func bindCryptoFiat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoFiatABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoFiat *CryptoFiatRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CryptoFiat.Contract.CryptoFiatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoFiat *CryptoFiatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoFiat.Contract.CryptoFiatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoFiat *CryptoFiatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoFiat.Contract.CryptoFiatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoFiat *CryptoFiatCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CryptoFiat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoFiat *CryptoFiatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoFiat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoFiat *CryptoFiatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoFiat.Contract.contract.Transact(opts, method, params...)
}

// ContractActive is a free data retrieval call binding the contract method 0xe814861e.
//
// Solidity: function contractActive(addr address) constant returns(bool)
func (_CryptoFiat *CryptoFiatCaller) ContractActive(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CryptoFiat.contract.Call(opts, out, "contractActive", addr)
	return *ret0, err
}

// ContractActive is a free data retrieval call binding the contract method 0xe814861e.
//
// Solidity: function contractActive(addr address) constant returns(bool)
func (_CryptoFiat *CryptoFiatSession) ContractActive(addr common.Address) (bool, error) {
	return _CryptoFiat.Contract.ContractActive(&_CryptoFiat.CallOpts, addr)
}

// ContractActive is a free data retrieval call binding the contract method 0xe814861e.
//
// Solidity: function contractActive(addr address) constant returns(bool)
func (_CryptoFiat *CryptoFiatCallerSession) ContractActive(addr common.Address) (bool, error) {
	return _CryptoFiat.Contract.ContractActive(&_CryptoFiat.CallOpts, addr)
}

// ContractAddress is a free data retrieval call binding the contract method 0x13c01368.
//
// Solidity: function contractAddress( uint256) constant returns(address)
func (_CryptoFiat *CryptoFiatCaller) ContractAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoFiat.contract.Call(opts, out, "contractAddress", arg0)
	return *ret0, err
}

// ContractAddress is a free data retrieval call binding the contract method 0x13c01368.
//
// Solidity: function contractAddress( uint256) constant returns(address)
func (_CryptoFiat *CryptoFiatSession) ContractAddress(arg0 *big.Int) (common.Address, error) {
	return _CryptoFiat.Contract.ContractAddress(&_CryptoFiat.CallOpts, arg0)
}

// ContractAddress is a free data retrieval call binding the contract method 0x13c01368.
//
// Solidity: function contractAddress( uint256) constant returns(address)
func (_CryptoFiat *CryptoFiatCallerSession) ContractAddress(arg0 *big.Int) (common.Address, error) {
	return _CryptoFiat.Contract.ContractAddress(&_CryptoFiat.CallOpts, arg0)
}

// ContractId is a free data retrieval call binding the contract method 0x874c3473.
//
// Solidity: function contractId( address) constant returns(uint256)
func (_CryptoFiat *CryptoFiatCaller) ContractId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CryptoFiat.contract.Call(opts, out, "contractId", arg0)
	return *ret0, err
}

// ContractId is a free data retrieval call binding the contract method 0x874c3473.
//
// Solidity: function contractId( address) constant returns(uint256)
func (_CryptoFiat *CryptoFiatSession) ContractId(arg0 common.Address) (*big.Int, error) {
	return _CryptoFiat.Contract.ContractId(&_CryptoFiat.CallOpts, arg0)
}

// ContractId is a free data retrieval call binding the contract method 0x874c3473.
//
// Solidity: function contractId( address) constant returns(uint256)
func (_CryptoFiat *CryptoFiatCallerSession) ContractId(arg0 common.Address) (*big.Int, error) {
	return _CryptoFiat.Contract.ContractId(&_CryptoFiat.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x474da79a.
//
// Solidity: function contracts( uint256) constant returns(address)
func (_CryptoFiat *CryptoFiatCaller) Contracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoFiat.contract.Call(opts, out, "contracts", arg0)
	return *ret0, err
}

// Contracts is a free data retrieval call binding the contract method 0x474da79a.
//
// Solidity: function contracts( uint256) constant returns(address)
func (_CryptoFiat *CryptoFiatSession) Contracts(arg0 *big.Int) (common.Address, error) {
	return _CryptoFiat.Contract.Contracts(&_CryptoFiat.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x474da79a.
//
// Solidity: function contracts( uint256) constant returns(address)
func (_CryptoFiat *CryptoFiatCallerSession) Contracts(arg0 *big.Int) (common.Address, error) {
	return _CryptoFiat.Contract.Contracts(&_CryptoFiat.CallOpts, arg0)
}

// ContractsLength is a free data retrieval call binding the contract method 0x3fad74ad.
//
// Solidity: function contractsLength() constant returns(uint256)
func (_CryptoFiat *CryptoFiatCaller) ContractsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CryptoFiat.contract.Call(opts, out, "contractsLength")
	return *ret0, err
}

// ContractsLength is a free data retrieval call binding the contract method 0x3fad74ad.
//
// Solidity: function contractsLength() constant returns(uint256)
func (_CryptoFiat *CryptoFiatSession) ContractsLength() (*big.Int, error) {
	return _CryptoFiat.Contract.ContractsLength(&_CryptoFiat.CallOpts)
}

// ContractsLength is a free data retrieval call binding the contract method 0x3fad74ad.
//
// Solidity: function contractsLength() constant returns(uint256)
func (_CryptoFiat *CryptoFiatCallerSession) ContractsLength() (*big.Int, error) {
	return _CryptoFiat.Contract.ContractsLength(&_CryptoFiat.CallOpts)
}

// MasterAccount is a free data retrieval call binding the contract method 0x9afd453c.
//
// Solidity: function masterAccount() constant returns(address)
func (_CryptoFiat *CryptoFiatCaller) MasterAccount(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoFiat.contract.Call(opts, out, "masterAccount")
	return *ret0, err
}

// MasterAccount is a free data retrieval call binding the contract method 0x9afd453c.
//
// Solidity: function masterAccount() constant returns(address)
func (_CryptoFiat *CryptoFiatSession) MasterAccount() (common.Address, error) {
	return _CryptoFiat.Contract.MasterAccount(&_CryptoFiat.CallOpts)
}

// MasterAccount is a free data retrieval call binding the contract method 0x9afd453c.
//
// Solidity: function masterAccount() constant returns(address)
func (_CryptoFiat *CryptoFiatCallerSession) MasterAccount() (common.Address, error) {
	return _CryptoFiat.Contract.MasterAccount(&_CryptoFiat.CallOpts)
}

// AppointMasterAccount is a paid mutator transaction binding the contract method 0x5db4380d.
//
// Solidity: function appointMasterAccount(next address) returns()
func (_CryptoFiat *CryptoFiatTransactor) AppointMasterAccount(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _CryptoFiat.contract.Transact(opts, "appointMasterAccount", next)
}

// AppointMasterAccount is a paid mutator transaction binding the contract method 0x5db4380d.
//
// Solidity: function appointMasterAccount(next address) returns()
func (_CryptoFiat *CryptoFiatSession) AppointMasterAccount(next common.Address) (*types.Transaction, error) {
	return _CryptoFiat.Contract.AppointMasterAccount(&_CryptoFiat.TransactOpts, next)
}

// AppointMasterAccount is a paid mutator transaction binding the contract method 0x5db4380d.
//
// Solidity: function appointMasterAccount(next address) returns()
func (_CryptoFiat *CryptoFiatTransactorSession) AppointMasterAccount(next common.Address) (*types.Transaction, error) {
	return _CryptoFiat.Contract.AppointMasterAccount(&_CryptoFiat.TransactOpts, next)
}

// Upgrade is a paid mutator transaction binding the contract method 0x028f4e47.
//
// Solidity: function upgrade(id uint256, next address) returns()
func (_CryptoFiat *CryptoFiatTransactor) Upgrade(opts *bind.TransactOpts, id *big.Int, next common.Address) (*types.Transaction, error) {
	return _CryptoFiat.contract.Transact(opts, "upgrade", id, next)
}

// Upgrade is a paid mutator transaction binding the contract method 0x028f4e47.
//
// Solidity: function upgrade(id uint256, next address) returns()
func (_CryptoFiat *CryptoFiatSession) Upgrade(id *big.Int, next common.Address) (*types.Transaction, error) {
	return _CryptoFiat.Contract.Upgrade(&_CryptoFiat.TransactOpts, id, next)
}

// Upgrade is a paid mutator transaction binding the contract method 0x028f4e47.
//
// Solidity: function upgrade(id uint256, next address) returns()
func (_CryptoFiat *CryptoFiatTransactorSession) Upgrade(id *big.Int, next common.Address) (*types.Transaction, error) {
	return _CryptoFiat.Contract.Upgrade(&_CryptoFiat.TransactOpts, id, next)
}

// DataABI is the input ABI used to generate the binding from.
const DataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"bucket\",\"type\":\"uint256\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bucket\",\"type\":\"uint256\"},{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// DataBin is the compiled bytecode used for deploying new contracts.
const DataBin = `0x6060604052341561000f57600080fd5b6040516020806103c7833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b610374806100536000396000f300606060405263ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663295f36d7811461005e578063461b09c014610089578063516c4b84146100a7578063fb55a055146100e3575b600080fd5b341561006957600080fd5b610077600435602435610111565b60405190815260200160405180910390f35b341561009457600080fd5b6100a560043560243560443561014b565b005b34156100b257600080fd5b6100ba610236565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100ee57600080fd5b6100a573ffffffffffffffffffffffffffffffffffffffff60043516610252565b005b6000600160008484604051918252602082015260409081019051908190039020815260208101919091526040016000205490505b92915050565b6000805473ffffffffffffffffffffffffffffffffffffffff169063e814861e903390604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff841602815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381600087803b15156101d957600080fd5b6102c65a03f115156101ea57600080fd5b5050506040518051905015156101ff57600080fd5b8060016000858560405191825260208201526040908101905190819003902081526020810191909152604001600020555b5b505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156102c657600080fd5b6102c65a03f115156102d757600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561030357600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b505600a165627a7a723058202cf6dc14646b90fe4f1a372c4186dbe1f838bb093662ef88225bd8fab3c4bce20029`

// DeployData deploys a new Ethereum contract, binding an instance of Data to it.
func DeployData(auth *bind.TransactOpts, backend bind.ContractBackend, _cryptoFiat common.Address) (common.Address, *types.Transaction, *Data, error) {
	parsed, err := abi.JSON(strings.NewReader(DataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataBin), backend, _cryptoFiat)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Data{DataCaller: DataCaller{contract: contract}, DataTransactor: DataTransactor{contract: contract}}, nil
}

// Data is an auto generated Go binding around an Ethereum contract.
type Data struct {
	DataCaller     // Read-only binding to the contract
	DataTransactor // Write-only binding to the contract
}

// DataCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataSession struct {
	Contract     *Data             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataCallerSession struct {
	Contract *DataCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataTransactorSession struct {
	Contract     *DataTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataRaw struct {
	Contract *Data // Generic contract binding to access the raw methods on
}

// DataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataCallerRaw struct {
	Contract *DataCaller // Generic read-only contract binding to access the raw methods on
}

// DataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataTransactorRaw struct {
	Contract *DataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewData creates a new instance of Data, bound to a specific deployed contract.
func NewData(address common.Address, backend bind.ContractBackend) (*Data, error) {
	contract, err := bindData(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Data{DataCaller: DataCaller{contract: contract}, DataTransactor: DataTransactor{contract: contract}}, nil
}

// NewDataCaller creates a new read-only instance of Data, bound to a specific deployed contract.
func NewDataCaller(address common.Address, caller bind.ContractCaller) (*DataCaller, error) {
	contract, err := bindData(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DataCaller{contract: contract}, nil
}

// NewDataTransactor creates a new write-only instance of Data, bound to a specific deployed contract.
func NewDataTransactor(address common.Address, transactor bind.ContractTransactor) (*DataTransactor, error) {
	contract, err := bindData(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DataTransactor{contract: contract}, nil
}

// bindData binds a generic wrapper to an already deployed contract.
func bindData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Data *DataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Data.Contract.DataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Data *DataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Data.Contract.DataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Data *DataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Data.Contract.DataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Data *DataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Data.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Data *DataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Data.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Data *DataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Data.Contract.contract.Transact(opts, method, params...)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Data *DataCaller) CryptoFiat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "cryptoFiat")
	return *ret0, err
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Data *DataSession) CryptoFiat() (common.Address, error) {
	return _Data.Contract.CryptoFiat(&_Data.CallOpts)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Data *DataCallerSession) CryptoFiat() (common.Address, error) {
	return _Data.Contract.CryptoFiat(&_Data.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x295f36d7.
//
// Solidity: function get(bucket uint256, key bytes32) constant returns(bytes32)
func (_Data *DataCaller) Get(opts *bind.CallOpts, bucket *big.Int, key [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Data.contract.Call(opts, out, "get", bucket, key)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x295f36d7.
//
// Solidity: function get(bucket uint256, key bytes32) constant returns(bytes32)
func (_Data *DataSession) Get(bucket *big.Int, key [32]byte) ([32]byte, error) {
	return _Data.Contract.Get(&_Data.CallOpts, bucket, key)
}

// Get is a free data retrieval call binding the contract method 0x295f36d7.
//
// Solidity: function get(bucket uint256, key bytes32) constant returns(bytes32)
func (_Data *DataCallerSession) Get(bucket *big.Int, key [32]byte) ([32]byte, error) {
	return _Data.Contract.Get(&_Data.CallOpts, bucket, key)
}

// Set is a paid mutator transaction binding the contract method 0x461b09c0.
//
// Solidity: function set(bucket uint256, key bytes32, value bytes32) returns()
func (_Data *DataTransactor) Set(opts *bind.TransactOpts, bucket *big.Int, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Data.contract.Transact(opts, "set", bucket, key, value)
}

// Set is a paid mutator transaction binding the contract method 0x461b09c0.
//
// Solidity: function set(bucket uint256, key bytes32, value bytes32) returns()
func (_Data *DataSession) Set(bucket *big.Int, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Data.Contract.Set(&_Data.TransactOpts, bucket, key, value)
}

// Set is a paid mutator transaction binding the contract method 0x461b09c0.
//
// Solidity: function set(bucket uint256, key bytes32, value bytes32) returns()
func (_Data *DataTransactorSession) Set(bucket *big.Int, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Data.Contract.Set(&_Data.TransactOpts, bucket, key, value)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Data *DataTransactor) SwitchCryptoFiat(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Data.contract.Transact(opts, "switchCryptoFiat", next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Data *DataSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Data.Contract.SwitchCryptoFiat(&_Data.TransactOpts, next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Data *DataTransactorSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Data.Contract.SwitchCryptoFiat(&_Data.TransactOpts, next)
}

// DelegationABI is the input ABI used to generate the binding from.
const DelegationABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\"},{\"name\":\"transfers\",\"type\":\"bytes\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"multitransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"nonceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// DelegationBin is the compiled bytecode used for deploying new contracts.
const DelegationBin = `0x6060604052341561000f57600080fd5b604051602080610ece833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b610e7b806100536000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305bafaa48114610069578063516c4b84146100d9578063e218e6d214610115578063ed2a2d64146101a9578063fb55a055146101e7575b600080fd5b341561007457600080fd5b6100d7600480359060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505050923573ffffffffffffffffffffffffffffffffffffffff169250610215915050565b005b34156100e457600080fd5b6100ec61039d565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561012057600080fd5b6100d760048035906024803573ffffffffffffffffffffffffffffffffffffffff169160443591606435919060a49060843590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505050923573ffffffffffffffffffffffffffffffffffffffff1692506103b9915050565b005b34156101b457600080fd5b6101d573ffffffffffffffffffffffffffffffffffffffff6004351661059f565b60405190815260200160405180910390f35b34156101f257600080fd5b6100d773ffffffffffffffffffffffffffffffffffffffff600435166105b2565b005b600061021f610e21565b600091505b848210156103955761023684836106a8565b9050610245816020015161081d565b8051610254826020015161086d565b1061025e57600080fd5b61026d8160200151825161092d565b610285816020015182608001518360600151016109e5565b61029781604001518260600151610a11565b806040015173ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836060015160405190815260200160405180910390a36000816080015111156103895761031f838260800151610a11565b8273ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836080015160405190815260200160405180910390a35b5b600190910190610224565b5b5050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000856103c581610a40565b156103cf57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156103f157600080fd5b826103fb81610a40565b1561040557600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561042757600080fd5b61047c8989898960405193845273ffffffffffffffffffffffffffffffffffffffff929092166c0100000000000000000000000002602084015260348301526054820152607401604051809103902086610a58565b92506104878361081d565b886104918461086d565b1061049b57600080fd5b6104a5838a61092d565b6104b1838789016109e5565b6104bb8888610a11565b8773ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8960405190815260200160405180910390a360008611156105915761052f8487610a11565b8373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8860405190815260200160405180910390a35b5b5b505b5050505050505050565b60006105aa8261086d565b90505b919050565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561062657600080fd5b6102c65a03f1151561063757600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561066357600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6106b0610e21565b60008060008060008060008060006106c6610e21565b60b58c029c8d019c995060208d0151985060348d0151975060548d0151965060748d0151955060948d0151945060b48d0151935060ff60b58e0151169250601b8360ff16101561071757601b830192505b8888888860405193845273ffffffffffffffffffffffffffffffffffffffff929092166c010000000000000000000000000260208401526034830152605482015260740160405190819003902089825291506001828487876040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f115156107ce57600080fd5b50506020604051035173ffffffffffffffffffffffffffffffffffffffff90811660208301528816604082015260608101879052608081018690529950895b5050505050505050505092915050565b8061082781610b0e565b151561083257600080fd5b61083b81610b26565b1561084557600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561086757600080fd5b5b5b5050565b6000610877610b3e565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760038473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561090b57600080fd5b6102c65a03f1151561091c57600080fd5b50505060405180519150505b919050565b610935610b3e565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060038473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156109cc57600080fd5b6102c65a03f1151561039557600080fd5b5050505b5050565b60006109f083610b50565b9050818110156109ff57600080fd5b610a0b83838303610c10565b5b505050565b6000610a1c83610b50565b905081810181901015610a2e57600080fd5b610a0b83838301610c10565b5b505050565b6000600280610a4e84610cc8565b161490505b919050565b6000806000808451604114610a6c57600080fd5b602085015192506040850151915060ff6041860151169050601b8160ff161015610a9457601b015b6001868285856040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f11515610af957600080fd5b50506020604051035193505b50505092915050565b6000600180610a4e84610cc8565b161490505b919050565b6000600480610a4e84610cc8565b161490505b919050565b6000610b4a6001610d88565b90505b90565b6000610b5a610b3e565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760028473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561090b57600080fd5b6102c65a03f1151561091c57600080fd5b50505060405180519150505b919050565b610c18610b3e565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060028473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156109cc57600080fd5b6102c65a03f1151561039557600080fd5b5050505b5050565b6000610cd2610b3e565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760018473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561090b57600080fd5b6102c65a03f1151561091c57600080fd5b50505060405180519150505b919050565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561090b57600080fd5b6102c65a03f1151561091c57600080fd5b50505060405180519150505b919050565b60a06040519081016040908152600080835260208301819052908201819052606082018190526080820152905600a165627a7a72305820d3441e3e0e0c5387ea204412f83f13443c59111c719da89715875a159332ef8d0029`

// DeployDelegation deploys a new Ethereum contract, binding an instance of Delegation to it.
func DeployDelegation(auth *bind.TransactOpts, backend bind.ContractBackend, _cryptoFiat common.Address) (common.Address, *types.Transaction, *Delegation, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DelegationBin), backend, _cryptoFiat)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Delegation{DelegationCaller: DelegationCaller{contract: contract}, DelegationTransactor: DelegationTransactor{contract: contract}}, nil
}

// Delegation is an auto generated Go binding around an Ethereum contract.
type Delegation struct {
	DelegationCaller     // Read-only binding to the contract
	DelegationTransactor // Write-only binding to the contract
}

// DelegationCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegationSession struct {
	Contract     *Delegation       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegationCallerSession struct {
	Contract *DelegationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DelegationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegationTransactorSession struct {
	Contract     *DelegationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DelegationRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegationRaw struct {
	Contract *Delegation // Generic contract binding to access the raw methods on
}

// DelegationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegationCallerRaw struct {
	Contract *DelegationCaller // Generic read-only contract binding to access the raw methods on
}

// DelegationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegationTransactorRaw struct {
	Contract *DelegationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegation creates a new instance of Delegation, bound to a specific deployed contract.
func NewDelegation(address common.Address, backend bind.ContractBackend) (*Delegation, error) {
	contract, err := bindDelegation(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Delegation{DelegationCaller: DelegationCaller{contract: contract}, DelegationTransactor: DelegationTransactor{contract: contract}}, nil
}

// NewDelegationCaller creates a new read-only instance of Delegation, bound to a specific deployed contract.
func NewDelegationCaller(address common.Address, caller bind.ContractCaller) (*DelegationCaller, error) {
	contract, err := bindDelegation(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DelegationCaller{contract: contract}, nil
}

// NewDelegationTransactor creates a new write-only instance of Delegation, bound to a specific deployed contract.
func NewDelegationTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegationTransactor, error) {
	contract, err := bindDelegation(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DelegationTransactor{contract: contract}, nil
}

// bindDelegation binds a generic wrapper to an already deployed contract.
func bindDelegation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegation *DelegationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Delegation.Contract.DelegationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegation *DelegationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegation.Contract.DelegationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegation *DelegationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegation.Contract.DelegationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegation *DelegationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Delegation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegation *DelegationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegation *DelegationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegation.Contract.contract.Transact(opts, method, params...)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Delegation *DelegationCaller) CryptoFiat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Delegation.contract.Call(opts, out, "cryptoFiat")
	return *ret0, err
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Delegation *DelegationSession) CryptoFiat() (common.Address, error) {
	return _Delegation.Contract.CryptoFiat(&_Delegation.CallOpts)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Delegation *DelegationCallerSession) CryptoFiat() (common.Address, error) {
	return _Delegation.Contract.CryptoFiat(&_Delegation.CallOpts)
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(account address) constant returns(uint256)
func (_Delegation *DelegationCaller) NonceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Delegation.contract.Call(opts, out, "nonceOf", account)
	return *ret0, err
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(account address) constant returns(uint256)
func (_Delegation *DelegationSession) NonceOf(account common.Address) (*big.Int, error) {
	return _Delegation.Contract.NonceOf(&_Delegation.CallOpts, account)
}

// NonceOf is a free data retrieval call binding the contract method 0xed2a2d64.
//
// Solidity: function nonceOf(account address) constant returns(uint256)
func (_Delegation *DelegationCallerSession) NonceOf(account common.Address) (*big.Int, error) {
	return _Delegation.Contract.NonceOf(&_Delegation.CallOpts, account)
}

// Multitransfer is a paid mutator transaction binding the contract method 0x05bafaa4.
//
// Solidity: function multitransfer(count uint256, transfers bytes, delegate address) returns()
func (_Delegation *DelegationTransactor) Multitransfer(opts *bind.TransactOpts, count *big.Int, transfers []byte, delegate common.Address) (*types.Transaction, error) {
	return _Delegation.contract.Transact(opts, "multitransfer", count, transfers, delegate)
}

// Multitransfer is a paid mutator transaction binding the contract method 0x05bafaa4.
//
// Solidity: function multitransfer(count uint256, transfers bytes, delegate address) returns()
func (_Delegation *DelegationSession) Multitransfer(count *big.Int, transfers []byte, delegate common.Address) (*types.Transaction, error) {
	return _Delegation.Contract.Multitransfer(&_Delegation.TransactOpts, count, transfers, delegate)
}

// Multitransfer is a paid mutator transaction binding the contract method 0x05bafaa4.
//
// Solidity: function multitransfer(count uint256, transfers bytes, delegate address) returns()
func (_Delegation *DelegationTransactorSession) Multitransfer(count *big.Int, transfers []byte, delegate common.Address) (*types.Transaction, error) {
	return _Delegation.Contract.Multitransfer(&_Delegation.TransactOpts, count, transfers, delegate)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Delegation *DelegationTransactor) SwitchCryptoFiat(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Delegation.contract.Transact(opts, "switchCryptoFiat", next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Delegation *DelegationSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Delegation.Contract.SwitchCryptoFiat(&_Delegation.TransactOpts, next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Delegation *DelegationTransactorSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Delegation.Contract.SwitchCryptoFiat(&_Delegation.TransactOpts, next)
}

// Transfer is a paid mutator transaction binding the contract method 0xe218e6d2.
//
// Solidity: function transfer(nonce uint256, destination address, amount uint256, fee uint256, signature bytes, delegate address) returns()
func (_Delegation *DelegationTransactor) Transfer(opts *bind.TransactOpts, nonce *big.Int, destination common.Address, amount *big.Int, fee *big.Int, signature []byte, delegate common.Address) (*types.Transaction, error) {
	return _Delegation.contract.Transact(opts, "transfer", nonce, destination, amount, fee, signature, delegate)
}

// Transfer is a paid mutator transaction binding the contract method 0xe218e6d2.
//
// Solidity: function transfer(nonce uint256, destination address, amount uint256, fee uint256, signature bytes, delegate address) returns()
func (_Delegation *DelegationSession) Transfer(nonce *big.Int, destination common.Address, amount *big.Int, fee *big.Int, signature []byte, delegate common.Address) (*types.Transaction, error) {
	return _Delegation.Contract.Transfer(&_Delegation.TransactOpts, nonce, destination, amount, fee, signature, delegate)
}

// Transfer is a paid mutator transaction binding the contract method 0xe218e6d2.
//
// Solidity: function transfer(nonce uint256, destination address, amount uint256, fee uint256, signature bytes, delegate address) returns()
func (_Delegation *DelegationTransactorSession) Transfer(nonce *big.Int, destination common.Address, amount *big.Int, fee *big.Int, signature []byte, delegate common.Address) (*types.Transaction, error) {
	return _Delegation.Contract.Transfer(&_Delegation.TransactOpts, nonce, destination, amount, fee, signature, delegate)
}

// EnforcementABI is the input ABI used to generate the binding from.
const EnforcementABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"account\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lawEnforcer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"unfreezeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountDesignator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointAccountDesignator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointLawEnforcer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"designateAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"freezeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_lawEnforcer\",\"type\":\"address\"},{\"name\":\"_enforcementAccountDesignator\",\"type\":\"address\"},{\"name\":\"_enforcementAccount\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// EnforcementBin is the compiled bytecode used for deploying new contracts.
const EnforcementBin = `0x6060604052341561000f57600080fd5b604051608080610cd8833981016040528080519190602001805191906020018051919060200180519150505b60008054600160a060020a03808716600160a060020a031992831617909255600180548684169083161790556002805485841690831617905560038054928416929091169190911790555b505050505b610c3e8061009a6000396000f300606060405236156100ac5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663516c4b8481146100b15780635dab2420146100ed57806372cfc9dc14610129578063788649ea1461016557806385a0f2821461019357806390f28e74146101cf578063b10725e8146101fd578063b9b0330f1461022b578063f26c159f14610259578063f3fef3a314610287578063fb55a055146102b8575b600080fd5b34156100bc57600080fd5b6100c46102e6565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100f857600080fd5b6100c4610302565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561013457600080fd5b6100c461031e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561017057600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff6004351661033a565b005b341561019e57600080fd5b6100c46103e8565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156101da57600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff60043516610404565b005b341561020857600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff60043516610471565b005b341561023657600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff600435166104de565b005b341561026457600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff60043516610583565b005b341561029257600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff60043516602435610612565b005b34156102c357600080fd5b61019173ffffffffffffffffffffffffffffffffffffffff6004351661070e565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6001543373ffffffffffffffffffffffffffffffffffffffff90811691161461036257600080fd5b610396817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb61039082610804565b166108c4565b8073ffffffffffffffffffffffffffffffffffffffff167fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee6000604051901515815260200160405180910390a25b5b50565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b6002543373ffffffffffffffffffffffffffffffffffffffff90811691161461042c57600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6001543373ffffffffffffffffffffffffffffffffffffffff90811691161461049957600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6002543373ffffffffffffffffffffffffffffffffffffffff90811691161461050657600080fd5b806105108161097c565b1561051a57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561053c57600080fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b5b505b50565b6001543373ffffffffffffffffffffffffffffffffffffffff9081169116146105ab57600080fd5b6105c08160046105ba84610804565b176108c4565b8073ffffffffffffffffffffffffffffffffffffffff167fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee6001604051901515815260200160405180910390a25b5b50565b6001543373ffffffffffffffffffffffffffffffffffffffff90811691161461063a57600080fd5b60035473ffffffffffffffffffffffffffffffffffffffff1661065c8161097c565b1561066657600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561068857600080fd5b6106928383610994565b6003546106b59073ffffffffffffffffffffffffffffffffffffffff16836109c0565b60035473ffffffffffffffffffffffffffffffffffffffff9081169084167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35b5b505b5050565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561078257600080fd5b6102c65a03f1151561079357600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff161415156107bf57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b600061080e6109ef565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760018473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156108a257600080fd5b6102c65a03f115156108b357600080fd5b50505060405180519150505b919050565b6108cc6109ef565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060018473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561096357600080fd5b6102c65a03f1151561097457600080fd5b5050505b5050565b600060028061098a84610804565b161490505b919050565b600061099f83610a01565b9050818110156109ae57600080fd5b61070783838303610ac1565b5b505050565b60006109cb83610a01565b9050818101819010156109dd57600080fd5b61070783838301610ac1565b5b505050565b60006109fb6001610b79565b90505b90565b6000610a0b6109ef565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760028473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156108a257600080fd5b6102c65a03f115156108b357600080fd5b50505060405180519150505b919050565b610ac96109ef565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060028473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561096357600080fd5b6102c65a03f1151561097457600080fd5b5050505b5050565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b15156108a257600080fd5b6102c65a03f115156108b357600080fd5b50505060405180519150505b9190505600a165627a7a72305820292a12740e45c08d841e9c2fe6027549605d4be7f81f45f522f03401c93cd4d20029`

// DeployEnforcement deploys a new Ethereum contract, binding an instance of Enforcement to it.
func DeployEnforcement(auth *bind.TransactOpts, backend bind.ContractBackend, _cryptoFiat common.Address, _lawEnforcer common.Address, _enforcementAccountDesignator common.Address, _enforcementAccount common.Address) (common.Address, *types.Transaction, *Enforcement, error) {
	parsed, err := abi.JSON(strings.NewReader(EnforcementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EnforcementBin), backend, _cryptoFiat, _lawEnforcer, _enforcementAccountDesignator, _enforcementAccount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Enforcement{EnforcementCaller: EnforcementCaller{contract: contract}, EnforcementTransactor: EnforcementTransactor{contract: contract}}, nil
}

// Enforcement is an auto generated Go binding around an Ethereum contract.
type Enforcement struct {
	EnforcementCaller     // Read-only binding to the contract
	EnforcementTransactor // Write-only binding to the contract
}

// EnforcementCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnforcementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnforcementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnforcementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnforcementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnforcementSession struct {
	Contract     *Enforcement      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnforcementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnforcementCallerSession struct {
	Contract *EnforcementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EnforcementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnforcementTransactorSession struct {
	Contract     *EnforcementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EnforcementRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnforcementRaw struct {
	Contract *Enforcement // Generic contract binding to access the raw methods on
}

// EnforcementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnforcementCallerRaw struct {
	Contract *EnforcementCaller // Generic read-only contract binding to access the raw methods on
}

// EnforcementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnforcementTransactorRaw struct {
	Contract *EnforcementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnforcement creates a new instance of Enforcement, bound to a specific deployed contract.
func NewEnforcement(address common.Address, backend bind.ContractBackend) (*Enforcement, error) {
	contract, err := bindEnforcement(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Enforcement{EnforcementCaller: EnforcementCaller{contract: contract}, EnforcementTransactor: EnforcementTransactor{contract: contract}}, nil
}

// NewEnforcementCaller creates a new read-only instance of Enforcement, bound to a specific deployed contract.
func NewEnforcementCaller(address common.Address, caller bind.ContractCaller) (*EnforcementCaller, error) {
	contract, err := bindEnforcement(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EnforcementCaller{contract: contract}, nil
}

// NewEnforcementTransactor creates a new write-only instance of Enforcement, bound to a specific deployed contract.
func NewEnforcementTransactor(address common.Address, transactor bind.ContractTransactor) (*EnforcementTransactor, error) {
	contract, err := bindEnforcement(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EnforcementTransactor{contract: contract}, nil
}

// bindEnforcement binds a generic wrapper to an already deployed contract.
func bindEnforcement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnforcementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Enforcement *EnforcementRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Enforcement.Contract.EnforcementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Enforcement *EnforcementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enforcement.Contract.EnforcementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Enforcement *EnforcementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Enforcement.Contract.EnforcementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Enforcement *EnforcementCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Enforcement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Enforcement *EnforcementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Enforcement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Enforcement *EnforcementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Enforcement.Contract.contract.Transact(opts, method, params...)
}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() constant returns(address)
func (_Enforcement *EnforcementCaller) Account(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Enforcement.contract.Call(opts, out, "account")
	return *ret0, err
}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() constant returns(address)
func (_Enforcement *EnforcementSession) Account() (common.Address, error) {
	return _Enforcement.Contract.Account(&_Enforcement.CallOpts)
}

// Account is a free data retrieval call binding the contract method 0x5dab2420.
//
// Solidity: function account() constant returns(address)
func (_Enforcement *EnforcementCallerSession) Account() (common.Address, error) {
	return _Enforcement.Contract.Account(&_Enforcement.CallOpts)
}

// AccountDesignator is a free data retrieval call binding the contract method 0x85a0f282.
//
// Solidity: function accountDesignator() constant returns(address)
func (_Enforcement *EnforcementCaller) AccountDesignator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Enforcement.contract.Call(opts, out, "accountDesignator")
	return *ret0, err
}

// AccountDesignator is a free data retrieval call binding the contract method 0x85a0f282.
//
// Solidity: function accountDesignator() constant returns(address)
func (_Enforcement *EnforcementSession) AccountDesignator() (common.Address, error) {
	return _Enforcement.Contract.AccountDesignator(&_Enforcement.CallOpts)
}

// AccountDesignator is a free data retrieval call binding the contract method 0x85a0f282.
//
// Solidity: function accountDesignator() constant returns(address)
func (_Enforcement *EnforcementCallerSession) AccountDesignator() (common.Address, error) {
	return _Enforcement.Contract.AccountDesignator(&_Enforcement.CallOpts)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Enforcement *EnforcementCaller) CryptoFiat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Enforcement.contract.Call(opts, out, "cryptoFiat")
	return *ret0, err
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Enforcement *EnforcementSession) CryptoFiat() (common.Address, error) {
	return _Enforcement.Contract.CryptoFiat(&_Enforcement.CallOpts)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Enforcement *EnforcementCallerSession) CryptoFiat() (common.Address, error) {
	return _Enforcement.Contract.CryptoFiat(&_Enforcement.CallOpts)
}

// LawEnforcer is a free data retrieval call binding the contract method 0x72cfc9dc.
//
// Solidity: function lawEnforcer() constant returns(address)
func (_Enforcement *EnforcementCaller) LawEnforcer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Enforcement.contract.Call(opts, out, "lawEnforcer")
	return *ret0, err
}

// LawEnforcer is a free data retrieval call binding the contract method 0x72cfc9dc.
//
// Solidity: function lawEnforcer() constant returns(address)
func (_Enforcement *EnforcementSession) LawEnforcer() (common.Address, error) {
	return _Enforcement.Contract.LawEnforcer(&_Enforcement.CallOpts)
}

// LawEnforcer is a free data retrieval call binding the contract method 0x72cfc9dc.
//
// Solidity: function lawEnforcer() constant returns(address)
func (_Enforcement *EnforcementCallerSession) LawEnforcer() (common.Address, error) {
	return _Enforcement.Contract.LawEnforcer(&_Enforcement.CallOpts)
}

// AppointAccountDesignator is a paid mutator transaction binding the contract method 0x90f28e74.
//
// Solidity: function appointAccountDesignator(next address) returns()
func (_Enforcement *EnforcementTransactor) AppointAccountDesignator(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Enforcement.contract.Transact(opts, "appointAccountDesignator", next)
}

// AppointAccountDesignator is a paid mutator transaction binding the contract method 0x90f28e74.
//
// Solidity: function appointAccountDesignator(next address) returns()
func (_Enforcement *EnforcementSession) AppointAccountDesignator(next common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.AppointAccountDesignator(&_Enforcement.TransactOpts, next)
}

// AppointAccountDesignator is a paid mutator transaction binding the contract method 0x90f28e74.
//
// Solidity: function appointAccountDesignator(next address) returns()
func (_Enforcement *EnforcementTransactorSession) AppointAccountDesignator(next common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.AppointAccountDesignator(&_Enforcement.TransactOpts, next)
}

// AppointLawEnforcer is a paid mutator transaction binding the contract method 0xb10725e8.
//
// Solidity: function appointLawEnforcer(next address) returns()
func (_Enforcement *EnforcementTransactor) AppointLawEnforcer(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Enforcement.contract.Transact(opts, "appointLawEnforcer", next)
}

// AppointLawEnforcer is a paid mutator transaction binding the contract method 0xb10725e8.
//
// Solidity: function appointLawEnforcer(next address) returns()
func (_Enforcement *EnforcementSession) AppointLawEnforcer(next common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.AppointLawEnforcer(&_Enforcement.TransactOpts, next)
}

// AppointLawEnforcer is a paid mutator transaction binding the contract method 0xb10725e8.
//
// Solidity: function appointLawEnforcer(next address) returns()
func (_Enforcement *EnforcementTransactorSession) AppointLawEnforcer(next common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.AppointLawEnforcer(&_Enforcement.TransactOpts, next)
}

// DesignateAccount is a paid mutator transaction binding the contract method 0xb9b0330f.
//
// Solidity: function designateAccount(_account address) returns()
func (_Enforcement *EnforcementTransactor) DesignateAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Enforcement.contract.Transact(opts, "designateAccount", _account)
}

// DesignateAccount is a paid mutator transaction binding the contract method 0xb9b0330f.
//
// Solidity: function designateAccount(_account address) returns()
func (_Enforcement *EnforcementSession) DesignateAccount(_account common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.DesignateAccount(&_Enforcement.TransactOpts, _account)
}

// DesignateAccount is a paid mutator transaction binding the contract method 0xb9b0330f.
//
// Solidity: function designateAccount(_account address) returns()
func (_Enforcement *EnforcementTransactorSession) DesignateAccount(_account common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.DesignateAccount(&_Enforcement.TransactOpts, _account)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xf26c159f.
//
// Solidity: function freezeAccount(target address) returns()
func (_Enforcement *EnforcementTransactor) FreezeAccount(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Enforcement.contract.Transact(opts, "freezeAccount", target)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xf26c159f.
//
// Solidity: function freezeAccount(target address) returns()
func (_Enforcement *EnforcementSession) FreezeAccount(target common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.FreezeAccount(&_Enforcement.TransactOpts, target)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xf26c159f.
//
// Solidity: function freezeAccount(target address) returns()
func (_Enforcement *EnforcementTransactorSession) FreezeAccount(target common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.FreezeAccount(&_Enforcement.TransactOpts, target)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Enforcement *EnforcementTransactor) SwitchCryptoFiat(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Enforcement.contract.Transact(opts, "switchCryptoFiat", next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Enforcement *EnforcementSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.SwitchCryptoFiat(&_Enforcement.TransactOpts, next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Enforcement *EnforcementTransactorSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.SwitchCryptoFiat(&_Enforcement.TransactOpts, next)
}

// UnfreezeAccount is a paid mutator transaction binding the contract method 0x788649ea.
//
// Solidity: function unfreezeAccount(target address) returns()
func (_Enforcement *EnforcementTransactor) UnfreezeAccount(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Enforcement.contract.Transact(opts, "unfreezeAccount", target)
}

// UnfreezeAccount is a paid mutator transaction binding the contract method 0x788649ea.
//
// Solidity: function unfreezeAccount(target address) returns()
func (_Enforcement *EnforcementSession) UnfreezeAccount(target common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.UnfreezeAccount(&_Enforcement.TransactOpts, target)
}

// UnfreezeAccount is a paid mutator transaction binding the contract method 0x788649ea.
//
// Solidity: function unfreezeAccount(target address) returns()
func (_Enforcement *EnforcementTransactorSession) UnfreezeAccount(target common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.UnfreezeAccount(&_Enforcement.TransactOpts, target)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(from address, amount uint256) returns()
func (_Enforcement *EnforcementTransactor) Withdraw(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enforcement.contract.Transact(opts, "withdraw", from, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(from address, amount uint256) returns()
func (_Enforcement *EnforcementSession) Withdraw(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enforcement.Contract.Withdraw(&_Enforcement.TransactOpts, from, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(from address, amount uint256) returns()
func (_Enforcement *EnforcementTransactorSession) Withdraw(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Enforcement.Contract.Withdraw(&_Enforcement.TransactOpts, from, amount)
}

// ReserveABI is the input ABI used to generate the binding from.
const ReserveABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"reserveBank\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointReserveBank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_reserveBank\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// ReserveBin is the compiled bytecode used for deploying new contracts.
const ReserveBin = `0x6060604052341561000f57600080fd5b604051604080610b6783398101604052808051919060200180519150505b60008054600160a060020a03808516600160a060020a03199283161790925560018054928416929091169190911790555b50505b610af7806100706000396000f300606060405236156100805763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302946804811461008557806318160ddd146100c1578063516c4b84146100e657806398e52f9a14610122578063b921e1631461013a578063ddf05f5914610152578063fb55a05514610180575b600080fd5b341561009057600080fd5b6100986101ae565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100cc57600080fd5b6100d46101ca565b60405190815260200160405180910390f35b34156100f157600080fd5b6100986101da565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561012d57600080fd5b6101386004356101f6565b005b341561014557600080fd5b610138600435610355565b005b341561015d57600080fd5b61013873ffffffffffffffffffffffffffffffffffffffff600435166104a1565b005b341561018b57600080fd5b61013873ffffffffffffffffffffffffffffffffffffffff6004351661050e565b005b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60006101d4610604565b90505b90565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6001546000903373ffffffffffffffffffffffffffffffffffffffff90811691161461022157600080fd5b60015473ffffffffffffffffffffffffffffffffffffffff16610243816106a9565b151561024e57600080fd5b610257816106c1565b1561026157600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561028357600080fd5b61028b610604565b91508282101561029a57600080fd5b82820391506102a8826106d9565b6001546102cb9073ffffffffffffffffffffffffffffffffffffffff1684610775565b60015460009073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405190815260200160405180910390a37ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e98260405190815260200160405180910390a15b5b505b5050565b6001546000903373ffffffffffffffffffffffffffffffffffffffff90811691161461038057600080fd5b60015473ffffffffffffffffffffffffffffffffffffffff166103a2816107a1565b156103ac57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156103ce57600080fd5b6103d66101ca565b9150828201829010156103e857600080fd5b908201906103f5826106d9565b6001546104189073ffffffffffffffffffffffffffffffffffffffff16846107b9565b60015473ffffffffffffffffffffffffffffffffffffffff1660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405190815260200160405180910390a37ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e98260405190815260200160405180910390a15b5b505b5050565b6001543373ffffffffffffffffffffffffffffffffffffffff9081169116146104c957600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561058257600080fd5b6102c65a03f1151561059357600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff161415156105bf57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b600061060e6107e8565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d76005600080604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561068957600080fd5b6102c65a03f1151561069a57600080fd5b50505060405180519150505b90565b60006001806106b7846107fa565b161490505b919050565b60006004806106b7846107fa565b161490505b919050565b6106e16107e8565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060056000846040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561075d57600080fd5b6102c65a03f1151561076e57600080fd5b5050505b50565b6000610780836108ba565b90508181101561078f57600080fd5b61034e8383830361097a565b5b505050565b60006002806106b7846107fa565b161490505b919050565b60006107c4836108ba565b9050818101819010156107d657600080fd5b61034e8383830161097a565b5b505050565b60006101d46001610a32565b90505b90565b60006108046107e8565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760018473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561089857600080fd5b6102c65a03f115156108a957600080fd5b50505060405180519150505b919050565b60006108c46107e8565b73ffffffffffffffffffffffffffffffffffffffff1663295f36d760028473ffffffffffffffffffffffffffffffffffffffff166001026000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561089857600080fd5b6102c65a03f115156108a957600080fd5b50505060405180519150505b919050565b6109826107e8565b73ffffffffffffffffffffffffffffffffffffffff1663461b09c060028473ffffffffffffffffffffffffffffffffffffffff16600102846001026040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b1515610a1957600080fd5b6102c65a03f11515610a2a57600080fd5b5050505b5050565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561089857600080fd5b6102c65a03f115156108a957600080fd5b50505060405180519150505b9190505600a165627a7a723058202a365205e7fed71551eab7f444ee7914391b0f49e96aa0cc4519a82bc67b0b770029`

// DeployReserve deploys a new Ethereum contract, binding an instance of Reserve to it.
func DeployReserve(auth *bind.TransactOpts, backend bind.ContractBackend, _cryptoFiat common.Address, _reserveBank common.Address) (common.Address, *types.Transaction, *Reserve, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReserveBin), backend, _cryptoFiat, _reserveBank)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reserve{ReserveCaller: ReserveCaller{contract: contract}, ReserveTransactor: ReserveTransactor{contract: contract}}, nil
}

// Reserve is an auto generated Go binding around an Ethereum contract.
type Reserve struct {
	ReserveCaller     // Read-only binding to the contract
	ReserveTransactor // Write-only binding to the contract
}

// ReserveCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReserveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReserveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReserveSession struct {
	Contract     *Reserve          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReserveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReserveCallerSession struct {
	Contract *ReserveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ReserveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReserveTransactorSession struct {
	Contract     *ReserveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReserveRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReserveRaw struct {
	Contract *Reserve // Generic contract binding to access the raw methods on
}

// ReserveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReserveCallerRaw struct {
	Contract *ReserveCaller // Generic read-only contract binding to access the raw methods on
}

// ReserveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReserveTransactorRaw struct {
	Contract *ReserveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReserve creates a new instance of Reserve, bound to a specific deployed contract.
func NewReserve(address common.Address, backend bind.ContractBackend) (*Reserve, error) {
	contract, err := bindReserve(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reserve{ReserveCaller: ReserveCaller{contract: contract}, ReserveTransactor: ReserveTransactor{contract: contract}}, nil
}

// NewReserveCaller creates a new read-only instance of Reserve, bound to a specific deployed contract.
func NewReserveCaller(address common.Address, caller bind.ContractCaller) (*ReserveCaller, error) {
	contract, err := bindReserve(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveCaller{contract: contract}, nil
}

// NewReserveTransactor creates a new write-only instance of Reserve, bound to a specific deployed contract.
func NewReserveTransactor(address common.Address, transactor bind.ContractTransactor) (*ReserveTransactor, error) {
	contract, err := bindReserve(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ReserveTransactor{contract: contract}, nil
}

// bindReserve binds a generic wrapper to an already deployed contract.
func bindReserve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reserve *ReserveRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reserve.Contract.ReserveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reserve *ReserveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.Contract.ReserveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reserve *ReserveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reserve.Contract.ReserveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reserve *ReserveCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Reserve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reserve *ReserveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reserve *ReserveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reserve.Contract.contract.Transact(opts, method, params...)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Reserve *ReserveCaller) CryptoFiat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "cryptoFiat")
	return *ret0, err
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Reserve *ReserveSession) CryptoFiat() (common.Address, error) {
	return _Reserve.Contract.CryptoFiat(&_Reserve.CallOpts)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Reserve *ReserveCallerSession) CryptoFiat() (common.Address, error) {
	return _Reserve.Contract.CryptoFiat(&_Reserve.CallOpts)
}

// ReserveBank is a free data retrieval call binding the contract method 0x02946804.
//
// Solidity: function reserveBank() constant returns(address)
func (_Reserve *ReserveCaller) ReserveBank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "reserveBank")
	return *ret0, err
}

// ReserveBank is a free data retrieval call binding the contract method 0x02946804.
//
// Solidity: function reserveBank() constant returns(address)
func (_Reserve *ReserveSession) ReserveBank() (common.Address, error) {
	return _Reserve.Contract.ReserveBank(&_Reserve.CallOpts)
}

// ReserveBank is a free data retrieval call binding the contract method 0x02946804.
//
// Solidity: function reserveBank() constant returns(address)
func (_Reserve *ReserveCallerSession) ReserveBank() (common.Address, error) {
	return _Reserve.Contract.ReserveBank(&_Reserve.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Reserve *ReserveCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Reserve *ReserveSession) TotalSupply() (*big.Int, error) {
	return _Reserve.Contract.TotalSupply(&_Reserve.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Reserve *ReserveCallerSession) TotalSupply() (*big.Int, error) {
	return _Reserve.Contract.TotalSupply(&_Reserve.CallOpts)
}

// AppointReserveBank is a paid mutator transaction binding the contract method 0xddf05f59.
//
// Solidity: function appointReserveBank(next address) returns()
func (_Reserve *ReserveTransactor) AppointReserveBank(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "appointReserveBank", next)
}

// AppointReserveBank is a paid mutator transaction binding the contract method 0xddf05f59.
//
// Solidity: function appointReserveBank(next address) returns()
func (_Reserve *ReserveSession) AppointReserveBank(next common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.AppointReserveBank(&_Reserve.TransactOpts, next)
}

// AppointReserveBank is a paid mutator transaction binding the contract method 0xddf05f59.
//
// Solidity: function appointReserveBank(next address) returns()
func (_Reserve *ReserveTransactorSession) AppointReserveBank(next common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.AppointReserveBank(&_Reserve.TransactOpts, next)
}

// DecreaseSupply is a paid mutator transaction binding the contract method 0x98e52f9a.
//
// Solidity: function decreaseSupply(amount uint256) returns()
func (_Reserve *ReserveTransactor) DecreaseSupply(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "decreaseSupply", amount)
}

// DecreaseSupply is a paid mutator transaction binding the contract method 0x98e52f9a.
//
// Solidity: function decreaseSupply(amount uint256) returns()
func (_Reserve *ReserveSession) DecreaseSupply(amount *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.DecreaseSupply(&_Reserve.TransactOpts, amount)
}

// DecreaseSupply is a paid mutator transaction binding the contract method 0x98e52f9a.
//
// Solidity: function decreaseSupply(amount uint256) returns()
func (_Reserve *ReserveTransactorSession) DecreaseSupply(amount *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.DecreaseSupply(&_Reserve.TransactOpts, amount)
}

// IncreaseSupply is a paid mutator transaction binding the contract method 0xb921e163.
//
// Solidity: function increaseSupply(amount uint256) returns()
func (_Reserve *ReserveTransactor) IncreaseSupply(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "increaseSupply", amount)
}

// IncreaseSupply is a paid mutator transaction binding the contract method 0xb921e163.
//
// Solidity: function increaseSupply(amount uint256) returns()
func (_Reserve *ReserveSession) IncreaseSupply(amount *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.IncreaseSupply(&_Reserve.TransactOpts, amount)
}

// IncreaseSupply is a paid mutator transaction binding the contract method 0xb921e163.
//
// Solidity: function increaseSupply(amount uint256) returns()
func (_Reserve *ReserveTransactorSession) IncreaseSupply(amount *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.IncreaseSupply(&_Reserve.TransactOpts, amount)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Reserve *ReserveTransactor) SwitchCryptoFiat(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "switchCryptoFiat", next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Reserve *ReserveSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.SwitchCryptoFiat(&_Reserve.TransactOpts, next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Reserve *ReserveTransactorSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.SwitchCryptoFiat(&_Reserve.TransactOpts, next)
}
