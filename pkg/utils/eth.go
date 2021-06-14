package utils

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

// SignPayload ...
func SignPayload(recipient common.Address,
	amount *big.Int,
	privateKey *ecdsa.PrivateKey) ([]byte, []byte, error) {

	hash := solsha3.SoliditySHA3(
		solsha3.Address(recipient),
		solsha3.Uint256(amount),
	)
	prefixedHash := crypto.Keccak256Hash(
		[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(hash))),
		hash,
	)
	hash = prefixedHash.Bytes()
	sign, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, nil, err
	}
	return sign, hash, nil
}
