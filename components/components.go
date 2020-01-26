package components

import (
	"fmt"
	"log"
	"reflect"

	"github.com/moethu/gocodegraph/node"
)

// typeregistry for available node types
var typeRegistry = make(map[string]reflect.Type)

// InitTypeRegistry loads all available nodes into the registry
// TODO: should use reflection to get implemented types
func InitTypeRegistry() {
	myTypes := []interface{}{Addition{}, Multiplication{}, Number{}}
	for _, v := range myTypes {
		log.Println("Loading", reflect.TypeOf(v))
		typeRegistry[fmt.Sprintf("%T", v)] = reflect.TypeOf(v)
	}
}

// MakeInstance creates a component by name
func MakeInstance(name string) node.Node {
	v := reflect.New(typeRegistry[name])
	n := v.Interface().(node.Node)
	return n
}

// GetComponents returns a list of all available components and their properties
func GetComponents() []Comp {
	keys := make([]Comp, len(typeRegistry))

	i := 0
	for k, v := range typeRegistry {
		instance := reflect.New(v)
		n := instance.Interface().(node.Node)
		n.Init(nil, "")
		c := Comp{Name: k}
		for _, p := range n.GetInputs() {
			io := Io{Name: p.Name, Type: p.Type.String()}
			c.I = append(c.I, io)
		}
		for _, p := range n.GetOutputs() {
			io := Io{Name: p.Name, Type: p.Type.String()}
			c.O = append(c.O, io)
		}
		keys[i] = c
		i++
	}
	return keys
}

type Comp struct {
	Name string
	I    []Io
	O    []Io
}

type Io struct {
	Name string
	Type string
}
