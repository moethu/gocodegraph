package components

import (
	"fmt"
	"reflect"

	"github.com/moethu/gocodegraph/node"
)

var typeRegistry = make(map[string]reflect.Type)

func InitTypeRegistry() {
	myTypes := []interface{}{Addition{}, Multiplication{}, Number{}}
	for _, v := range myTypes {
		typeRegistry[fmt.Sprintf("%T", v)] = reflect.TypeOf(v)
	}
}

func MakeInstance(name string) node.Node {
	v := reflect.New(typeRegistry[name])
	n := v.Interface().(node.Node)
	return n
}

func GetComponents() []Comp {
	keys := make([]Comp, len(typeRegistry))

	i := 0
	for k, v := range typeRegistry {
		instance := reflect.New(v)
		n := instance.Interface().(node.Node)
		n.Init()
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
