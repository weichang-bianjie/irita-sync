package monitor

import (
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/monitor/metrics"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	NodeStatusNotReachable = 0
	NodeStatusSyncing      = 1
	NodeStatusCatchingUp   = 2

	SyncTaskFollowing  = 1
	SyncTaskCatchingUp = 0
)

type clientNode struct {
	nodeHeight  metrics.Guage
	dbHeight    metrics.Guage
	nodeTimeGap metrics.Guage
	syncWorkWay metrics.Guage
}

func NewMetricNode(server metrics.Monitor) clientNode {
	dbHeightMetric := metrics.NewGuage(
		"sync",
		"status",
		"db_height",
		"sync system database max block height",
		nil,
	)
	nodeTimeGapMetric := metrics.NewGuage(
		"sync",
		"status",
		"node_seconds_gap",
		"the seconds gap between node block time with sync db block time",
		nil,
	)

	server.RegisterMetrics(dbHeightMetric, nodeTimeGapMetric)
	dbHeight, _ := metrics.CovertGuage(dbHeightMetric)
	nodeTimeGap, _ := metrics.CovertGuage(nodeTimeGapMetric)
	return clientNode{
		dbHeight:    dbHeight,
		nodeTimeGap: nodeTimeGap,
	}
}

func (node *clientNode) Report() {
	for {
		t := time.NewTimer(time.Duration(10) * time.Second)
		select {
		case <-t.C:
			node.nodeStatusReport()
		}
	}
}
func (node *clientNode) nodeStatusReport() {

	block, err := new(models.Block).GetMaxBlockHeight()
	if err != nil {
		logger.Error("query block exception", logger.String("error", err.Error()))
	}

	node.dbHeight.Set(float64(block.Height))

	timeGap := time.Now().Unix() - block.Time
	node.nodeTimeGap.Set(float64(timeGap))

	return
}

func Start() {
	c := make(chan os.Signal)
	//monitor system signal
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// start monitor
	server := metrics.NewMonitor(models.GetSrvConf().PromethousPort)
	node := NewMetricNode(server)

	server.Report(func() {
		go node.Report()
	})
	<-c
}
