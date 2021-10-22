package utils

import (
	"crypto/ecdsa"
	"latoken/relayer-smart-contract/src/models"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// // CalculateSwapID ...
// func calculateSwapID(randomNumberHash, sender, laRecipientAddr []byte) []byte {
// 	data := randomNumberHash
// 	data = append(data, []byte(sender)...)
// 	data = append(data, []byte(laRecipientAddr)...)
// 	return tmhash.Sum(data)
// }

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
	copy(byteArr[:], b)
	return byteArr
}

func BytesToBytes8(b []byte) [8]byte {
	var byteArr [8]byte
	copy(byteArr[:], b)
	return byteArr
}
