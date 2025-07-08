package clients

import (
	"uno-node-software/pkg/utils/constants"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewEthClient(url string) (*ethclient.Client, error) {
	var client *ethclient.Client
	var err error
	operation := func() error {
		client, err = ethclient.Dial(url)
		return err
	}
	err = backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), constants.MaxRetries))
	return client, err
}
