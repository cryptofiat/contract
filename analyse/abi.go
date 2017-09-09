// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
const AccountRecoveryBin = `0x6060604052341561000f57600080fd5b6040516020806107d6833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b610783806100536000396000f300606060405263ffffffff60e060020a6000350416632d1c5ff88114610045578063516c4b841461006c578063f1375f381461009b578063fb55a055146100bc575b600080fd5b341561005057600080fd5b61006a600160a060020a03600435811690602435166100dd565b005b341561007757600080fd5b61007f610228565b604051600160a060020a03909116815260200160405180910390f35b34156100a657600080fd5b61006a600160a060020a0360043516610237565b005b34156100c757600080fd5b61006a600160a060020a0360043516610245565b005b6000826100e9816102f0565b15156100f457600080fd5b6100fd81610308565b1561010757600080fd5b600160a060020a038116151561011c57600080fd5b8261012681610320565b1561013057600080fd5b600160a060020a038116151561014557600080fd5b61014e85610338565b600160a060020a031633600160a060020a031614151561016d57600080fd5b61018285600261017c886103c5565b17610452565b84600160a060020a03167fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de60405160405180910390a26101c1856104d7565b92506101cd8584610564565b6101d78484610590565b83600160a060020a031685600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405190815260200160405180910390a35b5b505b50505050565b600054600160a060020a031681565b61024133826105bf565b5b50565b60008054600160a060020a0333811692911690639afd453c90604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561029357600080fd5b6102c65a03f115156102a457600080fd5b50505060405180519050600160a060020a03161415156102c357600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60006001806102fe846103c5565b161490505b919050565b60006004806102fe846103c5565b161490505b919050565b60006002806102fe846103c5565b161490505b919050565b600061034261064d565b600160a060020a031663295f36d7600484600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156103a357600080fd5b6102c65a03f115156103b457600080fd5b50505060405180519150505b919050565b60006103cf61064d565b600160a060020a031663295f36d7600184600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156103a357600080fd5b6102c65a03f115156103b457600080fd5b50505060405180519150505b919050565b61045a61064d565b600160a060020a031663461b09c0600184600160a060020a03166001028460010260405160e060020a63ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156104be57600080fd5b6102c65a03f1151561021f57600080fd5b5050505b5050565b60006104e161064d565b600160a060020a031663295f36d7600284600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156103a357600080fd5b6102c65a03f115156103b457600080fd5b50505060405180519150505b919050565b600061056f836104d7565b90508181101561057e57600080fd5b61058a8383830361065f565b5b505050565b600061059b836104d7565b9050818101819010156105ad57600080fd5b61058a8383830161065f565b5b505050565b6105c761064d565b600160a060020a031663461b09c0600484600160a060020a031660010284600160a060020a031660010260405160e060020a63ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156104be57600080fd5b6102c65a03f1151561021f57600080fd5b5050505b5050565b600061065960016106e4565b90505b90565b61066761064d565b600160a060020a031663461b09c0600284600160a060020a03166001028460010260405160e060020a63ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156104be57600080fd5b6102c65a03f1151561021f57600080fd5b5050505b5050565b60008054600160a060020a03166313c0136883836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156103a357600080fd5b6102c65a03f115156103b457600080fd5b50505060405180519150505b9190505600a165627a7a72305820cc89cb449f415a8d95051593c71504aeadc3cdd3b314a34df6651d4b8eb826040029`

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
const AccountsBin = `0x6060604052341561000f57600080fd5b604051602080610715833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b6106c2806100536000396000f300606060405236156100725763ffffffff60e060020a600035041663516c4b848114610077578063673448dd146100a65780636943b017146100d957806370a082311461010c57806397a5d5b51461013d578063a9059cbb1461016e578063e583983614610192578063fb55a055146101c5575b600080fd5b341561008257600080fd5b61008a6101e6565b604051600160a060020a03909116815260200160405180910390f35b34156100b157600080fd5b6100c5600160a060020a03600435166101f5565b604051901515815260200160405180910390f35b34156100e457600080fd5b6100c5600160a060020a0360043516610208565b604051901515815260200160405180910390f35b341561011757600080fd5b61012b600160a060020a036004351661021b565b60405190815260200160405180910390f35b341561014857600080fd5b61012b600160a060020a036004351661022e565b60405190815260200160405180910390f35b341561017957600080fd5b610190600160a060020a0360043516602435610241565b005b341561019d57600080fd5b6100c5600160a060020a0360043516610311565b604051901515815260200160405180910390f35b34156101d057600080fd5b610190600160a060020a0360043516610324565b005b600054600160a060020a031681565b6000610200826103cf565b90505b919050565b6000610200826103e7565b90505b919050565b6000610200826103ff565b90505b919050565b60006102008261048c565b90505b919050565b60003361024d816103cf565b151561025857600080fd5b61026181610519565b1561026b57600080fd5b600160a060020a038116151561028057600080fd5b8361028a816103e7565b1561029457600080fd5b600160a060020a03811615156102a957600080fd5b3392506102b68385610531565b6102c0858561055d565b84600160a060020a031683600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8660405190815260200160405180910390a35b5b505b50505050565b600061020082610519565b90505b919050565b60008054600160a060020a0333811692911690639afd453c90604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561037257600080fd5b6102c65a03f1151561038357600080fd5b50505060405180519050600160a060020a03161415156103a257600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60006001806103dd8461048c565b161490505b919050565b60006002806103dd8461048c565b161490505b919050565b600061040961058c565b600160a060020a031663295f36d7600284600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561046a57600080fd5b6102c65a03f1151561047b57600080fd5b50505060405180519150505b919050565b600061049661058c565b600160a060020a031663295f36d7600184600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561046a57600080fd5b6102c65a03f1151561047b57600080fd5b50505060405180519150505b919050565b60006004806103dd8461048c565b161490505b919050565b600061053c836103ff565b90508181101561054b57600080fd5b6105578383830361059e565b5b505050565b6000610568836103ff565b90508181018190101561057a57600080fd5b6105578383830161059e565b5b505050565b60006105986001610623565b90505b90565b6105a661058c565b600160a060020a031663461b09c0600284600160a060020a03166001028460010260405160e060020a63ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b151561060a57600080fd5b6102c65a03f1151561030857600080fd5b5050505b5050565b60008054600160a060020a03166313c0136883836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561046a57600080fd5b6102c65a03f1151561047b57600080fd5b50505060405180519150505b9190505600a165627a7a72305820878c01a552cab39f8cdafdc837522cc2dcc232954a6bd9f16b6dfe5e0eab6e990029`

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
const ApprovingBin = `0x6060604052341561000f57600080fd5b6040516040806105f183398101604052808051919060200180519150505b60008054600160a060020a03808516600160a060020a03199283161790925560018054928416929091169190911790555b50505b610581806100706000396000f300606060405236156100675763ffffffff60e060020a600035041663071a8b53811461006c57806307a385e6146100bd578063516c4b84146100ec578063c8b091091461011b578063dd336b941461013c578063f89f4e771461015d578063fb55a0551461017e575b600080fd5b341561007757600080fd5b6100bb600460248135818101908301358060208181020160405190810160405280939291908181526020018383602002808284375094965061019f95505050505050565b005b34156100c857600080fd5b6100d06101d7565b604051600160a060020a03909116815260200160405180910390f35b34156100f757600080fd5b6100d06101e6565b604051600160a060020a03909116815260200160405180910390f35b341561012657600080fd5b6100bb600160a060020a03600435166101f5565b005b341561014757600080fd5b6100bb600160a060020a036004351661023d565b005b341561016857600080fd5b6100bb600160a060020a03600435166102a8565b005b341561018957600080fd5b6100bb600160a060020a0360043516610313565b005b60005b81518110156101d2576101c98282815181106101ba57fe5b906020019060200201516102a8565b5b6001016101a2565b5b5050565b600154600160a060020a031681565b600054600160a060020a031681565b60015433600160a060020a0390811691161461021057600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60015433600160a060020a0390811691161461025857600080fd5b61026d816002610267846103be565b1761044b565b80600160a060020a03167fa29911196d428d7968f8bde7515181a391bfa16e26042f789f3f2da7665e25de60405160405180910390a25b5b50565b60015433600160a060020a039081169116146102c357600080fd5b6102d8816001610267846103be565b1761044b565b80600160a060020a03167fbc7abdf8533487db28f8c616affbb4e122d90c5ab8deb258fd21b09cee59573060405160405180910390a25b5b50565b60008054600160a060020a0333811692911690639afd453c90604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561036157600080fd5b6102c65a03f1151561037257600080fd5b50505060405180519050600160a060020a031614151561039157600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60006103c86104d0565b600160a060020a031663295f36d7600184600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561042957600080fd5b6102c65a03f1151561043a57600080fd5b50505060405180519150505b919050565b6104536104d0565b600160a060020a031663461b09c0600184600160a060020a03166001028460010260405160e060020a63ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156104b757600080fd5b6102c65a03f115156104c857600080fd5b5050505b5050565b60006104dc60016104e2565b90505b90565b60008054600160a060020a03166313c0136883836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561042957600080fd5b6102c65a03f1151561043a57600080fd5b50505060405180519150505b9190505600a165627a7a72305820c83b17966f6cac2cb976f84addf292e188e2dd971b85860d460a21b1a48c8f760029`

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

// ConstantsABI is the input ABI used to generate the binding from.
const ConstantsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// ConstantsBin is the compiled bytecode used for deploying new contracts.
const ConstantsBin = `0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a723058200ba86195aa19a4eb72c070c114d70fe334d1dca580074de0615c8c85abe0352f0029`

// DeployConstants deploys a new Ethereum contract, binding an instance of Constants to it.
func DeployConstants(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Constants, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstantsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConstantsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Constants{ConstantsCaller: ConstantsCaller{contract: contract}, ConstantsTransactor: ConstantsTransactor{contract: contract}}, nil
}

// Constants is an auto generated Go binding around an Ethereum contract.
type Constants struct {
	ConstantsCaller     // Read-only binding to the contract
	ConstantsTransactor // Write-only binding to the contract
}

// ConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstantsSession struct {
	Contract     *Constants        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstantsCallerSession struct {
	Contract *ConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstantsTransactorSession struct {
	Contract     *ConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstantsRaw struct {
	Contract *Constants // Generic contract binding to access the raw methods on
}

// ConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstantsCallerRaw struct {
	Contract *ConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// ConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstantsTransactorRaw struct {
	Contract *ConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConstants creates a new instance of Constants, bound to a specific deployed contract.
func NewConstants(address common.Address, backend bind.ContractBackend) (*Constants, error) {
	contract, err := bindConstants(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Constants{ConstantsCaller: ConstantsCaller{contract: contract}, ConstantsTransactor: ConstantsTransactor{contract: contract}}, nil
}

// NewConstantsCaller creates a new read-only instance of Constants, bound to a specific deployed contract.
func NewConstantsCaller(address common.Address, caller bind.ContractCaller) (*ConstantsCaller, error) {
	contract, err := bindConstants(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantsCaller{contract: contract}, nil
}

// NewConstantsTransactor creates a new write-only instance of Constants, bound to a specific deployed contract.
func NewConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstantsTransactor, error) {
	contract, err := bindConstants(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ConstantsTransactor{contract: contract}, nil
}

// bindConstants binds a generic wrapper to an already deployed contract.
func bindConstants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstantsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constants *ConstantsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Constants.Contract.ConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constants *ConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constants.Contract.ConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constants *ConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constants.Contract.ConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Constants *ConstantsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Constants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Constants *ConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Constants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Constants *ConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Constants.Contract.contract.Transact(opts, method, params...)
}

// CryptoFiatABI is the input ABI used to generate the binding from.
const CryptoFiatABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointMasterAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"masterAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"contractActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"next\",\"type\":\"address\"}],\"name\":\"ContractUpgraded\",\"type\":\"event\"}]"

// CryptoFiatBin is the compiled bytecode used for deploying new contracts.
const CryptoFiatBin = `0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a0316179055600380546001810161003d8382610070565b916000526020600020900160005b8154600160a060020a033081166101009390930a92830292021916179055505b6100bb565b8154818355818115116100945760008381526020902061009491810190830161009a565b5b505050565b6100b891905b808211156100b457600081556001016100a0565b5090565b90565b6104c5806100ca6000396000f3006060604052361561008b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663028f4e47811461009057806313c01368146100b45780633fad74ad146100e6578063474da79a1461010b5780635db4380d1461013d578063874c34731461015e5780639afd453c1461018f578063e814861e146101be575b600080fd5b341561009b57600080fd5b6100b2600435600160a060020a03602435166101f1565b005b34156100bf57600080fd5b6100ca600435610371565b604051600160a060020a03909116815260200160405180910390f35b34156100f157600080fd5b6100f961038c565b60405190815260200160405180910390f35b341561011657600080fd5b6100ca600435610393565b604051600160a060020a03909116815260200160405180910390f35b341561014857600080fd5b6100b2600160a060020a03600435166103c5565b005b341561016957600080fd5b6100f9600160a060020a036004351661040d565b60405190815260200160405180910390f35b341561019a57600080fd5b6100ca61041f565b604051600160a060020a03909116815260200160405180910390f35b34156101c957600080fd5b6101dd600160a060020a036004351661042e565b604051901515815260200160405180910390f35b60008083151561020057600080fd5b600084815260016020526040902054600160a060020a039081169250831682141561022a57600080fd5b60005433600160a060020a0390811691161480610258575081600160a060020a031633600160a060020a0316145b905080151561026657600080fd5b61026f8361042e565b1561027957600080fd5b600160a060020a03828116600090815260026020908152604080832083905587835260019091529020805473ffffffffffffffffffffffffffffffffffffffff19169185169182179055156102e457600160a060020a03831660009081526002602052604090208490555b837fdc69b57038334451ee12fd1742228917cea7f40dbd33cda5162e7e5754acee1c8385604051600160a060020a039283168152911660208201526040908101905180910390a2600380546001810161033d838261044e565b916000526020600020900160005b8154600160a060020a038088166101009390930a92830292021916179055505b50505050565b600160205260009081526040902054600160a060020a031681565b6003545b90565b60038054829081106103a157fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b60005433600160a060020a039081169116146103e057600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60026020526000908152604090205481565b600054600160a060020a031681565b600160a060020a038116600090815260026020526040812054115b919050565b81548183558181151161047257600083815260209020610472918101908301610478565b5b505050565b61039091905b80821115610492576000815560010161047e565b5090565b905600a165627a7a723058202c9892aae452336a4531d49810434fd17d07a4444fe00a70b404023dfa58d8910029`

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
const DataBin = `0x6060604052341561000f57600080fd5b604051602080610309833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b6102b6806100536000396000f300606060405263ffffffff60e060020a600035041663295f36d78114610045578063461b09c014610070578063516c4b841461008e578063fb55a055146100bd575b600080fd5b341561005057600080fd5b61005e6004356024356100de565b60405190815260200160405180910390f35b341561007b57600080fd5b61008c600435602435604435610118565b005b341561009957600080fd5b6100a16101d0565b604051600160a060020a03909116815260200160405180910390f35b34156100c857600080fd5b61008c600160a060020a03600435166101df565b005b6000600160008484604051918252602082015260409081019051908190039020815260208101919091526040016000205490505b92915050565b60008054600160a060020a03169063e814861e9033906040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561017357600080fd5b6102c65a03f1151561018457600080fd5b50505060405180519050151561019957600080fd5b8060016000858560405191825260208201526040908101905190819003902081526020810191909152604001600020555b5b505050565b600054600160a060020a031681565b60008054600160a060020a0333811692911690639afd453c90604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561022d57600080fd5b6102c65a03f1151561023e57600080fd5b50505060405180519050600160a060020a031614151561025d57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b505600a165627a7a72305820caef31e9447b23ddb42e9e9d0c03a2cbccf9ad3c79fd3f61cdc67f2277c070e80029`

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
const DelegationBin = `0x6060604052341561000f57600080fd5b604051602080610c08833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b610bb5806100536000396000f300606060405263ffffffff60e060020a60003504166305bafaa48114610050578063516c4b84146100b3578063e218e6d2146100e2578063ed2a2d641461015c578063fb55a0551461018d575b600080fd5b341561005b57600080fd5b6100b1600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050509235600160a060020a031692506101ae915050565b005b34156100be57600080fd5b6100c66102de565b604051600160a060020a03909116815260200160405180910390f35b34156100ed57600080fd5b6100b1600480359060248035600160a060020a03169160443591606435919060a49060843590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050509235600160a060020a031692506102ed915050565b005b341561016757600080fd5b61017b600160a060020a0360043516610454565b60405190815260200160405180910390f35b341561019857600080fd5b6100b1600160a060020a0360043516610467565b005b60006101b8610b3b565b600091505b848210156102d6576101cf8483610512565b90506101de8160200151610669565b80516101ed82602001516106ac565b106101f757600080fd5b61020681602001518251610739565b61021e816020015182608001518360600151016107be565b610230816040015182606001516107ea565b8060400151600160a060020a03168160200151600160a060020a0316600080516020610b6a833981519152836060015160405190815260200160405180910390a36000816080015111156102ca5761028c8382608001516107ea565b82600160a060020a03168160200151600160a060020a0316600080516020610b6a833981519152836080015160405190815260200160405180910390a35b5b6001909101906101bd565b5b5050505050565b600054600160a060020a031681565b6000856102f981610819565b1561030357600080fd5b600160a060020a038116151561031857600080fd5b8261032281610819565b1561032c57600080fd5b600160a060020a038116151561034157600080fd5b61038989898989604051938452600160a060020a03929092166c0100000000000000000000000002602084015260348301526054820152607401604051809103902086610831565b925061039483610669565b8861039e846106ac565b106103a857600080fd5b6103b2838a610739565b6103be838789016107be565b6103c888886107ea565b87600160a060020a031683600160a060020a0316600080516020610b6a8339815191528960405190815260200160405180910390a360008611156104465761041084876107ea565b83600160a060020a031683600160a060020a0316600080516020610b6a8339815191528860405190815260200160405180910390a35b5b5b505b5050505050505050565b600061045f826106ac565b90505b919050565b60008054600160a060020a0333811692911690639afd453c90604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156104b557600080fd5b6102c65a03f115156104c657600080fd5b50505060405180519050600160a060020a03161415156104e557600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b61051a610b3b565b6000806000806000806000806000610530610b3b565b8b60c102995060208d0151985060408d0151975060608d0151965060808d0151955060a08d0151945060c08d0151935060ff60c18e0151169250601b8360ff16101561057d57601b830192505b88888888604051938452600160a060020a03929092166c010000000000000000000000000260208401526034830152605482015260740160405190819003902089825291506001828487876040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f1151561062757600080fd5b505060206040510351600160a060020a0390811660208301528816604082015260608101879052608081018690529950895b5050505050505050505092915050565b80610673816108e7565b151561067e57600080fd5b610687816108ff565b1561069157600080fd5b600160a060020a03811615156106a657600080fd5b5b5b5050565b60006106b6610917565b600160a060020a031663295f36d7600384600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561071757600080fd5b6102c65a03f1151561072857600080fd5b50505060405180519150505b919050565b610741610917565b600160a060020a031663461b09c0600384600160a060020a03166001028460010260405160e060020a63ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156107a557600080fd5b6102c65a03f115156102d657600080fd5b5050505b5050565b60006107c983610929565b9050818110156107d857600080fd5b6107e4838383036109b6565b5b505050565b60006107f583610929565b90508181018190101561080757600080fd5b6107e4838383016109b6565b5b505050565b600060028061082784610a3b565b161490505b919050565b600080600080845160411461084557600080fd5b602085015192506040850151915060ff6041860151169050601b8160ff16101561086d57601b015b6001868285856040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f115156108d257600080fd5b50506020604051035193505b50505092915050565b600060018061082784610a3b565b161490505b919050565b600060048061082784610a3b565b161490505b919050565b60006109236001610ac8565b90505b90565b6000610933610917565b600160a060020a031663295f36d7600284600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561071757600080fd5b6102c65a03f1151561072857600080fd5b50505060405180519150505b919050565b6109be610917565b600160a060020a031663461b09c0600284600160a060020a03166001028460010260405160e060020a63ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156107a557600080fd5b6102c65a03f115156102d657600080fd5b5050505b5050565b6000610a45610917565b600160a060020a031663295f36d7600184600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561071757600080fd5b6102c65a03f1151561072857600080fd5b50505060405180519150505b919050565b60008054600160a060020a03166313c0136883836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561071757600080fd5b6102c65a03f1151561072857600080fd5b50505060405180519150505b919050565b60a06040519081016040908152600080835260208301819052908201819052606082018190526080820152905600ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa165627a7a72305820d304f9704f0ef077ecbb5b0244d02eeece9d1672c686cdb99ab106780f5b9fc10029`

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
const EnforcementBin = `0x6060604052341561000f57600080fd5b6040516080806109b0833981016040528080519190602001805191906020018051919060200180519150505b60008054600160a060020a03808716600160a060020a031992831617909255600180548684169083161790556002805485841690831617905560038054928416929091169190911790555b505050505b6109168061009a6000396000f300606060405236156100935763ffffffff60e060020a600035041663516c4b8481146100985780635dab2420146100c757806372cfc9dc146100f6578063788649ea1461012557806385a0f2821461014657806390f28e7414610175578063b10725e814610196578063b9b0330f146101b7578063f26c159f146101d8578063f3fef3a3146101f9578063fb55a0551461021d575b600080fd5b34156100a357600080fd5b6100ab61023e565b604051600160a060020a03909116815260200160405180910390f35b34156100d257600080fd5b6100ab61024d565b604051600160a060020a03909116815260200160405180910390f35b341561010157600080fd5b6100ab61025c565b604051600160a060020a03909116815260200160405180910390f35b341561013057600080fd5b610144600160a060020a036004351661026b565b005b341561015157600080fd5b6100ab6102e1565b604051600160a060020a03909116815260200160405180910390f35b341561018057600080fd5b610144600160a060020a03600435166102f0565b005b34156101a157600080fd5b610144600160a060020a0360043516610338565b005b34156101c257600080fd5b610144600160a060020a0360043516610380565b005b34156101e357600080fd5b610144600160a060020a03600435166103f3565b005b341561020457600080fd5b610144600160a060020a0360043516602435610468565b005b341561022857600080fd5b610144600160a060020a0360043516610523565b005b600054600160a060020a031681565b600354600160a060020a031681565b600154600160a060020a031681565b60015433600160a060020a0390811691161461028657600080fd5b61029c81600419610296826105ce565b1661065b565b80600160a060020a03167fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee6000604051901515815260200160405180910390a25b5b50565b600254600160a060020a031681565b60025433600160a060020a0390811691161461030b57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60015433600160a060020a0390811691161461035357600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60025433600160a060020a0390811691161461039b57600080fd5b806103a5816106e0565b156103af57600080fd5b600160a060020a03811615156103c457600080fd5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790555b5b505b50565b60015433600160a060020a0390811691161461040e57600080fd5b61042381600461041d846105ce565b1761065b565b80600160a060020a03167fc0a52010de04a4a5a920bfbaa006102b1014b44a1e1f7315f03903cbcf5318ee6001604051901515815260200160405180910390a25b5b50565b60015433600160a060020a0390811691161461048357600080fd5b600354600160a060020a0316610498816106e0565b156104a257600080fd5b600160a060020a03811615156104b757600080fd5b6104c183836106f8565b6003546104d790600160a060020a031683610724565b600354600160a060020a039081169084167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35b5b505b5050565b60008054600160a060020a0333811692911690639afd453c90604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561057157600080fd5b6102c65a03f1151561058257600080fd5b50505060405180519050600160a060020a03161415156105a157600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60006105d8610753565b600160a060020a031663295f36d7600184600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561063957600080fd5b6102c65a03f1151561064a57600080fd5b50505060405180519150505b919050565b610663610753565b600160a060020a031663461b09c0600184600160a060020a03166001028460010260405160e060020a63ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156106c757600080fd5b6102c65a03f115156106d857600080fd5b5050505b5050565b60006002806106ee846105ce565b161490505b919050565b600061070383610765565b90508181101561071257600080fd5b61051c838383036107f2565b5b505050565b600061072f83610765565b90508181018190101561074157600080fd5b61051c838383016107f2565b5b505050565b600061075f6001610877565b90505b90565b600061076f610753565b600160a060020a031663295f36d7600284600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561063957600080fd5b6102c65a03f1151561064a57600080fd5b50505060405180519150505b919050565b6107fa610753565b600160a060020a031663461b09c0600284600160a060020a03166001028460010260405160e060020a63ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156106c757600080fd5b6102c65a03f115156106d857600080fd5b5050505b5050565b60008054600160a060020a03166313c0136883836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561063957600080fd5b6102c65a03f1151561064a57600080fd5b50505060405180519150505b9190505600a165627a7a72305820a562deac332826e8ceb76184859df381d263b0e3678b2d00bc2397e2eb9833c10029`

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

// InternalDataABI is the input ABI used to generate the binding from.
const InternalDataABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// InternalDataBin is the compiled bytecode used for deploying new contracts.
const InternalDataBin = `0x6060604052341561000f57600080fd5b5b6101978061001f6000396000f300606060405263ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663516c4b848114610048578063fb55a05514610077575b600080fd5b341561005357600080fd5b61005b610098565b604051600160a060020a03909116815260200160405180910390f35b341561008257600080fd5b610096600160a060020a03600435166100a7565b005b600054600160a060020a031681565b60008054600160a060020a0333811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561010e57600080fd5b6102c65a03f1151561011f57600080fd5b50505060405180519050600160a060020a031614151561013e57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b505600a165627a7a72305820b133303da7ad0962d57e0faed7dd1cee9f2b9538d7fecfe0e81a897906759a970029`

// DeployInternalData deploys a new Ethereum contract, binding an instance of InternalData to it.
func DeployInternalData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InternalData, error) {
	parsed, err := abi.JSON(strings.NewReader(InternalDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InternalDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InternalData{InternalDataCaller: InternalDataCaller{contract: contract}, InternalDataTransactor: InternalDataTransactor{contract: contract}}, nil
}

// InternalData is an auto generated Go binding around an Ethereum contract.
type InternalData struct {
	InternalDataCaller     // Read-only binding to the contract
	InternalDataTransactor // Write-only binding to the contract
}

// InternalDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type InternalDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InternalDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InternalDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InternalDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InternalDataSession struct {
	Contract     *InternalData     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InternalDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InternalDataCallerSession struct {
	Contract *InternalDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// InternalDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InternalDataTransactorSession struct {
	Contract     *InternalDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InternalDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type InternalDataRaw struct {
	Contract *InternalData // Generic contract binding to access the raw methods on
}

// InternalDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InternalDataCallerRaw struct {
	Contract *InternalDataCaller // Generic read-only contract binding to access the raw methods on
}

// InternalDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InternalDataTransactorRaw struct {
	Contract *InternalDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInternalData creates a new instance of InternalData, bound to a specific deployed contract.
func NewInternalData(address common.Address, backend bind.ContractBackend) (*InternalData, error) {
	contract, err := bindInternalData(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InternalData{InternalDataCaller: InternalDataCaller{contract: contract}, InternalDataTransactor: InternalDataTransactor{contract: contract}}, nil
}

// NewInternalDataCaller creates a new read-only instance of InternalData, bound to a specific deployed contract.
func NewInternalDataCaller(address common.Address, caller bind.ContractCaller) (*InternalDataCaller, error) {
	contract, err := bindInternalData(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &InternalDataCaller{contract: contract}, nil
}

// NewInternalDataTransactor creates a new write-only instance of InternalData, bound to a specific deployed contract.
func NewInternalDataTransactor(address common.Address, transactor bind.ContractTransactor) (*InternalDataTransactor, error) {
	contract, err := bindInternalData(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &InternalDataTransactor{contract: contract}, nil
}

// bindInternalData binds a generic wrapper to an already deployed contract.
func bindInternalData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InternalDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InternalData *InternalDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InternalData.Contract.InternalDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InternalData *InternalDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InternalData.Contract.InternalDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InternalData *InternalDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InternalData.Contract.InternalDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InternalData *InternalDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InternalData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InternalData *InternalDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InternalData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InternalData *InternalDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InternalData.Contract.contract.Transact(opts, method, params...)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_InternalData *InternalDataCaller) CryptoFiat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _InternalData.contract.Call(opts, out, "cryptoFiat")
	return *ret0, err
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_InternalData *InternalDataSession) CryptoFiat() (common.Address, error) {
	return _InternalData.Contract.CryptoFiat(&_InternalData.CallOpts)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_InternalData *InternalDataCallerSession) CryptoFiat() (common.Address, error) {
	return _InternalData.Contract.CryptoFiat(&_InternalData.CallOpts)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_InternalData *InternalDataTransactor) SwitchCryptoFiat(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _InternalData.contract.Transact(opts, "switchCryptoFiat", next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_InternalData *InternalDataSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _InternalData.Contract.SwitchCryptoFiat(&_InternalData.TransactOpts, next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_InternalData *InternalDataTransactorSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _InternalData.Contract.SwitchCryptoFiat(&_InternalData.TransactOpts, next)
}

// RelayABI is the input ABI used to generate the binding from.
const RelayABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// RelayBin is the compiled bytecode used for deploying new contracts.
const RelayBin = `0x6060604052341561000f57600080fd5b5b6101978061001f6000396000f300606060405263ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663516c4b848114610048578063fb55a05514610077575b600080fd5b341561005357600080fd5b61005b610098565b604051600160a060020a03909116815260200160405180910390f35b341561008257600080fd5b610096600160a060020a03600435166100a7565b005b600054600160a060020a031681565b60008054600160a060020a0333811692911690639afd453c90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561010e57600080fd5b6102c65a03f1151561011f57600080fd5b50505060405180519050600160a060020a031614151561013e57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b505600a165627a7a7230582056902239fd3785a7cc72902208106c220b593e44ead97a5fee0cec2f5be69aa90029`

// DeployRelay deploys a new Ethereum contract, binding an instance of Relay to it.
func DeployRelay(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Relay, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RelayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Relay{RelayCaller: RelayCaller{contract: contract}, RelayTransactor: RelayTransactor{contract: contract}}, nil
}

// Relay is an auto generated Go binding around an Ethereum contract.
type Relay struct {
	RelayCaller     // Read-only binding to the contract
	RelayTransactor // Write-only binding to the contract
}

// RelayCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelaySession struct {
	Contract     *Relay            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayCallerSession struct {
	Contract *RelayCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RelayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayTransactorSession struct {
	Contract     *RelayTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayRaw struct {
	Contract *Relay // Generic contract binding to access the raw methods on
}

// RelayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayCallerRaw struct {
	Contract *RelayCaller // Generic read-only contract binding to access the raw methods on
}

// RelayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayTransactorRaw struct {
	Contract *RelayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelay creates a new instance of Relay, bound to a specific deployed contract.
func NewRelay(address common.Address, backend bind.ContractBackend) (*Relay, error) {
	contract, err := bindRelay(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relay{RelayCaller: RelayCaller{contract: contract}, RelayTransactor: RelayTransactor{contract: contract}}, nil
}

// NewRelayCaller creates a new read-only instance of Relay, bound to a specific deployed contract.
func NewRelayCaller(address common.Address, caller bind.ContractCaller) (*RelayCaller, error) {
	contract, err := bindRelay(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RelayCaller{contract: contract}, nil
}

// NewRelayTransactor creates a new write-only instance of Relay, bound to a specific deployed contract.
func NewRelayTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayTransactor, error) {
	contract, err := bindRelay(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RelayTransactor{contract: contract}, nil
}

// bindRelay binds a generic wrapper to an already deployed contract.
func bindRelay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relay *RelayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relay.Contract.RelayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relay *RelayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relay.Contract.RelayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relay *RelayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relay.Contract.RelayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relay *RelayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relay *RelayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relay *RelayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relay.Contract.contract.Transact(opts, method, params...)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Relay *RelayCaller) CryptoFiat(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relay.contract.Call(opts, out, "cryptoFiat")
	return *ret0, err
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Relay *RelaySession) CryptoFiat() (common.Address, error) {
	return _Relay.Contract.CryptoFiat(&_Relay.CallOpts)
}

// CryptoFiat is a free data retrieval call binding the contract method 0x516c4b84.
//
// Solidity: function cryptoFiat() constant returns(address)
func (_Relay *RelayCallerSession) CryptoFiat() (common.Address, error) {
	return _Relay.Contract.CryptoFiat(&_Relay.CallOpts)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Relay *RelayTransactor) SwitchCryptoFiat(opts *bind.TransactOpts, next common.Address) (*types.Transaction, error) {
	return _Relay.contract.Transact(opts, "switchCryptoFiat", next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Relay *RelaySession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Relay.Contract.SwitchCryptoFiat(&_Relay.TransactOpts, next)
}

// SwitchCryptoFiat is a paid mutator transaction binding the contract method 0xfb55a055.
//
// Solidity: function switchCryptoFiat(next address) returns()
func (_Relay *RelayTransactorSession) SwitchCryptoFiat(next common.Address) (*types.Transaction, error) {
	return _Relay.Contract.SwitchCryptoFiat(&_Relay.TransactOpts, next)
}

// ReserveABI is the input ABI used to generate the binding from.
const ReserveABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"reserveBank\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cryptoFiat\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"appointReserveBank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"next\",\"type\":\"address\"}],\"name\":\"switchCryptoFiat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_cryptoFiat\",\"type\":\"address\"},{\"name\":\"_reserveBank\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"AccountFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"SupplyChanged\",\"type\":\"event\"}]"

// ReserveBin is the compiled bytecode used for deploying new contracts.
const ReserveBin = `0x6060604052341561000f57600080fd5b60405160408061090383398101604052808051919060200180519150505b60008054600160a060020a03808516600160a060020a03199283161790925560018054928416929091169190911790555b50505b610893806100706000396000f300606060405236156100675763ffffffff60e060020a60003504166302946804811461006c57806318160ddd1461009b578063516c4b84146100c057806398e52f9a146100ef578063b921e16314610107578063ddf05f591461011f578063fb55a05514610140575b600080fd5b341561007757600080fd5b61007f610161565b604051600160a060020a03909116815260200160405180910390f35b34156100a657600080fd5b6100ae610170565b60405190815260200160405180910390f35b34156100cb57600080fd5b61007f610180565b604051600160a060020a03909116815260200160405180910390f35b34156100fa57600080fd5b61010560043561018f565b005b341561011257600080fd5b6101056004356102ad565b005b341561012a57600080fd5b610105600160a060020a03600435166103b8565b005b341561014b57600080fd5b610105600160a060020a0360043516610400565b005b600154600160a060020a031681565b600061017a6104ab565b90505b90565b600054600160a060020a031681565b60015460009033600160a060020a039081169116146101ad57600080fd5b600154600160a060020a03166101c28161052a565b15156101cd57600080fd5b6101d681610542565b156101e057600080fd5b600160a060020a03811615156101f557600080fd5b6101fd6104ab565b91508282101561020c57600080fd5b828203915061021a8261055a565b60015461023090600160a060020a0316846105d0565b600154600090600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405190815260200160405180910390a37ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e98260405190815260200160405180910390a15b5b505b5050565b60015460009033600160a060020a039081169116146102cb57600080fd5b600154600160a060020a03166102e0816105fc565b156102ea57600080fd5b600160a060020a03811615156102ff57600080fd5b610307610170565b91508282018290101561031957600080fd5b908201906103268261055a565b60015461033c90600160a060020a031684610614565b600154600160a060020a031660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405190815260200160405180910390a37ff71f9c3841c0bab7774017ffe585aeab36b5438d148506067901d47c5fa6f7e98260405190815260200160405180910390a15b5b505b5050565b60015433600160a060020a039081169116146103d357600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60008054600160a060020a0333811692911690639afd453c90604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561044e57600080fd5b6102c65a03f1151561045f57600080fd5b50505060405180519050600160a060020a031614151561047e57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b60006104b5610643565b600160a060020a031663295f36d760056000806040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561050a57600080fd5b6102c65a03f1151561051b57600080fd5b50505060405180519150505b90565b600060018061053884610655565b161490505b919050565b600060048061053884610655565b161490505b919050565b610562610643565b600160a060020a031663461b09c0600560008460405160e060020a63ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156105b857600080fd5b6102c65a03f115156105c957600080fd5b5050505b50565b60006105db836106e2565b9050818110156105ea57600080fd5b6102a68383830361076f565b5b505050565b600060028061053884610655565b161490505b919050565b600061061f836106e2565b90508181018190101561063157600080fd5b6102a68383830161076f565b5b505050565b600061017a60016107f4565b90505b90565b600061065f610643565b600160a060020a031663295f36d7600184600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156106c057600080fd5b6102c65a03f115156106d157600080fd5b50505060405180519150505b919050565b60006106ec610643565b600160a060020a031663295f36d7600284600160a060020a031660010260006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156106c057600080fd5b6102c65a03f115156106d157600080fd5b50505060405180519150505b919050565b610777610643565b600160a060020a031663461b09c0600284600160a060020a03166001028460010260405160e060020a63ffffffff8616028152600481019390935260248301919091526044820152606401600060405180830381600087803b15156107db57600080fd5b6102c65a03f115156107ec57600080fd5b5050505b5050565b60008054600160a060020a03166313c0136883836040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156106c057600080fd5b6102c65a03f115156106d157600080fd5b50505060405180519150505b9190505600a165627a7a723058208552427c3025902f692877013f3b80929a0d31703f9edfb455899762b2fb40570029`

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
