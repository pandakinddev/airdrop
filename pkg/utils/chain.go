package utils

import (
	"context"
	"errors"

	"github.com/HungryPandaHome/airdrop/pkg/defs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetStage ...
func GetStage(ctx context.Context,
	client *ethclient.Client,
	account common.Address,
	contract *string) (defs.Stage, error) {

	if contract == nil {
		return defs.StageUnknown, nil
	}
	address := common.HexToAddress(*contract)
	_ = address
	// instance, err := airdrop.NewTimedCaller(address, client)
	// if err != nil {
	// 	return defs.StageUnknown, err
	// }
	// status, err := instance.Status(&bind.CallOpts{
	// 	From:    account,
	// 	Context: ctx,
	// })
	// if err != nil {
	// 	return defs.StageUnknown, err
	// }
	switch defs.StageRegistration {
	case 0:
		return defs.StageRegistration, nil
	case 1:
		return defs.StageRewards, nil
	case 2:
		return defs.StageEnded, nil
	default:
		return defs.StageUnknown, errors.New("unexpected status")
	}
}
