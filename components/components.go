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

func GetComponents() []string {
	keys := make([]string, len(typeRegistry))

	i := 0
	for k, _ := range typeRegistry {
		keys[i] = k
		i++
	}
	return keys
}
