package input

type Input interface {
	GetPosition() int
	Read() int
}
