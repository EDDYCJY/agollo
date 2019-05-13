package agollo

import (
	"sync"
	"encoding/json"
	"context"
)

type AbsComponent interface {
	Start(context.Context)
}


func StartRefreshConfig(ctx context.Context, component AbsComponent)  {
	component.Start(ctx)
}

type ApolloConnConfig struct {
	AppId string `json:"appId"`
	Cluster string `json:"cluster"`
	NamespaceName string `json:"namespaceName"`
	ReleaseKey string `json:"releaseKey"`
	sync.RWMutex
}

type ApolloConfig struct {
	ApolloConnConfig
	Configurations map[string]string `json:"configurations"`
}

func createApolloConfigWithJson(b []byte) (*ApolloConfig,error) {
	apolloConfig:=&ApolloConfig{}
	err:=json.Unmarshal(b,apolloConfig)
	if isNotNil(err) {
		return nil,err
	}
	return apolloConfig,nil
}