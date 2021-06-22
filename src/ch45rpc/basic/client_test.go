package main

import (
	"ch45rpc/basic/string-service"
	"fmt"
	"log"
	"net/rpc"
	"testing"
)

func TestClient(t *testing.T) {
	client, err := rpc.DialHTTP("tcp","127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//Synchronous call
	stringReq := &service.StringRequest{"A","B"}
	var reply string
	err = client.Call("StringService.Concat",stringReq,&reply)
	if err != nil {
		log.Fatal("StringService error:", err)
	}
	fmt.Printf("StringService Concat : %s concat %s = %s\n",stringReq.A,stringReq.B,reply)


	//async call
	stringReq = &service.StringRequest{"ACD","BDF"}
	call := client.Go("StringService.Diff",stringReq,&reply,nil)
	_ = <-call.Done
	fmt.Printf("StringService Diff : %s diff %s = %s\n",stringReq.A,stringReq.B,reply)
}