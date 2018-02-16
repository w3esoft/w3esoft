package token

import (
	"errors"
)

type Token struct {
	Index int
	Value interface{}
	Name  string
}


type Position struct {
	Len int
	Offset int
}

func New(tokenIndex int, value interface{} , pos Position) *Token {
	var token Token = Token{}
	token.Index = tokenIndex
	token.Value = value
	token.Name = NAMES[tokenIndex]
	return &token
}

func (token *Token) Is(tokenIndexs []int, value interface{}) bool {
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
func (token *Token) IsNot(tokenIndexs []int, value interface{}) bool {
	return !token.Is(tokenIndexs, value)
}
func (token *Token) Expected(tokenIndexs []int, value interface{}) (r bool, err error) {
	b := token.Is(tokenIndexs, value)
	if !b {
		return b, errors.New("unexpected token " + token.Name)
	}
	return b, nil
}
func (token *Token) Unexpected(tokenIndexs []int, value interface{}) (r bool, err error) {
	b := token.IsNot(tokenIndexs, value)
	if !b {
		return b, errors.New("unexpected token " + token.Name)
	}
	return b, nil
}
