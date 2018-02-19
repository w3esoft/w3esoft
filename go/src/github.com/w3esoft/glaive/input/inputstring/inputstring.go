package inputstring


type InputString struct {
	value string
	pos   int
}
func  New(value string) *InputString {
	i := InputString{}
	i.pos = 0
	i.value=value;
	return &i
}
func (input *InputString ) Read() byte{
	l :=len(input.value)
	if l > input.pos {
		value  :=  input.value[input.pos]
		input.pos =input.pos+1
		return value;
	} else{
		return 0;
	}
}
func (input *InputString ) GetPosition() int{
	l :=len(input.value)
	if l > input.pos{
		return input.pos;
	} else{
		return l;
	}
}
