package operator

import (
	"uno-node-software/pkg/utils/constants"
	"uno-node-software/pkg/web3/events/polling"
)

func Start() {
	operator := loadOperator()
	poll := polling.NewPollingService(operator.RpcUrl, constants.DefaultPollingInterval)
	operator.logger.Info("starting operator")
	poll.Poll(operator.PollEpochCallback, operator.logger)
}
