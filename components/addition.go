package components

import (
	"reflect"

	"github.com/moethu/gocodegraph/node"
	uuid "github.com/satori/go.uuid"
)

func (n *Addition) Init() {
	n.Id = uuid.NewV4().String()
	p1 := node.NewPort(n, "a", reflect.Int)
	p2 := node.NewPort(n, "b", reflect.Int)
	p3 := node.NewPort(n, "result", reflect.Int)
	n.Inputs = []node.Port{p1, p2}
	n.Outputs = []node.Port{p3}
}

type Addition struct {
	Inputs   []node.Port
	Outputs  []node.Port
	Id       string
	Position node.Location
}

func (n *Addition) GetPosition() node.Location {
	return n.Position
}

func (n *Addition) Solve() {
	a := n.Inputs[0].GetValue().(int)
	b := n.Inputs[1].GetValue().(int)
	n.Outputs[0].SetValue(a + b)
}

func (n *Addition) GetId() string {
	return n.Id
}
func (n *Addition) GetInputs() []node.Port {
	return n.Inputs
}

func (n *Addition) GetOutputs() []node.Port {
	return n.Outputs
}
func (n *Addition) GetInput(i int) *node.Port {
	return &n.Inputs[i]
}

func (n *Addition) GetOutput(i int) *node.Port {
	return &n.Outputs[i]
}
