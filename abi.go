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
const AccountRecoveryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"into\",\"type\":\"address\"}],\"name\":\"recoverBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recoveryAccount\",\"type\":\"address\"}],\"name\":\"designateRecoveryAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// AccountRecoveryBin is the compiled bytecode used for deploying new contracts.
const AccountRecoveryBin = `0x6060604052341561000f57600080fd5b604051602080610c2d833981016040528080519150505b60008054600160a060020a031916600160a060020a03831617905561005664010000000061005d81026109d01704565b5b5061012d565b6100746001640100000000610a2c6100a182021704565b60018054600160a060020a031916600160a060020a03928316179081905516151561009e57600080fd5b5b565b60008054600160a060020a03166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561010b57600080fd5b6102c65a03f1151561011c57600080fd5b50505060405180519150505b919050565b610af18061013c6000396000f300606060405236156100755763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632d1c5ff8811461007a5780633363375c146100ae578063516c4b84146100dc57806373d4a13a14610118578063f1375f3814610154578063fb55a05514610182575b600080fd5b341561008557600080fd5b6100ac73ffffffffffffffffffffffffffffffffffffffff600435811690602435166101b0565b005b34156100b957600080fd5b6100ac73ffffffffffffffffffffffffffffffffffffffff6004351661035c565b005b34156100e757600080fd5b6100ef610452565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561012357600080fd5b6100ef61046e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561015f57600080fd5b6100ac73ffffffffffffffffffffffffffffffffffffffff6004351661048a565b005b341561018d57600080fd5b6100ac73ffffffffffffffffffffffffffffffffffffffff60043516610498565b005b60008260006101be8261058e565b90506001808216146101cf57600080fd5b600481811614156101df57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8216151561020157600080fd5b8361020b81610635565b1561021557600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561023757600080fd5b6102408661064d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561027957600080fd5b61028e8660026102888961058e565b176106f5565b8573ffffffffffffffffffffffffffffffffffffffff167fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de60405160405180910390a26102da86610790565b93506102e68685610838565b6102f08585610864565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8660405190815260200160405180910390a35b5b505b5050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156103d057600080fd5b6102c65a03f115156103e157600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561040d57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b61044e3382610893565b5b50565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561050c57600080fd5b6102c65a03f1151561051d57600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561054957600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6001805460009173ffffffffffffffffffffffffffffffffffffffff9182169163295f36d79190851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561061357600080fd5b6102c65a03f1151561062457600080fd5b50505060405180519150505b919050565b60006002806106438461058e565b161490505b919050565b60015460009073ffffffffffffffffffffffffffffffffffffffff9081169063295f36d790600490851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561061357600080fd5b6102c65a03f1151561062457600080fd5b50505060405180519150505b919050565b6001805473ffffffffffffffffffffffffffffffffffffffff9081169163461b09c0918516846040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561077757600080fd5b6102c65a03f1151561035557600080fd5b5050505b5050565b60015460009073ffffffffffffffffffffffffffffffffffffffff9081169063295f36d790600290851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561061357600080fd5b6102c65a03f1151561062457600080fd5b50505060405180519150505b919050565b600061084383610790565b90508181101561085257600080fd5b61085e83838303610933565b5b505050565b600061086f83610790565b90508181018190101561088157600080fd5b61085e83838301610933565b5b505050565b60015473ffffffffffffffffffffffffffffffffffffffff9081169063461b09c0906004908581169085166040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561077757600080fd5b6102c65a03f1151561035557600080fd5b5050505b5050565b60015473ffffffffffffffffffffffffffffffffffffffff9081169063461b09c0906002908516846040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561077757600080fd5b6102c65a03f1151561035557600080fd5b5050505b5050565b6109da6001610a2c565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9283161790819055161515610a2957600080fd5b5b565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561061357600080fd5b6102c65a03f1151561062457600080fd5b50505060405180519150505b9190505600a165627a7a7230582072b77ae890bb44997d90588124da71e166c1ce1e1b935e9acba8c4647b53e2100029`

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

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_AccountRecovery *AccountRecoveryCaller) Data(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountRecovery.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_AccountRecovery *AccountRecoverySession) Data() (common.Address, error) {
	return _AccountRecovery.Contract.Data(&_AccountRecovery.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_AccountRecovery *AccountRecoveryCallerSession) Data() (common.Address, error) {
	return _AccountRecovery.Contract.Data(&_AccountRecovery.CallOpts)
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

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_AccountRecovery *AccountRecoveryTransactor) SwitchData(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _AccountRecovery.contract.Transact(opts, "switchData", next)
}

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_AccountRecovery *AccountRecoverySession) SwitchData(next common.Address) (*types.Transaction, error) {
	return _AccountRecovery.Contract.SwitchData(&_AccountRecovery.TransactOpts, next)
}

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_AccountRecovery *AccountRecoveryTransactorSession) SwitchData(next common.Address) (*types.Transaction, error) {
	return _AccountRecovery.Contract.SwitchData(&_AccountRecovery.TransactOpts, next)
}

// AccountsABI is the input ABI used to generate the binding from.
const AccountsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"statusOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFrozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// AccountsBin is the compiled bytecode used for deploying new contracts.
const AccountsBin = `0x6060604052341561000f57600080fd5b604051602080610b60833981016040528080519150505b60008054600160a060020a031916600160a060020a03831617905561005664010000000061005d81026109031704565b5b5061012d565b610074600164010000000061095f6100a182021704565b60018054600160a060020a031916600160a060020a03928316179081905516151561009e57600080fd5b5b565b60008054600160a060020a03166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561010b57600080fd5b6102c65a03f1151561011c57600080fd5b50505060405180519150505b919050565b610a248061013c6000396000f300606060405236156100a15763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633363375c81146100a6578063516c4b84146100d4578063673448dd146101105780636943b0171461015057806370a082311461019057806373d4a13a146101ce57806397a5d5b51461020a578063a9059cbb14610248578063e583983614610279578063fb55a055146102b9575b600080fd5b34156100b157600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff600435166102e7565b005b34156100df57600080fd5b6100e76103dd565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561011b57600080fd5b61013c73ffffffffffffffffffffffffffffffffffffffff600435166103f9565b604051901515815260200160405180910390f35b341561015b57600080fd5b61013c73ffffffffffffffffffffffffffffffffffffffff6004351661040c565b604051901515815260200160405180910390f35b341561019b57600080fd5b6101bc73ffffffffffffffffffffffffffffffffffffffff6004351661041f565b60405190815260200160405180910390f35b34156101d957600080fd5b6100e7610432565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561021557600080fd5b6101bc73ffffffffffffffffffffffffffffffffffffffff6004351661044e565b60405190815260200160405180910390f35b341561025357600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff60043516602435610461565b005b341561028457600080fd5b61013c73ffffffffffffffffffffffffffffffffffffffff6004351661056b565b604051901515815260200160405180910390f35b34156102c457600080fd5b6100d273ffffffffffffffffffffffffffffffffffffffff6004351661057e565b005b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561035b57600080fd5b6102c65a03f1151561036c57600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561039857600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600061040482610674565b90505b919050565b60006104048261068c565b90505b919050565b6000610404826106a4565b90505b919050565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60006104048261074c565b90505b919050565b600033600061046f8261074c565b905060018082161461048057600080fd5b6004818116141561049057600080fd5b73ffffffffffffffffffffffffffffffffffffffff821615156104b257600080fd5b846104bc8161068c565b156104c657600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156104e857600080fd5b3393506104f584866107f3565b6104ff868661081f565b8573ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8760405190815260200160405180910390a35b5b505b5050505050565b60006104048261084e565b90505b919050565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156105f257600080fd5b6102c65a03f1151561060357600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561062f57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60006001806106828461074c565b161490505b919050565b60006002806106828461074c565b161490505b919050565b60015460009073ffffffffffffffffffffffffffffffffffffffff9081169063295f36d790600290851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561072a57600080fd5b6102c65a03f1151561073b57600080fd5b50505060405180519150505b919050565b6001805460009173ffffffffffffffffffffffffffffffffffffffff9182169163295f36d79190851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561072a57600080fd5b6102c65a03f1151561073b57600080fd5b50505060405180519150505b919050565b60006107fe836106a4565b90508181101561080d57600080fd5b61081983838303610866565b5b505050565b600061082a836106a4565b90508181018190101561083c57600080fd5b61081983838301610866565b5b505050565b60006004806106828461074c565b161490505b919050565b60015473ffffffffffffffffffffffffffffffffffffffff9081169063461b09c0906002908516846040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156108ea57600080fd5b6102c65a03f1151561056457600080fd5b5050505b5050565b61090d600161095f565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff928316179081905516151561095c57600080fd5b5b565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561072a57600080fd5b6102c65a03f1151561073b57600080fd5b50505060405180519150505b9190505600a165627a7a723058206fa401b0ad8e624cffcfc7cc46272f349da97cc4fd129f1246e41971e0c201de0029`

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

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Accounts *AccountsCaller) Data(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Accounts *AccountsSession) Data() (common.Address, error) {
	return _Accounts.Contract.Data(&_Accounts.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Accounts *AccountsCallerSession) Data() (common.Address, error) {
	return _Accounts.Contract.Data(&_Accounts.CallOpts)
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

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Accounts *AccountsTransactor) SwitchData(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "switchData", next)
}

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Accounts *AccountsSession) SwitchData(next common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SwitchData(&_Accounts.TransactOpts, next)
}

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Accounts *AccountsTransactorSession) SwitchData(next common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SwitchData(&_Accounts.TransactOpts, next)
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
const ApprovingABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"approveAccounts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountApprover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointAccountApprover\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"closeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"approveAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_accountApprover\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// ApprovingBin is the compiled bytecode used for deploying new contracts.
const ApprovingBin = `0x6060604052341561000f57600080fd5b604051604080610a3783398101604052808051919060200180519150505b60008054600160a060020a031916600160a060020a03841617905561005d64010000000061008081026107b71704565b60028054600160a060020a031916600160a060020a0383161790555b5050610150565b61009760016401000000006108136100c482021704565b60018054600160a060020a031916600160a060020a0392831617908190551615156100c157600080fd5b5b565b60008054600160a060020a03166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561012e57600080fd5b6102c65a03f1151561013f57600080fd5b50505060405180519150505b919050565b6108d88061015f6000396000f300606060405236156100965763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663071a8b53811461009b57806307a385e6146100ec5780633363375c14610128578063516c4b841461015657806373d4a13a14610192578063c8b09109146101ce578063dd336b94146101fc578063f89f4e771461022a578063fb55a05514610258575b600080fd5b34156100a657600080fd5b6100ea600460248135818101908301358060208181020160405190810160405280939291908181526020018383602002808284375094965061028695505050505050565b005b34156100f757600080fd5b6100ff6102be565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561013357600080fd5b6100ea73ffffffffffffffffffffffffffffffffffffffff600435166102da565b005b341561016157600080fd5b6100ff6103d0565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561019d57600080fd5b6100ff6103ec565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156101d957600080fd5b6100ea73ffffffffffffffffffffffffffffffffffffffff60043516610408565b005b341561020757600080fd5b6100ea73ffffffffffffffffffffffffffffffffffffffff60043516610475565b005b341561023557600080fd5b6100ea73ffffffffffffffffffffffffffffffffffffffff600435166104fa565b005b341561026357600080fd5b6100ea73ffffffffffffffffffffffffffffffffffffffff6004351661057f565b005b60005b81518110156102b9576102b08282815181106102a157fe5b906020019060200201516104fa565b5b600101610289565b5b5050565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561034e57600080fd5b6102c65a03f1151561035f57600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561038b57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6002543373ffffffffffffffffffffffffffffffffffffffff90811691161461043057600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6002543373ffffffffffffffffffffffffffffffffffffffff90811691161461049d57600080fd5b6104b28160026104ac84610675565b1761071c565b8073ffffffffffffffffffffffffffffffffffffffff167fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de60405160405180910390a25b5b50565b6002543373ffffffffffffffffffffffffffffffffffffffff90811691161461052257600080fd5b6105378160016104ac84610675565b1761071c565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7abdf8533487db28f8c616affbb4e122d90c5ab8deb258fd21b09cee59573060405160405180910390a25b5b50565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156105f357600080fd5b6102c65a03f1151561060457600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561063057600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6001805460009173ffffffffffffffffffffffffffffffffffffffff9182169163295f36d79190851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156106fa57600080fd5b6102c65a03f1151561070b57600080fd5b50505060405180519150505b919050565b6001805473ffffffffffffffffffffffffffffffffffffffff9081169163461b09c0918516846040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561079e57600080fd5b6102c65a03f115156107af57600080fd5b5050505b5050565b6107c16001610813565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff928316179081905516151561081057600080fd5b5b565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b15156106fa57600080fd5b6102c65a03f1151561070b57600080fd5b50505060405180519150505b9190505600a165627a7a72305820bc85b14b021bb8a7ee0dafb922ea1a79b4ad8a388e937c6d022df422e54542fb0029`

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

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Approving *ApprovingCaller) Data(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Approving.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Approving *ApprovingSession) Data() (common.Address, error) {
	return _Approving.Contract.Data(&_Approving.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Approving *ApprovingCallerSession) Data() (common.Address, error) {
	return _Approving.Contract.Data(&_Approving.CallOpts)
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

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Approving *ApprovingTransactor) SwitchData(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Approving.contract.Transact(opts, "switchData", next)
}

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Approving *ApprovingSession) SwitchData(next common.Address) (*types.Transaction, error) {
	return _Approving.Contract.SwitchData(&_Approving.TransactOpts, next)
}

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Approving *ApprovingTransactorSession) SwitchData(next common.Address) (*types.Transaction, error) {
	return _Approving.Contract.SwitchData(&_Approving.TransactOpts, next)
}

// CryptoFiatABI is the input ABI used to generate the binding from.
const CryptoFiatABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointMasterAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"masterAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"contractActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"next\",\"type\":\"address\"}],\"name\":\"ContractUpgraded\",\"type\":\"event\"}]"

// CryptoFiatBin is the compiled bytecode used for deploying new contracts.
const CryptoFiatBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a0316179055600380546001810161003d8382610070565b916000526020600020900160005b8154600160a060020a033081166101009390930a92830292021916179055505b6100bb565b8154818355818115116100945760008381526020902061009491810190830161009a565b5b505050565b6100b891905b808211156100b457600081556001016100a0565b5090565b90565b6105ec806100ca6000396000f3006060604052361561008b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663028f4e47811461009057806313c01368146100c15780633fad74ad14610100578063474da79a146101255780635db4380d14610164578063874c3473146101925780639afd453c146101d0578063e814861e1461020c575b600080fd5b341561009b57600080fd5b6100bf60043573ffffffffffffffffffffffffffffffffffffffff6024351661024c565b005b34156100cc57600080fd5b6100d760043561043f565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561010b57600080fd5b610113610467565b60405190815260200160405180910390f35b341561013057600080fd5b6100d760043561046e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561016f57600080fd5b6100bf73ffffffffffffffffffffffffffffffffffffffff600435166104ad565b005b341561019d57600080fd5b61011373ffffffffffffffffffffffffffffffffffffffff6004351661051a565b60405190815260200160405180910390f35b34156101db57600080fd5b6100d761052c565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561021757600080fd5b61023873ffffffffffffffffffffffffffffffffffffffff60043516610548565b604051901515815260200160405180910390f35b60008083151561025b57600080fd5b60008481526001602052604090205473ffffffffffffffffffffffffffffffffffffffff9081169250831682141561029257600080fd5b6000543373ffffffffffffffffffffffffffffffffffffffff908116911614806102e757508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b90508015156102f557600080fd5b6102fe83610548565b1561030857600080fd5b73ffffffffffffffffffffffffffffffffffffffff82811660009081526002602090815260408083208390558783526001909152902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169185169182179055156103985773ffffffffffffffffffffffffffffffffffffffff831660009081526002602052604090208490555b837fdc69b57038334451ee12fd1742228917cea7f40dbd33cda5162e7e5754acee1c838560405173ffffffffffffffffffffffffffffffffffffffff9283168152911660208201526040908101905180910390a260038054600181016103fe8382610575565b916000526020600020900160005b815473ffffffffffffffffffffffffffffffffffffffff8088166101009390930a92830292021916179055505b50505050565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003545b90565b600380548290811061047c57fe5b906000526020600020900160005b915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146104d557600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60026020526000908152604090205481565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260026020526040812054115b919050565b8154818355818115116105995760008381526020902061059991810190830161059f565b5b505050565b61046b91905b808211156105b957600081556001016105a5565b5090565b905600a165627a7a72305820e52de72a9943af1841c3c32f993a1f70e767a0d10e8cfb1e84aeb580538ade3b0029`

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
const DataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"bucket\",\"type\":\"uint256\"},{\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"bucket\",\"type\":\"uint256\"},{\"name\":\"key\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DataBin is the compiled bytecode used for deploying new contracts.
const DataBin = `0x6060604052341561000f57600080fd5b604051602080610298833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b610245806100536000396000f300606060405263ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663295f36d78114610053578063461b09c01461007e578063516c4b841461009c575b600080fd5b341561005e57600080fd5b61006c6004356024356100d8565b60405190815260200160405180910390f35b341561008957600080fd5b61009a600435602435604435610112565b005b34156100a757600080fd5b6100af6101fd565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6000600160008484604051918252602082015260409081019051908190039020815260208101919091526040016000205490505b92915050565b6000805473ffffffffffffffffffffffffffffffffffffffff169063e814861e903390604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff841602815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381600087803b15156101a057600080fd5b6102c65a03f115156101b157600080fd5b5050506040518051905015156101c657600080fd5b8060016000858560405191825260208201526040908101905190819003902081526020810191909152604001600020555b5b505050565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820100f73b52163d03d5d794fe74d405cd7d1c58e82dc9674b4b483705ecd70f88d0029`

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

// DelegationABI is the input ABI used to generate the binding from.
const DelegationABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\"},{\"name\":\"transfers\",\"type\":\"bytes\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"multitransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"nonceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// DelegationBin is the compiled bytecode used for deploying new contracts.
const DelegationBin = `0x6060604052341561000f57600080fd5b604051602080611171833981016040528080519150505b60008054600160a060020a031916600160a060020a03831617905561005664010000000061005d8102610f141704565b5b5061012d565b6100746001640100000000610f706100a182021704565b60018054600160a060020a031916600160a060020a03928316179081905516151561009e57600080fd5b5b565b60008054600160a060020a03166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561010b57600080fd5b6102c65a03f1151561011c57600080fd5b50505060405180519150505b919050565b6110358061013c6000396000f300606060405236156100805763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305bafaa481146100855780633363375c146100f5578063516c4b841461012357806373d4a13a1461015f578063e218e6d21461019b578063ed2a2d641461022f578063fb55a0551461026d575b600080fd5b341561009057600080fd5b6100f3600480359060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505050923573ffffffffffffffffffffffffffffffffffffffff16925061029b915050565b005b341561010057600080fd5b6100f373ffffffffffffffffffffffffffffffffffffffff60043516610468565b005b341561012e57600080fd5b61013661055e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561016a57600080fd5b61013661057a565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156101a657600080fd5b6100f360048035906024803573ffffffffffffffffffffffffffffffffffffffff169160443591606435919060a49060843590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496505050923573ffffffffffffffffffffffffffffffffffffffff169250610596915050565b005b341561023a57600080fd5b61025b73ffffffffffffffffffffffffffffffffffffffff6004351661077c565b60405190815260200160405180910390f35b341561027857600080fd5b6100f373ffffffffffffffffffffffffffffffffffffffff6004351661078f565b005b60006102a5610ee6565b826102af81610885565b156102b957600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156102db57600080fd5b600092505b8583101561045e576102f2858461089d565b91506103018260200151610a12565b61030e8260400151610a68565b815161031d8360200151610aa4565b1061032757600080fd5b61033682602001518351610b4c565b61034e82602001518360800151846060015101610be9565b61036082604001518360600151610c15565b816040015173ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846060015160405190815260200160405180910390a3600082608001511115610452576103e8848360800151610c15565b8373ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846080015160405190815260200160405180910390a35b5b6001909201916102e0565b5b5b505050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156104dc57600080fd5b6102c65a03f115156104ed57600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561051957600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6000856105a281610885565b156105ac57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156105ce57600080fd5b826105d881610885565b156105e257600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561060457600080fd5b6106598989898960405193845273ffffffffffffffffffffffffffffffffffffffff929092166c0100000000000000000000000002602084015260348301526054820152607401604051809103902086610c44565b925061066483610a12565b8861066e84610aa4565b1061067857600080fd5b610682838a610b4c565b61068e83878901610be9565b6106988888610c15565b8773ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8960405190815260200160405180910390a3600086111561076e5761070c8487610c15565b8373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8860405190815260200160405180910390a35b5b5b505b5050505050505050565b600061078782610aa4565b90505b919050565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561080357600080fd5b6102c65a03f1151561081457600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561084057600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b600060028061089384610cfa565b161490505b919050565b6108a5610ee6565b60008060008060008060008060006108bb610ee6565b60b58c029c8d019c995060208d0151985060348d0151975060548d0151965060748d0151955060948d0151945060b48d0151935060ff60b58e0151169250601b8360ff16101561090c57601b830192505b8888888860405193845273ffffffffffffffffffffffffffffffffffffffff929092166c010000000000000000000000000260208401526034830152605482015260740160405190819003902089825291506001828487876040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f115156109c357600080fd5b50506020604051035173ffffffffffffffffffffffffffffffffffffffff90811660208301528816604082015260608101879052608081018690529950895b5050505050505050505092915050565b806000610a1e82610cfa565b9050600180821614610a2f57600080fd5b60048181161415610a3f57600080fd5b73ffffffffffffffffffffffffffffffffffffffff82161515610a6157600080fd5b5b5b505050565b80610a7281610885565b15610a7c57600080fd5b73ffffffffffffffffffffffffffffffffffffffff81161515610a9e57600080fd5b5b5b5050565b60015460009073ffffffffffffffffffffffffffffffffffffffff9081169063295f36d790600390851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b1515610b2a57600080fd5b6102c65a03f11515610b3b57600080fd5b50505060405180519150505b919050565b60015473ffffffffffffffffffffffffffffffffffffffff9081169063461b09c0906003908516846040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b1515610bd057600080fd5b6102c65a03f11515610be157600080fd5b5050505b5050565b6000610bf483610da1565b905081811015610c0357600080fd5b610a6183838303610e49565b5b505050565b6000610c2083610da1565b905081810181901015610c3257600080fd5b610a6183838301610e49565b5b505050565b6000806000808451604114610c5857600080fd5b602085015192506040850151915060ff6041860151169050601b8160ff161015610c8057601b015b6001868285856040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f11515610ce557600080fd5b50506020604051035193505b50505092915050565b6001805460009173ffffffffffffffffffffffffffffffffffffffff9182169163295f36d79190851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b1515610b2a57600080fd5b6102c65a03f11515610b3b57600080fd5b50505060405180519150505b919050565b60015460009073ffffffffffffffffffffffffffffffffffffffff9081169063295f36d790600290851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b1515610b2a57600080fd5b6102c65a03f11515610b3b57600080fd5b50505060405180519150505b919050565b60015473ffffffffffffffffffffffffffffffffffffffff9081169063461b09c0906002908516846040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b1515610bd057600080fd5b6102c65a03f11515610be157600080fd5b5050505b5050565b60a0604051908101604090815260008083526020830181905290820181905260608201819052608082015290565b610f1e6001610f70565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9283161790819055161515610f6d57600080fd5b5b565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b1515610b2a57600080fd5b6102c65a03f11515610b3b57600080fd5b50505060405180519150505b9190505600a165627a7a72305820a301c3bcf47591d5deebc8b3cfb626bd28e8d9e124e7b7c3c8b936ad704962120029`

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

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Delegation *DelegationCaller) Data(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Delegation.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Delegation *DelegationSession) Data() (common.Address, error) {
	return _Delegation.Contract.Data(&_Delegation.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Delegation *DelegationCallerSession) Data() (common.Address, error) {
	return _Delegation.Contract.Data(&_Delegation.CallOpts)
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

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Delegation *DelegationTransactor) SwitchData(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Delegation.contract.Transact(opts, "switchData", next)
}

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Delegation *DelegationSession) SwitchData(next common.Address) (*types.Transaction, error) {
	return _Delegation.Contract.SwitchData(&_Delegation.TransactOpts, next)
}

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Delegation *DelegationTransactorSession) SwitchData(next common.Address) (*types.Transaction, error) {
	return _Delegation.Contract.SwitchData(&_Delegation.TransactOpts, next)
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
const EnforcementABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"account\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lawEnforcer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"unfreezeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountDesignator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointAccountDesignator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointLawEnforcer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"designateAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"freezeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_lawEnforcer\",\"type\":\"address\"},{\"name\":\"_enforcementAccountDesignator\",\"type\":\"address\"},{\"name\":\"_enforcementAccount\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// EnforcementBin is the compiled bytecode used for deploying new contracts.
const EnforcementBin = `0x6060604052341561000f57600080fd5b604051608080610f42833981016040528080519190602001805191906020018051919060200180519150505b60008054600160a060020a031916600160a060020a03861617905561006b6401000000006100b28102610c901704565b60028054600160a060020a03808616600160a060020a0319928316179092556003805485841690831617905560048054928416929091169190911790555b50505050610182565b6100c96001640100000000610cec6100f682021704565b60018054600160a060020a031916600160a060020a0392831617908190551615156100f357600080fd5b5b565b60008054600160a060020a03166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561016057600080fd5b6102c65a03f1151561017157600080fd5b50505060405180519150505b919050565b610db1806101916000396000f300606060405236156100c25763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633363375c81146100c7578063516c4b84146100f55780635dab24201461013157806372cfc9dc1461016d57806373d4a13a146101a9578063788649ea146101e557806385a0f2821461021357806390f28e741461024f578063b10725e81461027d578063b9b0330f146102ab578063f26c159f146102d9578063f3fef3a314610307578063fb55a05514610338575b600080fd5b34156100d257600080fd5b6100f373ffffffffffffffffffffffffffffffffffffffff60043516610366565b005b341561010057600080fd5b61010861045c565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561013c57600080fd5b610108610478565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561017857600080fd5b610108610494565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156101b457600080fd5b6101086104b0565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156101f057600080fd5b6100f373ffffffffffffffffffffffffffffffffffffffff600435166104cc565b005b341561021e57600080fd5b61010861057a565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561025a57600080fd5b6100f373ffffffffffffffffffffffffffffffffffffffff60043516610596565b005b341561028857600080fd5b6100f373ffffffffffffffffffffffffffffffffffffffff60043516610603565b005b34156102b657600080fd5b6100f373ffffffffffffffffffffffffffffffffffffffff60043516610670565b005b34156102e457600080fd5b6100f373ffffffffffffffffffffffffffffffffffffffff60043516610715565b005b341561031257600080fd5b6100f373ffffffffffffffffffffffffffffffffffffffff600435166024356107a4565b005b341561034357600080fd5b6100f373ffffffffffffffffffffffffffffffffffffffff600435166108a0565b005b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156103da57600080fd5b6102c65a03f115156103eb57600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561041757600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6002543373ffffffffffffffffffffffffffffffffffffffff9081169116146104f457600080fd5b610528817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb61052282610996565b16610a3d565b8073ffffffffffffffffffffffffffffffffffffffff167fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee6000604051901515815260200160405180910390a25b5b50565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b6003543373ffffffffffffffffffffffffffffffffffffffff9081169116146105be57600080fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6002543373ffffffffffffffffffffffffffffffffffffffff90811691161461062b57600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6003543373ffffffffffffffffffffffffffffffffffffffff90811691161461069857600080fd5b806106a281610ad8565b156106ac57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156106ce57600080fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b5b505b50565b6002543373ffffffffffffffffffffffffffffffffffffffff90811691161461073d57600080fd5b61075281600461074c84610996565b17610a3d565b8073ffffffffffffffffffffffffffffffffffffffff167fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee6001604051901515815260200160405180910390a25b5b50565b6002543373ffffffffffffffffffffffffffffffffffffffff9081169116146107cc57600080fd5b60045473ffffffffffffffffffffffffffffffffffffffff166107ee81610ad8565b156107f857600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561081a57600080fd5b6108248383610af0565b6004546108479073ffffffffffffffffffffffffffffffffffffffff1683610b1c565b60045473ffffffffffffffffffffffffffffffffffffffff9081169084167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35b5b505b5050565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561091457600080fd5b6102c65a03f1151561092557600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561095157600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6001805460009173ffffffffffffffffffffffffffffffffffffffff9182169163295f36d79190851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b1515610a1b57600080fd5b6102c65a03f11515610a2c57600080fd5b50505060405180519150505b919050565b6001805473ffffffffffffffffffffffffffffffffffffffff9081169163461b09c0918516846040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b1515610abf57600080fd5b6102c65a03f11515610ad057600080fd5b5050505b5050565b6000600280610ae684610996565b161490505b919050565b6000610afb83610b4b565b905081811015610b0a57600080fd5b61089983838303610bf3565b5b505050565b6000610b2783610b4b565b905081810181901015610b3957600080fd5b61089983838301610bf3565b5b505050565b60015460009073ffffffffffffffffffffffffffffffffffffffff9081169063295f36d790600290851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b1515610a1b57600080fd5b6102c65a03f11515610a2c57600080fd5b50505060405180519150505b919050565b60015473ffffffffffffffffffffffffffffffffffffffff9081169063461b09c0906002908516846040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b1515610abf57600080fd5b6102c65a03f11515610ad057600080fd5b5050505b5050565b610c9a6001610cec565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9283161790819055161515610ce957600080fd5b5b565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b1515610a1b57600080fd5b6102c65a03f11515610a2c57600080fd5b50505060405180519150505b9190505600a165627a7a723058202ede946e64c97255e2caa91b1f9ff541c50045682ce7ce089af4102d1f3e50db0029`

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

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Enforcement *EnforcementCaller) Data(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Enforcement.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Enforcement *EnforcementSession) Data() (common.Address, error) {
	return _Enforcement.Contract.Data(&_Enforcement.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Enforcement *EnforcementCallerSession) Data() (common.Address, error) {
	return _Enforcement.Contract.Data(&_Enforcement.CallOpts)
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

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Enforcement *EnforcementTransactor) SwitchData(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Enforcement.contract.Transact(opts, "switchData", next)
}

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Enforcement *EnforcementSession) SwitchData(next common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.SwitchData(&_Enforcement.TransactOpts, next)
}

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Enforcement *EnforcementTransactorSession) SwitchData(next common.Address) (*types.Transaction, error) {
	return _Enforcement.Contract.SwitchData(&_Enforcement.TransactOpts, next)
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
const ReserveABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"reserveBank\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointReserveBank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_reserveBank\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// ReserveBin is the compiled bytecode used for deploying new contracts.
const ReserveBin = `0x6060604052341561000f57600080fd5b604051604080610db283398101604052808051919060200180519150505b60008054600160a060020a031916600160a060020a03841617905561005d6401000000006100808102610b321704565b60028054600160a060020a031916600160a060020a0383161790555b5050610150565b6100976001640100000000610b8e6100c482021704565b60018054600160a060020a031916600160a060020a0392831617908190551615156100c157600080fd5b5b565b60008054600160a060020a03166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b151561012e57600080fd5b6102c65a03f1151561013f57600080fd5b50505060405180519150505b919050565b610c538061015f6000396000f300606060405236156100965763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302946804811461009b57806318160ddd146100d75780633363375c146100fc578063516c4b841461012a57806373d4a13a1461016657806398e52f9a146101a2578063b921e163146101ba578063ddf05f59146101d2578063fb55a05514610200575b600080fd5b34156100a657600080fd5b6100ae61022e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100e257600080fd5b6100ea61024a565b60405190815260200160405180910390f35b341561010757600080fd5b61012873ffffffffffffffffffffffffffffffffffffffff6004351661025a565b005b341561013557600080fd5b6100ae610350565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561017157600080fd5b6100ae61036c565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156101ad57600080fd5b610128600435610388565b005b34156101c557600080fd5b6101286004356104ed565b005b34156101dd57600080fd5b61012873ffffffffffffffffffffffffffffffffffffffff60043516610639565b005b341561020b57600080fd5b61012873ffffffffffffffffffffffffffffffffffffffff600435166106a6565b005b60025473ffffffffffffffffffffffffffffffffffffffff1681565b600061025461079c565b90505b90565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156102ce57600080fd5b6102c65a03f115156102df57600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561030b57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6002546000903373ffffffffffffffffffffffffffffffffffffffff9081169116146103b357600080fd5b60025473ffffffffffffffffffffffffffffffffffffffff1660006103d78261083c565b90506001808216146103e857600080fd5b600481811614156103f857600080fd5b73ffffffffffffffffffffffffffffffffffffffff8216151561041a57600080fd5b61042261079c565b92508383101561043157600080fd5b838303925061043f836108e3565b6002546104629073ffffffffffffffffffffffffffffffffffffffff168561097a565b60025460009073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8660405190815260200160405180910390a37ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e98360405190815260200160405180910390a15b5b50505b5050565b6002546000903373ffffffffffffffffffffffffffffffffffffffff90811691161461051857600080fd5b60025473ffffffffffffffffffffffffffffffffffffffff1661053a816109a6565b1561054457600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561056657600080fd5b61056e61024a565b91508282018290101561058057600080fd5b9082019061058d826108e3565b6002546105b09073ffffffffffffffffffffffffffffffffffffffff16846109be565b60025473ffffffffffffffffffffffffffffffffffffffff1660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405190815260200160405180910390a37ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e98260405190815260200160405180910390a15b5b505b5050565b6002543373ffffffffffffffffffffffffffffffffffffffff90811691161461066157600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b6000805473ffffffffffffffffffffffffffffffffffffffff33811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561071a57600080fd5b6102c65a03f1151561072b57600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff1614151561075757600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5b50565b60015460009073ffffffffffffffffffffffffffffffffffffffff1663295f36d760058380604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561081c57600080fd5b6102c65a03f1151561082d57600080fd5b50505060405180519150505b90565b6001805460009173ffffffffffffffffffffffffffffffffffffffff9182169163295f36d79190851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156108c157600080fd5b6102c65a03f115156108d257600080fd5b50505060405180519150505b919050565b60015473ffffffffffffffffffffffffffffffffffffffff1663461b09c060056000846040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561096257600080fd5b6102c65a03f115156104e557600080fd5b5050505b50565b6000610985836109ed565b90508181101561099457600080fd5b61063283838303610a95565b5b505050565b60006002806109b48461083c565b161490505b919050565b60006109c9836109ed565b9050818101819010156109db57600080fd5b61063283838301610a95565b5b505050565b60015460009073ffffffffffffffffffffffffffffffffffffffff9081169063295f36d790600290851684604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156108c157600080fd5b6102c65a03f115156108d257600080fd5b50505060405180519150505b919050565b60015473ffffffffffffffffffffffffffffffffffffffff9081169063461b09c0906002908516846040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b1515610b1957600080fd5b6102c65a03f11515610b2a57600080fd5b5050505b5050565b610b3c6001610b8e565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9283161790819055161515610b8b57600080fd5b5b565b6000805473ffffffffffffffffffffffffffffffffffffffff166313c013688383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401602060405180830381600087803b15156108c157600080fd5b6102c65a03f115156108d257600080fd5b50505060405180519150505b9190505600a165627a7a72305820a41b07828dc6e81519212018b9b0c86bdc331a39648a4d3730fa8803d75b6e630029`

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

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Reserve *ReserveCaller) Data(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Reserve.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Reserve *ReserveSession) Data() (common.Address, error) {
	return _Reserve.Contract.Data(&_Reserve.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_Reserve *ReserveCallerSession) Data() (common.Address, error) {
	return _Reserve.Contract.Data(&_Reserve.CallOpts)
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

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Reserve *ReserveTransactor) SwitchData(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "switchData", next)
}

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Reserve *ReserveSession) SwitchData(next common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.SwitchData(&_Reserve.TransactOpts, next)
}

// SwitchData is a paid mutator transaction binding the contract method 0x3363375c.
//
// Solidity: function switchData(next address) returns()
func (_Reserve *ReserveTransactorSession) SwitchData(next common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.SwitchData(&_Reserve.TransactOpts, next)
}
