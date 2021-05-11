package main

import (
	"fmt"
	"log"

	"github.com/goldenbean/netty-pb/pkg/gen"
	"google.golang.org/protobuf/proto"
)

func main() {

	fmt.Println("Hello")

	req := &gen.SubscribeReq{
		SubReqID:    1,
		UserName:    "bob",
		ProductName: "netty book for protobuf",
		Address:     []string{"Beijing, NanJing, ShenZhen"},
	}

	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatal("Marshal error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	req2 := &gen.SubscribeReq{}
	err = proto.Unmarshal(data, req2)
	if err != nil {
		log.Fatal("Unmarshal error: ", err)
	}


	fmt.Println(req2.GetUserName())
	fmt.Println(req2.GetProductName())
}
