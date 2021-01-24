package token

import (
    "github.com/Block-Chen/blocksdk-go/base"
)

type Token struct{
    Base *base.Base
}

func New(api_token string) *Token {
	b := base.New(api_token)
	return &Token{
		Base : b,
	}
}


func (m *Token)GetUsages(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/token/usage",request)
}