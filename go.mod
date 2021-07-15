module github.com/bianjieai/irita-sync

go 1.15

require (
	github.com/kaifei-bianjie/msg-parser v0.0.0-20210205025444-68411682d9a1
	github.com/spf13/viper v1.7.1
	github.com/tendermint/tendermint v0.34.3
	github.com/weichang-bianjie/metric-sdk v1.0.0
	go.uber.org/zap v1.15.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
