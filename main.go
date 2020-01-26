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
	components.InitTypeRegistry()

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
	router.Any("/solver", serveWebsocket)
	router.GET("/components", nodes)
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

// home rendering main page
func home(c *gin.Context) {
	viewertemplate := template.Must(template.ParseFiles("templates/index.html"))
	viewertemplate.Execute(c.Writer, "http://localhost:8000")
}

// mapOperators maps json operators to graph components
func mapOperators(c chan node.Result, data interface{}) map[string]node.Node {
	nodes := make(map[string]node.Node)
	d := data.(map[string]interface{})
	for id, operator := range d {
		op := operator.(map[string]interface{})
		prop := op["properties"].(map[string]interface{})
		in := prop["inputs"].(map[string]interface{})
		out := prop["outputs"].(map[string]interface{})

		n := components.MakeInstance("components." + prop["title"].(string))

		n.Init(c, id)
		mapPorts(in, n, true)
		mapPorts(out, n, false)
		nodes[id] = n
	}
	return nodes
}

// mapPorts maps json ports to graph ports
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

// mapLinks maps json links to graph links by creating edges
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

// nodes returns a list of available components
func nodes(c *gin.Context) {
	list := components.GetComponents()
	res, _ := json.Marshal(list)
	st := string(res)
	c.JSON(200, st)
}
