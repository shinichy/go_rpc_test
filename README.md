### Start RPC server

```
% touch /tmp/rpc_benchmark.sock
% go run server.go
```

### Run benchmark test

```
% go test -bench .
BenchmarkFuncCall	2000000000	         1.31 ns/op
BenchmarkRpcViaTCP	   50000	     45962 ns/op
BenchmarkRpcViaUnixDomainSocket	   50000	     30653 ns/op
ok  	_/Users/shinichi/Code/go_rpc_test	7.713s
```
