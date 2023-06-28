package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}
type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// fill reply pointer to sent the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// create a new RPC server
	timeserver := new(TimeServer)
	// Register RPC server
	rpc.Register(timeserver)
	rpc.HandleHTTP()
	// listen for request on port 1234
	l, e := net.Listen("tcp", ":2233")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
