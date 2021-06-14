package utils_test

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/HungryPandaHome/airdrop/pkg/airdrop"
	"github.com/HungryPandaHome/airdrop/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	currentFile, _ = utils.SourceFilePath()
	pkgSuffix      = filepath.Join("pkg", "utils")
	projectRoot    = strings.TrimSuffix(filepath.Dir(currentFile), pkgSuffix)
)

func init() {
	_ = godotenv.Load(filepath.Join(projectRoot, ".env"))
	err := godotenv.Load(filepath.Join(projectRoot, fmt.Sprintf("%s.env", os.Getenv("ENV"))))
	if err != nil {
		log.Println(err)
	}
}

func TestStatus(t *testing.T) {

	address := os.Getenv("TEST_CONTRACT")
	key := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(key)
	assert.NoError(t, err)
	adminAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	assert.NoError(t, err)

	stage, err := utils.GetStage(context.Background(), client, adminAddress, &address)
	assert.NoError(t, err)
	log.Println(stage)
}

func TestSignPayload(t *testing.T) {
	address := os.Getenv("TEST_CONTRACT")
	key := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(key)
	assert.NoError(t, err)
	adminAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	log.Println(adminAddress.Hex())
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	assert.NoError(t, err)

	airdropContract, err := airdrop.NewAirdrop(common.HexToAddress(address), client)
	assert.NoError(t, err)
	recipient := common.HexToAddress("0x798f3a9ded0460b2FF4727a59e34f97955486111")
	amount := big.NewInt(100)

	sign, hash, err := utils.SignPayload(recipient, amount, privateKey)
	assert.NoError(t, err)
	log.Printf("RECIPIENT: 0x%x\n", recipient.Bytes())
	log.Printf("HASH: %s\n", hexutil.Encode(hash))
	log.Printf("SIGNATURE: %s\n", hexutil.Encode(sign))
	pubKey, err := crypto.Ecrecover(hash, sign)
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(crypto.FromECDSAPub(&privateKey.PublicKey), pubKey))
	log.Println("LEN", len(sign), len(pubKey))
	valid := crypto.VerifySignature(pubKey, hash, sign)
	assert.True(t, valid)
	contractAdmin, err := airdropContract.Admin(&bind.CallOpts{Context: context.Background()})
	assert.NoError(t, err)
	assert.Equal(t, contractAdmin.Hex(), adminAddress.Hex())

	tx, err := airdropContract.ClaimTokens(&bind.TransactOpts{
		From:    recipient,
		Context: context.Background(),
	}, amount, []byte(hexutil.Encode(sign)))
	assert.NoError(t, err)
	data, err := tx.MarshalJSON()
	assert.NoError(t, err)
	log.Println(string(data))
}
