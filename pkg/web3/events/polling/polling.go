package polling

import (
	"errors"
	"fmt"
	"time"
	"uno-node-software/pkg/web3/clients"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type PollingService struct {
	url             string
	client          *ethclient.Client
	pollingInterval *time.Ticker
}

func NewPollingService(_url string, interval time.Duration) *PollingService {
	return &PollingService{
		url:             _url,
		pollingInterval: time.NewTicker(interval),
	}
}

func (s *PollingService) Poll(operation func(client *ethclient.Client) error, logger *zap.Logger) {
	defer s.pollingInterval.Stop()
	logger.Info("Polling service started at ", zap.String("time", time.Now().String()))
	err := s.pollWithRetry(operation)
	if err != nil {
		logger.Error("Polling operation failed", zap.Error(err))
	}
	for {
		select {
		case <-s.pollingInterval.C:
			err = s.pollWithRetry(operation)
			if err != nil {
				logger.Error("Polling operation failed", zap.Error(err))
			}
			logger.Info("Awaiting next poll...")
		}
	}
}

func (s *PollingService) pollWithRetry(operation func(client *ethclient.Client) error) error {
	if !clients.IsConnectionAlive(s.client) {
		client, err := clients.NewEthClient(s.url)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to create rpc client: %v", err))
		}
		s.client = client
	}

	defer s.client.Close()

	return operation(s.client)
}
