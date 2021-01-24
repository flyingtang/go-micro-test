package main

import (
	"github.com/flyingtang/go-micro/bak/serviceimpl"
	"github.com/flyingtang/go-micro/bak/services"
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
	)

	service.Init()
	services.RegisterProdServiceHandler(service.Server(), new(serviceimpl.ProdService))
	service.Run()
}
