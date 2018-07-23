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
const AccountRecoveryBin = `0x608060405234801561001057600080fd5b50604051602080610ae7833981016040525160008054600160a060020a031916600160a060020a03831617905561004e640100000000610054810204565b5061012c565b6100676001640100000000610093810204565b60018054600160a060020a031916600160a060020a03928316179081905516151561009157600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b1580156100fa57600080fd5b505af115801561010e573d6000803e3d6000fd5b505050506040513d602081101561012457600080fd5b505192915050565b6109ac8061013b6000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632d1c5ff8811461007c5780633363375c146100b2578063516c4b84146100e057806373d4a13a1461011e578063f1375f3814610133578063fb55a05514610161575b600080fd5b34801561008857600080fd5b506100b073ffffffffffffffffffffffffffffffffffffffff6004358116906024351661018f565b005b3480156100be57600080fd5b506100b073ffffffffffffffffffffffffffffffffffffffff60043516610322565b3480156100ec57600080fd5b506100f5610451565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561012a57600080fd5b506100f561046d565b34801561013f57600080fd5b506100b073ffffffffffffffffffffffffffffffffffffffff60043516610489565b34801561016d57600080fd5b506100b073ffffffffffffffffffffffffffffffffffffffff60043516610496565b600082600061019d826105c5565b90506001808216146101ae57600080fd5b600481811614156101be57600080fd5b73ffffffffffffffffffffffffffffffffffffffff821615156101e057600080fd5b836101ea81610678565b156101f457600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561021657600080fd5b61021f8661068e565b73ffffffffffffffffffffffffffffffffffffffff16331461024057600080fd5b61025586600261024f896105c5565b1761070c565b60405173ffffffffffffffffffffffffffffffffffffffff8716907fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de90600090a261029f866107aa565b93506102ab8685610829565b6102b58585610854565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef866040518082815260200191505060405180910390a3505050505050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156103be57600080fd5b505af11580156103d2573d6000803e3d6000fd5b505050506040513d60208110156103e857600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461040a57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b610493338261087c565b50565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561053257600080fd5b505af1158015610546573d6000803e3d6000fd5b505050506040513d602081101561055c57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461057e57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60018054604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff84811660248501529051600093919092169163295f36d79160448082019260209290919082900301818787803b15801561064657600080fd5b505af115801561065a573d6000803e3d6000fd5b505050506040513d602081101561067057600080fd5b505192915050565b6000600280610686846105c5565b161492915050565b600154604080517f295f36d700000000000000000000000000000000000000000000000000000000815260048181015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b15801561064657600080fd5b60018054604080517f461b09c0000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff858116602485015260448401859052905191169163461b09c091606480830192600092919082900301818387803b15801561078e57600080fd5b505af11580156107a2573d6000803e3d6000fd5b505050505050565b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b15801561064657600080fd5b6000610834836107aa565b90508181101561084357600080fd5b61084f838383036108fe565b505050565b600061085f836107aa565b905081810181111561087057600080fd5b61084f838383016108fe565b600154604080517f461b09c000000000000000000000000000000000000000000000000000000000815260048181015273ffffffffffffffffffffffffffffffffffffffff858116602483015284811660448301529151919092169163461b09c091606480830192600092919082900301818387803b15801561078e57600080fd5b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff8581166024830152604482018590529151919092169163461b09c091606480830192600092919082900301818387803b15801561078e57600080fd00a165627a7a723058200188215b9ff1ee8d796476ad18cd5b823e46677f13aaeb488c532983cc0d65c00029`

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
const AccountsBin = `0x608060405234801561001057600080fd5b50604051602080610a30833981016040525160008054600160a060020a031916600160a060020a03831617905561004e640100000000610054810204565b5061012c565b6100676001640100000000610093810204565b60018054600160a060020a031916600160a060020a03928316179081905516151561009157600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b1580156100fa57600080fd5b505af115801561010e573d6000803e3d6000fd5b505050506040513d602081101561012457600080fd5b505192915050565b6108f58061013b6000396000f3006080604052600436106100a35763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633363375c81146100a8578063516c4b84146100d8578063673448dd146101165780636943b0171461015857806370a082311461018657806373d4a13a146101c657806397a5d5b5146101db578063a9059cbb14610209578063e58398361461023a578063fb55a05514610268575b600080fd5b3480156100b457600080fd5b506100d673ffffffffffffffffffffffffffffffffffffffff60043516610296565b005b3480156100e457600080fd5b506100ed6103c5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561012257600080fd5b5061014473ffffffffffffffffffffffffffffffffffffffff600435166103e1565b604080519115158252519081900360200190f35b34801561016457600080fd5b5061014473ffffffffffffffffffffffffffffffffffffffff600435166103f2565b34801561019257600080fd5b506101b473ffffffffffffffffffffffffffffffffffffffff600435166103fd565b60408051918252519081900360200190f35b3480156101d257600080fd5b506100ed610408565b3480156101e757600080fd5b506101b473ffffffffffffffffffffffffffffffffffffffff60043516610424565b34801561021557600080fd5b506100d673ffffffffffffffffffffffffffffffffffffffff6004351660243561042f565b34801561024657600080fd5b5061014473ffffffffffffffffffffffffffffffffffffffff6004351661053a565b34801561027457600080fd5b506100d673ffffffffffffffffffffffffffffffffffffffff60043516610545565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561033257600080fd5b505af1158015610346573d6000803e3d6000fd5b505050506040513d602081101561035c57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461037e57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60006103ec82610674565b92915050565b60006103ec8261068a565b60006103ec82610698565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60006103ec82610749565b600033600061043d82610749565b905060018082161461044e57600080fd5b6004818116141561045e57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8216151561048057600080fd5b8461048a8161068a565b1561049457600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156104b657600080fd5b3393506104c384866107ca565b6104cd86866107f5565b8573ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef876040518082815260200191505060405180910390a3505050505050565b60006103ec8261081d565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156105e157600080fd5b505af11580156105f5573d6000803e3d6000fd5b505050506040513d602081101561060b57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461062d57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600060018061068284610749565b161492915050565b600060028061068284610749565b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b15801561071757600080fd5b505af115801561072b573d6000803e3d6000fd5b505050506040513d602081101561074157600080fd5b505192915050565b60018054604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff84811660248501529051600093919092169163295f36d79160448082019260209290919082900301818787803b15801561071757600080fd5b60006107d583610698565b9050818110156107e457600080fd5b6107f08383830361082b565b505050565b600061080083610698565b905081810181111561081157600080fd5b6107f08383830161082b565b600060048061068284610749565b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff8581166024830152604482018590529151919092169163461b09c091606480830192600092919082900301818387803b1580156108ad57600080fd5b505af11580156108c1573d6000803e3d6000fd5b5050505050505600a165627a7a72305820a8a2953760b304b29f6a6fa8db54be1e13c3305a8662e705f94396e3a4c1e0cf0029`

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
const ApprovingBin = `0x608060405234801561001057600080fd5b5060405160408061096a83398101604052805160209091015160008054600160a060020a031916600160a060020a03841617905561005564010000000061007b810204565b60028054600160a060020a031916600160a060020a039290921691909117905550610153565b61008e60016401000000006100ba810204565b60018054600160a060020a031916600160a060020a0392831617908190551615156100b857600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b15801561012157600080fd5b505af1158015610135573d6000803e3d6000fd5b505050506040513d602081101561014b57600080fd5b505192915050565b610808806101626000396000f3006080604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663071a8b53811461009d57806307a385e6146100f45780633363375c14610132578063516c4b841461016057806373d4a13a14610175578063c8b091091461018a578063dd336b94146101b8578063f89f4e77146101e6578063fb55a05514610214575b600080fd5b3480156100a957600080fd5b50604080516020600480358082013583810280860185019096528085526100f2953695939460249493850192918291850190849080828437509497506102429650505050505050565b005b34801561010057600080fd5b5061010961027a565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561013e57600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff60043516610296565b34801561016c57600080fd5b506101096103c5565b34801561018157600080fd5b506101096103e1565b34801561019657600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff600435166103fd565b3480156101c457600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff60043516610468565b3480156101f257600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff600435166104e5565b34801561022057600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff6004351661055c565b60005b81518110156102765761026e828281518110151561025f57fe5b906020019060200201516104e5565b600101610245565b5050565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561033257600080fd5b505af1158015610346573d6000803e3d6000fd5b505050506040513d602081101561035c57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461037e57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60025473ffffffffffffffffffffffffffffffffffffffff16331461042157600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60025473ffffffffffffffffffffffffffffffffffffffff16331461048c57600080fd5b6104a181600261049b8461068b565b1761073e565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de90600090a250565b60025473ffffffffffffffffffffffffffffffffffffffff16331461050957600080fd5b61051881600161049b8461068b565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7abdf8533487db28f8c616affbb4e122d90c5ab8deb258fd21b09cee59573090600090a250565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156105f857600080fd5b505af115801561060c573d6000803e3d6000fd5b505050506040513d602081101561062257600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461064457600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60018054604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff84811660248501529051600093919092169163295f36d79160448082019260209290919082900301818787803b15801561070c57600080fd5b505af1158015610720573d6000803e3d6000fd5b505050506040513d602081101561073657600080fd5b505192915050565b60018054604080517f461b09c0000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff858116602485015260448401859052905191169163461b09c091606480830192600092919082900301818387803b1580156107c057600080fd5b505af11580156107d4573d6000803e3d6000fd5b5050505050505600a165627a7a7230582085224779143f6669d52695d45d40cbcebb9f30cc878327219164147116ffcee00029`

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
const CryptoFiatBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a0319908116331782556003805460018101825592527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9091018054909116301790556105468061006d6000396000f30060806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663028f4e47811461009257806313c01368146100c55780633fad74ad14610106578063474da79a1461012d5780635db4380d14610145578063874c3473146101735780639afd453c146101a1578063e814861e146101b6575b600080fd5b34801561009e57600080fd5b506100c360043573ffffffffffffffffffffffffffffffffffffffff602435166101f8565b005b3480156100d157600080fd5b506100dd6004356103f5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561011257600080fd5b5061011b61041d565b60408051918252519081900360200190f35b34801561013957600080fd5b506100dd600435610423565b34801561015157600080fd5b506100c373ffffffffffffffffffffffffffffffffffffffff60043516610458565b34801561017f57600080fd5b5061011b73ffffffffffffffffffffffffffffffffffffffff600435166104c3565b3480156101ad57600080fd5b506100dd6104d5565b3480156101c257600080fd5b506101e473ffffffffffffffffffffffffffffffffffffffff600435166104f1565b604080519115158252519081900360200190f35b60008083151561020757600080fd5b60008481526001602052604090205473ffffffffffffffffffffffffffffffffffffffff9081169250831682141561023e57600080fd5b60005473ffffffffffffffffffffffffffffffffffffffff1633148061027957503373ffffffffffffffffffffffffffffffffffffffff8316145b905080151561028757600080fd5b610290836104f1565b1561029a57600080fd5b73ffffffffffffffffffffffffffffffffffffffff82811660009081526002602090815260408083208390558783526001909152902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001691851691821790551561032a5773ffffffffffffffffffffffffffffffffffffffff831660009081526002602052604090208490555b6040805173ffffffffffffffffffffffffffffffffffffffff808516825285166020820152815186927fdc69b57038334451ee12fd1742228917cea7f40dbd33cda5162e7e5754acee1c928290030190a25050600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60035490565b600380548290811061043157fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60005473ffffffffffffffffffffffffffffffffffffffff16331461047c57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60026020526000908152604090205481565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b73ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604081205411905600a165627a7a7230582038bd3cc8ec8ee41f7b2e7468c55293150e27e0c7ab678e5e4162233dbef3f4360029`

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
const DataBin = `0x608060405234801561001057600080fd5b506040516020806103d8833981016040525160008054600160a060020a03909216600160a060020a0319909216919091179055610386806100526000396000f3006080604052600436106100565763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663295f36d7811461005b578063461b09c014610088578063516c4b84146100a8575b600080fd5b34801561006757600080fd5b506100766004356024356100e6565b60408051918252519081900360200190f35b34801561009457600080fd5b506100a66004356024356044356101bc565b005b3480156100b457600080fd5b506100bd61033e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60408051602080820185905281830184905282518083038401815260609092019283905281516000936001938593909282918401908083835b6020831061015c57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0909201916020918201910161011f565b51815160209384036101000a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01801990921691161790526040805192909401829003909120865285019590955292909201600020549695505050505050565b60008054604080517fe814861e000000000000000000000000000000000000000000000000000000008152336004820152905173ffffffffffffffffffffffffffffffffffffffff9092169263e814861e926024808401936020939083900390910190829087803b15801561023057600080fd5b505af1158015610244573d6000803e3d6000fd5b505050506040513d602081101561025a57600080fd5b5051151561026757600080fd5b60408051602080820186905281830185905282518083038401815260609092019283905281518493600193600093909282918401908083835b602083106102dd57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe090920191602091820191016102a0565b51815160209384036101000a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0180199092169116179052604080519290940182900390912086528501959095529290920160002093909355505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820831ad6bd7960b490d17b231075877e71ec8d74a792e3731e6ea67035ff8586900029`

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
const DelegationBin = `0x608060405234801561001057600080fd5b50604051602080611213833981016040525160008054600160a060020a031916600160a060020a03831617905561004e640100000000610054810204565b5061012c565b6100676001640100000000610093810204565b60018054600160a060020a031916600160a060020a03928316179081905516151561009157600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b1580156100fa57600080fd5b505af115801561010e573d6000803e3d6000fd5b505050506040513d602081101561012457600080fd5b505192915050565b6110d88061013b6000396000f3006080604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305bafaa481146100875780633363375c146100ff578063516c4b841461012d57806373d4a13a1461016b578063e218e6d214610180578063ed2a2d6414610219578063fb55a05514610259575b600080fd5b34801561009357600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100fd9583359536956044949193909101919081908401838280828437509497505050923573ffffffffffffffffffffffffffffffffffffffff16935061028792505050565b005b34801561010b57600080fd5b506100fd73ffffffffffffffffffffffffffffffffffffffff6004351661045d565b34801561013957600080fd5b5061014261058c565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561017757600080fd5b506101426105a8565b34801561018c57600080fd5b50604080516020600460843581810135601f81018490048402850184019095528484526100fd948235946024803573ffffffffffffffffffffffffffffffffffffffff169560443595606435953695919460a49490939101919081908401838280828437509497505050923573ffffffffffffffffffffffffffffffffffffffff1693506105c492505050565b34801561022557600080fd5b5061024773ffffffffffffffffffffffffffffffffffffffff60043516610833565b60408051918252519081900360200190f35b34801561026557600080fd5b506100fd73ffffffffffffffffffffffffffffffffffffffff60043516610844565b6000610291611050565b8261029b81610973565b156102a557600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156102c757600080fd5b600092505b85831015610455576102de8584610989565b91506102ed8260200151610bc8565b6102fa8260400151610c1c565b8151602083015161030a90610c56565b1061031457600080fd5b61032682602001518360000151610d07565b61033e82602001518360800151846060015101610d9d565b61035082604001518360600151610dc3565b816040015173ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84606001516040518082815260200191505060405180910390a360008260800151111561044a576103dc848360800151610dc3565b8373ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84608001516040518082815260200191505060405180910390a35b6001909201916102cc565b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156104f957600080fd5b505af115801561050d573d6000803e3d6000fd5b505050506040513d602081101561052357600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461054557600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b6000856105d081610973565b156105da57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156105fc57600080fd5b8261060681610973565b1561061057600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561063257600080fd5b6040805160208082018c90526c0100000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8c160282840152605482018a905260748083018a905283518084039091018152609490920192839052815161070b93918291908401908083835b602083106106d857805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0909201916020918201910161069b565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902086610deb565b925061071683610bc8565b8861072084610c56565b1061072a57600080fd5b610734838a610d07565b61074083878901610d9d565b61074a8888610dc3565b8773ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef896040518082815260200191505060405180910390a36000861115610828576107c28487610dc3565b8373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef886040518082815260200191505060405180910390a35b505050505050505050565b600061083e82610c56565b92915050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156108e057600080fd5b505af11580156108f4573d6000803e3d6000fd5b505050506040513d602081101561090a57600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461092c57600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600060028061098184610ece565b161492915050565b610991611050565b60008060008060008060008060006109a7611050565b8b60b5029950898d019c5060208d0151985060348d0151975060548d0151965060748d0151955060948d0151945060b48d0151935060ff60b58e0151169250601b8360ff1610156109f957601b830192505b6040805160208082018c90526c0100000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8c160282840152605482018a905260748083018a905283518084039091018152609490920192839052815191929182918401908083835b60208310610a9c57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101610a5f565b5181517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209485036101000a0190811690199190911617905260408051949092018490039093208e87528151600080825281860180855283905260ff8b1682850152606082018d9052608082018c905292519198506001965060a080820196507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083019450908290030191865af1158015610b5c573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015173ffffffffffffffffffffffffffffffffffffffff90811660208501528a1690830152506060810187905260808101869052995050505050505050505092915050565b806000610bd482610ece565b9050600180821614610be557600080fd5b60048181161415610bf557600080fd5b73ffffffffffffffffffffffffffffffffffffffff82161515610c1757600080fd5b505050565b80610c2681610973565b15610c3057600080fd5b73ffffffffffffffffffffffffffffffffffffffff81161515610c5257600080fd5b5050565b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526003600482015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b158015610cd557600080fd5b505af1158015610ce9573d6000803e3d6000fd5b505050506040513d6020811015610cff57600080fd5b505192915050565b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526003600482015273ffffffffffffffffffffffffffffffffffffffff8581166024830152604482018590529151919092169163461b09c091606480830192600092919082900301818387803b158015610d8957600080fd5b505af1158015610455573d6000803e3d6000fd5b6000610da883610f4f565b905081811015610db757600080fd5b610c1783838303610fce565b6000610dce83610f4f565b9050818101811115610ddf57600080fd5b610c1783838301610fce565b60008060008084516041141515610e0157600080fd5b50505060208201516040830151604184015160ff16601b811015610e2357601b015b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a08085019491937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0840193928390039091019190865af1158015610e9b573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00151979650505050505050565b60018054604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff84811660248501529051600093919092169163295f36d79160448082019260209290919082900301818787803b158015610cd557600080fd5b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b158015610cd557600080fd5b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff8581166024830152604482018590529151919092169163461b09c091606480830192600092919082900301818387803b158015610d8957600080fd5b60a06040519081016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815250905600a165627a7a723058208df1c8c23e4a032647a53cb610af0d9c7ba3e2433e6499858f341d399d6efca90029`

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
const EnforcementBin = `0x608060405234801561001057600080fd5b50604051608080610dce833981016040908152815160208301519183015160609093015160008054600160a060020a031916600160a060020a0384161790559092906100636401000000006100a6810204565b60028054600160a060020a03948516600160a060020a0319918216179091556003805493851693821693909317909255600480549190931691161790555061017e565b6100b960016401000000006100e5810204565b60018054600160a060020a031916600160a060020a0392831617908190551615156100e357600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b15801561014c57600080fd5b505af1158015610160573d6000803e3d6000fd5b505050506040513d602081101561017657600080fd5b505192915050565b610c418061018d6000396000f3006080604052600436106100c45763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633363375c81146100c9578063516c4b84146100f95780635dab24201461013757806372cfc9dc1461014c57806373d4a13a14610161578063788649ea1461017657806385a0f282146101a457806390f28e74146101b9578063b10725e8146101e7578063b9b0330f14610215578063f26c159f14610243578063f3fef3a314610271578063fb55a055146102a2575b600080fd5b3480156100d557600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff600435166102d0565b005b34801561010557600080fd5b5061010e6103ff565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561014357600080fd5b5061010e61041b565b34801561015857600080fd5b5061010e610437565b34801561016d57600080fd5b5061010e610453565b34801561018257600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff6004351661046f565b3480156101b057600080fd5b5061010e610517565b3480156101c557600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff60043516610533565b3480156101f357600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff6004351661059e565b34801561022157600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff60043516610609565b34801561024f57600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff600435166106ab565b34801561027d57600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff60043516602435610734565b3480156102ae57600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff6004351661082b565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561036c57600080fd5b505af1158015610380573d6000803e3d6000fd5b505050506040513d602081101561039657600080fd5b505173ffffffffffffffffffffffffffffffffffffffff16146103b857600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60025473ffffffffffffffffffffffffffffffffffffffff16331461049357600080fd5b6104c7817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb6104c18261095a565b16610a0d565b6040805160008152905173ffffffffffffffffffffffffffffffffffffffff8316917fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee919081900360200190a250565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff16331461055757600080fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60025473ffffffffffffffffffffffffffffffffffffffff1633146105c257600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60035473ffffffffffffffffffffffffffffffffffffffff16331461062d57600080fd5b8061063781610aab565b1561064157600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561066357600080fd5b50600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60025473ffffffffffffffffffffffffffffffffffffffff1633146106cf57600080fd5b6106e48160046106de8461095a565b17610a0d565b6040805160018152905173ffffffffffffffffffffffffffffffffffffffff8316917fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee919081900360200190a250565b60025473ffffffffffffffffffffffffffffffffffffffff16331461075857600080fd5b60045473ffffffffffffffffffffffffffffffffffffffff1661077a81610aab565b1561078457600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156107a657600080fd5b6107b08383610ac1565b6004546107d39073ffffffffffffffffffffffffffffffffffffffff1683610aec565b60045460408051848152905173ffffffffffffffffffffffffffffffffffffffff928316928616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a3505050565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156108c757600080fd5b505af11580156108db573d6000803e3d6000fd5b505050506040513d60208110156108f157600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461091357600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60018054604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff84811660248501529051600093919092169163295f36d79160448082019260209290919082900301818787803b1580156109db57600080fd5b505af11580156109ef573d6000803e3d6000fd5b505050506040513d6020811015610a0557600080fd5b505192915050565b60018054604080517f461b09c0000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff858116602485015260448401859052905191169163461b09c091606480830192600092919082900301818387803b158015610a8f57600080fd5b505af1158015610aa3573d6000803e3d6000fd5b505050505050565b6000600280610ab98461095a565b161492915050565b6000610acc83610b14565b905081811015610adb57600080fd5b610ae783838303610b93565b505050565b6000610af783610b14565b9050818101811115610b0857600080fd5b610ae783838301610b93565b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b1580156109db57600080fd5b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff8581166024830152604482018590529151919092169163461b09c091606480830192600092919082900301818387803b158015610a8f57600080fd00a165627a7a72305820a17e0428ce3a8ae44793fc733a66c6af15d663b98a45b15cb96119b2dcac6fba0029`

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
const ReserveBin = `0x608060405234801561001057600080fd5b50604051604080610ccc83398101604052805160209091015160008054600160a060020a031916600160a060020a03841617905561005564010000000061007b810204565b60028054600160a060020a031916600160a060020a039290921691909117905550610153565b61008e60016401000000006100ba810204565b60018054600160a060020a031916600160a060020a0392831617908190551615156100b857600080fd5b565b60008054604080517f13c01368000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916313c013689160248082019260209290919082900301818787803b15801561012157600080fd5b505af1158015610135573d6000803e3d6000fd5b505050506040513d602081101561014b57600080fd5b505192915050565b610b6a806101626000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166302946804811461009d57806318160ddd146100db5780633363375c14610102578063516c4b841461013257806373d4a13a1461014757806398e52f9a1461015c578063b921e16314610174578063ddf05f591461018c578063fb55a055146101ba575b600080fd5b3480156100a957600080fd5b506100b26101e8565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156100e757600080fd5b506100f0610204565b60408051918252519081900360200190f35b34801561010e57600080fd5b5061013073ffffffffffffffffffffffffffffffffffffffff60043516610213565b005b34801561013e57600080fd5b506100b2610342565b34801561015357600080fd5b506100b261035e565b34801561016857600080fd5b5061013060043561037a565b34801561018057600080fd5b506101306004356104da565b34801561019857600080fd5b5061013073ffffffffffffffffffffffffffffffffffffffff60043516610623565b3480156101c657600080fd5b5061013073ffffffffffffffffffffffffffffffffffffffff6004351661068e565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b600061020e6107bd565b905090565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156102af57600080fd5b505af11580156102c3573d6000803e3d6000fd5b505050506040513d60208110156102d957600080fd5b505173ffffffffffffffffffffffffffffffffffffffff16146102fb57600080fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60025460009073ffffffffffffffffffffffffffffffffffffffff1633146103a157600080fd5b60025473ffffffffffffffffffffffffffffffffffffffff1660006103c58261086a565b90506001808216146103d657600080fd5b600481811614156103e657600080fd5b73ffffffffffffffffffffffffffffffffffffffff8216151561040857600080fd5b6104106107bd565b92508383101561041f57600080fd5b838303925061042d8361091d565b6002546104509073ffffffffffffffffffffffffffffffffffffffff16856109b8565b60025460408051868152905160009273ffffffffffffffffffffffffffffffffffffffff16917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a36040805184815290517ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e99181900360200190a150505050565b60025460009073ffffffffffffffffffffffffffffffffffffffff16331461050157600080fd5b60025473ffffffffffffffffffffffffffffffffffffffff16610523816109e3565b1561052d57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561054f57600080fd5b610557610204565b915082820182111561056857600080fd5b908201906105758261091d565b6002546105989073ffffffffffffffffffffffffffffffffffffffff16846109f9565b60025460408051858152905173ffffffffffffffffffffffffffffffffffffffff909216916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a36040805183815290517ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e99181900360200190a1505050565b60025473ffffffffffffffffffffffffffffffffffffffff16331461064757600080fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b3373ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639afd453c6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561072a57600080fd5b505af115801561073e573d6000803e3d6000fd5b505050506040513d602081101561075457600080fd5b505173ffffffffffffffffffffffffffffffffffffffff161461077657600080fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600154604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600560048201526000602482018190529151919273ffffffffffffffffffffffffffffffffffffffff169163295f36d79160448082019260209290919082900301818787803b15801561083957600080fd5b505af115801561084d573d6000803e3d6000fd5b505050506040513d602081101561086357600080fd5b5051905090565b60018054604080517f295f36d7000000000000000000000000000000000000000000000000000000008152600481019390935273ffffffffffffffffffffffffffffffffffffffff84811660248501529051600093919092169163295f36d79160448082019260209290919082900301818787803b1580156108eb57600080fd5b505af11580156108ff573d6000803e3d6000fd5b505050506040513d602081101561091557600080fd5b505192915050565b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526005600482015260006024820181905260448201859052915173ffffffffffffffffffffffffffffffffffffffff9093169263461b09c09260648084019391929182900301818387803b15801561099d57600080fd5b505af11580156109b1573d6000803e3d6000fd5b5050505050565b60006109c383610a21565b9050818110156109d257600080fd5b6109de83838303610aa0565b505050565b60006002806109f18461086a565b161492915050565b6000610a0483610a21565b9050818101811115610a1557600080fd5b6109de83838301610aa0565b600154604080517f295f36d70000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff84811660248301529151600093929092169163295f36d79160448082019260209290919082900301818787803b1580156108eb57600080fd5b600154604080517f461b09c00000000000000000000000000000000000000000000000000000000081526002600482015273ffffffffffffffffffffffffffffffffffffffff8581166024830152604482018590529151919092169163461b09c091606480830192600092919082900301818387803b158015610b2257600080fd5b505af1158015610b36573d6000803e3d6000fd5b5050505050505600a165627a7a72305820c31c813b4344ce6b585169a838e95f32e023d0b734a4d6f9c10daaba58d10e6b0029`

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
