package main

import (
	"fmt"
	"reflect"
)

func main() {

	//Explore a new type
	var x1 float32 = 5.8
	v := reflect.ValueOf(x1)
	fmt.Println("type:", v.Type())
	fmt.Println("Kind is float64?", v.Kind() == reflect.Float64)
	fmt.Println("Value:", v.Float())

	//we can use Interface() to get the interface value back from the reflect object
	interfaceValue := v.Interface()
	switch t := interfaceValue.(type) {
	case float32:
		fmt.Println("Float32 value", t)
	case float64:
		fmt.Println("float64 value", t)
	}

	//v.SetFloat(2.2)
	fmt.Println("v settable?", v.CanSet())

	//get a reflect value object that represents a pointer to x1
	vp := reflect.ValueOf(&x1)
	fmt.Println("Type of vp", vp.Type(), "settable?", vp.CanSet())

	//get a reflect value object representing the value that the pointer refers to
	vpElement := vp.Elem()
	fmt.Println("Type of vpElement", vpElement.Type(), "settable?", vpElement.CanSet())
	vpElement.SetFloat(8.98)
	fmt.Println(x1)

	//For simplicity, the getter provides the largest type of a value. So instead of float32 it returns float64
	x2 := v.Float()
	fmt.Println(reflect.TypeOf(x2))

	//Reflection gets the underlying type instead of the custom/static type
	type myFloat float64
	var x3 myFloat = 5.87
	v = reflect.ValueOf(x3)
	fmt.Println(v.Type(), "kind is float64?", v.Kind() == reflect.Float64)

}

func inspectValue(i interface{}) {
	v := reflect.ValueOf(i)
	fmt.Println("type:", v.Type())
	fmt.Println("Kind is float64?", v.Kind() == reflect.Float64)
	fmt.Println("Value:", v.Float())
}
