package market

import (
    "github.com/Block-Chen/blocksdk-go/base"
)

type Market struct{
    Base *base.Base
}

func New(api_token string) *Market {
	b := base.New(api_token)
	return &Market{
		Base : b,
	}
}

func (m *Market)GetExchanges() map[string]interface{} {
	return m.Base.Request("GET","/market/exchanges",nil)
}
func (m *Market)GetTrades(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/market/trades",request)
}
func (m *Market)GetRates(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/market/rates/" + m.Base.ToString(request["from"]),request)
}
func (m *Market)GetExchangeTrades(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/market/trades/" + m.Base.ToString(request["exchage_id"]),request)
}
func (m *Market)GetExchangeRates(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/market/rates/" + m.Base.ToString(request["exchage_id"]) + "/" + m.Base.ToString(request["from"]),request)
}
