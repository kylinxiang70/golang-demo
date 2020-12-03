/**
 * @author xiangqilin
 * @date 2020-07-12
**/
package client

import (
	"fmt"
	"log"
	"net/rpc"
)

type Quotient struct {
	Quo, Rem int
}

type Args struct {
	A, B int
}

func main() {
	args := Args{5, 6}

	var quo Quotient
	client, err := rpc.DialHTTP("tcp", "localhost:18888")
	if err != nil {
		log.Fatal(err.Error())
	}

	// async call
	call := client.Go("Math.Div", &args, &quo, nil)
	fmt.Println("Async test text...")
	// wait async call finish
	<-call.Done
	fmt.Printf("quotient: %v, remainder: %v.\n", quo.Quo, quo.Rem)
}
