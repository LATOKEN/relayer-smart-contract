// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbr

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthbrMetaData contains all meta data concerning the Ethbr contract.
var EthbrMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExtraFeeSupplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_backendSrvAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_chainID\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_executedProposals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_isInitialised\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_nativeResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"adminCollectFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceIDOwner\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"resourceIDSpender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"approveSpending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToLA\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"handlers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"chainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initBackendSrvAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"internalDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ownableInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBackendSrv\",\"type\":\"address\"}],\"name\":\"setBackendSrv\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"setNativeResourceID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EthbrABI is the input ABI used to generate the binding from.
// Deprecated: Use EthbrMetaData.ABI instead.
var EthbrABI = EthbrMetaData.ABI

// Ethbr is an auto generated Go binding around an Ethereum contract.
type Ethbr struct {
	EthbrCaller     // Read-only binding to the contract
	EthbrTransactor // Write-only binding to the contract
	EthbrFilterer   // Log filterer for contract events
}

// EthbrCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthbrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthbrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthbrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthbrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthbrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthbrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthbrSession struct {
	Contract     *Ethbr            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthbrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthbrCallerSession struct {
	Contract *EthbrCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthbrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthbrTransactorSession struct {
	Contract     *EthbrTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthbrRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthbrRaw struct {
	Contract *Ethbr // Generic contract binding to access the raw methods on
}

// EthbrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthbrCallerRaw struct {
	Contract *EthbrCaller // Generic read-only contract binding to access the raw methods on
}

// EthbrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthbrTransactorRaw struct {
	Contract *EthbrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthbr creates a new instance of Ethbr, bound to a specific deployed contract.
func NewEthbr(address common.Address, backend bind.ContractBackend) (*Ethbr, error) {
	contract, err := bindEthbr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethbr{EthbrCaller: EthbrCaller{contract: contract}, EthbrTransactor: EthbrTransactor{contract: contract}, EthbrFilterer: EthbrFilterer{contract: contract}}, nil
}

// NewEthbrCaller creates a new read-only instance of Ethbr, bound to a specific deployed contract.
func NewEthbrCaller(address common.Address, caller bind.ContractCaller) (*EthbrCaller, error) {
	contract, err := bindEthbr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthbrCaller{contract: contract}, nil
}

// NewEthbrTransactor creates a new write-only instance of Ethbr, bound to a specific deployed contract.
func NewEthbrTransactor(address common.Address, transactor bind.ContractTransactor) (*EthbrTransactor, error) {
	contract, err := bindEthbr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthbrTransactor{contract: contract}, nil
}

// NewEthbrFilterer creates a new log filterer instance of Ethbr, bound to a specific deployed contract.
func NewEthbrFilterer(address common.Address, filterer bind.ContractFilterer) (*EthbrFilterer, error) {
	contract, err := bindEthbr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthbrFilterer{contract: contract}, nil
}

// bindEthbr binds a generic wrapper to an already deployed contract.
func bindEthbr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthbrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethbr *EthbrRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethbr.Contract.EthbrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethbr *EthbrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethbr.Contract.EthbrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethbr *EthbrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethbr.Contract.EthbrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethbr *EthbrCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethbr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethbr *EthbrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethbr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethbr *EthbrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethbr.Contract.contract.Transact(opts, method, params...)
}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_Ethbr *EthbrCaller) BackendSrvAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "_backendSrvAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_Ethbr *EthbrSession) BackendSrvAddress() (common.Address, error) {
	return _Ethbr.Contract.BackendSrvAddress(&_Ethbr.CallOpts)
}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_Ethbr *EthbrCallerSession) BackendSrvAddress() (common.Address, error) {
	return _Ethbr.Contract.BackendSrvAddress(&_Ethbr.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_Ethbr *EthbrCaller) ChainID(opts *bind.CallOpts) ([8]byte, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "_chainID")

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_Ethbr *EthbrSession) ChainID() ([8]byte, error) {
	return _Ethbr.Contract.ChainID(&_Ethbr.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_Ethbr *EthbrCallerSession) ChainID() ([8]byte, error) {
	return _Ethbr.Contract.ChainID(&_Ethbr.CallOpts)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_Ethbr *EthbrCaller) DepositCounts(opts *bind.CallOpts, arg0 [8]byte) (uint64, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "_depositCounts", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_Ethbr *EthbrSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _Ethbr.Contract.DepositCounts(&_Ethbr.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_Ethbr *EthbrCallerSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _Ethbr.Contract.DepositCounts(&_Ethbr.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_Ethbr *EthbrCaller) DepositRecords(opts *bind.CallOpts, arg0 uint64, arg1 [8]byte) ([]byte, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_Ethbr *EthbrSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _Ethbr.Contract.DepositRecords(&_Ethbr.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_Ethbr *EthbrCallerSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _Ethbr.Contract.DepositRecords(&_Ethbr.CallOpts, arg0, arg1)
}

// ExecutedProposals is a free data retrieval call binding the contract method 0x9c1650e0.
//
// Solidity: function _executedProposals(bytes32 , bytes32 ) view returns(bool)
func (_Ethbr *EthbrCaller) ExecutedProposals(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "_executedProposals", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExecutedProposals is a free data retrieval call binding the contract method 0x9c1650e0.
//
// Solidity: function _executedProposals(bytes32 , bytes32 ) view returns(bool)
func (_Ethbr *EthbrSession) ExecutedProposals(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _Ethbr.Contract.ExecutedProposals(&_Ethbr.CallOpts, arg0, arg1)
}

// ExecutedProposals is a free data retrieval call binding the contract method 0x9c1650e0.
//
// Solidity: function _executedProposals(bytes32 , bytes32 ) view returns(bool)
func (_Ethbr *EthbrCallerSession) ExecutedProposals(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _Ethbr.Contract.ExecutedProposals(&_Ethbr.CallOpts, arg0, arg1)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Ethbr *EthbrCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Ethbr *EthbrSession) Fee() (*big.Int, error) {
	return _Ethbr.Contract.Fee(&_Ethbr.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Ethbr *EthbrCallerSession) Fee() (*big.Int, error) {
	return _Ethbr.Contract.Fee(&_Ethbr.CallOpts)
}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_Ethbr *EthbrCaller) IsInitialised(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "_isInitialised")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_Ethbr *EthbrSession) IsInitialised() (bool, error) {
	return _Ethbr.Contract.IsInitialised(&_Ethbr.CallOpts)
}

// IsInitialised is a free data retrieval call binding the contract method 0xdd2e8ec3.
//
// Solidity: function _isInitialised() view returns(bool)
func (_Ethbr *EthbrCallerSession) IsInitialised() (bool, error) {
	return _Ethbr.Contract.IsInitialised(&_Ethbr.CallOpts)
}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_Ethbr *EthbrCaller) NativeResourceID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "_nativeResourceID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_Ethbr *EthbrSession) NativeResourceID() ([32]byte, error) {
	return _Ethbr.Contract.NativeResourceID(&_Ethbr.CallOpts)
}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_Ethbr *EthbrCallerSession) NativeResourceID() ([32]byte, error) {
	return _Ethbr.Contract.NativeResourceID(&_Ethbr.CallOpts)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Ethbr *EthbrCaller) ResourceIDToHandlerAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "_resourceIDToHandlerAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Ethbr *EthbrSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Ethbr.Contract.ResourceIDToHandlerAddress(&_Ethbr.CallOpts, arg0)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Ethbr *EthbrCallerSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Ethbr.Contract.ResourceIDToHandlerAddress(&_Ethbr.CallOpts, arg0)
}

// Handlers is a free data retrieval call binding the contract method 0x1a21c0bc.
//
// Solidity: function handlers(address ) view returns(bool)
func (_Ethbr *EthbrCaller) Handlers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "handlers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Handlers is a free data retrieval call binding the contract method 0x1a21c0bc.
//
// Solidity: function handlers(address ) view returns(bool)
func (_Ethbr *EthbrSession) Handlers(arg0 common.Address) (bool, error) {
	return _Ethbr.Contract.Handlers(&_Ethbr.CallOpts, arg0)
}

// Handlers is a free data retrieval call binding the contract method 0x1a21c0bc.
//
// Solidity: function handlers(address ) view returns(bool)
func (_Ethbr *EthbrCallerSession) Handlers(arg0 common.Address) (bool, error) {
	return _Ethbr.Contract.Handlers(&_Ethbr.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethbr *EthbrCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethbr *EthbrSession) Owner() (common.Address, error) {
	return _Ethbr.Contract.Owner(&_Ethbr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethbr *EthbrCallerSession) Owner() (common.Address, error) {
	return _Ethbr.Contract.Owner(&_Ethbr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Ethbr *EthbrCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Ethbr *EthbrSession) Paused() (bool, error) {
	return _Ethbr.Contract.Paused(&_Ethbr.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Ethbr *EthbrCallerSession) Paused() (bool, error) {
	return _Ethbr.Contract.Paused(&_Ethbr.CallOpts)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_Ethbr *EthbrCaller) Sub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ethbr.contract.Call(opts, &out, "sub", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_Ethbr *EthbrSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Ethbr.Contract.Sub(&_Ethbr.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_Ethbr *EthbrCallerSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Ethbr.Contract.Sub(&_Ethbr.CallOpts, a, b)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_Ethbr *EthbrTransactor) AdminCollectFees(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "adminCollectFees", recipient, amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_Ethbr *EthbrSession) AdminCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ethbr.Contract.AdminCollectFees(&_Ethbr.TransactOpts, recipient, amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_Ethbr *EthbrTransactorSession) AdminCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ethbr.Contract.AdminCollectFees(&_Ethbr.TransactOpts, recipient, amount)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Ethbr *EthbrTransactor) AdminPauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "adminPauseTransfers")
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Ethbr *EthbrSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Ethbr.Contract.AdminPauseTransfers(&_Ethbr.TransactOpts)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Ethbr *EthbrTransactorSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Ethbr.Contract.AdminPauseTransfers(&_Ethbr.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Ethbr *EthbrTransactor) AdminUnpauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "adminUnpauseTransfers")
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Ethbr *EthbrSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Ethbr.Contract.AdminUnpauseTransfers(&_Ethbr.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Ethbr *EthbrTransactorSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Ethbr.Contract.AdminUnpauseTransfers(&_Ethbr.TransactOpts)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Ethbr *EthbrTransactor) AdminWithdraw(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "adminWithdraw", handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Ethbr *EthbrSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Ethbr.Contract.AdminWithdraw(&_Ethbr.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Ethbr *EthbrTransactorSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Ethbr.Contract.AdminWithdraw(&_Ethbr.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// ApproveSpending is a paid mutator transaction binding the contract method 0xde771b31.
//
// Solidity: function approveSpending(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_Ethbr *EthbrTransactor) ApproveSpending(opts *bind.TransactOpts, resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "approveSpending", resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// ApproveSpending is a paid mutator transaction binding the contract method 0xde771b31.
//
// Solidity: function approveSpending(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_Ethbr *EthbrSession) ApproveSpending(resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Ethbr.Contract.ApproveSpending(&_Ethbr.TransactOpts, resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// ApproveSpending is a paid mutator transaction binding the contract method 0xde771b31.
//
// Solidity: function approveSpending(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_Ethbr *EthbrTransactorSession) ApproveSpending(resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Ethbr.Contract.ApproveSpending(&_Ethbr.TransactOpts, resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_Ethbr *EthbrTransactor) ChangeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "changeFee", newFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_Ethbr *EthbrSession) ChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Ethbr.Contract.ChangeFee(&_Ethbr.TransactOpts, newFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_Ethbr *EthbrTransactorSession) ChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Ethbr.Contract.ChangeFee(&_Ethbr.TransactOpts, newFee)
}

// Deposit is a paid mutator transaction binding the contract method 0xe4e21fce.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, uint256 amountToLA, bytes params) payable returns()
func (_Ethbr *EthbrTransactor) Deposit(opts *bind.TransactOpts, destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, amountToLA *big.Int, params []byte) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "deposit", destinationChainID, resourceID, amount, recipientAddress, amountToLA, params)
}

// Deposit is a paid mutator transaction binding the contract method 0xe4e21fce.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, uint256 amountToLA, bytes params) payable returns()
func (_Ethbr *EthbrSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, amountToLA *big.Int, params []byte) (*types.Transaction, error) {
	return _Ethbr.Contract.Deposit(&_Ethbr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress, amountToLA, params)
}

// Deposit is a paid mutator transaction binding the contract method 0xe4e21fce.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, uint256 amountToLA, bytes params) payable returns()
func (_Ethbr *EthbrTransactorSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, amountToLA *big.Int, params []byte) (*types.Transaction, error) {
	return _Ethbr.Contract.Deposit(&_Ethbr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress, amountToLA, params)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Ethbr *EthbrTransactor) DepositFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "depositFunds")
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Ethbr *EthbrSession) DepositFunds() (*types.Transaction, error) {
	return _Ethbr.Contract.DepositFunds(&_Ethbr.TransactOpts)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Ethbr *EthbrTransactorSession) DepositFunds() (*types.Transaction, error) {
	return _Ethbr.Contract.DepositFunds(&_Ethbr.TransactOpts)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_Ethbr *EthbrTransactor) ExecuteProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "executeProposal", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_Ethbr *EthbrSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Ethbr.Contract.ExecuteProposal(&_Ethbr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_Ethbr *EthbrTransactorSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Ethbr.Contract.ExecuteProposal(&_Ethbr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x92c6ecf2.
//
// Solidity: function initialize(bytes8 chainID, uint256 fee, address initBackendSrvAddress) returns()
func (_Ethbr *EthbrTransactor) Initialize(opts *bind.TransactOpts, chainID [8]byte, fee *big.Int, initBackendSrvAddress common.Address) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "initialize", chainID, fee, initBackendSrvAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x92c6ecf2.
//
// Solidity: function initialize(bytes8 chainID, uint256 fee, address initBackendSrvAddress) returns()
func (_Ethbr *EthbrSession) Initialize(chainID [8]byte, fee *big.Int, initBackendSrvAddress common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.Initialize(&_Ethbr.TransactOpts, chainID, fee, initBackendSrvAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x92c6ecf2.
//
// Solidity: function initialize(bytes8 chainID, uint256 fee, address initBackendSrvAddress) returns()
func (_Ethbr *EthbrTransactorSession) Initialize(chainID [8]byte, fee *big.Int, initBackendSrvAddress common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.Initialize(&_Ethbr.TransactOpts, chainID, fee, initBackendSrvAddress)
}

// InternalDeposit is a paid mutator transaction binding the contract method 0x8ece6f27.
//
// Solidity: function internalDeposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) returns()
func (_Ethbr *EthbrTransactor) InternalDeposit(opts *bind.TransactOpts, destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "internalDeposit", destinationChainID, resourceID, amount, recipientAddress)
}

// InternalDeposit is a paid mutator transaction binding the contract method 0x8ece6f27.
//
// Solidity: function internalDeposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) returns()
func (_Ethbr *EthbrSession) InternalDeposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.InternalDeposit(&_Ethbr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress)
}

// InternalDeposit is a paid mutator transaction binding the contract method 0x8ece6f27.
//
// Solidity: function internalDeposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) returns()
func (_Ethbr *EthbrTransactorSession) InternalDeposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.InternalDeposit(&_Ethbr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner) returns()
func (_Ethbr *EthbrTransactor) OwnableInit(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "ownableInit", owner)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner) returns()
func (_Ethbr *EthbrSession) OwnableInit(owner common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.OwnableInit(&_Ethbr.TransactOpts, owner)
}

// OwnableInit is a paid mutator transaction binding the contract method 0xea439b2b.
//
// Solidity: function ownableInit(address owner) returns()
func (_Ethbr *EthbrTransactorSession) OwnableInit(owner common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.OwnableInit(&_Ethbr.TransactOpts, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethbr *EthbrTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethbr *EthbrSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ethbr.Contract.RenounceOwnership(&_Ethbr.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ethbr *EthbrTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ethbr.Contract.RenounceOwnership(&_Ethbr.TransactOpts)
}

// SetBackendSrv is a paid mutator transaction binding the contract method 0x0bf6096f.
//
// Solidity: function setBackendSrv(address newBackendSrv) returns()
func (_Ethbr *EthbrTransactor) SetBackendSrv(opts *bind.TransactOpts, newBackendSrv common.Address) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "setBackendSrv", newBackendSrv)
}

// SetBackendSrv is a paid mutator transaction binding the contract method 0x0bf6096f.
//
// Solidity: function setBackendSrv(address newBackendSrv) returns()
func (_Ethbr *EthbrSession) SetBackendSrv(newBackendSrv common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.SetBackendSrv(&_Ethbr.TransactOpts, newBackendSrv)
}

// SetBackendSrv is a paid mutator transaction binding the contract method 0x0bf6096f.
//
// Solidity: function setBackendSrv(address newBackendSrv) returns()
func (_Ethbr *EthbrTransactorSession) SetBackendSrv(newBackendSrv common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.SetBackendSrv(&_Ethbr.TransactOpts, newBackendSrv)
}

// SetBurnable is a paid mutator transaction binding the contract method 0xdf0fc133.
//
// Solidity: function setBurnable(address handlerAddress, address tokenAddress) returns()
func (_Ethbr *EthbrTransactor) SetBurnable(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "setBurnable", handlerAddress, tokenAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0xdf0fc133.
//
// Solidity: function setBurnable(address handlerAddress, address tokenAddress) returns()
func (_Ethbr *EthbrSession) SetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.SetBurnable(&_Ethbr.TransactOpts, handlerAddress, tokenAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0xdf0fc133.
//
// Solidity: function setBurnable(address handlerAddress, address tokenAddress) returns()
func (_Ethbr *EthbrTransactorSession) SetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.SetBurnable(&_Ethbr.TransactOpts, handlerAddress, tokenAddress)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool value) returns()
func (_Ethbr *EthbrTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, value bool) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "setHandler", _handler, value)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool value) returns()
func (_Ethbr *EthbrSession) SetHandler(_handler common.Address, value bool) (*types.Transaction, error) {
	return _Ethbr.Contract.SetHandler(&_Ethbr.TransactOpts, _handler, value)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool value) returns()
func (_Ethbr *EthbrTransactorSession) SetHandler(_handler common.Address, value bool) (*types.Transaction, error) {
	return _Ethbr.Contract.SetHandler(&_Ethbr.TransactOpts, _handler, value)
}

// SetNativeResourceID is a paid mutator transaction binding the contract method 0x21016ed8.
//
// Solidity: function setNativeResourceID(bytes32 resourceID) returns()
func (_Ethbr *EthbrTransactor) SetNativeResourceID(opts *bind.TransactOpts, resourceID [32]byte) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "setNativeResourceID", resourceID)
}

// SetNativeResourceID is a paid mutator transaction binding the contract method 0x21016ed8.
//
// Solidity: function setNativeResourceID(bytes32 resourceID) returns()
func (_Ethbr *EthbrSession) SetNativeResourceID(resourceID [32]byte) (*types.Transaction, error) {
	return _Ethbr.Contract.SetNativeResourceID(&_Ethbr.TransactOpts, resourceID)
}

// SetNativeResourceID is a paid mutator transaction binding the contract method 0x21016ed8.
//
// Solidity: function setNativeResourceID(bytes32 resourceID) returns()
func (_Ethbr *EthbrTransactorSession) SetNativeResourceID(resourceID [32]byte) (*types.Transaction, error) {
	return _Ethbr.Contract.SetNativeResourceID(&_Ethbr.TransactOpts, resourceID)
}

// SetResource is a paid mutator transaction binding the contract method 0xa737be4f.
//
// Solidity: function setResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Ethbr *EthbrTransactor) SetResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "setResource", handlerAddress, resourceID, tokenAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xa737be4f.
//
// Solidity: function setResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Ethbr *EthbrSession) SetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.SetResource(&_Ethbr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xa737be4f.
//
// Solidity: function setResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Ethbr *EthbrTransactorSession) SetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.SetResource(&_Ethbr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethbr *EthbrTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ethbr.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethbr *EthbrSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.TransferOwnership(&_Ethbr.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ethbr *EthbrTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ethbr.Contract.TransferOwnership(&_Ethbr.TransactOpts, newOwner)
}

// EthbrDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Ethbr contract.
type EthbrDepositIterator struct {
	Event *EthbrDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthbrDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthbrDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthbrDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthbrDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthbrDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthbrDeposit represents a Deposit event raised by the Ethbr contract.
type EthbrDeposit struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	ResourceID         [32]byte
	DepositNonce       uint64
	Depositor          common.Address
	RecipientAddress   common.Address
	TokenAddress       common.Address
	Amount             *big.Int
	DataHash           [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x370525803ffa9a7c0e6adb3868e393dca45d8b42b2f62fd1f23ecfe99f6ce8fc.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes32 dataHash)
func (_Ethbr *EthbrFilterer) FilterDeposit(opts *bind.FilterOpts, destinationChainID [][8]byte, resourceID [][32]byte, depositNonce []uint64) (*EthbrDepositIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _Ethbr.contract.FilterLogs(opts, "Deposit", destinationChainIDRule, resourceIDRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return &EthbrDepositIterator{contract: _Ethbr.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x370525803ffa9a7c0e6adb3868e393dca45d8b42b2f62fd1f23ecfe99f6ce8fc.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes32 dataHash)
func (_Ethbr *EthbrFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EthbrDeposit, destinationChainID [][8]byte, resourceID [][32]byte, depositNonce []uint64) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _Ethbr.contract.WatchLogs(opts, "Deposit", destinationChainIDRule, resourceIDRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthbrDeposit)
				if err := _Ethbr.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x370525803ffa9a7c0e6adb3868e393dca45d8b42b2f62fd1f23ecfe99f6ce8fc.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes32 dataHash)
func (_Ethbr *EthbrFilterer) ParseDeposit(log types.Log) (*EthbrDeposit, error) {
	event := new(EthbrDeposit)
	if err := _Ethbr.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthbrExtraFeeSuppliedIterator is returned from FilterExtraFeeSupplied and is used to iterate over the raw logs and unpacked data for ExtraFeeSupplied events raised by the Ethbr contract.
type EthbrExtraFeeSuppliedIterator struct {
	Event *EthbrExtraFeeSupplied // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthbrExtraFeeSuppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthbrExtraFeeSupplied)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthbrExtraFeeSupplied)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthbrExtraFeeSuppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthbrExtraFeeSuppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthbrExtraFeeSupplied represents a ExtraFeeSupplied event raised by the Ethbr contract.
type EthbrExtraFeeSupplied struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	DepositNonce       uint64
	ResourceID         [32]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExtraFeeSupplied is a free log retrieval operation binding the contract event 0x525223e7c9e63747e47dd4558940766054da3d0378f4006848d2a201545f55a4.
//
// Solidity: event ExtraFeeSupplied(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount)
func (_Ethbr *EthbrFilterer) FilterExtraFeeSupplied(opts *bind.FilterOpts) (*EthbrExtraFeeSuppliedIterator, error) {

	logs, sub, err := _Ethbr.contract.FilterLogs(opts, "ExtraFeeSupplied")
	if err != nil {
		return nil, err
	}
	return &EthbrExtraFeeSuppliedIterator{contract: _Ethbr.contract, event: "ExtraFeeSupplied", logs: logs, sub: sub}, nil
}

// WatchExtraFeeSupplied is a free log subscription operation binding the contract event 0x525223e7c9e63747e47dd4558940766054da3d0378f4006848d2a201545f55a4.
//
// Solidity: event ExtraFeeSupplied(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount)
func (_Ethbr *EthbrFilterer) WatchExtraFeeSupplied(opts *bind.WatchOpts, sink chan<- *EthbrExtraFeeSupplied) (event.Subscription, error) {

	logs, sub, err := _Ethbr.contract.WatchLogs(opts, "ExtraFeeSupplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthbrExtraFeeSupplied)
				if err := _Ethbr.contract.UnpackLog(event, "ExtraFeeSupplied", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtraFeeSupplied is a log parse operation binding the contract event 0x525223e7c9e63747e47dd4558940766054da3d0378f4006848d2a201545f55a4.
//
// Solidity: event ExtraFeeSupplied(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount)
func (_Ethbr *EthbrFilterer) ParseExtraFeeSupplied(log types.Log) (*EthbrExtraFeeSupplied, error) {
	event := new(EthbrExtraFeeSupplied)
	if err := _Ethbr.contract.UnpackLog(event, "ExtraFeeSupplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthbrOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ethbr contract.
type EthbrOwnershipTransferredIterator struct {
	Event *EthbrOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthbrOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthbrOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthbrOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthbrOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthbrOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthbrOwnershipTransferred represents a OwnershipTransferred event raised by the Ethbr contract.
type EthbrOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethbr *EthbrFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthbrOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ethbr.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthbrOwnershipTransferredIterator{contract: _Ethbr.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethbr *EthbrFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthbrOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ethbr.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthbrOwnershipTransferred)
				if err := _Ethbr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ethbr *EthbrFilterer) ParseOwnershipTransferred(log types.Log) (*EthbrOwnershipTransferred, error) {
	event := new(EthbrOwnershipTransferred)
	if err := _Ethbr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthbrPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Ethbr contract.
type EthbrPausedIterator struct {
	Event *EthbrPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthbrPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthbrPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthbrPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthbrPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthbrPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthbrPaused represents a Paused event raised by the Ethbr contract.
type EthbrPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Ethbr *EthbrFilterer) FilterPaused(opts *bind.FilterOpts) (*EthbrPausedIterator, error) {

	logs, sub, err := _Ethbr.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EthbrPausedIterator{contract: _Ethbr.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Ethbr *EthbrFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EthbrPaused) (event.Subscription, error) {

	logs, sub, err := _Ethbr.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthbrPaused)
				if err := _Ethbr.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Ethbr *EthbrFilterer) ParsePaused(log types.Log) (*EthbrPaused, error) {
	event := new(EthbrPaused)
	if err := _Ethbr.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthbrProposalEventIterator is returned from FilterProposalEvent and is used to iterate over the raw logs and unpacked data for ProposalEvent events raised by the Ethbr contract.
type EthbrProposalEventIterator struct {
	Event *EthbrProposalEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthbrProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthbrProposalEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthbrProposalEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthbrProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthbrProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthbrProposalEvent represents a ProposalEvent event raised by the Ethbr contract.
type EthbrProposalEvent struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	DepositNonce       uint64
	Status             uint8
	ResourceID         [32]byte
	DataHash           [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProposalEvent is a free log retrieval operation binding the contract event 0x9686dcabd0450cad86a88df15a9d35b08b35d1b08a19008df37cf8538c467516.
//
// Solidity: event ProposalEvent(bytes8 indexed originChainID, bytes8 indexed destinationChainID, address indexed recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_Ethbr *EthbrFilterer) FilterProposalEvent(opts *bind.FilterOpts, originChainID [][8]byte, destinationChainID [][8]byte, recipientAddress []common.Address) (*EthbrProposalEventIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var recipientAddressRule []interface{}
	for _, recipientAddressItem := range recipientAddress {
		recipientAddressRule = append(recipientAddressRule, recipientAddressItem)
	}

	logs, sub, err := _Ethbr.contract.FilterLogs(opts, "ProposalEvent", originChainIDRule, destinationChainIDRule, recipientAddressRule)
	if err != nil {
		return nil, err
	}
	return &EthbrProposalEventIterator{contract: _Ethbr.contract, event: "ProposalEvent", logs: logs, sub: sub}, nil
}

// WatchProposalEvent is a free log subscription operation binding the contract event 0x9686dcabd0450cad86a88df15a9d35b08b35d1b08a19008df37cf8538c467516.
//
// Solidity: event ProposalEvent(bytes8 indexed originChainID, bytes8 indexed destinationChainID, address indexed recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_Ethbr *EthbrFilterer) WatchProposalEvent(opts *bind.WatchOpts, sink chan<- *EthbrProposalEvent, originChainID [][8]byte, destinationChainID [][8]byte, recipientAddress []common.Address) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var recipientAddressRule []interface{}
	for _, recipientAddressItem := range recipientAddress {
		recipientAddressRule = append(recipientAddressRule, recipientAddressItem)
	}

	logs, sub, err := _Ethbr.contract.WatchLogs(opts, "ProposalEvent", originChainIDRule, destinationChainIDRule, recipientAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthbrProposalEvent)
				if err := _Ethbr.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalEvent is a log parse operation binding the contract event 0x9686dcabd0450cad86a88df15a9d35b08b35d1b08a19008df37cf8538c467516.
//
// Solidity: event ProposalEvent(bytes8 indexed originChainID, bytes8 indexed destinationChainID, address indexed recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_Ethbr *EthbrFilterer) ParseProposalEvent(log types.Log) (*EthbrProposalEvent, error) {
	event := new(EthbrProposalEvent)
	if err := _Ethbr.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthbrUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Ethbr contract.
type EthbrUnpausedIterator struct {
	Event *EthbrUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthbrUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthbrUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthbrUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthbrUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthbrUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthbrUnpaused represents a Unpaused event raised by the Ethbr contract.
type EthbrUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Ethbr *EthbrFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EthbrUnpausedIterator, error) {

	logs, sub, err := _Ethbr.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EthbrUnpausedIterator{contract: _Ethbr.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Ethbr *EthbrFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthbrUnpaused) (event.Subscription, error) {

	logs, sub, err := _Ethbr.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthbrUnpaused)
				if err := _Ethbr.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Ethbr *EthbrFilterer) ParseUnpaused(log types.Log) (*EthbrUnpaused, error) {
	event := new(EthbrUnpaused)
	if err := _Ethbr.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
