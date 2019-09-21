package node

type Node interface {
	Solve(done chan bool)
	GetInputs() []Port
	GetOutputs() []Port
	GetId() string
	GetPosition() Location
	Init()
}
