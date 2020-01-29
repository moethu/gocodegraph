package node

import (
	"reflect"
	"testing"
)

func TestNewEdge(t *testing.T) {
	c := make(chan Result)
	p1 := NewPort(nil, "a", reflect.Int64, c)
	p2 := NewPort(nil, "b", reflect.Int64, c)
	NewEdge(&p1, &p2)
	if p1.ResultChannel != c {
		t.Error("channel hasn't been set to port")
	}
	if p2.ResultChannel != c {
		t.Error("channel hasn't been set to port")
	}
	e := p2.GetIncomingEdges()
	if len(e) == 0 {
		t.Error("Edge hasn't been applied correctly")
	}
	if e[0].Channel == nil {
		t.Error("channel hasn't been created for edge")
	}
}
