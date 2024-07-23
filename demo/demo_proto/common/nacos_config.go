package nacos_config

import (
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/conf"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func GetNacosConfig() (iClient naming_client.INamingClient, err error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}

	nacos := conf.GetConf().Nacos

	cc := constant.ClientConfig{
		NamespaceId:         nacos.Namespace,
		TimeoutMs:           nacos.TimeoutMs,
		NotLoadCacheAtStart: nacos.NotLoadCacheAtStart,
		LogDir:              nacos.LogDir,
		CacheDir:            nacos.CacheDir,
		LogLevel:            nacos.LogLevel,
		Username:            nacos.Username,
		Password:            nacos.Password,
	}

	return clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

}
