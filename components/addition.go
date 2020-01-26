package components

import (
	"fmt"
	"log"
	"reflect"

	"github.com/moethu/gocodegraph/node"
	uuid "github.com/satori/go.uuid"
)

func (n *Addition) Init(c chan node.Result) {
	n.Id = uuid.NewV4().String()
	p1 := node.NewPort(n, "a", reflect.Int, c)
	p2 := node.NewPort(n, "b", reflect.Int, c)
	p3 := node.NewPort(n, "result", reflect.Int, c)
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
	log.Println("Awaiting Values from Channel")
	_, c1 := n.Inputs[0].GetIncomingChannel(0)
	_, c2 := n.Inputs[1].GetIncomingChannel(0)
	a, b := <-c1, <-c2
	log.Println(fmt.Sprintf("Solving %v + %v", a, b))
	n.Outputs[0].SetValue(a.(int) + b.(int))
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
