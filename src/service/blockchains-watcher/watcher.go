package watcher

import (
	"fmt"
	"strings"
	"time"

	"latoken/relayer-smart-contract/src/service/storage"
	"latoken/relayer-smart-contract/src/service/workers"

	"github.com/sirupsen/logrus"
)

// WatcherSRV ...
type WatcherSRV struct {
	logger  *logrus.Logger
	storage *storage.DataBase
	Workers map[string]workers.IWorker
}

// CreateNewWatcherSRV ...
func CreateNewWatcherSRV(logger *logrus.Logger, db *storage.DataBase, workers map[string]workers.IWorker) *WatcherSRV {
	return &WatcherSRV{
		logger:  logger,
		storage: db,
		Workers: workers,
	}
}

// Run ...
func (w *WatcherSRV) Run() {
	for name, worker := range w.Workers {
		w.logger.Infoln("watcher | worker name: ", name)
		go w.collector(worker, worker.GetFetchInterval(), worker.GetStartHeight())
		time.Sleep(100 * time.Millisecond)
	}

}

func (w *WatcherSRV) collector(worker workers.IWorker, threshold time.Duration, startHeight int64) {
	for {
		curBlockLog := w.storage.GetCurrentBlockLog(worker.GetChain())
		if curBlockLog.Height == 0 {
			w.logger.Warnf("%s current height: %d", worker.GetChain(), curBlockLog.Height)
		} else {
			w.logger.Infof("%s current height: %d", worker.GetChain(), curBlockLog.Height)
		}

		nextHeight := curBlockLog.Height + 1
		if curBlockLog.Height == 0 && startHeight != 0 {
			nextHeight = startHeight
		}

		if err := w.getBlock(worker, curBlockLog.Height, nextHeight, curBlockLog.BlockHash); err != nil {
			normalizedErr := strings.ToLower(err.Error())
			if strings.Contains(normalizedErr, "height must be less than or equal to the current blockchain height") ||
				strings.Contains(normalizedErr, "not found") ||
				strings.Contains(normalizedErr, "block number out of range") {
				w.logger.Infof("try to get ahead block, chain=%s, height=%d", worker.GetChain(), nextHeight)
			} else {
				w.logger.Error(normalizedErr)
			}
			time.Sleep(threshold)
		}

		time.Sleep(time.Second)
	}
}

func (w *WatcherSRV) getBlock(worker workers.IWorker, curHeight, nextHeight int64, curBlockHash string) error {
	blockAndTxLogs, err := worker.GetBlockAndTxs(nextHeight)
	if err != nil {
		return fmt.Errorf("get %s block info error, height=%d, err=%s", worker.GetChain(), nextHeight, err.Error())
	}

	parentHash := blockAndTxLogs.ParentBlockHash
	if curHeight != 0 && parentHash != curBlockHash {
		w.logger.Infof("delete %s block at height %d, hash=%s", worker.GetChain(), curHeight, curBlockHash)
		return w.storage.DeleteBlockAndTxs(worker.GetChain(), curHeight)
	}

	nextBlockLog := storage.BlockLog{
		Chain:      worker.GetChain(),
		BlockHash:  blockAndTxLogs.BlockHash,
		ParentHash: parentHash,
		Height:     blockAndTxLogs.Height,
		BlockTime:  blockAndTxLogs.BlockTime,
		Type:       storage.BlockTypeCurrent,
		CreateTime: time.Now().Unix(),
	}

	// put block header and block txs into database
	if err := w.storage.SaveBlockAndTxs(worker.GetChain(), &nextBlockLog, blockAndTxLogs.TxLogs); err != nil {
		return err
	}

	// update prev txs confirmed number(&& update_time)
	if err := w.storage.UpdateConfirmedNum(worker.GetChain(), nextBlockLog.Height); err != nil {
		return err
	}

	return nil
}
