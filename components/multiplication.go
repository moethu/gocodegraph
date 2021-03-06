package components

import (
	"reflect"

	"github.com/moethu/gocodegraph/node"
)

func (n *Multiplication) Init(c chan node.Result, id string) {
	n.Id = id
	p1 := node.NewPort(n, "a", reflect.Int, c)
	p2 := node.NewPort(n, "b", reflect.Int, c)
	p3 := node.NewPort(n, "result", reflect.Int, c)
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

func (n *Multiplication) Solve() {
	a := n.Inputs[0].AwaitValue().(int)
	b := n.Inputs[1].AwaitValue().(int)
	n.Outputs[0].SetValue(a * b)
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
func (n *Multiplication) GetInput(i int) *node.Port {
	return &n.Inputs[i]
}

func (n *Multiplication) GetOutput(i int) *node.Port {
	return &n.Outputs[i]
}
