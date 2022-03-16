package main

import (
	"context"
	"log"
	v1 "lotteryKratos/api/helloworld/v1"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	srcgrpc "google.golang.org/grpc"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(cli)

	connGRPC, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///lotteryKratos"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer connGRPC.Close()

	for {

		callGRPC(connGRPC)
		time.Sleep(time.Second)
	}
}

func callGRPC(conn *srcgrpc.ClientConn) {

	client := v1.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &v1.HelloRequest{Name: "kaiduo"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] SayHello %+v\n", reply)
}
