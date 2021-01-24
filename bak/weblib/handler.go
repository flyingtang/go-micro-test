package weblib

import (
	"github.com/flyingtang/go-micro/bak/prod_service"
	"github.com/flyingtang/go-micro/bak/services"
	"github.com/gin-gonic/gin"
)



func GetProdList(c *gin.Context) {
	var prodReq services.ProdListRequest
	if err := c.Bind(&prodReq); err != nil {
		c.JSON(500, gin.H{"status": err.Error()})
	}else{
		prod_service.NewProdList(5)
	}
	return
}