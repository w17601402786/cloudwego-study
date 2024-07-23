// @Author Adrian.Wang 2024/7/23 下午4:01:00
package main

import (
	"context"
	nacos_config "github.com/cloudwego/biz-demo/gomall/demo/demo_proto/common"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/api"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/api/echoservice"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"log"
	"time"
)

func main() {

	// 获取nacos配置
	cli, err := nacos_config.GetNacosConfig()

	r := resolver.NewNacosResolver(cli)

	if err != nil {
		log.Panic(err)
	}

	// 创建client
	echoClient, err := echoservice.NewClient(
		"demo_proto",
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*3),
	)

	resp, err := echoClient.Echo(context.Background(), &api.Request{
		Message: "hello  nacos-kitex",
	})

	if err != nil {
		log.Panic(err)
	}

	log.Println(resp.Message)
}
