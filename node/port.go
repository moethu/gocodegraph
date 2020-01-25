package node

import (
	"errors"
	"log"
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

func (p *Port) SetValue(val interface{}) {
	p.value = val
	p.HasValue = true
	for _, e := range p.Outgoing {
		e.Propagate(val)
	}
}

func (p *Port) GetIncomingEdges() []*Edge {
	return p.Incoming
}

func (p *Port) GetIncomingChannel(edge int) (error, chan interface{}) {
	if len(p.Incoming) > edge {
		return nil, p.Incoming[edge].Channel
	} else {
		return errors.New("Out of bounds"), nil
	}
}

func (p *Port) GetValue() interface{} {
	log.Println("awaiting value")
	return <-p.Incoming[0].Channel
}
