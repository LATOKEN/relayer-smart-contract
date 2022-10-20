package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/LATOKEN/relayer-smart-contract.git/src/models"

	vault "github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config ...
type Config interface {
	ReadServiceConfig() string
	ReadWorkersConfig() []*models.WorkerConfig
	ReadLachainConfig() *models.WorkerConfig
	ReadDBConfig() *models.StorageConfig
	ReadChains() []string
	GetString(key string) string
	GetStringMap(key string) map[string]string
	GetInt64(key string) int64
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetStringSlice(key string) []string
	GetInterface(key string) interface{}
	ReadVaultConfig()
	SetLogger(logger *logrus.Logger)
	Init()
}

type viperConfig struct {
	isDevMode bool
	logger    *logrus.Logger
	vault     *vault.Client
}

func (v *viperConfig) SetLogger(logger *logrus.Logger) {
	v.logger = logger
}

func (v *viperConfig) Init() {
	viper.AutomaticEnv()
	// viper.AddConfigPath("../config-files/")
	viper.AddConfigPath(os.Getenv("FILE_PATH"))
	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	// viper.SetConfigName("config.json")
	viper.SetConfigName(os.Getenv("FILE_NAME"))
	if _, err := os.Stat("./config.json.local"); !os.IsNotExist(err) {
		viper.SetConfigFile(`config.json.local`)
	}
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("read config error: %s", err))
	}
}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (v *viperConfig) GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func (v *viperConfig) GetStringMap(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func (v *viperConfig) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func (v *viperConfig) GetInterface(key string) interface{} {
	return viper.Get(key)
}

func (v *viperConfig) IsSet(key string) bool {
	return viper.IsSet(key)
}

// NewViperConfig creates new viper for reading config.json
func NewViperConfig(isDev bool) Config {
	v := &viperConfig{isDevMode: isDev}
	v.Init()
	return v
}
