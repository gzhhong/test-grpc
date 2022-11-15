## This is a test code using go micro 

### Purpose
When upgrading to go micro v4.6 the timestamppb.timestamp will cause errors and fail the test.

### How to run
```
go mod tidy

protoc --micro_out=. --go_out=. ./proto/hello.proto
```

then

```
go run ./server
```

and 

```
go run ./client
```

### Output
when using the go micro v4.5
The output of client will be:
```
Micro service started, with name: HelloClient
The response greeting is Hello James say hello without time
Success the test on SayHelloWithoutTime!
The response greeting is Hello James say hello with time at 2022-11-15T11:49:15+08:00
Success the test on SayHelloWithTime!
```

The output of server will be:
```
2022-11-15 11:49:08  file=v4@v4.5.0/service.go:206 level=info Starting [service] HelloServer
2022-11-15 11:49:08  file=server/rpc_server.go:820 level=info Transport [http] Listening on [::]:64502
2022-11-15 11:49:08  file=server/rpc_server.go:840 level=info Broker [http] Connected to 127.0.0.1:64503
2022-11-15 11:49:08  file=server/rpc_server.go:654 level=info Registry [mdns] Registering node: HelloServer-63ccb7b2-ae06-4915-94e2-adf2c37995b2
HelloServer_started
request name: James say hello without time
request name: James say hello with time
request time: 2022-11-15T11:49:15+08:00
```

When using go micro v4.6
The output of client will be:
```
Micro service started, with name: HelloClient
The response greeting is Hello James say hello without time
Success the test on SayHelloWithoutTime!
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x28 pc=0x4774083]

goroutine 1 [running]:
main.main()
	/Users/zhihong/workspace/test-grpc/client/main.go:35 +0x343
exit status 2
```

and output of server will be:
```
2022-11-15 11:52:43  file=v4@v4.6.0/service.go:206 level=info Starting [service] HelloServer
2022-11-15 11:52:43  file=server/rpc_server.go:820 level=info Transport [http] Listening on [::]:64550
2022-11-15 11:52:43  file=server/rpc_server.go:840 level=info Broker [http] Connected to 127.0.0.1:64551
2022-11-15 11:52:43  file=server/rpc_server.go:654 level=info Registry [mdns] Registering node: HelloServer-a2718b68-a8a2-4eff-afa5-7bb435d46cad
HelloServer_started
request name: James say hello without time
```