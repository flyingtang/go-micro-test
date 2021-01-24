package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
	"log"
)

func callApi(addr string, path string, method string) {
	//myHttp.
}

func callApi2(s selector.Selector) {
	myClient := http.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)

	fmt.Println(s.String())
	req := myClient.NewRequest("prodservice", "/v1/prods",
		ser.ProdListRequest{Size: 5})
	var resp models.ProdListResponse
	err := myClient.Call(context.Background(), req, &resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Data)
}

func main() {
	conReg := consul.NewRegistry(registry.Addrs("192.168.2.112:8500"))
	selector := selector.NewSelector(
		selector.Registry(conReg),
		selector.SetStrategy(selector.Random),
	)


	callApi2(selector)

	//
	//for{
	//	proService, err  := conReg.GetService("prodservice")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	next := selector.RoundRobin(proService)
	//	node , err := next()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(node)
	//	time.Sleep(time.Second)
	//}

}
