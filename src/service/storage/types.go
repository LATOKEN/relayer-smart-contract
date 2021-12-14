package storage

var (
	LaChain string = "LA"
	// POSChain - Polygon chain
	POSChain string = "POS"
	// BSCChain - Binance chain
	BSCChain string = "BSC"
	// BtcChain - Bitcoin chain
	BtcChain string = "BTC"
)

// BlockType ...
type BlockType string

const (
	BlockTypeCurrent BlockType = "CURRENT"
	BlockTypeParent  BlockType = "PARENT"
)

// TxStatus ...
type TxStatus string

const (
	TxSentStatusInit     TxStatus = "INIT"
	TxSentStatusNotFound TxStatus = "NOT_FOUND"
	TxSentStatusPending  TxStatus = "PENDING"
	TxSentStatusFailed   TxStatus = "FAILED"
	TxSentStatusSuccess  TxStatus = "SUCCESS"
	TxSentStatusLost     TxStatus = "LOST"
)

//
type TxType string

const (
	TxTypeVote    TxType = "VOTE"
	TxTypeDeposit TxType = "DEPOSIT"
	TxTypePassed  TxType = "PASSED"
	TxTypeSpend   TxType = "SPEND"
)

type SwapType string

const (
	SwapTypeBind   SwapType = "BIND"
	SwapTypeUnbind SwapType = "UNBIND"
)

type SwapStatus string

const (
	// CLAIM
	SwapStatusClaimSent       SwapStatus = "CLAIM_SENT"
	SwapStatusClaimConfirmed  SwapStatus = "CLAIM_CONFIRMED"
	SwapStatusClaimSentFailed SwapStatus = "CLAIM_SENT_FAILED"

	//Deposit
	SwapStatusDepositConfirmed SwapStatus = "DEPOSIT_CONFIRMED"
	SwapStatusDepositFailed    SwapStatus = "DEPOSIT_FAILED"
	SwapStatusPassedConfirmed  SwapStatus = "PASSED_CONFIRMED"
	SwapStatusSpendSent        SwapStatus = "SPEND_SENT"
	SwapStatusSpendConfirmed   SwapStatus = "SPEND_CONFIRMED"

	SwapStatusRejected SwapStatus = "REJECTED"
)

// TxLogStatus ...
type TxLogStatus string

const (
	TxStatusInit      TxLogStatus = "INIT"
	TxStatusConfirmed TxLogStatus = "CONFIRMED"
)