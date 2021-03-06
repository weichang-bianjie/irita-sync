package server

import (
	constant "github.com/bianjieai/irita-sync/confs"
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/utils"
	"os"
	"strings"
)

type ServerConf struct {
	NodeUrls                []string
	NetWork                 string
	WorkerNumCreateTask     int
	WorkerNumExecuteTask    int
	WorkerMaxSleepTime      int
	BlockNumPerWorkerHandle int

	MaxConnectionNum  int
	InitConnectionNum int
}

var (
	SvrConf *ServerConf

	nodeUrls                = []string{"tcp://192.168.150.31:31557"}
	network                 = constant.Testnet
	workerNumExecuteTask    = 30
	workerMaxSleepTime      = 2 * 60
	blockNumPerWorkerHandle = 100
)

// get value of env var
func init() {
	if v, ok := os.LookupEnv(constant.EnvNameSerNetworkFullNode); ok {
		nodeUrls = strings.Split(v, ",")
	}

	if v, ok := os.LookupEnv(constant.EnvNameBlockChainNetwork); ok {
		switch v {
		case constant.Testnet, constant.Mainnet:
			network = v
		default:
			logger.Fatal("unknown network", logger.String(constant.EnvNameBlockChainNetwork, v))
		}
	}

	if v, ok := os.LookupEnv(constant.EnvNameWorkerNumExecuteTask); ok {
		if n, err := utils.ConvStrToInt(v); err != nil {
			logger.Fatal("convert str to int fail", logger.String(constant.EnvNameWorkerNumExecuteTask, v))
		} else {
			workerNumExecuteTask = n
		}
	}

	if v, ok := os.LookupEnv(constant.EnvNameWorkerMaxSleepTime); ok {
		if n, err := utils.ConvStrToInt(v); err != nil {
			logger.Fatal("convert str to int fail", logger.String(constant.EnvNameWorkerMaxSleepTime, v))
		} else {
			workerMaxSleepTime = n
		}
	}

	if v, ok := os.LookupEnv(constant.EnvNameBlockNumPerWorkerHandle); ok {
		if n, err := utils.ConvStrToInt(v); err != nil {
			logger.Fatal("convert str to int fail", logger.String(constant.EnvNameBlockNumPerWorkerHandle, v))
		} else {
			blockNumPerWorkerHandle = n
		}
	}

	SvrConf = &ServerConf{
		NodeUrls:                nodeUrls,
		NetWork:                 network,
		WorkerNumCreateTask:     1,
		WorkerNumExecuteTask:    workerNumExecuteTask,
		WorkerMaxSleepTime:      workerMaxSleepTime,
		BlockNumPerWorkerHandle: blockNumPerWorkerHandle,

		MaxConnectionNum:  100,
		InitConnectionNum: 5,
	}

	logger.Debug("print server config", logger.String("serverConf", utils.MarshalJsonIgnoreErr(SvrConf)))
}
