package main

import "fmt"

type Visitor interface {
	Visit(VisitorFunc) error
}

type VisitorFunc func() error

type VisitorList []Visitor

func (l VisitorList) Visit(f VisitorFunc) error {
	for i := range l {
		if err := l[i].Visit(func() error {
			fmt.Println("In visit list before f")
			f()
			fmt.Println("In visit list after f")
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

type Visitor1 struct {
}

func (v Visitor1) Visit(f VisitorFunc) error {
	fmt.Println("In Visit1 before f")
	f()
	fmt.Println("In Visit1 after f")
	return nil
}

type Visitor2 struct {
	visitor Visitor
}

func (v Visitor2) Visit(f VisitorFunc) error {
	v.visitor.Visit(func() error {
		fmt.Println("In visitor 2 before f")
		f()
		fmt.Println("In visitor 2 before f")
		return nil
	})
	return nil
}

type Visitor3 struct {
	visitor Visitor2
}

func (v Visitor3) Visit(f VisitorFunc) error {
	v.visitor.Visit(func() error {
		fmt.Println("In visitor3 before f")
		f()
		fmt.Println("In visitor3 before f")
		return nil
	})
	return nil
}

func main() {
	var visitor1 Visitor1
	var visitor2 Visitor2
	var visitor3 Visitor3
	var visitors []Visitor

	visitors = append(visitors, visitor1)
	visitorList := VisitorList(visitors)
	visitor2 = Visitor2{visitorList}
	visitor3 = Visitor3{visitor2}
	

	visitorFunc := func() error{
		fmt.Println("In func")
		return nil
	}

	visitor3.Visit(visitorFunc)
}
