package errors

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

func HandleTransactionError(client *ethclient.Client, auth *bind.TransactOpts, err error) error {
	if strings.Contains(err.Error(), "Nonce too low") || strings.Contains(err.Error(), "Nonce too high") {
		nonce, ok := extractNonceFromError(err.Error())
		var diff uint64
		if ok {
			diff = nonce - auth.Nonce.Uint64()
		} else {
			pending, err := client.PendingNonceAt(context.Background(), auth.From)
			if err != nil {
				return err
			}
			diff = pending - auth.Nonce.Uint64()
		}

		auth.Nonce = new(big.Int).Add(auth.Nonce, new(big.Int).SetUint64(diff))
	}
	return nil
}

func extractNonceFromError(err string) (uint64, bool) {
	regex := regexp.MustCompile(`nonce too low: next nonce (\d+), tx nonce \d+`)
	matches := regex.FindStringSubmatch(err)
	if len(matches) > 1 {
		nonce, err := strconv.ParseUint(matches[1], 10, 64)
		return nonce, err == nil
	}
	return 0, false
}
