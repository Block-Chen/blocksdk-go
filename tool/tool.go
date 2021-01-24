package tool



import (
    "github.com/Block-Chen/blocksdk-go/base"
)

type Tool struct{
    Base *base.Base
}

func New(api_token string) *Tool {
	b := base.New(api_token)
	return &Tool{
		Base : b,
	}
}


func (m *Tool)GetHashType(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/tools/hash-type/" + m.Base.ToString(request["hash"]),nil)
}