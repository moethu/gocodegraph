package main

import (
	"github.com/moethu/gocodegraph/components"
	"github.com/moethu/gocodegraph/core"
	"github.com/moethu/gocodegraph/node"
)

func main() {

	// setup some components
	add := components.Addition{}
	add2 := components.Multiplication{}
	num1 := components.Number{Value: 3}
	num2 := components.Number{Value: 2}
	var nodes = []node.Node{&add, &num1, &num2, &add2}
	add2.Position = node.Location{X: 0, Y: 0}

	// initialize all nodes
	core.Init(nodes)

	// draw edges to connect components
	node.NewEdge(&num1.Outputs[0], &add.Inputs[0])
	node.NewEdge(&num2.Outputs[0], &add.Inputs[1])
	node.NewEdge(&add.Outputs[0], &add2.Inputs[1])
	node.NewEdge(&num2.Outputs[0], &add2.Inputs[0])

	// solve graph
	core.Solve(nodes, true)
}
