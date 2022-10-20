package config

import (
	"context"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"os"
)

func (v *viperConfig) ReadVaultConfig() {
	config := vault.DefaultConfig()
	if v.IsSet("vault.address") {
		config.Address = v.GetString("vault.address")
	}
	client, err := vault.NewClient(config)
	if err != nil {
		panic(fmt.Sprintf("unable to initialize Vault client: %v", err))
	}
	if v.IsSet("vault.token") {
		client.SetToken(v.GetString("vault.token"))
	}
	v.vault = client
}

func (v *viperConfig) ReadPvtKeyForChain(Chain string) string {
	var keypath string
	if v.IsSet("vault.path") {
		keypath = v.GetString("vault.path")
	} else {
		keypath = os.Getenv("VAULT_PATH")
	}
	secret, err := v.vault.KVv2("secret").Get(context.Background(), keypath)
	if err != nil {
		panic(fmt.Sprintf("unable to read secret: %v", err))
	}

	value, ok := secret.Data[Chain].(string)
	if !ok {
		panic(fmt.Sprintf("pvt key for %s.%s not set", keypath, Chain))
	}
	return value
}
