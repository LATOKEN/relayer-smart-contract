package config

import (
	"fmt"
	"os"
	"strings"

	"latoken/relayer-smart-contract/src/models"

	"github.com/spf13/viper"
)

// Config ...
type Config interface {
	ReadServiceConfig() string
	ReadWorkersConfig() (*models.WorkerConfig, *models.WorkerConfig)
	ReadLachainConfig() *models.WorkerConfig
	ReadDBConfig() *models.StorageConfig
	GetString(key string) string
	GetStringMap(key string) map[string]string
	GetInt64(key string) int64
	GetBool(key string) bool
	GetFloat64(key string) float64
	Init()
}

type viperConfig struct {
}

func (v *viperConfig) Init() {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(`config.json`)
	if _, err := os.Stat("./config.json.local"); !os.IsNotExist(err) {
		viper.SetConfigFile(`config.json.local`)
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
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

// NewViperConfig creates new viper for reading config.json
func NewViperConfig() Config {
	v := &viperConfig{}
	v.Init()
	return v
}
