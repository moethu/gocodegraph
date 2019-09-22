package node

type Node interface {
	Solve()
	GetInputs() []Port
	GetOutputs() []Port
	GetInput(i int) *Port
	GetOutput(i int) *Port
	GetId() string
	GetPosition() Location
	Init()
}
