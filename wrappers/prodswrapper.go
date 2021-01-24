package wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/flyingtang/go-micro/serviceimpl"
	"github.com/flyingtang/go-micro/services"
	"github.com/micro/go-micro/client"
	"strconv"
)

type ProdsWrapper struct {
	client.Client
}

func NewProdsWrapper (c client.Client) client.Client{
	return &ProdsWrapper{Client:c}
}


func defaultProds(rsp interface{}) {
	models := make([]*services.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		models = append(models, serviceimpl.NewProd(20+1, "prodname"+strconv.Itoa(20+int(i))))
	}
	result := rsp.(*services.ProdListResponse)
	result.Data = models
}


func (this *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout: 1000,
	}
	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return this.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		defaultProds(rsp)
		return nil
	})
}
