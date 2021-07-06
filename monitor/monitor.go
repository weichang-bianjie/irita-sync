package monitor

import (
	"context"
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/libs/pool"
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
)

type clientNode struct {
	nodeStatus  metrics.Guage
	nodeHeight  metrics.Guage
	dbHeight    metrics.Guage
	nodeTimeGap metrics.Guage
}

func NewMetricNode(server metrics.Monitor) clientNode {
	nodeHeightMetric := metrics.NewGuage(
		"sync",
		"status",
		"node_height",
		"full node latest block height",
		nil,
	)
	dbHeightMetric := metrics.NewGuage(
		"sync",
		"status",
		"db_height",
		"sync system database max block height",
		nil,
	)
	nodeStatusMetric := metrics.NewGuage(
		"sync",
		"status",
		"node_status",
		"full node status(0:NotReachable,1:Syncing,2:CatchingUp)",
		nil,
	)
	nodeTimeGapMetric := metrics.NewGuage(
		"sync",
		"status",
		"node_seconds_gap",
		"the seconds gap between node block time with sync db block time",
		nil,
	)
	server.RegisterMetrics(nodeHeightMetric, dbHeightMetric, nodeStatusMetric, nodeTimeGapMetric)
	nodeHeight, _ := metrics.CovertGuage(nodeHeightMetric)
	dbHeight, _ := metrics.CovertGuage(dbHeightMetric)
	nodeStatus, _ := metrics.CovertGuage(nodeStatusMetric)
	nodeTimeGap, _ := metrics.CovertGuage(nodeTimeGapMetric)
	return clientNode{
		nodeStatus:  nodeStatus,
		nodeHeight:  nodeHeight,
		dbHeight:    dbHeight,
		nodeTimeGap: nodeTimeGap,
	}
}

func (node *clientNode) Report() {
	for {
		t := time.NewTimer(time.Duration(5) * time.Second)
		select {
		case <-t.C:
			node.nodeStatusReport()
		}
	}
}
func (node *clientNode) nodeStatusReport() {
	client, err := pool.GetClientWithTimeout(10 * time.Second)
	if err != nil {
		logger.Error("rpc node connection exception", logger.String("error", err.Error()))
		node.nodeStatus.Set(float64(NodeStatusNotReachable))
		return
	}
	defer func() {
		client.Release()
	}()

	block, err := new(models.Block).GetMaxBlockHeight()
	if err != nil {
		logger.Error("query block exception", logger.String("error", err.Error()))
	}
	node.dbHeight.Set(float64(block.Height))
	status, err := client.Status(context.Background())
	if err != nil {
		logger.Error("rpc node connection exception", logger.String("error", err.Error()))
		node.nodeStatus.Set(float64(NodeStatusNotReachable))
		return
	}

	if status.SyncInfo.CatchingUp {
		node.nodeStatus.Set(float64(NodeStatusCatchingUp))
	} else {
		node.nodeStatus.Set(float64(NodeStatusSyncing))
	}
	node.nodeHeight.Set(float64(status.SyncInfo.LatestBlockHeight))
	timeGap := status.SyncInfo.LatestBlockTime.Unix() - block.Time
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