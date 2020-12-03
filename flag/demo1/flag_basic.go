/**
 * @author xiangqilin
 * @date 2020-07-24
**/
package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	// 第 1 个参数是用于存储该命令参数值的地址，具体到这里就是在前面声明的变量name的地址了，由表达式&name表示。
	// 第 2 个参数是为了指定该命令参数的名称，这里是name。
	// 第 3 个参数是为了指定在未追加该命令参数时的默认值，这里是everyone。
	// 第 4 个函数参数，即是该命令参数的简短说明了。
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello %s.\n", name)
}
