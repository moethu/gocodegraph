package core

import (
	"log"
	"reflect"
	"time"

	"github.com/moethu/gocodegraph/node"
)

// Init initializes all nodes
func Init(nodes []node.Node) {
	for _, n := range nodes {
		n.Init()
	}
}

// printNode helps to print the definition and values of a node
func printNode(n node.Node) {
	log.Println(" -- ", n.GetId(), reflect.TypeOf(n))
	log.Println("     Position", n.GetPosition())
	log.Println("     Inputs")
	for _, p := range n.GetInputs() {
		log.Println("       ", p.Name, p.Type, p.GetValue())
	}
	log.Println("     Outputs")
	for _, p := range n.GetOutputs() {
		log.Println("       ", p.Name, p.Type, p.GetValue())
	}
}

// Solve solves the entire graph
func Solve(nodes []node.Node, verbose bool) {
	start := time.Now()

	solve(nodes, verbose)

	duration := time.Now().Sub(start)
	if verbose {
		log.Println(duration)
	}
}

// solve is an internal function first solving all nodes awaiting inputs
// while those nodes are opening channels to await input from nodes which don't require inputs
func solve(nodes []node.Node, verbose bool) {
	rest := []node.Node{}
	run := []node.Node{}

	for _, n := range nodes {
		ready := true
		connected := true

		for _, p := range n.GetInputs() {
			if !p.HasValue {
				ready = false
			}
			if len(p.Incoming) == 0 && !p.Optional {
				connected = false
			}
		}

		if ready {
			run = append(run, n)
		} else {
			if connected {
				rest = append(rest, n)
			}
		}
	}

	for _, n := range rest {
		go n.Solve()
	}

	// TODO: should check if all nodes are initialized and awaiting input

	for _, n := range run {
		go n.Solve()
	}
}
