package inputfile

import (
	"os"
	"io"
)

type InputFile struct {
	File  *os.File
	isClose bool
	pos  int
}

func  New(value string) (*InputFile , error) {
	file, err := os.Open(value)
	input := InputFile{};
	input.File=file
	input.isClose =false
	input.pos=0
	return &input ,err
}


func (input *InputFile ) Read() byte{
	if input.isClose{
		return 0
	}
	buf := make([]byte, 1)
	_, err := input.File.Read(buf)
	if err == io.EOF {
		input.File.Close();
		input.isClose =true
		return 0
	}
	if err != nil {
		input.File.Close();
		input.isClose =true
		return 0
	}
	input.pos++;
	return buf[0];
}
func (input *InputFile ) GetPosition() int{
	return input.pos;
}
