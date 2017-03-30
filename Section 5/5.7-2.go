package main

import (
	"fmt"
	"reflect"
)

type Printer interface {
	Print(s string)
}

type pStruct struct {
	s string
}

func (p *pStruct) Print(s string) {
	p.s = s
	fmt.Println(s)
}

func main() {
	p := new(pStruct)
	inspectType(p)
}

func inspectType(obj interface{}) {

	v := reflect.ValueOf(obj)
	t := v.Type()
	myInterface := reflect.TypeOf((*Printer)(nil)).Elem()

	fmt.Println("obj implements Printer?", t.Implements(myInterface))

	if t.Implements(myInterface) {
		printFunc := v.MethodByName("Print")
		args := []reflect.Value{reflect.ValueOf("Printing Hello")}
		printFunc.Call(args)
	}

}
