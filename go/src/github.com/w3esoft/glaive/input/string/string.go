package string


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
func (input *InputString ) Read() uint8{
	if (input.pos> len(input.value)){
		return 0;
	} else{
		i:=input.pos
		input.pos =input.pos+1;
		return input.value[i];
	}
}
func (input *InputString ) GetPosition() int{
	if !(input.pos> len(input.value)){
		return len(input.value);
	} else{
		return input.pos;
	}
}
