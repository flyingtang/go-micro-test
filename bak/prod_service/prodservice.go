package prod_service

import (
	"strconv"
)

type ProdService struct {
	ProdId int64 `json:"prod_id"`
	ProdName string `json:"prod_name"`
}

func NewProd(id int64, name string) *ProdService {
	return &ProdService{
		ProdId:   id,
		ProdName: name,
	}
}

func NewProdList(n int) (list []ProdService){
	for i := 0 ; i < n; i++ {
		list =append(list, ProdService{
			ProdId:   int64(100+i),
			ProdName: "prodname"+ strconv.Itoa(100+i),
		})
	}
	return
}