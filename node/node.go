package node

type Node interface {
	Solve()
	GetInputs() []Port
	GetOutputs() []Port
	GetId() string
	GetPosition() Location
	Init()
}
