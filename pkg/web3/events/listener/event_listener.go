package listener

import (
	"context"
	"log"
	"time"
	"uno-node-software/pkg/utils/constants"
	"uno-node-software/pkg/web3/clients"
	common2 "uno-node-software/pkg/web3/events/common"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EventListener struct {
	addresses      []common.Address
	topics         [][]common.Hash
	ws             string
	pacemaker      *time.Ticker
	backoffHandler *common2.BackoffHandler
}

func NewEventListener(wsUrl string, addrs []common.Address, tps [][]common.Hash) *EventListener {
	return &EventListener{
		addresses:      addrs,
		topics:         tps,
		ws:             wsUrl,
		pacemaker:      time.NewTicker(constants.DefaultPacemakerInterval),
		backoffHandler: common2.NewBackoffHandler(constants.DefaultBackoffInitial, constants.DefaultBackoffMax),
	}
}

func (e *EventListener) StartListening(callback func(*ethclient.Client, types.Log)) {
	query := ethereum.FilterQuery{
		Addresses: e.addresses,
		Topics:    e.topics,
	}

	go e.listen(query, callback)
}

func (e *EventListener) listen(query ethereum.FilterQuery, callback func(client *ethclient.Client, log types.Log)) {
	for {
		client, err := clients.NewEthClient(e.ws)
		if err != nil {
			// Allow a cooldown period - rely on polling
			log.Printf("Failed to connect to WebSocket: %v", err)
			time.Sleep(constants.DefaultBackoffMax)
			continue
		}

		logs := make(chan types.Log)
		sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)

		if err != nil {
			time.Sleep(e.backoffHandler.Next())
			client.Close()
			continue
		}

		e.backoffHandler.Reset() // Reset backoff on successful connection

		for {
			select {
			case <-sub.Err():
				log.Printf("Subscription error: %v", err)
				sub.Unsubscribe()
				client.Close()
				time.Sleep(e.backoffHandler.Next())
				break
			case vLog := <-logs:
				callback(client, vLog)
			case <-e.pacemaker.C:
				if !clients.IsConnectionAlive(client) {
					log.Println("Connection lost. Attempting to reconnect...")
					sub.Unsubscribe()
					client.Close()
					time.Sleep(e.backoffHandler.Next())
					break
				}
			}
		}
	}
}
