package main

import (
	"github.com/flyingtang/go-micro/bak/services"
	"github.com/flyingtang/go-micro/bak/weblib"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry(registry.Addrs("192.168.2.112:8500"))
	myservice := micro.NewService(micro.Name("prodservice.client"))
	prodService := services.NewProdService("prodservice", myservice.Client())

	server := web.NewService(
		web.Name("httpprodservice"),
		web.Version("1.0"),
		web.Metadata(map[string]string{"protocol": "http"}),
		web.Address(":8001"),
		web.Handler(weblib.NewGinRouter(prodService)),
		web.Registry(consulReg),
	)

	server.Init()

	server.Run()
}
