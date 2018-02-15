package input

type Input interface {
	GetPosition() int
	Read() uint8
}
