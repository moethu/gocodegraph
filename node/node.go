package node

type Node interface {
	Solve() bool
	GetInputs() []Port
	GetOutputs() []Port
	GetId() string
	GetPosition() Location
	Init()
}
