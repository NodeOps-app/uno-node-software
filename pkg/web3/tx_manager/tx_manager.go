package tx_manager

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"sync"
	"uno-node-software/pkg/web3/auth"
	"uno-node-software/pkg/web3/query"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TxManager struct {
	Opts       *bind.TransactOpts
	privateKey *ecdsa.PrivateKey
	mu         sync.Mutex
}

func NewTxManager(pk *ecdsa.PrivateKey) *TxManager {
	return &TxManager{
		privateKey: pk,
	}
}

func (tm *TxManager) GetPrivateKey() *ecdsa.PrivateKey {
	return tm.privateKey
}

func (tm *TxManager) LockAccount(client *ethclient.Client) {
	tm.mu.Lock()
	tm.UpdateAuth(client)
}

func (tm *TxManager) UnlockAccount() {
	tm.mu.Unlock()
}

func (tm *TxManager) UpdateAuth(client *ethclient.Client) {
	tm.Opts, _ = auth.SetupAuth(client, tm.privateKey)
}

func (tm *TxManager) SecureExec(client *ethclient.Client, op func(opts *bind.TransactOpts) (*types.Transaction, error), maxRetries int) (*types.Transaction, error) {
	var err error
	var tx *types.Transaction
	tx, err = op(tm.Opts)
	if err != nil {
		return nil, err
	}

	receipt, err := query.GetReceiptWithRetry(context.Background(), client, tx.Hash())
	if err != nil {
		return nil, err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return nil, errors.New(fmt.Sprintf("Transaction failed: ", query.GetTxRevertReason(tm.Opts.From, client, tx)))
	}

	return tx, nil
}
