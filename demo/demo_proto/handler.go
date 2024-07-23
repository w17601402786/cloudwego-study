package main

import (
	"context"
	api "github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/api"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/service"
)

// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// Echo implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}
