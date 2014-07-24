package main

import (
	. "./common"
	"log"
	"net/rpc"
	"net"
	"net/http"
	"io/ioutil"
	"os"
)

func serverListener_unix() (net.Listener, error) {
	tf, err := ioutil.TempFile("", "packer-plugin")
	if err != nil {
		return nil, err
	}
	path := "/tmp/rpc_benchmark.sock"
	println(path)

	// Close the file and remove it because it has to not exist for
	// the domain socket.
	if err := tf.Close(); err != nil {
		return nil, err
	}
	if err := os.Remove(path); err != nil {
		return nil, err
	}

	return net.Listen("unix", path)
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	ul, e := serverListener_unix()
	if e != nil {
		log.Fatal("listen error:", e)
	}
	func() {
		for {
			conn, _ := ul.Accept()
			go rpc.ServeConn(conn)
		}
	}()
}
