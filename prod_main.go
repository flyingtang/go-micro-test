package main

import (
	"context"
	"github.com/flyingtang/go-micro/services"
	"github.com/flyingtang/go-micro/wrappers"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {

	consulReg := consul.NewRegistry(registry.Addrs("192.168.2.112:8500"))
	ginRouter := gin.Default()
	httpService := web.NewService(
		web.Name("httpprodservice"),
		web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
		web.Version("1.0"),
		//web.Metadata(map[string]string{"protocol": "http"}),
	)

	myService := micro.NewService(
		micro.Name("prodservice.client"),
		micro.WrapClient(wrappers.NewProdsWrapper),
	)
	prodService := services.NewProdService("prodservice", myService.Client())

	v1 := ginRouter.Group("/v1")
	v1.POST("prods", func(c *gin.Context) {
		var par services.ProdListRequest
		if err := c.Bind(&par); err != nil {
			c.JSON(400, gin.H{"status": err.Error()})
		} else {

			prodRes, _ := prodService.GetProdsList(context.Background(), &par)
			// 永遠不會執行
			//if err != nil {
			//	c.JSON(400, gin.H{"status": err.Error()})
			//	return
			//}
			c.JSON(200, gin.H{"data": prodRes.GetData()})

			// 熔斷
			//configA := hystrix.CommandConfig{
			//	Timeout: 2000,
			//}
			//var prodRes *services.ProdListResponse
			//hystrix.ConfigureCommand("getprods", configA)
			//err := hystrix.Do("getprods", func() error {
			//	prodRes, err = prodService.GetProdsList(context.Background(), &par)
			//	log.Println(err)
			//	return err
			//}, func(err error) error {
			//
			//	prodRes  = defaultProds()
			//	return  nil
			//})
			//if err != nil {
			//	c.JSON(400, gin.H{"status": err.Error()})
			//	return
			//}
			//c.JSON(200, gin.H{"data": prodRes.GetData()})
			//return
		}
	})
	httpService.Init()
	httpService.Run()
}
