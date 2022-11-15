package main


import (
	"fmt"
	"context"
	"go-micro.dev/v4"
	proto "test-grpc/proto"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
)

func main () {
	serviceName := "HelloClient"
	ms := micro.NewService(
		micro.Name(serviceName),
	)
	ms.Init()
	fmt.Println("Micro service started, with name: "+serviceName)

	helloService := proto.NewHelloService("HelloServer", ms.Client())
	respWithoutTime, err := helloService.SayHelloWithoutTime(context.Background(), &proto.HelloRequestWithoutTime{Name: "James"})
	if err != nil {
		fmt.Errorf("Error happened with detail %v", err)
	}

	fmt.Printf("The response greeting is %v", respWithoutTime.Greeting)
	// fmt.Printf("The response time is %v", time.Unix(resp.End.Seconds, int64(resp.End.Nanos)).Format(time.RFC3339))
	fmt.Println("Success the test on SayHelloWithoutTime!")

	respWithTime, err := helloService.SayHelloWithTime(context.Background(), &proto.HelloRequestWithTime{Name: "James", Start: timestamp.Now()})
	if err != nil {
		fmt.Errorf("Error happened with detail %v", err)   
	}

	fmt.Printf("The response greeting is %v", respWithTime.Greeting)
	// fmt.Printf("The response time is %v", time.Unix(resp.End.Seconds, int64(resp.End.Nanos)).Format(time.RFC3339))
	fmt.Println("Success the test on SayHelloWithTime!")
}