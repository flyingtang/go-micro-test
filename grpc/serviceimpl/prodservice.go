package serviceimpl

import (
	"context"
	"go-micro/grpc/services"
	"strconv"
)

func NewProd(id int32, name string) *services.ProdModel {
	return &services.ProdModel{
		ProdId:   id,
		ProdName: name,
	}
}

type ProdService struct{}

func (s *ProdService) GetProdsList(context context.Context, in *services.ProdListRequest, res *services.ProdListResponse) error {
	models := make([]*services.ProdModel, 0)

	var i int32
	for i = 0; i < in.Size; i++ {
		models = append(models, NewProd(100+1, "prodname"+strconv.Itoa(100+int(i))))
	}
	res.Data = models
	return nil
}
