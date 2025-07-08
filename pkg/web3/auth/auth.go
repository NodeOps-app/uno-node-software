package auth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

const defaultGasLimit = uint64(1000000)

func SetupAuth(client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.New("failed to get chainId: " + err.Error())
	}
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, errors.New("failed to get nonce: " + err.Error())
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, errors.New("failed to create authorized transactor: " + err.Error())
	}

	auth.Nonce = new(big.Int).SetUint64(nonce)
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = defaultGasLimit
	auth.From = address

	if err = UpdateGasPrice(context.Background(), auth, client, 1); err != nil {
		auth.GasPrice = new(big.Int).SetUint64(defaultGasLimit)
	}

	return auth, nil
}

func UpdateGasPrice(ctx context.Context, auth *bind.TransactOpts, client *ethclient.Client, multiplier int) (err error) {
	auth.GasFeeCap, auth.GasTipCap, err = SuggestEIP1559Fees(ctx, client, int64(multiplier))
	return err
}

// SuggestEIP1559Fees suggests fees for EIP-1559 transactions.
func SuggestEIP1559Fees(ctx context.Context, client *ethclient.Client, multiplier int64) (maxFeePerGas, maxPriorityFeePerGas *big.Int, err error) {
	// Suggest a gas tip (priority fee) using the eth_maxPriorityFeePerGas method
	tip, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, err
	}

	// Retrieve the base fee for the next block
	head, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	baseFee := head.BaseFee

	// Calculate maxFeePerGas: baseFee * 2 + priorityFee
	maxFeePerGas = new(big.Int).Mul(baseFee, big.NewInt(multiplier))
	maxFeePerGas.Add(maxFeePerGas, new(big.Int).Mul(tip, big.NewInt(multiplier)))

	return maxFeePerGas, tip, nil
}
