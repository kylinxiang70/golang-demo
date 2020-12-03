/**
 * @author xiangqilin
 * @date 2020-11-17
**/
package main

import (
	"fmt"
	"reflect"
)

type Interface interface {
	GetName() string
	GetId() string
}

type Person struct {
	ID   string
	Name string
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetId() string {
	return p.Name
}

func main() {
	p := &Person{ID: "123456", Name: "kylin"}
	t := reflect.TypeOf(p)
	fmt.Println(t.Kind().String())
	fmt.Println(t.Name())
	//fmt.Println(t.Elem())
}
