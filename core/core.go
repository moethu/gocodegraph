package core

import (
	"log"
	"reflect"
	"time"

	"github.com/moethu/gocodegraph/node"
)

func Init(nodes []node.Node) {
	for _, n := range nodes {
		n.Init()
	}
}

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

func Solve(nodes []node.Node, verbose bool) {
	start := time.Now()

	solve(nodes, verbose)

	duration := time.Now().Sub(start)
	if verbose {
		log.Println(duration)
	}
}

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
		time.Sleep(3 * time.Second)
	}

	for _, n := range run {
		go n.Solve()
	}

}
