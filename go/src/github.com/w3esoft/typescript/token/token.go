package token

import (
	"errors"
)

type Token struct {
	Index int
	Value *string
	Name  string
	Position * Position
	Valid bool
	Id int
}


type Position struct {
	Len int
	Offset int
}

var Id =0

func New(index int, value *string , position * Position , valid bool) *Token {
	var token = Token{}
	token.Index = index
	token.Id = Id ;
	Id++
	token.Value = value
	token.Position = position
	token.Name = NAMES[index]
	token.Valid = valid
	return &token
}

func (token *Token) Is(tokenIndexs []int, value *string, valid bool) bool {
	if(token.Valid!=valid){
		return false
	}
	for _,tokenIndex := range tokenIndexs {
		b1 := token.Index == tokenIndex
		b2 := true;
		if value!=nil{
			b2 =token.Value == value
		}
		if b1 && b2 {
			return true
		}
	}
	return false
}
func (token *Token) IsNot(tokenIndexs []int, value *string, valid bool) bool {
	return !token.Is(tokenIndexs, value,valid)
}
func (token *Token) Expected(tokenIndexs []int, value *string, valid bool) (r bool, err error) {
	b := token.Is(tokenIndexs, value,valid)
	if !b {
		return b, errors.New("unexpected token " + token.Name)
	}
	return b, nil
}
func (token *Token) Unexpected(tokenIndexs []int, value *string, valid bool) (r bool, err error) {
	b := token.IsNot(tokenIndexs, value,valid)
	if !b {
		return b, errors.New("unexpected token " + token.Name)
	}
	return b, nil
}
