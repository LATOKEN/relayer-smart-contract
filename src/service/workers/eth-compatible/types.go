package eth

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/LATOKEN/relayer-smart-contract.git/src/service/storage"
	"github.com/LATOKEN/relayer-smart-contract.git/src/service/workers/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

////
// EVENTS NAMES
////

const (
	// DepositEventName ...
	DepositEventName  = "Deposit"
	ProposalEventName = "ProposalEvent"
)

////
// EVENTS HASHES  | web3.utils.sha3('HTLT(types,...)');
////

var (
	// DepositEventHash
	DepositEventHash = common.HexToHash("0x3cdf0bc4e2723a2132944314ba37022e8f01ee627cbbc3c834065f80f8b2b04f")

	// ProposalEventHash
	ProposalEventHash = common.HexToHash("0x98515ff66d46eef043e6e17beb65b19f71802dc829ff974ca92d66d61019286d")
)

// ContractEvent ...
type ContractEvent interface {
	ToTxLog() *storage.TxLog
}

// DepositEvent represents a Deposit event raised by the Bridge.sol contract.
type DepositEvent struct {
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

type ProposalEvent struct {
	OriginChainID      [8]byte
	DestinationChainID [8]byte
	RecipientAddress   common.Address
	Amount             *big.Int
	DepositNonce       uint64
	Status             uint8
	ResourceID         [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// ToTxLog ...
func (ev DepositEvent) ToTxLog() *storage.TxLog {
	return &storage.TxLog{
		TxType:             storage.TxTypeDeposit,
		DestinationChainID: common.Bytes2Hex(ev.DestinationChainID[:]),
		OriginСhainID:      common.Bytes2Hex(ev.OriginChainID[:]),
		SwapID:             utils.CalcutateSwapID(common.Bytes2Hex(ev.OriginChainID[:]), common.Bytes2Hex(ev.DestinationChainID[:]), fmt.Sprint(ev.DepositNonce)),
		ResourceID:         common.Bytes2Hex(ev.ResourceID[:]),
		StepIndex:          ev.Params[0],
		DepositNonce:       ev.DepositNonce,
		SenderAddr:         ev.Depositor.Hex(),
		ReceiverAddr:       ev.RecipientAddress.Hex(),
		InTokenAddr:        ev.TokenAddress.Hex(),
		OutAmount:          ev.Amount.String(),
		SwapStatus:         storage.SwapStatusDepositInit,
	}
}

func (ev ProposalEvent) ToTxLog() *storage.TxLog {
	txlog := &storage.TxLog{
		TxType:             storage.TxTypeVote,
		DestinationChainID: common.Bytes2Hex(ev.DestinationChainID[:]),
		OriginСhainID:      common.Bytes2Hex(ev.OriginChainID[:]),
		SwapID:             utils.CalcutateSwapID(common.Bytes2Hex(ev.OriginChainID[:]), common.Bytes2Hex(ev.DestinationChainID[:]), fmt.Sprint(ev.DepositNonce)),
		ResourceID:         common.Bytes2Hex(ev.ResourceID[:]),
		DepositNonce:       ev.DepositNonce,
		ReceiverAddr:       ev.RecipientAddress.Hex(),
		OutAmount:          ev.Amount.String(),
		SwapStatus:         storage.SwapStatusClaimConfirmed,
	}

	if ev.Status == uint8(2) {
		txlog.TxType = storage.TxTypePassed
		txlog.SwapStatus = storage.SwapStatusPassedConfirmed
	} else if ev.Status == uint8(3) {
		txlog.TxType = storage.TxTypeSpend
		txlog.SwapStatus = storage.SwapStatusSpendConfirmed
	} else if ev.Status == uint8(4) {
		txlog.TxType = storage.TxTypeExpired
		txlog.SwapStatus = storage.SwapStatusExpiredConfirmed
	}

	return txlog
}

// ParseEvent ...
func (w *Erc20Worker) parseEvent(log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], DepositEventHash[:]) {
		if w.chainName == "LA" {
			return w.ParseLaDepositEvent(log)
		} else {
			return w.ParseEthDepositEvent(log)
		}
	}
	if bytes.Equal(log.Topics[0][:], ProposalEventHash[:]) {
		if w.chainName == "LA" {
			return w.ParseLaProposalEvent(log)
		} else {
			return w.ParseEthProposalEvent(log)
		}
	}
	return nil, nil
}

type Header struct {
	Hash       common.Hash    `json:"hash"`
	ParentHash common.Hash    `json:"parentHash"       gencodec:"required"`
	Time       hexutil.Uint64 `json:"timestamp"        gencodec:"required"`
	Number     hexutil.Uint64 `json:"number"					 gencodec:"required"`
}
