package main

import "fmt"

type magicStore struct {
	value interface{}
	name  string
}

func (ms *magicStore) SetValue(v interface{}) {
	ms.value = v
}

func (ms *magicStore) GetValue() interface{} {
	return ms.value
}

func NewMagicStore() *magicStore {
	return new(magicStore)
}

func main() {
	mstore := NewMagicStore()
	mstore.SetValue("Hello")
	fmt.Println(mstore.GetValue())

}
