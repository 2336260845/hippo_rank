package server

import (
	"context"
	"fmt"
	"github.com/2336260845/hippo_rank/config"
	"github.com/2336260845/hippo_search/gen-go/rank"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/sirupsen/logrus"
)

type RankServer struct{}

func (qas *RankServer) Rank(ctx context.Context, req *rank.RankParam) (r []*rank.Doc, err error) {
	logrus.Infof("Rank param=%+v", req.Docs)
	if len(req.Docs) == 0 {
		return req.Docs, nil
	}

	if config.GetConfig().Debug {
		for i, v := range req.Docs {
			v.RankScore = float64(100-i)
		}
	}

	//TODO 实现rank函数
	return req.Docs, nil
}

func ThriftInit(conf *config.Config) {
	transport, err := thrift.NewTServerSocket(conf.ThriftAddress)
	if err != nil {
		panic(fmt.Sprintf("ThriftInit NewTServerSocket error, err=%+v", err.Error()))
	}

	handler := &RankServer{}
	processor := rank.NewRankServiceProcessor(handler)
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	if err := server.Serve(); err != nil {
		panic(fmt.Sprintf("ThriftInit thrift Serve error, err=%+v", err.Error()))
	}
}
