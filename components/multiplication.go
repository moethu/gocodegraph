package components

import (
	"math/rand"
	"reflect"
	"strconv"

	"../node"
)

func (n *Multiplication) Init() {
	n.Id = strconv.Itoa(rand.Intn(100))
	p1 := node.NewPort(n, "a", reflect.Int)
	p2 := node.NewPort(n, "b", reflect.Int)
	p3 := node.NewPort(n, "result", reflect.Int)
	n.Inputs = []node.Port{p1, p2}
	n.Outputs = []node.Port{p3}
}

type Multiplication struct {
	Inputs   []node.Port
	Outputs  []node.Port
	Id       string
	Position node.Location
}

func (n *Multiplication) GetPosition() node.Location {
	return n.Position
}

func (n *Multiplication) Solve() bool {
	a := n.Inputs[0].GetValue().(int)
	b := n.Inputs[1].GetValue().(int)
	n.Outputs[0].SetValue(a * b)
	return true
}

func (n *Multiplication) GetId() string {
	return n.Id
}
func (n *Multiplication) GetInputs() []node.Port {
	return n.Inputs
}

func (n *Multiplication) GetOutputs() []node.Port {
	return n.Outputs
}
