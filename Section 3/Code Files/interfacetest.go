package main

import "fmt"

type Node interface {
	SetValue(v int)
	GetValue() int	
}

//type SLLNode
type SLLNode struct {
	next  *SLLNode
	value int
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSLLNode()*SLLNode{
	return new(SLLNode)
}

//type PowerNode
type PowerNode struct {
	next  *PowerNode
	value int
}

func (sNode *PowerNode) SetValue(v int) {
	sNode.value = v * 10
}

func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

func NewPowerNode()*PowerNode{
	return new(PowerNode)
}

func main() {
	var node Node
	node = NewSLLNode()
	node.SetValue(4)
	fmt.Println("Node is of value ",node.GetValue())
	
	
	node = NewPowerNode()
	node.SetValue(5)
	fmt.Println("Node is of value ",node.GetValue())
	
	if n,ok := node.(*PowerNode); ok {
		fmt.Println(n.value)
	}
}