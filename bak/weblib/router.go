package weblib

import (
	"github.com/flyingtang/go-micro/bak/services"
	"github.com/gin-gonic/gin"
)

func NewGinRouter(prodService services.ProdService) *gin.Engine {
	g := gin.Default()
	g.Use(InitMiddleware(prodService))
	v1 := g.Group("/v1")
	v1.POST("/prods", GetProdList)
	return g
}
