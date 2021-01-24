package weblib

import (
	"github.com/flyingtang/go-micro/bak/services"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(prodService services.ProdService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["prodservice"] = prodService
		c.Next()
	}
}