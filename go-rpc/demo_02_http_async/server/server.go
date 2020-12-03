/**
 * @author xiangqilin
 * @date 2020-07-12
**/
package server

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

type Quotient struct {
	Quo, Rem int
}

type Args struct {
	A, B int
}

type Math struct{}

func (d *Math) Div(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("Error: zero divisor.")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	rpc.RegisterName("Math", new(Math))
	rpc.HandleHTTP()

	err := http.ListenAndServe("localhost:18888", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
