/**
 * @author xiangqilin
 * @date 2020-07-12
**/
package main

import (
	"fmt"
	"time"
)

func WaitChannel(ch <-chan int, timeout time.Duration) (int, bool) {
	t := time.NewTimer(timeout)
	select {
	case x, ok := <-ch:
		t.Stop()
		return x, ok
	case <-t.C:
		fmt.Println("timeout....")
		return 0, false
	}
}

func main() {
	ch1 := make(chan int)
	go func() {
		time.Sleep(2 * time.Millisecond)
		ch1 <- 1
	}()
	x, ok := WaitChannel(ch1, 2*time.Millisecond)
	fmt.Printf("%v, %v\n", x, ok)
}
