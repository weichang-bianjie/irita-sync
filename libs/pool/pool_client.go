//init client from clientPool.
//client is httpClient of tendermint

package pool

import (
	"context"
	svrConf "github.com/bianjieai/irita-sync/confs/server"
	"github.com/bianjieai/irita-sync/libs/logger"
	commonPool "github.com/jolestar/go-commons-pool"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"sync"
	"time"
)

var (
	poolObject  *commonPool.ObjectPool
	poolFactory PoolFactory
	ctx         = context.Background()
)

func init() {
	var (
		syncMap sync.Map
	)
	conf := svrConf.SvrConf
	for _, url := range conf.NodeUrls {
		key := generateId(url)
		endPoint := EndPoint{
			Address:   url,
			Network:   conf.NetWork,
			Available: true,
		}

		syncMap.Store(key, endPoint)
	}

	poolFactory = PoolFactory{
		peersMap: syncMap,
	}

	config := commonPool.NewDefaultPoolConfig()
	config.MaxTotal = conf.MaxConnectionNum
	config.MaxIdle = conf.InitConnectionNum
	config.MinIdle = conf.InitConnectionNum
	config.TestOnBorrow = true
	config.TestOnCreate = true
	config.TestWhileIdle = true

	poolObject = commonPool.NewObjectPool(ctx, &poolFactory, config)
	poolObject.PreparePool(ctx)
}

// get client from pool
func GetClient() *Client {
	c, err := poolObject.BorrowObject(ctx)
	for err != nil {
		logger.Error("GetClient failed,will try again after 3 seconds", logger.String("err", err.Error()))
		time.Sleep(3 * time.Second)
		c, err = poolObject.BorrowObject(ctx)
	}

	return c.(*Client)
}

// release client
func (c *Client) Release() {
	err := poolObject.ReturnObject(ctx, c)
	if err != nil {
		logger.Error(err.Error())
	}
}

func (c *Client) HeartBeat() error {
	http := c.Client.(*rpcclient.HTTP)
	_, err := http.Health()
	return err
}

func ClosePool() {
	poolObject.Close(ctx)
}
