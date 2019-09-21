package node

type Edge struct {
	From *Port
	To   *Port
}

// NewEdge creates a new edge between two ports
func NewEdge(from *Port, to *Port) {
	e := Edge{From: from, To: to}
	e.From.AddOutgoingEdge(e)
	e.To.AddIncomingEdge(&e)
}

// Propagates the value from the start to the end of the edge
func (e *Edge) Propagate() {
	val := e.From.GetValue()
	e.To.SetValue(val)
}
