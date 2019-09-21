package components

import (
	"reflect"

	"github.com/satori/go.uuid"
	"github.com/moethu/gocodegraph/node"
)

type Number struct {
	Inputs   []node.Port
	Outputs  []node.Port
	Value    int
	Id       string
	Position node.Location
}

func (n *Number) Solve() {
	n.Outputs[0].SetValue(n.Value)
}

func (n *Number) Init() {
	n.Id = uuid.NewV4().String()
	p3 := node.NewPort(n, "constant", reflect.Int)
	n.Inputs = []node.Port{}
	n.Outputs = []node.Port{p3}
}

func (n *Number) GetId() string {
	return n.Id
}

func (n *Number) GetPosition() node.Location {
	return n.Position
}

func (n *Number) GetInputs() []node.Port {
	return n.Inputs
}

func (n *Number) GetOutputs() []node.Port {
	return n.Outputs
}
