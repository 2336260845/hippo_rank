package main

import (
	"github.com/2336260845/hippo_rank/config"
	"github.com/2336260845/hippo_rank/es"
	"github.com/2336260845/hippo_rank/server"
)

func init() {
	config.InitConfig("")
	cf := config.GetConfig()
	es.InitEsClient(cf)
}

func main() {
	cf := config.GetConfig()
	server.ThriftInit(cf)
}
