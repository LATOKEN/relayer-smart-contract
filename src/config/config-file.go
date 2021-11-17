package config

import (
	"fmt"
	"math/big"
	"strings"

	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/models"
	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service/storage"

	"github.com/ethereum/go-ethereum/common"
)

// Reads Service params from config.json
func (v *viperConfig) ReadServiceConfig() string {
	return fmt.Sprintf("%s:%s", v.GetString("service.host"), v.GetString("service.port"))
}

// ReadLachainConfig reads lachain chain params from config.json
func (v *viperConfig) ReadLachainConfig() *models.WorkerConfig {
	return v.readWorkerConfig(storage.LaChain)
}

// ReadEthWorkerConfig reads ethereum chain worker params from config.json
func (v *viperConfig) ReadWorkersConfig() (eth *models.WorkerConfig, btc *models.WorkerConfig) {
	return v.readWorkerConfig(storage.EthChain), v.readWorkerConfig(storage.BtcChain)
}

// readETHWorkerConfig reads ethereum chain worker params from config.json
func (v *viperConfig) readWorkerConfig(name string) *models.WorkerConfig {
	return &models.WorkerConfig{
		NetworkType:       v.GetString(fmt.Sprintf("workers.%s.type", name)),
		ChainName:         strings.ToUpper(name),
		ChainID:           v.GetInt64(fmt.Sprintf("workers.%s.chain_id", name)),
		User:              v.GetString(fmt.Sprintf("workers.%s.user", name)),
		Password:          v.GetString(fmt.Sprintf("workers.%s.password", name)),
		WorkerAddr:        common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.worker_addr", name))),
		PrivateKey:        v.GetString(fmt.Sprintf("workers.%s.private_key", name)),
		Provider:          v.GetString(fmt.Sprintf("workers.%s.provider", name)),
		ContractAddr:      common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.contract_addr", name))),
		TokenContractAddr: common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.token_addr", name))),
		GasLimit:          v.GetInt64(fmt.Sprintf("workers.%s.gas_limit", name)),
		GasPrice:          big.NewInt(v.GetInt64(fmt.Sprintf("workers.%s.gas_price", name))),
		FetchInterval:     v.GetInt64(fmt.Sprintf("workers.%s.fetch_interval", name)),
		ConfirmNum:        v.GetInt64(fmt.Sprintf("workers.%s.confirm_num", name)),
		StartBlockHeight:  v.GetInt64(fmt.Sprintf("workers.%s.start_block_height", name)),
	}
}

// Reads storage params from config.json
func (v *viperConfig) ReadDBConfig() *models.StorageConfig {
	return &models.StorageConfig{
		URL:        v.GetString("storage.url"),
		DBDriver:   v.GetString("storage.driver"),
		DBHOST:     v.GetString("storage.host"),
		DBPORT:     v.GetInt64("storage.port"),
		DBSSL:      v.GetString("storage.ssl_mode"),
		DBName:     v.GetString("storage.db_name"),
		DBUser:     v.GetString("storage.user"),
		DBPassword: v.GetString("storage.password"),
	}
}
