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

	mapOperators(payload["operators"])
	mapLinks(payload["links"])
}

func mapOperators(data interface{}) []node.Node {
	nodes := []node.Node{}
	d := data.(map[string]interface{})
	for id, operator := range d {
		log.Println(id)
		op := operator.(map[string]interface{})
		prop := op["properties"].(map[string]interface{})
		in := prop["inputs"].(map[string]interface{})
		out := prop["outputs"].(map[string]interface{})
		log.Println(prop["title"], in, out)
		mapPorts(in)
		mapPorts(out)

		switch prop["title"] {
		case "Addition":
			add := components.Addition{}
			nodes = append(nodes, &add)
		case "Number":
			n := components.Number{Value: 3}
			nodes = append(nodes, &n)
		}
	}
	return nodes
}

func mapPorts(data map[string]interface{}) {
	for id, _ := range data {
		log.Println(id)
	}
}

func mapLinks(data interface{}) {
	d := data.(map[string]interface{})
	for _, link := range d {
		//op := operator.(map[string]interface{})
		//node.NewEdge(&num1.Outputs[0], &add.Inputs[0])
		log.Println(link)
	}
}
