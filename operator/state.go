package operator

import (
	"math/big"
	"uno-node-software/contracts/orchestrator"
	"uno-node-software/logger"
	"uno-node-software/pkg/utils/constants"
	"uno-node-software/pkg/utils/env"
	"uno-node-software/pkg/web3/clients"
	"uno-node-software/pkg/web3/query"
	"uno-node-software/pkg/web3/tx_manager"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Operator struct {
	epochId         uint64
	license         uint64
	RpcUrl          string
	txManager       *tx_manager.TxManager
	ContractAddress common.Address
	logger          *zap.Logger
}

func loadOperator() *Operator {
	godotenv.Load(".env")
	operator := &Operator{}
	operator.RpcUrl = env.MustLoadString("RPC_URL")
	operator.ContractAddress = common.HexToAddress(env.MustLoadString("CONTRACT_ADDRESS"))
	pk := env.MustLoadPrivateKey("PRIVATE_KEY")
	operator.txManager = tx_manager.NewTxManager(pk)
	operator.logger = logger.InitLogger()

	client, err := clients.NewEthClient(operator.RpcUrl)
	if err != nil {
		operator.logger.Fatal("failed to create rpc client", zap.Error(err))
	}

	contract, err := orchestrator.NewOrchestrator(operator.ContractAddress, client)
	if err != nil {
		operator.logger.Fatal("failed to create contract", zap.Error(err))
	}

	info, err := contract.GetOperatorInfo(&bind.CallOpts{}, crypto.PubkeyToAddress(pk.PublicKey))

	if err != nil {
		operator.logger.Fatal("failed to query license", zap.Error(err))
	}

	license := info.License

	if license.Uint64() == 0 {
		operator.logger.Fatal("no license delegated to operator")
	}

	operator.license = license.Uint64()

	epoch, err := query.QueryWithRetry[*big.Int](contract.CurrentEpoch, constants.MaxRetries)

	if err != nil {
		operator.logger.Fatal("failed to query register", zap.Error(err))
	}

	operator.epochId = epoch.Uint64()

	return operator
}
