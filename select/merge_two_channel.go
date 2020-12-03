/**
 * @author xiangqilin
 * @date 2020-07-11
**/
package main

import "fmt"

func mergeChannel(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case x, ok := <-ch1:
				if !ok {
					ch1 = nil
				} else {
					out <- x
				}
			case x, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				} else {
					out <- x
				}
			}

			if ch1 == nil && ch2 == nil {
				break
			}
		}
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 3; i < 6; i++ {
			ch2 <- i
		}
	}()

	out := mergeChannel(ch1, ch2)

	for x := range out {
		fmt.Println(x)
	}
}
