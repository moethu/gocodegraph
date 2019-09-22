package node

import (
	"reflect"
)

// Generic Port Interface
type Port struct {
	Parent   Node
	Name     string
	Type     reflect.Kind
	HasValue bool
	value    interface{}
	Outgoing []Edge
	Incoming []*Edge
	Optional bool
}

func NewPort(parent Node, name string, kind reflect.Kind) Port {
	p := Port{Parent: parent, Name: name, Type: kind, HasValue: false}
	p.Outgoing = []Edge{}
	p.Incoming = []*Edge{}
	return p
}

func (p *Port) AddOutgoingEdge(e Edge) {
	p.Outgoing = append(p.Outgoing, e)
}

func (p *Port) AddIncomingEdge(e *Edge) {
	p.Incoming = append(p.Incoming, e)
}

func resolve(n Node) {
	n.Solve()
	for _, p := range n.GetOutputs() {
		for _, o := range p.Outgoing {
			o.Propagate()
			resolve(o.To.Parent)
		}
	}
}

func (p *Port) SetValue(val interface{}) {
	if p.HasValue {
		// invalidate all dependencies
		for _, o := range p.Outgoing {
			o.Propagate()
			resolve(o.To.Parent)
		}
	}
	p.value = val
	p.HasValue = true
}

func (p *Port) GetValue() interface{} {
	return p.value
}
