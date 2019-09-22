package main

import (
	"context"
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moethu/gocodegraph/components"
	"github.com/moethu/gocodegraph/core"
	"github.com/moethu/gocodegraph/node"
)

func main() {

	// // setup some components
	// add := components.Addition{}
	// add2 := components.Multiplication{}
	// num1 := components.Number{Value: 3}
	// num2 := components.Number{Value: 2}
	// var nodes = []node.Node{&add, &num1, &num2, &add2}
	// add2.Position = node.Location{X: 0, Y: 0}

	// // initialize all nodes
	// core.Init(nodes)

	// // draw edges to connect components
	// node.NewEdge(&num1.Outputs[0], &add.Inputs[0])
	// node.NewEdge(&num2.Outputs[0], &add.Inputs[1])
	// node.NewEdge(&num1.Outputs[0], &add2.Inputs[1])
	// node.NewEdge(&num2.Outputs[0], &add2.Inputs[0])

	// // solve graph
	// core.Solve(nodes, true)

	flag.Parse()
	log.SetFlags(0)

	router := gin.Default()
	port := ":8000"
	srv := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  600 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	router.Static("/static/", "./static/")
	router.GET("/", home)
	router.POST("/solve", solve)
	log.Println("Starting HTTP Server on Port 8000")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}

// Home route, loading template and serving it
func home(c *gin.Context) {
	viewertemplate := template.Must(template.ParseFiles("templates/index.html"))
	viewertemplate.Execute(c.Writer, "http://localhost:8000")
}

func solve(c *gin.Context) {
	var payload map[string]interface{}
	bdata, _ := c.GetRawData()
	err := json.Unmarshal(bdata, &payload)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "error deserializing json"})
		return
	}

	nodes := mapOperators(payload["operators"])
	mapLinks(payload["links"], nodes)
	ns := []node.Node{}
	for _, value := range nodes {
		ns = append(ns, value)
	}
	core.Solve(ns, true)
	// TODO collect result map from nodes and return
	c.JSON(200, gin.H{"status": "OK"})
}

func mapOperators(data interface{}) map[string]node.Node {
	nodes := make(map[string]node.Node)
	d := data.(map[string]interface{})
	for id, operator := range d {
		op := operator.(map[string]interface{})
		prop := op["properties"].(map[string]interface{})
		in := prop["inputs"].(map[string]interface{})
		out := prop["outputs"].(map[string]interface{})

		var n node.Node
		switch prop["title"] {
		case "Addition":
			n = &components.Addition{}
		case "Number":
			n = &components.Number{Value: 3}
		}
		n.Init()
		mapPorts(in, n, true)
		mapPorts(out, n, false)
		nodes[id] = n
	}
	return nodes
}

func mapPorts(data map[string]interface{}, node node.Node, in bool) {
	ctr := 0
	for id, _ := range data {
		if in {
			node.GetInput(ctr).Name = id
		} else {
			node.GetOutput(ctr).Name = id
		}
		ctr++
	}
}

func mapLinks(data interface{}, nodes map[string]node.Node) {
	d := data.(map[string]interface{})
	for _, link := range d {
		op := link.(map[string]interface{})
		n1 := nodes[op["fromOperator"].(string)]
		n2 := nodes[op["toOperator"].(string)]
		var p1 *node.Port
		var p2 *node.Port
		for i, port := range n1.GetOutputs() {
			if port.Name == op["fromConnector"] {
				p1 = n1.GetOutput(i)
			}
		}
		for i, port := range n2.GetInputs() {
			if port.Name == op["toConnector"] {
				p2 = n2.GetInput(i)
			}
		}
		node.NewEdge(p1, p2)
	}
}
