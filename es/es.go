package es

import (
	"fmt"
	"github.com/2336260845/hippo_rank/config"
	es "github.com/elastic/go-elasticsearch"
)

var esClient *es.Client

const (
	AnalyzerIkSmart   = "ik_smart"
	AnalyzerIkMaxWord = "ik_max_word"
)

func InitEsClient(conf *config.Config) {
	cf := es.Config{Addresses: []string{conf.EsAddress}}
	client, err := es.NewClient(cf)
	if err != nil {
		panic(fmt.Sprintf("InitEsClient NewDefaultClient error, err=%+v", err.Error()))
	}

	esClient = client
}
