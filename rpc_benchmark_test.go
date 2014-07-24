package main

import (
	. "./common"
	"log"
	"net/rpc"
	"fmt"
	"testing"
)

func BenchmarkFuncCall(b *testing.B) {
	var arith = new(Arith)
	args := &Args{7, 8}
	var reply int
	for i := 0; i < b.N; i++ {
		err := arith.Multiply(args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
	}
}

func BenchmarkRpcViaTCP(b *testing.B) {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{7, 8}
	var reply int
	for i := 0; i < b.N; i++ {
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
	}
}

func BenchmarkRpcViaUnixDomainSocket(b *testing.B) {
	client, err := rpc.Dial("unix", "/tmp/rpc_benchmark.sock")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{7, 8}
	var reply int
	for i := 0; i < b.N; i++ {
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
	}
}

func TestRpc(*testing.T) {
	client, err := rpc.Dial("unix", "/tmp/rpc_benchmark.sock")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{7, 8}
	var reply int
	for i := 0; i < 10000; i++ {
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
	}
}
