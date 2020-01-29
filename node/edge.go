package node

type Edge struct {
	From    *Port
	To      *Port
	Channel chan interface{}
}

// NewEdge creates a new edge between two ports
func NewEdge(from *Port, to *Port) {
	e := Edge{From: from, To: to}
	e.Channel = make(chan interface{})
	e.From.AddOutgoingEdge(e)
	e.To.AddIncomingEdge(&e)
}

// Propagates the value from the start to the end of the edge
func (e *Edge) Propagate(val interface{}) {
	// push value down the channel
	e.Channel <- val
}
