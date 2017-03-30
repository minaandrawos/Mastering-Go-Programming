package main

import (
	"errors"
	"fmt"
)

var ErrInvalidNode = errors.New("Node is not valid")

type Node interface {
	SetValue(v int) error
	GetValue() int
}

//type SLLNode
type SLLNode struct {
	next         *SLLNode
	value        int
	SNodeMessage string
}

func (sNode *SLLNode) SetValue(v int) error {
	if sNode == nil {
		return ErrInvalidNode
	}
	sNode.value = v
	return nil
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSLLNode() *SLLNode {
	return &SLLNode{SNodeMessage: "This is a message from the normal Node"}
}

//type PowerNode
type PowerNode struct {
	next         *PowerNode
	value        int
	PNodeMessage string
}

func (sNode *PowerNode) SetValue(v int) error {
	if sNode == nil {
		return ErrInvalidNode
	}
	sNode.value = v * 10
	return nil

}

func (sNode *PowerNode) GetValue() int {

	return sNode.value
}

func NewPowerNode() *PowerNode {
	return &PowerNode{PNodeMessage: "This is a message from the power Node"}
}

func main() {

	/*
		var n Node
		var sllnode *SLLNode
		n = sllnode
		fmt.Println(n.SetValue(4))
	*/

	n := createNode(5)

	switch n := n.(type) {
	case *SLLNode:
		fmt.Println("Type is SLLNode , message:", n.SNodeMessage)
	case *PowerNode:
		fmt.Println("Type is PowerNode, message:", n.PNodeMessage)

	}

	//sNode := &SLLNode{}
	//n = sNode

}

func createNode(v int) Node {
	sn := NewSLLNode()
	sn.SetValue(v)
	return sn
}
