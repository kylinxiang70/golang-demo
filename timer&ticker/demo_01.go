/**
 * @author xiangqilin
 * @date 2020-07-11
**/

package main

import (
	"fmt"
	"time"
)

func main() {
	//初始化定时器
	t := time.NewTimer(2 * time.Second)
	//当前时间
	now := time.Now()
	fmt.Printf("Type: %T, Now time : %v.\n", now, now)

	expire := <-t.C
	fmt.Printf("Type %T, Expiration time: %v.\n", expire, expire)
}
