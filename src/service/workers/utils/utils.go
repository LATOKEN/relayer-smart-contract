package utils

import (
	"crypto/ecdsa"
	"errors"
	"math"
	"math/big"

	"github.com/LATOKEN/relayer-smart-contract.git/src/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetPrivateKey(config *models.WorkerConfig) (*ecdsa.PrivateKey, error) {
	privKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

// QuoBigInt ...
func QuoBigInt(a *big.Int, b *big.Int) *big.Float {
	fl := new(big.Float).SetInt(a)
	fl.Quo(fl, new(big.Float).SetInt(b))
	return fl
}

// GetBigIntForDecimal ...
func GetBigIntForDecimal(decimal int) *big.Int {
	floatDecimal := big.NewFloat(math.Pow10(decimal))
	bigIntDecimal := new(big.Int)
	floatDecimal.Int(bigIntDecimal)

	return bigIntDecimal
}

// CalcActualOutAmount ...
func CalcActualOutAmount(amount *big.Int, ratio *big.Float, fixedFee *big.Int) *big.Int {
	res := new(big.Float).SetInt(amount)
	res.Mul(res, ratio)

	amountInt := new(big.Int)
	res.Int(amountInt)

	amountInt.Sub(amountInt, fixedFee)
	return amountInt
}

func BytesToBytes32(b []byte) [32]byte {
	var byteArr [32]byte
	copy(byteArr[:], b)
	return byteArr
}

func StringToBytes32(b string) [32]byte {
	var byteArr [32]byte
	copy(byteArr[:], common.RightPadBytes(common.Hex2Bytes(b), 32))
	return byteArr
}

func StringToBytes8(b string) [8]byte {
	var byteArr [8]byte
	copy(byteArr[:], common.RightPadBytes(common.Hex2Bytes(b), 8))
	return byteArr
}

func BytesToBytes8(b []byte) [8]byte {
	var byteArr [8]byte
	copy(byteArr[:], b)
	return byteArr
}

func CalcutateSwapID(originChainID, destChainID, nonce string) string {
	return originChainID + destChainID + nonce
}

func ConvertDecimals(originDecimals, destDecimals uint8, amount string) string {
	origin := new(big.Int).SetInt64(int64(math.Pow10(int(originDecimals))))
	dest := new(big.Int).SetInt64(int64(math.Pow10(int(destDecimals))))
	amountInFloat, _ := new(big.Int).SetString(amount, 10)
	// conversion := new(big.Int).Quo(origin, dest)
	return new(big.Int).Quo(new(big.Int).Mul(amountInFloat, dest), origin).String()
}

func GetCurrentStep(resourceID string, stepIndex uint8) string {
	return resourceID[16*stepIndex:16*stepIndex+8] + "00000000000000000000000000000000000000000000000000000000"
}

func PvtKeyToAddress(pvtKey string) (string, error) {
	privateKey, err := crypto.HexToECDSA(pvtKey)
	if err != nil {
		return "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex(), nil
}
