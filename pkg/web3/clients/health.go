package clients

import (
	"context"
	"uno-node-software/pkg/utils/constants"

	"github.com/ethereum/go-ethereum/ethclient"
)

func IsConnectionAlive(client *ethclient.Client) bool {
	if client == nil {
		return false
	}
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultConnectionTimeout)
	defer cancel()

	_, err := client.NetworkID(ctx)
	if err != nil {
		return false
	}
	return true
}
