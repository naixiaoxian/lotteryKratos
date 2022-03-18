package main

import (
	"context"
	"log"
	activity "lotteryKratos/api/activity/v1"
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
	callGRPCActivity(connGRPC)
	for {
		//callGRPC(connGRPC)
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

func callGRPCActivity(conn *srcgrpc.ClientConn) {
	activityClient := activity.NewActivityClient(conn)
	res, err := activityClient.GetActivity(context.Background(), &activity.GetActivityRequest{Id: 4})
	if err != nil {
		log.Fatal(err)
	}
	if res.Activity != nil {
		log.Println("data find", res)
	} else {
		rep, err := activityClient.CreateActivity(context.Background(), &activity.CreateActivityRequest{
			ActivityName:  "kd",
			ActivityDesc:  "kddsec",
			BeginDateTime: "2020-01-01 00:00:00",
			EndDateTime:   "2020-01-01 00:00:00",
			StockCount:    10,
			TakeCount:     0,
			State:         0,
			Creator:       "kaid",
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[grpc] activity %+v\n", rep)
	}

}
