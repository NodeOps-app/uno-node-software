package query

import (
	"context"
	"fmt"
	"strings"
	"uno-node-software/pkg/utils/constants"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func QueryWithRetry[K any](op func(*bind.CallOpts) (K, error), maxRetries uint64) (K, error) {
	var val K
	var err error
	err = backoff.Retry(
		func() error {
			val, err = op(&bind.CallOpts{})
			return err
		},
		backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxRetries),
	)
	return val, err
}

func GetReceiptWithRetry(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	var receipt *types.Receipt
	var err error
	operation := func() error {
		receipt, err = client.TransactionReceipt(ctx, txHash)
		return err
	}
	err = backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), constants.MaxRetries))
	return receipt, err
}

func QueryAny[K any](call func() (K, error)) (K, error) {
	var val K
	operation := func() error {
		var err error
		val, err = call()
		return err
	}
	// Use the retry package to execute the operation with backoff
	err := backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), constants.MaxRetries))
	if err != nil {
		return *new(K), err
	}
	return val, err
}

func GetTxRevertReason(from common.Address, client *ethclient.Client, tx *types.Transaction) []byte {
	var data []byte
	// Prepare a call message to simulate the transaction and get the revert reason
	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	// Perform a simulated call
	revertData, err := client.CallContract(context.Background(), msg, nil)

	if err != nil {
		if strings.Contains(err.Error(), "execution reverted") {
			data = revertData
		} else {
			data = []byte(fmt.Sprintf("Failed to call contract: %v", err))
		}
	} else {
		data = []byte(fmt.Sprintf("Transaction did not revert with a reason."))
	}
	return data
}
