package main

import (
	"github.com/flyingtang/go-micro/serviceimpl"
	"github.com/flyingtang/go-micro/services"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry(registry.Addrs("192.168.2.112:8500"))
	service := micro.NewService(
		micro.Name("prodservice"),
		micro.Address(":8011"),
		micro.Registry(consulReg),
		micro.Metadata(map[string]string{"protocol":"grpc"}),
	)

	service.Init()
	services.RegisterProdServiceHandler(service.Server(), new(serviceimpl.ProdService))
	service.Run()
}
