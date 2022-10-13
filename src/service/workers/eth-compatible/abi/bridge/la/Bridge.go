// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package labr

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

// IBridgeProposal is an auto generated low-level Go binding around an user-defined struct.
type IBridgeProposal struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	YesVotes      []common.Address
	NoVotes       []common.Address
	Status        uint8
	ProposedBlock *big.Int
}

// LabrMetaData contains all meta data concerning the Labr contract.
var LabrMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"ExtraFeeSupplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"ProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"ProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"RelayerThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardCollected\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_backendSrvAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_balancerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_chainID\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_dexAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_expiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_hasVotedOnProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_nativeResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_proposals\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_yesVotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_noVotes\",\"type\":\"address[]\"},{\"internalType\":\"enumIBridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIBridge.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_relayerHubAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_relayerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalProposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"adminCollectFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceIDOwner\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"resourceIDSpender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"approveSpending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"backendSrvCollectFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"cancelProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDEX\",\"type\":\"address\"}],\"name\":\"changeDEXAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"changeRelayerThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWETH\",\"type\":\"address\"}],\"name\":\"changeWETHAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"extraLATransferred\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getExtraLATransferred\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_yesVotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_noVotes\",\"type\":\"address[]\"},{\"internalType\":\"enumIBridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIBridge.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"chainID_\",\"type\":\"bytes8\"},{\"internalType\":\"uint256\",\"name\":\"relayerThreshold_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initBackendSrvAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initBalancerAddress_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayerCollectReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBackendSrv\",\"type\":\"address\"}],\"name\":\"setBackendSrvAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBalancer\",\"type\":\"address\"}],\"name\":\"setBalancerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBridgeDelegate\",\"type\":\"address\"}],\"name\":\"setBridgeDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"setNativeResourceID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRelayerHub\",\"type\":\"address\"}],\"name\":\"setRelayerHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"transferExtraFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"addresspayable\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToTransfer\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"enumIBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"updateExternalTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"voteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LabrABI is the input ABI used to generate the binding from.
// Deprecated: Use LabrMetaData.ABI instead.
var LabrABI = LabrMetaData.ABI

// Labr is an auto generated Go binding around an Ethereum contract.
type Labr struct {
	LabrCaller     // Read-only binding to the contract
	LabrTransactor // Write-only binding to the contract
	LabrFilterer   // Log filterer for contract events
}

// LabrCaller is an auto generated read-only Go binding around an Ethereum contract.
type LabrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LabrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LabrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LabrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LabrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LabrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LabrSession struct {
	Contract     *Labr             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LabrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LabrCallerSession struct {
	Contract *LabrCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LabrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LabrTransactorSession struct {
	Contract     *LabrTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LabrRaw is an auto generated low-level Go binding around an Ethereum contract.
type LabrRaw struct {
	Contract *Labr // Generic contract binding to access the raw methods on
}

// LabrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LabrCallerRaw struct {
	Contract *LabrCaller // Generic read-only contract binding to access the raw methods on
}

// LabrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LabrTransactorRaw struct {
	Contract *LabrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLabr creates a new instance of Labr, bound to a specific deployed contract.
func NewLabr(address common.Address, backend bind.ContractBackend) (*Labr, error) {
	contract, err := bindLabr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Labr{LabrCaller: LabrCaller{contract: contract}, LabrTransactor: LabrTransactor{contract: contract}, LabrFilterer: LabrFilterer{contract: contract}}, nil
}

// NewLabrCaller creates a new read-only instance of Labr, bound to a specific deployed contract.
func NewLabrCaller(address common.Address, caller bind.ContractCaller) (*LabrCaller, error) {
	contract, err := bindLabr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LabrCaller{contract: contract}, nil
}

// NewLabrTransactor creates a new write-only instance of Labr, bound to a specific deployed contract.
func NewLabrTransactor(address common.Address, transactor bind.ContractTransactor) (*LabrTransactor, error) {
	contract, err := bindLabr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LabrTransactor{contract: contract}, nil
}

// NewLabrFilterer creates a new log filterer instance of Labr, bound to a specific deployed contract.
func NewLabrFilterer(address common.Address, filterer bind.ContractFilterer) (*LabrFilterer, error) {
	contract, err := bindLabr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LabrFilterer{contract: contract}, nil
}

// bindLabr binds a generic wrapper to an already deployed contract.
func bindLabr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LabrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Labr *LabrRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Labr.Contract.LabrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Labr *LabrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.Contract.LabrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Labr *LabrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Labr.Contract.LabrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Labr *LabrCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Labr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Labr *LabrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Labr *LabrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Labr.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xe0af3616.
//
// Solidity: function _WETH() view returns(address)
func (_Labr *LabrCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xe0af3616.
//
// Solidity: function _WETH() view returns(address)
func (_Labr *LabrSession) WETH() (common.Address, error) {
	return _Labr.Contract.WETH(&_Labr.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xe0af3616.
//
// Solidity: function _WETH() view returns(address)
func (_Labr *LabrCallerSession) WETH() (common.Address, error) {
	return _Labr.Contract.WETH(&_Labr.CallOpts)
}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_Labr *LabrCaller) BackendSrvAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_backendSrvAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_Labr *LabrSession) BackendSrvAddress() (common.Address, error) {
	return _Labr.Contract.BackendSrvAddress(&_Labr.CallOpts)
}

// BackendSrvAddress is a free data retrieval call binding the contract method 0x7f5d6a5e.
//
// Solidity: function _backendSrvAddress() view returns(address)
func (_Labr *LabrCallerSession) BackendSrvAddress() (common.Address, error) {
	return _Labr.Contract.BackendSrvAddress(&_Labr.CallOpts)
}

// BalancerAddress is a free data retrieval call binding the contract method 0x8c682bbf.
//
// Solidity: function _balancerAddress() view returns(address)
func (_Labr *LabrCaller) BalancerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_balancerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalancerAddress is a free data retrieval call binding the contract method 0x8c682bbf.
//
// Solidity: function _balancerAddress() view returns(address)
func (_Labr *LabrSession) BalancerAddress() (common.Address, error) {
	return _Labr.Contract.BalancerAddress(&_Labr.CallOpts)
}

// BalancerAddress is a free data retrieval call binding the contract method 0x8c682bbf.
//
// Solidity: function _balancerAddress() view returns(address)
func (_Labr *LabrCallerSession) BalancerAddress() (common.Address, error) {
	return _Labr.Contract.BalancerAddress(&_Labr.CallOpts)
}

// BridgeDelegate is a free data retrieval call binding the contract method 0x20f77707.
//
// Solidity: function _bridgeDelegate() view returns(address)
func (_Labr *LabrCaller) BridgeDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_bridgeDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeDelegate is a free data retrieval call binding the contract method 0x20f77707.
//
// Solidity: function _bridgeDelegate() view returns(address)
func (_Labr *LabrSession) BridgeDelegate() (common.Address, error) {
	return _Labr.Contract.BridgeDelegate(&_Labr.CallOpts)
}

// BridgeDelegate is a free data retrieval call binding the contract method 0x20f77707.
//
// Solidity: function _bridgeDelegate() view returns(address)
func (_Labr *LabrCallerSession) BridgeDelegate() (common.Address, error) {
	return _Labr.Contract.BridgeDelegate(&_Labr.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_Labr *LabrCaller) ChainID(opts *bind.CallOpts) ([8]byte, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_chainID")

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_Labr *LabrSession) ChainID() ([8]byte, error) {
	return _Labr.Contract.ChainID(&_Labr.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_Labr *LabrCallerSession) ChainID() ([8]byte, error) {
	return _Labr.Contract.ChainID(&_Labr.CallOpts)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_Labr *LabrCaller) DepositCounts(opts *bind.CallOpts, arg0 [8]byte) (uint64, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_depositCounts", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_Labr *LabrSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _Labr.Contract.DepositCounts(&_Labr.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_Labr *LabrCallerSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _Labr.Contract.DepositCounts(&_Labr.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_Labr *LabrCaller) DepositRecords(opts *bind.CallOpts, arg0 uint64, arg1 [8]byte) ([]byte, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_Labr *LabrSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _Labr.Contract.DepositRecords(&_Labr.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_Labr *LabrCallerSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _Labr.Contract.DepositRecords(&_Labr.CallOpts, arg0, arg1)
}

// DexAddress is a free data retrieval call binding the contract method 0x98bdd27a.
//
// Solidity: function _dexAddress() view returns(address)
func (_Labr *LabrCaller) DexAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_dexAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DexAddress is a free data retrieval call binding the contract method 0x98bdd27a.
//
// Solidity: function _dexAddress() view returns(address)
func (_Labr *LabrSession) DexAddress() (common.Address, error) {
	return _Labr.Contract.DexAddress(&_Labr.CallOpts)
}

// DexAddress is a free data retrieval call binding the contract method 0x98bdd27a.
//
// Solidity: function _dexAddress() view returns(address)
func (_Labr *LabrCallerSession) DexAddress() (common.Address, error) {
	return _Labr.Contract.DexAddress(&_Labr.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Labr *LabrCaller) Expiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Labr *LabrSession) Expiry() (*big.Int, error) {
	return _Labr.Contract.Expiry(&_Labr.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Labr *LabrCallerSession) Expiry() (*big.Int, error) {
	return _Labr.Contract.Expiry(&_Labr.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Labr *LabrCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Labr *LabrSession) Fee() (*big.Int, error) {
	return _Labr.Contract.Fee(&_Labr.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Labr *LabrCallerSession) Fee() (*big.Int, error) {
	return _Labr.Contract.Fee(&_Labr.CallOpts)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x86c99f0e.
//
// Solidity: function _hasVotedOnProposal(bytes32 , bytes32 , address ) view returns(bool)
func (_Labr *LabrCaller) HasVotedOnProposal(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_hasVotedOnProposal", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x86c99f0e.
//
// Solidity: function _hasVotedOnProposal(bytes32 , bytes32 , address ) view returns(bool)
func (_Labr *LabrSession) HasVotedOnProposal(arg0 [32]byte, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _Labr.Contract.HasVotedOnProposal(&_Labr.CallOpts, arg0, arg1, arg2)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x86c99f0e.
//
// Solidity: function _hasVotedOnProposal(bytes32 , bytes32 , address ) view returns(bool)
func (_Labr *LabrCallerSession) HasVotedOnProposal(arg0 [32]byte, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _Labr.Contract.HasVotedOnProposal(&_Labr.CallOpts, arg0, arg1, arg2)
}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_Labr *LabrCaller) NativeResourceID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_nativeResourceID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_Labr *LabrSession) NativeResourceID() ([32]byte, error) {
	return _Labr.Contract.NativeResourceID(&_Labr.CallOpts)
}

// NativeResourceID is a free data retrieval call binding the contract method 0xa21b952e.
//
// Solidity: function _nativeResourceID() view returns(bytes32)
func (_Labr *LabrCallerSession) NativeResourceID() ([32]byte, error) {
	return _Labr.Contract.NativeResourceID(&_Labr.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x7bda741c.
//
// Solidity: function _proposals(bytes32 , bytes32 ) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Labr *LabrCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (IBridgeProposal, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_proposals", arg0, arg1)

	if err != nil {
		return *new(IBridgeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeProposal)).(*IBridgeProposal)

	return out0, err

}

// Proposals is a free data retrieval call binding the contract method 0x7bda741c.
//
// Solidity: function _proposals(bytes32 , bytes32 ) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Labr *LabrSession) Proposals(arg0 [32]byte, arg1 [32]byte) (IBridgeProposal, error) {
	return _Labr.Contract.Proposals(&_Labr.CallOpts, arg0, arg1)
}

// Proposals is a free data retrieval call binding the contract method 0x7bda741c.
//
// Solidity: function _proposals(bytes32 , bytes32 ) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Labr *LabrCallerSession) Proposals(arg0 [32]byte, arg1 [32]byte) (IBridgeProposal, error) {
	return _Labr.Contract.Proposals(&_Labr.CallOpts, arg0, arg1)
}

// RelayerHubAddress is a free data retrieval call binding the contract method 0xb6d97623.
//
// Solidity: function _relayerHubAddress() view returns(address)
func (_Labr *LabrCaller) RelayerHubAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_relayerHubAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayerHubAddress is a free data retrieval call binding the contract method 0xb6d97623.
//
// Solidity: function _relayerHubAddress() view returns(address)
func (_Labr *LabrSession) RelayerHubAddress() (common.Address, error) {
	return _Labr.Contract.RelayerHubAddress(&_Labr.CallOpts)
}

// RelayerHubAddress is a free data retrieval call binding the contract method 0xb6d97623.
//
// Solidity: function _relayerHubAddress() view returns(address)
func (_Labr *LabrCallerSession) RelayerHubAddress() (common.Address, error) {
	return _Labr.Contract.RelayerHubAddress(&_Labr.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Labr *LabrCaller) RelayerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_relayerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Labr *LabrSession) RelayerThreshold() (*big.Int, error) {
	return _Labr.Contract.RelayerThreshold(&_Labr.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Labr *LabrCallerSession) RelayerThreshold() (*big.Int, error) {
	return _Labr.Contract.RelayerThreshold(&_Labr.CallOpts)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Labr *LabrCaller) ResourceIDToHandlerAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_resourceIDToHandlerAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Labr *LabrSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Labr.Contract.ResourceIDToHandlerAddress(&_Labr.CallOpts, arg0)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Labr *LabrCallerSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Labr.Contract.ResourceIDToHandlerAddress(&_Labr.CallOpts, arg0)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Labr *LabrCaller) TotalProposals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "_totalProposals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Labr *LabrSession) TotalProposals() (*big.Int, error) {
	return _Labr.Contract.TotalProposals(&_Labr.CallOpts)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Labr *LabrCallerSession) TotalProposals() (*big.Int, error) {
	return _Labr.Contract.TotalProposals(&_Labr.CallOpts)
}

// ExtraLATransferred is a free data retrieval call binding the contract method 0x51fc8c58.
//
// Solidity: function extraLATransferred(bytes32 , bytes32 ) view returns(bool)
func (_Labr *LabrCaller) ExtraLATransferred(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "extraLATransferred", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExtraLATransferred is a free data retrieval call binding the contract method 0x51fc8c58.
//
// Solidity: function extraLATransferred(bytes32 , bytes32 ) view returns(bool)
func (_Labr *LabrSession) ExtraLATransferred(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _Labr.Contract.ExtraLATransferred(&_Labr.CallOpts, arg0, arg1)
}

// ExtraLATransferred is a free data retrieval call binding the contract method 0x51fc8c58.
//
// Solidity: function extraLATransferred(bytes32 , bytes32 ) view returns(bool)
func (_Labr *LabrCallerSession) ExtraLATransferred(arg0 [32]byte, arg1 [32]byte) (bool, error) {
	return _Labr.Contract.ExtraLATransferred(&_Labr.CallOpts, arg0, arg1)
}

// GetExtraLATransferred is a free data retrieval call binding the contract method 0xd4ae814c.
//
// Solidity: function getExtraLATransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) view returns()
func (_Labr *LabrCaller) GetExtraLATransferred(opts *bind.CallOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) error {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "getExtraLATransferred", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)

	if err != nil {
		return err
	}

	return err

}

// GetExtraLATransferred is a free data retrieval call binding the contract method 0xd4ae814c.
//
// Solidity: function getExtraLATransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) view returns()
func (_Labr *LabrSession) GetExtraLATransferred(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) error {
	return _Labr.Contract.GetExtraLATransferred(&_Labr.CallOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// GetExtraLATransferred is a free data retrieval call binding the contract method 0xd4ae814c.
//
// Solidity: function getExtraLATransferred(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) view returns()
func (_Labr *LabrCallerSession) GetExtraLATransferred(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) error {
	return _Labr.Contract.GetExtraLATransferred(&_Labr.CallOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount)
}

// GetProposal is a free data retrieval call binding the contract method 0xa241c0dd.
//
// Solidity: function getProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Labr *LabrCaller) GetProposal(opts *bind.CallOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (IBridgeProposal, error) {
	var out []interface{}
	err := _Labr.contract.Call(opts, &out, "getProposal", originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)

	if err != nil {
		return *new(IBridgeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeProposal)).(*IBridgeProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xa241c0dd.
//
// Solidity: function getProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Labr *LabrSession) GetProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (IBridgeProposal, error) {
	return _Labr.Contract.GetProposal(&_Labr.CallOpts, originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// GetProposal is a free data retrieval call binding the contract method 0xa241c0dd.
//
// Solidity: function getProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Labr *LabrCallerSession) GetProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (IBridgeProposal, error) {
	return _Labr.Contract.GetProposal(&_Labr.CallOpts, originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_Labr *LabrTransactor) AdminCollectFees(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminCollectFees", recipient, amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_Labr *LabrSession) AdminCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminCollectFees(&_Labr.TransactOpts, recipient, amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0x7a7eed77.
//
// Solidity: function adminCollectFees(address recipient, uint256 amount) returns()
func (_Labr *LabrTransactorSession) AdminCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminCollectFees(&_Labr.TransactOpts, recipient, amount)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Labr *LabrTransactor) AdminPauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminPauseTransfers")
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Labr *LabrSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Labr.Contract.AdminPauseTransfers(&_Labr.TransactOpts)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Labr *LabrTransactorSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Labr.Contract.AdminPauseTransfers(&_Labr.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Labr *LabrTransactor) AdminUnpauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminUnpauseTransfers")
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Labr *LabrSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Labr.Contract.AdminUnpauseTransfers(&_Labr.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Labr *LabrTransactorSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Labr.Contract.AdminUnpauseTransfers(&_Labr.TransactOpts)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Labr *LabrTransactor) AdminWithdraw(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "adminWithdraw", handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Labr *LabrSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminWithdraw(&_Labr.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Labr *LabrTransactorSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.AdminWithdraw(&_Labr.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// ApproveSpending is a paid mutator transaction binding the contract method 0xde771b31.
//
// Solidity: function approveSpending(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_Labr *LabrTransactor) ApproveSpending(opts *bind.TransactOpts, resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "approveSpending", resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// ApproveSpending is a paid mutator transaction binding the contract method 0xde771b31.
//
// Solidity: function approveSpending(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_Labr *LabrSession) ApproveSpending(resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.ApproveSpending(&_Labr.TransactOpts, resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// ApproveSpending is a paid mutator transaction binding the contract method 0xde771b31.
//
// Solidity: function approveSpending(bytes32 resourceIDOwner, bytes32 resourceIDSpender, uint256 amountOrTokenID) returns()
func (_Labr *LabrTransactorSession) ApproveSpending(resourceIDOwner [32]byte, resourceIDSpender [32]byte, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.ApproveSpending(&_Labr.TransactOpts, resourceIDOwner, resourceIDSpender, amountOrTokenID)
}

// BackendSrvCollectFees is a paid mutator transaction binding the contract method 0x4dd623fe.
//
// Solidity: function backendSrvCollectFees(address recipient, uint256 amount) returns()
func (_Labr *LabrTransactor) BackendSrvCollectFees(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "backendSrvCollectFees", recipient, amount)
}

// BackendSrvCollectFees is a paid mutator transaction binding the contract method 0x4dd623fe.
//
// Solidity: function backendSrvCollectFees(address recipient, uint256 amount) returns()
func (_Labr *LabrSession) BackendSrvCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.BackendSrvCollectFees(&_Labr.TransactOpts, recipient, amount)
}

// BackendSrvCollectFees is a paid mutator transaction binding the contract method 0x4dd623fe.
//
// Solidity: function backendSrvCollectFees(address recipient, uint256 amount) returns()
func (_Labr *LabrTransactorSession) BackendSrvCollectFees(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.BackendSrvCollectFees(&_Labr.TransactOpts, recipient, amount)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x8d39b457.
//
// Solidity: function cancelProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) returns()
func (_Labr *LabrTransactor) CancelProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "cancelProposal", originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x8d39b457.
//
// Solidity: function cancelProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) returns()
func (_Labr *LabrSession) CancelProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (*types.Transaction, error) {
	return _Labr.Contract.CancelProposal(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x8d39b457.
//
// Solidity: function cancelProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) returns()
func (_Labr *LabrTransactorSession) CancelProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (*types.Transaction, error) {
	return _Labr.Contract.CancelProposal(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, recipientAddress, amount, resourceID)
}

// ChangeDEXAddress is a paid mutator transaction binding the contract method 0x415d86d0.
//
// Solidity: function changeDEXAddress(address newDEX) returns()
func (_Labr *LabrTransactor) ChangeDEXAddress(opts *bind.TransactOpts, newDEX common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "changeDEXAddress", newDEX)
}

// ChangeDEXAddress is a paid mutator transaction binding the contract method 0x415d86d0.
//
// Solidity: function changeDEXAddress(address newDEX) returns()
func (_Labr *LabrSession) ChangeDEXAddress(newDEX common.Address) (*types.Transaction, error) {
	return _Labr.Contract.ChangeDEXAddress(&_Labr.TransactOpts, newDEX)
}

// ChangeDEXAddress is a paid mutator transaction binding the contract method 0x415d86d0.
//
// Solidity: function changeDEXAddress(address newDEX) returns()
func (_Labr *LabrTransactorSession) ChangeDEXAddress(newDEX common.Address) (*types.Transaction, error) {
	return _Labr.Contract.ChangeDEXAddress(&_Labr.TransactOpts, newDEX)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_Labr *LabrTransactor) ChangeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "changeFee", newFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_Labr *LabrSession) ChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.ChangeFee(&_Labr.TransactOpts, newFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 newFee) returns()
func (_Labr *LabrTransactorSession) ChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.ChangeFee(&_Labr.TransactOpts, newFee)
}

// ChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x2bceaea6.
//
// Solidity: function changeRelayerThreshold(uint256 newThreshold) returns()
func (_Labr *LabrTransactor) ChangeRelayerThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "changeRelayerThreshold", newThreshold)
}

// ChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x2bceaea6.
//
// Solidity: function changeRelayerThreshold(uint256 newThreshold) returns()
func (_Labr *LabrSession) ChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.ChangeRelayerThreshold(&_Labr.TransactOpts, newThreshold)
}

// ChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x2bceaea6.
//
// Solidity: function changeRelayerThreshold(uint256 newThreshold) returns()
func (_Labr *LabrTransactorSession) ChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Labr.Contract.ChangeRelayerThreshold(&_Labr.TransactOpts, newThreshold)
}

// ChangeWETHAddress is a paid mutator transaction binding the contract method 0x5fca9717.
//
// Solidity: function changeWETHAddress(address newWETH) returns()
func (_Labr *LabrTransactor) ChangeWETHAddress(opts *bind.TransactOpts, newWETH common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "changeWETHAddress", newWETH)
}

// ChangeWETHAddress is a paid mutator transaction binding the contract method 0x5fca9717.
//
// Solidity: function changeWETHAddress(address newWETH) returns()
func (_Labr *LabrSession) ChangeWETHAddress(newWETH common.Address) (*types.Transaction, error) {
	return _Labr.Contract.ChangeWETHAddress(&_Labr.TransactOpts, newWETH)
}

// ChangeWETHAddress is a paid mutator transaction binding the contract method 0x5fca9717.
//
// Solidity: function changeWETHAddress(address newWETH) returns()
func (_Labr *LabrTransactorSession) ChangeWETHAddress(newWETH common.Address) (*types.Transaction, error) {
	return _Labr.Contract.ChangeWETHAddress(&_Labr.TransactOpts, newWETH)
}

// Deposit is a paid mutator transaction binding the contract method 0x98ce7a18.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, bytes signature, bytes params) payable returns()
func (_Labr *LabrTransactor) Deposit(opts *bind.TransactOpts, destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, signature []byte, params []byte) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "deposit", destinationChainID, resourceID, amount, recipientAddress, signature, params)
}

// Deposit is a paid mutator transaction binding the contract method 0x98ce7a18.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, bytes signature, bytes params) payable returns()
func (_Labr *LabrSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, signature []byte, params []byte) (*types.Transaction, error) {
	return _Labr.Contract.Deposit(&_Labr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress, signature, params)
}

// Deposit is a paid mutator transaction binding the contract method 0x98ce7a18.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress, bytes signature, bytes params) payable returns()
func (_Labr *LabrTransactorSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address, signature []byte, params []byte) (*types.Transaction, error) {
	return _Labr.Contract.Deposit(&_Labr.TransactOpts, destinationChainID, resourceID, amount, recipientAddress, signature, params)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Labr *LabrTransactor) DepositFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "depositFunds")
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Labr *LabrSession) DepositFunds() (*types.Transaction, error) {
	return _Labr.Contract.DepositFunds(&_Labr.TransactOpts)
}

// DepositFunds is a paid mutator transaction binding the contract method 0xe2c41dbc.
//
// Solidity: function depositFunds() payable returns()
func (_Labr *LabrTransactorSession) DepositFunds() (*types.Transaction, error) {
	return _Labr.Contract.DepositFunds(&_Labr.TransactOpts)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_Labr *LabrTransactor) ExecuteProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "executeProposal", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_Labr *LabrSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Labr.Contract.ExecuteProposal(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x8d225943.
//
// Solidity: function executeProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_Labr *LabrTransactorSession) ExecuteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Labr.Contract.ExecuteProposal(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x37d8d3a8.
//
// Solidity: function initialize(bytes8 chainID_, uint256 relayerThreshold_, uint256 fee_, uint256 expiry_, address ownerAddress, address initBackendSrvAddress_, address initBalancerAddress_) returns()
func (_Labr *LabrTransactor) Initialize(opts *bind.TransactOpts, chainID_ [8]byte, relayerThreshold_ *big.Int, fee_ *big.Int, expiry_ *big.Int, ownerAddress common.Address, initBackendSrvAddress_ common.Address, initBalancerAddress_ common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "initialize", chainID_, relayerThreshold_, fee_, expiry_, ownerAddress, initBackendSrvAddress_, initBalancerAddress_)
}

// Initialize is a paid mutator transaction binding the contract method 0x37d8d3a8.
//
// Solidity: function initialize(bytes8 chainID_, uint256 relayerThreshold_, uint256 fee_, uint256 expiry_, address ownerAddress, address initBackendSrvAddress_, address initBalancerAddress_) returns()
func (_Labr *LabrSession) Initialize(chainID_ [8]byte, relayerThreshold_ *big.Int, fee_ *big.Int, expiry_ *big.Int, ownerAddress common.Address, initBackendSrvAddress_ common.Address, initBalancerAddress_ common.Address) (*types.Transaction, error) {
	return _Labr.Contract.Initialize(&_Labr.TransactOpts, chainID_, relayerThreshold_, fee_, expiry_, ownerAddress, initBackendSrvAddress_, initBalancerAddress_)
}

// Initialize is a paid mutator transaction binding the contract method 0x37d8d3a8.
//
// Solidity: function initialize(bytes8 chainID_, uint256 relayerThreshold_, uint256 fee_, uint256 expiry_, address ownerAddress, address initBackendSrvAddress_, address initBalancerAddress_) returns()
func (_Labr *LabrTransactorSession) Initialize(chainID_ [8]byte, relayerThreshold_ *big.Int, fee_ *big.Int, expiry_ *big.Int, ownerAddress common.Address, initBackendSrvAddress_ common.Address, initBalancerAddress_ common.Address) (*types.Transaction, error) {
	return _Labr.Contract.Initialize(&_Labr.TransactOpts, chainID_, relayerThreshold_, fee_, expiry_, ownerAddress, initBackendSrvAddress_, initBalancerAddress_)
}

// RelayerCollectReward is a paid mutator transaction binding the contract method 0xf408da91.
//
// Solidity: function relayerCollectReward() returns()
func (_Labr *LabrTransactor) RelayerCollectReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "relayerCollectReward")
}

// RelayerCollectReward is a paid mutator transaction binding the contract method 0xf408da91.
//
// Solidity: function relayerCollectReward() returns()
func (_Labr *LabrSession) RelayerCollectReward() (*types.Transaction, error) {
	return _Labr.Contract.RelayerCollectReward(&_Labr.TransactOpts)
}

// RelayerCollectReward is a paid mutator transaction binding the contract method 0xf408da91.
//
// Solidity: function relayerCollectReward() returns()
func (_Labr *LabrTransactorSession) RelayerCollectReward() (*types.Transaction, error) {
	return _Labr.Contract.RelayerCollectReward(&_Labr.TransactOpts)
}

// SetBackendSrvAddress is a paid mutator transaction binding the contract method 0x142852b5.
//
// Solidity: function setBackendSrvAddress(address newBackendSrv) returns()
func (_Labr *LabrTransactor) SetBackendSrvAddress(opts *bind.TransactOpts, newBackendSrv common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "setBackendSrvAddress", newBackendSrv)
}

// SetBackendSrvAddress is a paid mutator transaction binding the contract method 0x142852b5.
//
// Solidity: function setBackendSrvAddress(address newBackendSrv) returns()
func (_Labr *LabrSession) SetBackendSrvAddress(newBackendSrv common.Address) (*types.Transaction, error) {
	return _Labr.Contract.SetBackendSrvAddress(&_Labr.TransactOpts, newBackendSrv)
}

// SetBackendSrvAddress is a paid mutator transaction binding the contract method 0x142852b5.
//
// Solidity: function setBackendSrvAddress(address newBackendSrv) returns()
func (_Labr *LabrTransactorSession) SetBackendSrvAddress(newBackendSrv common.Address) (*types.Transaction, error) {
	return _Labr.Contract.SetBackendSrvAddress(&_Labr.TransactOpts, newBackendSrv)
}

// SetBalancerAddress is a paid mutator transaction binding the contract method 0xb4019dee.
//
// Solidity: function setBalancerAddress(address newBalancer) returns()
func (_Labr *LabrTransactor) SetBalancerAddress(opts *bind.TransactOpts, newBalancer common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "setBalancerAddress", newBalancer)
}

// SetBalancerAddress is a paid mutator transaction binding the contract method 0xb4019dee.
//
// Solidity: function setBalancerAddress(address newBalancer) returns()
func (_Labr *LabrSession) SetBalancerAddress(newBalancer common.Address) (*types.Transaction, error) {
	return _Labr.Contract.SetBalancerAddress(&_Labr.TransactOpts, newBalancer)
}

// SetBalancerAddress is a paid mutator transaction binding the contract method 0xb4019dee.
//
// Solidity: function setBalancerAddress(address newBalancer) returns()
func (_Labr *LabrTransactorSession) SetBalancerAddress(newBalancer common.Address) (*types.Transaction, error) {
	return _Labr.Contract.SetBalancerAddress(&_Labr.TransactOpts, newBalancer)
}

// SetBridgeDelegate is a paid mutator transaction binding the contract method 0xe0325208.
//
// Solidity: function setBridgeDelegate(address newBridgeDelegate) returns()
func (_Labr *LabrTransactor) SetBridgeDelegate(opts *bind.TransactOpts, newBridgeDelegate common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "setBridgeDelegate", newBridgeDelegate)
}

// SetBridgeDelegate is a paid mutator transaction binding the contract method 0xe0325208.
//
// Solidity: function setBridgeDelegate(address newBridgeDelegate) returns()
func (_Labr *LabrSession) SetBridgeDelegate(newBridgeDelegate common.Address) (*types.Transaction, error) {
	return _Labr.Contract.SetBridgeDelegate(&_Labr.TransactOpts, newBridgeDelegate)
}

// SetBridgeDelegate is a paid mutator transaction binding the contract method 0xe0325208.
//
// Solidity: function setBridgeDelegate(address newBridgeDelegate) returns()
func (_Labr *LabrTransactorSession) SetBridgeDelegate(newBridgeDelegate common.Address) (*types.Transaction, error) {
	return _Labr.Contract.SetBridgeDelegate(&_Labr.TransactOpts, newBridgeDelegate)
}

// SetBurnable is a paid mutator transaction binding the contract method 0xdf0fc133.
//
// Solidity: function setBurnable(address handlerAddress, address tokenAddress) returns()
func (_Labr *LabrTransactor) SetBurnable(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "setBurnable", handlerAddress, tokenAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0xdf0fc133.
//
// Solidity: function setBurnable(address handlerAddress, address tokenAddress) returns()
func (_Labr *LabrSession) SetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Labr.Contract.SetBurnable(&_Labr.TransactOpts, handlerAddress, tokenAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0xdf0fc133.
//
// Solidity: function setBurnable(address handlerAddress, address tokenAddress) returns()
func (_Labr *LabrTransactorSession) SetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Labr.Contract.SetBurnable(&_Labr.TransactOpts, handlerAddress, tokenAddress)
}

// SetNativeResourceID is a paid mutator transaction binding the contract method 0x21016ed8.
//
// Solidity: function setNativeResourceID(bytes32 resourceID) returns()
func (_Labr *LabrTransactor) SetNativeResourceID(opts *bind.TransactOpts, resourceID [32]byte) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "setNativeResourceID", resourceID)
}

// SetNativeResourceID is a paid mutator transaction binding the contract method 0x21016ed8.
//
// Solidity: function setNativeResourceID(bytes32 resourceID) returns()
func (_Labr *LabrSession) SetNativeResourceID(resourceID [32]byte) (*types.Transaction, error) {
	return _Labr.Contract.SetNativeResourceID(&_Labr.TransactOpts, resourceID)
}

// SetNativeResourceID is a paid mutator transaction binding the contract method 0x21016ed8.
//
// Solidity: function setNativeResourceID(bytes32 resourceID) returns()
func (_Labr *LabrTransactorSession) SetNativeResourceID(resourceID [32]byte) (*types.Transaction, error) {
	return _Labr.Contract.SetNativeResourceID(&_Labr.TransactOpts, resourceID)
}

// SetRelayerHub is a paid mutator transaction binding the contract method 0xce9e014a.
//
// Solidity: function setRelayerHub(address newRelayerHub) returns()
func (_Labr *LabrTransactor) SetRelayerHub(opts *bind.TransactOpts, newRelayerHub common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "setRelayerHub", newRelayerHub)
}

// SetRelayerHub is a paid mutator transaction binding the contract method 0xce9e014a.
//
// Solidity: function setRelayerHub(address newRelayerHub) returns()
func (_Labr *LabrSession) SetRelayerHub(newRelayerHub common.Address) (*types.Transaction, error) {
	return _Labr.Contract.SetRelayerHub(&_Labr.TransactOpts, newRelayerHub)
}

// SetRelayerHub is a paid mutator transaction binding the contract method 0xce9e014a.
//
// Solidity: function setRelayerHub(address newRelayerHub) returns()
func (_Labr *LabrTransactorSession) SetRelayerHub(newRelayerHub common.Address) (*types.Transaction, error) {
	return _Labr.Contract.SetRelayerHub(&_Labr.TransactOpts, newRelayerHub)
}

// SetResource is a paid mutator transaction binding the contract method 0xa737be4f.
//
// Solidity: function setResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Labr *LabrTransactor) SetResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "setResource", handlerAddress, resourceID, tokenAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xa737be4f.
//
// Solidity: function setResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Labr *LabrSession) SetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Labr.Contract.SetResource(&_Labr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xa737be4f.
//
// Solidity: function setResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Labr *LabrTransactorSession) SetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Labr.Contract.SetResource(&_Labr.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// TransferExtraFee is a paid mutator transaction binding the contract method 0xe70a54ce.
//
// Solidity: function transferExtraFee(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount, bytes params) returns()
func (_Labr *LabrTransactor) TransferExtraFee(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipient common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "transferExtraFee", originChainID, destinationChainID, depositNonce, resourceID, recipient, amount, params)
}

// TransferExtraFee is a paid mutator transaction binding the contract method 0xe70a54ce.
//
// Solidity: function transferExtraFee(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount, bytes params) returns()
func (_Labr *LabrSession) TransferExtraFee(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipient common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Labr.Contract.TransferExtraFee(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipient, amount, params)
}

// TransferExtraFee is a paid mutator transaction binding the contract method 0xe70a54ce.
//
// Solidity: function transferExtraFee(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipient, uint256 amount, bytes params) returns()
func (_Labr *LabrTransactorSession) TransferExtraFee(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipient common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Labr.Contract.TransferExtraFee(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipient, amount, params)
}

// UpdateExternalTx is a paid mutator transaction binding the contract method 0x00ecbc2c.
//
// Solidity: function updateExternalTx(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 proposalAmount, uint256 amountToTransfer, bytes params, uint8 status) returns()
func (_Labr *LabrTransactor) UpdateExternalTx(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, proposalAmount *big.Int, amountToTransfer *big.Int, params []byte, status uint8) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "updateExternalTx", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, proposalAmount, amountToTransfer, params, status)
}

// UpdateExternalTx is a paid mutator transaction binding the contract method 0x00ecbc2c.
//
// Solidity: function updateExternalTx(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 proposalAmount, uint256 amountToTransfer, bytes params, uint8 status) returns()
func (_Labr *LabrSession) UpdateExternalTx(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, proposalAmount *big.Int, amountToTransfer *big.Int, params []byte, status uint8) (*types.Transaction, error) {
	return _Labr.Contract.UpdateExternalTx(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, proposalAmount, amountToTransfer, params, status)
}

// UpdateExternalTx is a paid mutator transaction binding the contract method 0x00ecbc2c.
//
// Solidity: function updateExternalTx(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 proposalAmount, uint256 amountToTransfer, bytes params, uint8 status) returns()
func (_Labr *LabrTransactorSession) UpdateExternalTx(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, proposalAmount *big.Int, amountToTransfer *big.Int, params []byte, status uint8) (*types.Transaction, error) {
	return _Labr.Contract.UpdateExternalTx(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, proposalAmount, amountToTransfer, params, status)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x369ab5ac.
//
// Solidity: function voteProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_Labr *LabrTransactor) VoteProposal(opts *bind.TransactOpts, originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Labr.contract.Transact(opts, "voteProposal", originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x369ab5ac.
//
// Solidity: function voteProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_Labr *LabrSession) VoteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Labr.Contract.VoteProposal(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x369ab5ac.
//
// Solidity: function voteProposal(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, bytes params) returns()
func (_Labr *LabrTransactorSession) VoteProposal(originChainID [8]byte, destinationChainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _Labr.Contract.VoteProposal(&_Labr.TransactOpts, originChainID, destinationChainID, depositNonce, resourceID, recipientAddress, amount, params)
}

// LabrDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Labr contract.
type LabrDepositIterator struct {
	Event *LabrDeposit // Event containing the contract specifics and raw log

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
func (it *LabrDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrDeposit)
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
		it.Event = new(LabrDeposit)
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
func (it *LabrDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrDeposit represents a Deposit event raised by the Labr contract.
type LabrDeposit struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	ResourceID         [32]byte
	DepositNonce       uint64
	Depositor          common.Address
	RecipientAddress   common.Address
	TokenAddress       common.Address
	Amount             *big.Int
	Params             []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x3cdf0bc4e2723a2132944314ba37022e8f01ee627cbbc3c834065f80f8b2b04f.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 destinationChainID, bytes32 resourceID, uint64 depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes params)
func (_Labr *LabrFilterer) FilterDeposit(opts *bind.FilterOpts) (*LabrDepositIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &LabrDepositIterator{contract: _Labr.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x3cdf0bc4e2723a2132944314ba37022e8f01ee627cbbc3c834065f80f8b2b04f.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 destinationChainID, bytes32 resourceID, uint64 depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes params)
func (_Labr *LabrFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LabrDeposit) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrDeposit)
				if err := _Labr.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x3cdf0bc4e2723a2132944314ba37022e8f01ee627cbbc3c834065f80f8b2b04f.
//
// Solidity: event Deposit(bytes8 originChainID, bytes8 destinationChainID, bytes32 resourceID, uint64 depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount, bytes params)
func (_Labr *LabrFilterer) ParseDeposit(log types.Log) (*LabrDeposit, error) {
	event := new(LabrDeposit)
	if err := _Labr.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrExtraFeeSuppliedIterator is returned from FilterExtraFeeSupplied and is used to iterate over the raw logs and unpacked data for ExtraFeeSupplied events raised by the Labr contract.
type LabrExtraFeeSuppliedIterator struct {
	Event *LabrExtraFeeSupplied // Event containing the contract specifics and raw log

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
func (it *LabrExtraFeeSuppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrExtraFeeSupplied)
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
		it.Event = new(LabrExtraFeeSupplied)
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
func (it *LabrExtraFeeSuppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrExtraFeeSuppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrExtraFeeSupplied represents a ExtraFeeSupplied event raised by the Labr contract.
type LabrExtraFeeSupplied struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	DepositNonce       uint64
	ResourceID         [32]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	Status             uint8
	Params             []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExtraFeeSupplied is a free log retrieval operation binding the contract event 0xa111a4bf39fd61f7abcd239236bed67639dd6c74bf937b213b862b51397d65db.
//
// Solidity: event ExtraFeeSupplied(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, uint8 status, bytes params)
func (_Labr *LabrFilterer) FilterExtraFeeSupplied(opts *bind.FilterOpts) (*LabrExtraFeeSuppliedIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "ExtraFeeSupplied")
	if err != nil {
		return nil, err
	}
	return &LabrExtraFeeSuppliedIterator{contract: _Labr.contract, event: "ExtraFeeSupplied", logs: logs, sub: sub}, nil
}

// WatchExtraFeeSupplied is a free log subscription operation binding the contract event 0xa111a4bf39fd61f7abcd239236bed67639dd6c74bf937b213b862b51397d65db.
//
// Solidity: event ExtraFeeSupplied(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, uint8 status, bytes params)
func (_Labr *LabrFilterer) WatchExtraFeeSupplied(opts *bind.WatchOpts, sink chan<- *LabrExtraFeeSupplied) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "ExtraFeeSupplied")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrExtraFeeSupplied)
				if err := _Labr.contract.UnpackLog(event, "ExtraFeeSupplied", log); err != nil {
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

// ParseExtraFeeSupplied is a log parse operation binding the contract event 0xa111a4bf39fd61f7abcd239236bed67639dd6c74bf937b213b862b51397d65db.
//
// Solidity: event ExtraFeeSupplied(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount, uint8 status, bytes params)
func (_Labr *LabrFilterer) ParseExtraFeeSupplied(log types.Log) (*LabrExtraFeeSupplied, error) {
	event := new(LabrExtraFeeSupplied)
	if err := _Labr.contract.UnpackLog(event, "ExtraFeeSupplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrProposalEventIterator is returned from FilterProposalEvent and is used to iterate over the raw logs and unpacked data for ProposalEvent events raised by the Labr contract.
type LabrProposalEventIterator struct {
	Event *LabrProposalEvent // Event containing the contract specifics and raw log

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
func (it *LabrProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrProposalEvent)
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
		it.Event = new(LabrProposalEvent)
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
func (it *LabrProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrProposalEvent represents a ProposalEvent event raised by the Labr contract.
type LabrProposalEvent struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	DepositNonce       uint64
	Status             uint8
	ResourceID         [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProposalEvent is a free log retrieval operation binding the contract event 0x98515ff66d46eef043e6e17beb65b19f71802dc829ff974ca92d66d61019286d.
//
// Solidity: event ProposalEvent(bytes8 originChainID, bytes8 destinationChainID, address recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_Labr *LabrFilterer) FilterProposalEvent(opts *bind.FilterOpts) (*LabrProposalEventIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "ProposalEvent")
	if err != nil {
		return nil, err
	}
	return &LabrProposalEventIterator{contract: _Labr.contract, event: "ProposalEvent", logs: logs, sub: sub}, nil
}

// WatchProposalEvent is a free log subscription operation binding the contract event 0x98515ff66d46eef043e6e17beb65b19f71802dc829ff974ca92d66d61019286d.
//
// Solidity: event ProposalEvent(bytes8 originChainID, bytes8 destinationChainID, address recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_Labr *LabrFilterer) WatchProposalEvent(opts *bind.WatchOpts, sink chan<- *LabrProposalEvent) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "ProposalEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrProposalEvent)
				if err := _Labr.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
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

// ParseProposalEvent is a log parse operation binding the contract event 0x98515ff66d46eef043e6e17beb65b19f71802dc829ff974ca92d66d61019286d.
//
// Solidity: event ProposalEvent(bytes8 originChainID, bytes8 destinationChainID, address recipientAddress, uint256 amount, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_Labr *LabrFilterer) ParseProposalEvent(log types.Log) (*LabrProposalEvent, error) {
	event := new(LabrProposalEvent)
	if err := _Labr.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrProposalVoteIterator is returned from FilterProposalVote and is used to iterate over the raw logs and unpacked data for ProposalVote events raised by the Labr contract.
type LabrProposalVoteIterator struct {
	Event *LabrProposalVote // Event containing the contract specifics and raw log

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
func (it *LabrProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrProposalVote)
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
		it.Event = new(LabrProposalVote)
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
func (it *LabrProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrProposalVote represents a ProposalVote event raised by the Labr contract.
type LabrProposalVote struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	DepositNonce       uint64
	Relayer            common.Address
	Status             uint8
	ResourceID         [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProposalVote is a free log retrieval operation binding the contract event 0xd2015884433ce339f6e95e46f15bbfa85ea1012aa20db6296527b5f96592b584.
//
// Solidity: event ProposalVote(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address relayer, uint8 status, bytes32 resourceID)
func (_Labr *LabrFilterer) FilterProposalVote(opts *bind.FilterOpts) (*LabrProposalVoteIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "ProposalVote")
	if err != nil {
		return nil, err
	}
	return &LabrProposalVoteIterator{contract: _Labr.contract, event: "ProposalVote", logs: logs, sub: sub}, nil
}

// WatchProposalVote is a free log subscription operation binding the contract event 0xd2015884433ce339f6e95e46f15bbfa85ea1012aa20db6296527b5f96592b584.
//
// Solidity: event ProposalVote(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address relayer, uint8 status, bytes32 resourceID)
func (_Labr *LabrFilterer) WatchProposalVote(opts *bind.WatchOpts, sink chan<- *LabrProposalVote) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "ProposalVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrProposalVote)
				if err := _Labr.contract.UnpackLog(event, "ProposalVote", log); err != nil {
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

// ParseProposalVote is a log parse operation binding the contract event 0xd2015884433ce339f6e95e46f15bbfa85ea1012aa20db6296527b5f96592b584.
//
// Solidity: event ProposalVote(bytes8 originChainID, bytes8 destinationChainID, uint64 depositNonce, address relayer, uint8 status, bytes32 resourceID)
func (_Labr *LabrFilterer) ParseProposalVote(log types.Log) (*LabrProposalVote, error) {
	event := new(LabrProposalVote)
	if err := _Labr.contract.UnpackLog(event, "ProposalVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrRelayerThresholdChangedIterator is returned from FilterRelayerThresholdChanged and is used to iterate over the raw logs and unpacked data for RelayerThresholdChanged events raised by the Labr contract.
type LabrRelayerThresholdChangedIterator struct {
	Event *LabrRelayerThresholdChanged // Event containing the contract specifics and raw log

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
func (it *LabrRelayerThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrRelayerThresholdChanged)
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
		it.Event = new(LabrRelayerThresholdChanged)
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
func (it *LabrRelayerThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrRelayerThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrRelayerThresholdChanged represents a RelayerThresholdChanged event raised by the Labr contract.
type LabrRelayerThresholdChanged struct {
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelayerThresholdChanged is a free log retrieval operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 newThreshold)
func (_Labr *LabrFilterer) FilterRelayerThresholdChanged(opts *bind.FilterOpts) (*LabrRelayerThresholdChangedIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "RelayerThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &LabrRelayerThresholdChangedIterator{contract: _Labr.contract, event: "RelayerThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchRelayerThresholdChanged is a free log subscription operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 newThreshold)
func (_Labr *LabrFilterer) WatchRelayerThresholdChanged(opts *bind.WatchOpts, sink chan<- *LabrRelayerThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "RelayerThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrRelayerThresholdChanged)
				if err := _Labr.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
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

// ParseRelayerThresholdChanged is a log parse operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 newThreshold)
func (_Labr *LabrFilterer) ParseRelayerThresholdChanged(log types.Log) (*LabrRelayerThresholdChanged, error) {
	event := new(LabrRelayerThresholdChanged)
	if err := _Labr.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LabrRewardCollectedIterator is returned from FilterRewardCollected and is used to iterate over the raw logs and unpacked data for RewardCollected events raised by the Labr contract.
type LabrRewardCollectedIterator struct {
	Event *LabrRewardCollected // Event containing the contract specifics and raw log

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
func (it *LabrRewardCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LabrRewardCollected)
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
		it.Event = new(LabrRewardCollected)
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
func (it *LabrRewardCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LabrRewardCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LabrRewardCollected represents a RewardCollected event raised by the Labr contract.
type LabrRewardCollected struct {
	Relayer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardCollected is a free log retrieval operation binding the contract event 0xe8354b169cd993d5cdfad1036a9a3f1ea7ed77e430bccb279200fd088243f595.
//
// Solidity: event RewardCollected(address relayer, uint256 amount)
func (_Labr *LabrFilterer) FilterRewardCollected(opts *bind.FilterOpts) (*LabrRewardCollectedIterator, error) {

	logs, sub, err := _Labr.contract.FilterLogs(opts, "RewardCollected")
	if err != nil {
		return nil, err
	}
	return &LabrRewardCollectedIterator{contract: _Labr.contract, event: "RewardCollected", logs: logs, sub: sub}, nil
}

// WatchRewardCollected is a free log subscription operation binding the contract event 0xe8354b169cd993d5cdfad1036a9a3f1ea7ed77e430bccb279200fd088243f595.
//
// Solidity: event RewardCollected(address relayer, uint256 amount)
func (_Labr *LabrFilterer) WatchRewardCollected(opts *bind.WatchOpts, sink chan<- *LabrRewardCollected) (event.Subscription, error) {

	logs, sub, err := _Labr.contract.WatchLogs(opts, "RewardCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LabrRewardCollected)
				if err := _Labr.contract.UnpackLog(event, "RewardCollected", log); err != nil {
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

// ParseRewardCollected is a log parse operation binding the contract event 0xe8354b169cd993d5cdfad1036a9a3f1ea7ed77e430bccb279200fd088243f595.
//
// Solidity: event RewardCollected(address relayer, uint256 amount)
func (_Labr *LabrFilterer) ParseRewardCollected(log types.Log) (*LabrRewardCollected, error) {
	event := new(LabrRewardCollected)
	if err := _Labr.contract.UnpackLog(event, "RewardCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
