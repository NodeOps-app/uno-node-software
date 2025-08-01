package operator

import (
	"math/big"
	"uno-node-software/contracts/orchestrator"
	"uno-node-software/pkg/utils/constants"
	"uno-node-software/pkg/web3/query"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

func (o *Operator) PollEpochCallback(client *ethclient.Client) error {
	o.logger.Info("Polling epoch")
	contract, err := orchestrator.NewOrchestrator(o.ContractAddress, client)

	if err != nil {
		return err
	}

	epoch, err := query.QueryWithRetry[*big.Int](contract.CurrentEpoch, constants.MaxRetries)
	if err != nil {
		return err
	}

	if epoch.Uint64() > o.epochId {
		o.logger.Info("new epoch detected", zap.Uint64("epochId", epoch.Uint64()))
		o.epochId = epoch.Uint64()
		err = o.presentMaam(client, contract)
		if err == nil {
			o.epochId = epoch.Uint64()
			o.logger.Info("epoch updated", zap.Uint64("epochId", o.epochId))
		} else {
			return err
		}
	} else {
		o.logger.Info("epoch unchanged", zap.Uint64("epochId", epoch.Uint64()))
	}
	return nil
}

func (o *Operator) presentMaam(client *ethclient.Client, contract *orchestrator.Orchestrator) error {
	o.logger.Info("Preparing submission...")

	o.txManager.LockAccount(client)
	defer o.txManager.UnlockAccount()

	op := func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.Present(opts, big.NewInt(int64(o.epochId)))
	}
	o.logger.Info("Submitting transaction...")

	tx, err := o.txManager.SecureExec(client, op, constants.MaxRetries)

	if err != nil {
		return err
	}

	o.logger.Info("Submission sent with tx hash ", zap.String("txHash", tx.Hash().Hex()))
	return nil
}
