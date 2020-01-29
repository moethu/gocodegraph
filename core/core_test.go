package core

import (
	"log"
	"testing"
	"time"

	"github.com/moethu/gocodegraph/components"
	"github.com/moethu/gocodegraph/node"
)

func TestGraphModelWithNodes(t *testing.T) {
	results := make(chan node.Result)

	add := components.Addition{}
	add.Init(results, "+")
	mul := components.Multiplication{}
	mul.Init(results, "*")

	n1 := components.Number{}
	n1.Init(results, "a")
	n2 := components.Number{}
	n2.Init(results, "b")

	node.NewEdge(n1.GetOutput(0), add.GetInput(0))
	node.NewEdge(n2.GetOutput(0), add.GetInput(1))
	node.NewEdge(n1.GetOutput(0), mul.GetInput(0))
	node.NewEdge(n2.GetOutput(0), mul.GetInput(1))

	nodes := []node.Node{&n1, &n2, &add, &mul}
	Solve(nodes, false)

	expectations := make(map[string]int)
	expectations["*"] = 25
	expectations["+"] = 10
	expectations["a"] = 5
	expectations["b"] = 5
	awaitResult(t, results, expectations)
}

func awaitResult(t *testing.T, c chan node.Result, expectations map[string]int) {
	ticker := time.NewTicker(5 * time.Second)
	defer func() {
		ticker.Stop()
	}()

	found := 0
	for {
		select {

		case message, _ := <-c:
			if expected, ok := expectations[message.Id]; ok {
				if message.Value != expected {
					t.Error(message.Id, "Expected:", expected, "Received:", message.Value)
					return
				} else {
					log.Println(message.Id, "Expected:", expected, "Received:", message.Value)
					found++
					if found == len(expectations) {
						return
					}
				}
			}
		case <-ticker.C:
			t.Error("Test timed out")
			return
		}

	}
}
