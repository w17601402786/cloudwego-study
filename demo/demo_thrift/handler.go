package main

import (
	"context"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/biz/service"
	api "github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api"
)

// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// Echo implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Echo(ctx context.Context, request *api.Request) (resp *api.Response, err error) {
	resp, err = service.NewEchoService(ctx).Run(request)

	return resp, err
}
