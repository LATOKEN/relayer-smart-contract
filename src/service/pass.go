package rlr

import (
	"time"

	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/storage"
)

// emitChainSendClaim ...
func (r *RelayerSRV) emitChainSendPass() {
	for {
		swaps := r.storage.GetSwapsByTypeAndStatuses(
			[]storage.SwapStatus{storage.SwapStatusPassedConfirmed})

		for _, swap := range swaps {
			r.logger.Info("updating passed status")
			r.storage.UpdateSwapStatus(swap, storage.SwapStatusPassedSent, "")
			r.logger.Infof("send passed tx success | swap_ID=%s",
				swap.SwapID)
		}

		time.Sleep(2 * time.Second)
	}
}
