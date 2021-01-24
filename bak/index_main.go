package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"net/http"
)

func main() {


	consulReg := consul.NewRegistry(registry.Addrs("192.168.2.112:8500"))

	ginRouter := gin.Default()
	ginRouter.Handle(http.MethodGet, "/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "index",
		})
	})

	ginRouter.Handle(http.MethodGet, "/news", func(c *gin.Context) {
		c.Data(http.StatusOK, "", []byte("news"))
	})

	server := web.NewService(
		web.Name("prodservice"),
		web.Address(":8000"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)
	server.Run()
}
