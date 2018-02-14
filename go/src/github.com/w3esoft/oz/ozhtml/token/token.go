package token

import (
    "errors"
)

type Token struct {
    Index int
    Value *string
    Name  string
}

func New(index int, value* string) * Token {
    var token Token = Token{}
    token.Index = index;
    token.Value = value;
    //token.name = name;
    return  &token;
}
func ( token Token ) Is( indexs [] int, value *string) bool {
    for index := range indexs {
        b1 := token.Index ==index
        b2 :=true;
        if value!=nil{
            b2= token.Value== value
        }
        if b1 && b2 {
            return  true
        }
    }
    return  false
}
func (token Token) IsNot( indexs[] int, value *string) bool {
    return !token.Is( indexs, value)
}
func (token Token) Expected( indexs[] int, value *string) (r bool, err error) {
    b := token.Is( indexs, value)
    if b {
        return b, errors.New("unexpected token " + token.Name)
    }
    return b, nil
}
func (token Token) Unexpected(  indexs[] int, value *string) (r bool, err error) {
    b := token.IsNot( indexs, value)
    if b {
        return b, errors.New("unexpected token " + token.Name)
    }
    return b, nil
}
func (token Token) ToString() string {
    if token.Value == nil {
        return "Token." + token.Name + "();";
    } else {
        return "Token." + token.Name + "();";
    }
}