package monero



import (
    "github.com/Block-Chen/blocksdk-go/base"
)

type Monero struct{
    Base *base.Base
}

func New(api_token string) *Monero {
	b := base.New(api_token)
	return &Monero{
		Base : b,
	}
}

func (m *Monero)GetBlockChain() map[string]interface{} {
	return m.Base.Request("GET","/xmr/info",nil)
}


func (m *Monero)GetBlock(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/xmr/blocks/" + m.Base.ToString(request["block"]),request)
}


func (m *Monero)GetMemPool(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/xmr/mempool",request)
}


func (m *Monero)GetAddresses(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/xmr/addresses",request)
}

func (m *Monero)LoadAddress(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("POST","/xmr/addresses/" + m.Base.ToString(request["address_id"]) + "/load",request)
}

func (m *Monero)UnloadAddress(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("POST","/xmr/addresses/" + m.Base.ToString(request["address_id"]) + "/unload",request)
}

func (m *Monero)CreateAddress(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("POST","/xmr/addresses",request)
}

func (m *Monero)GetAddressInfo(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/xmr/addresses/" + m.Base.ToString(request["address_id"]),request)
}

func (m *Monero)GetAddressBalance(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/xmr/addresses/" + m.Base.ToString(request["address_id"]) + "/balance",request)
}

func (m *Monero)SendToAddress(request  map[string]interface{}) map[string]interface{} {

	return m.Base.Request("POST","/xmr/addresses/" + m.Base.ToString(request["address_id"]) + "/sendtoaddress",request)
}

func (m *Monero)SendTransaction(request map[string]interface{})  map[string]interface{} {
	return m.Base.Request("POST","/eth/transactions/send",request)
}

func (m *Monero)GetTransaction(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/xmr/transactions/" + m.Base.ToString(request["hash"]),request)
}







