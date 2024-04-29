package main

import (
	"context"
	"flag"
	"log"

	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", ":8972", "service address")
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	//serverConfig := []constant.ServerConfig{
	//	{
	//		IpAddr: "10.0.1.84",
	//		Port:   38848,
	//	},
	//}
	//clientConfig := constant.ClientConfig{
	//	NamespaceId:         "game01",
	//	TimeoutMs:           5000,
	//	NotLoadCacheAtStart: true,
	//}

	//nacoDiscovery, err := nclient.NewNacosDiscovery("game", "", "jerry-test", clientConfig, serverConfig)
	//if err != nil {
	//	panic(err)
	//}

	flag.Parse()

	*addr = "10002"

	peerDiscovery, _ := client.NewPeer2PeerDiscovery("tcp@172.27.0.1:"+*addr, "")

	xclient := client.NewXClient("game", client.Failtry, client.RandomSelect, peerDiscovery, client.DefaultOption)
	defer xclient.Close()

	//args := &Args{A: 10, B: 20}
	//var reply int

	var reply []byte

	err := xclient.Call(context.Background(), "Hello", "jerry", &reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	//fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)

	//args = &Args{50, 20}
	//var quo Quotient
	//err = xclient.Call(context.Background(), "Div", args, &quo)
	//if err != nil {
	//	log.Fatalf("failed to call: %v", err)
	//}

	//fmt.Printf("%d * %d = %d...%d\n", args.A, args.B, quo.Quo, quo.Rem)
}
