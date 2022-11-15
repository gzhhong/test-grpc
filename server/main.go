package main

import (
	"fmt"
	"context"
	"go-micro.dev/v4"
	proto "test-grpc/proto"
	"time"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"

)

type Social struct{}

func (s *Social) SayHelloWithoutTime(ctx context.Context, req *proto.HelloRequestWithoutTime, resp *proto.HelloResponseWithoutTime) error {
	fmt.Println("request name:", req.Name)
	resp.Greeting = "Hello " + req.Name
	return nil
}

func (s *Social) SayHelloWithTime(ctx context.Context, req *proto.HelloRequestWithTime, resp *proto.HelloResponseWithTime) error {
	fmt.Println("request name:", req.Name)
	atTime := time.Unix(req.Start.Seconds, int64(req.Start.Nanos))
	fmt.Println("request time:", atTime.Format(time.RFC3339))

	resp.Greeting = "Hello " + req.Name + " at " + atTime.Format(time.RFC3339)
	resp.End = timestamp.New(atTime.Add(2 * time.Second))
	return nil
}

func main() {
	serviceName := "HelloServer"

	// new version based on go-micro v4
	ms := micro.NewService(
		micro.Name(serviceName),
		micro.WrapHandler(),
		micro.AfterStart(func() error {
			fmt.Println(serviceName + "_started")
			return nil
		}),
		micro.BeforeStop(func() error {
			fmt.Println(serviceName + "stopping")
			return nil
		}),
		micro.AfterStop(func() error {
			fmt.Println(serviceName + "stopped")
			return nil
		}),
	)

	ms.Init()

	// Register services
	handler := &Social{}
	err := proto.RegisterHelloHandler(ms.Server(), handler)
	if err != nil {
		fmt.Errorf("Failed to register hello service %v", err)
	}

	// Start Server
	if err = ms.Run(); err != nil {
		fmt.Errorf("run hello service %v", err)
	}
}
