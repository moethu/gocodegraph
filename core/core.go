package core

import (
	"log"
	"reflect"
	"sync"
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

	if len(run) > 0 {
		var wg sync.WaitGroup

		for i, n := range run {
			wg.Add(1)
			go func(i int, n node.Node) {
				defer wg.Done()
				n.Solve()
				if verbose {
					printNode(n)
				}
				for _, p := range n.GetOutputs() {
					for _, e := range p.Outgoing {
						e.Propagate()
					}
				}
			}(i, n)
		}

		wg.Wait()
	}

	if len(rest) > 0 {
		solve(rest, verbose)
	}

}
