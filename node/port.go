package node

import (
	"errors"
	"log"
	"reflect"
)

// Generic Port Interface
type Port struct {
	Parent        Node
	Name          string
	Type          reflect.Kind
	HasValue      bool
	value         interface{}
	Outgoing      []Edge
	Incoming      []*Edge
	Optional      bool
	ResultChannel chan Result
}

type Result struct {
	Id    string
	Port  string
	Value interface{}
}

// NewPort creates a new port
func NewPort(parent Node, name string, kind reflect.Kind, c chan Result) Port {
	p := Port{Parent: parent, Name: name, Type: kind, HasValue: false}
	p.Outgoing = []Edge{}
	p.Incoming = []*Edge{}
	p.ResultChannel = c
	return p
}

// AddOutgoingEdge adds an edge to a port as outgoing
func (p *Port) AddOutgoingEdge(e Edge) {
	p.Outgoing = append(p.Outgoing, e)
}

// AddIncomingEdge adds and edge to a port as incoming
func (p *Port) AddIncomingEdge(e *Edge) {
	p.Incoming = append(p.Incoming, e)
}

// SetValue sets a ports value and propagates it down the line
func (p *Port) SetValue(val interface{}) {
	p.value = val
	p.HasValue = true
	p.ResultChannel <- Result{Id: p.Parent.GetId(), Port: p.Name, Value: val}
	for _, e := range p.Outgoing {
		e.Propagate(val)
	}
}

// GetIncomingEdges returns all incoming edges
func (p *Port) GetIncomingEdges() []*Edge {
	return p.Incoming
}

// GetIncomingChannel gets an incoming channel for a specific edge
func (p *Port) GetIncomingChannel(edge int) (error, chan interface{}) {
	if len(p.Incoming) > edge {
		return nil, p.Incoming[edge].Channel
	} else {
		return errors.New("Out of bounds"), nil
	}
}

// GetValue returns the ports value
func (p *Port) GetValue() interface{} {
	return p.value
}

// AwaitValue awaits the value of a channel
func (p *Port) AwaitValue() interface{} {
	err, c := p.GetIncomingChannel(0)
	if err != nil {
		log.Println(err)
		return nil
	}
	return <-c
}
