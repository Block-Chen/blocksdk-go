package monero



import (
    "github.com/Block-Chen/blocksdk-go/base"
)

type Monero struct{
    Base *base.Base
}

func NewMoneroClient(api_token string) *Monero {
	b := base.NewBase(api_token)
	return &Monero{
		Base : b,
	}
}

func (m *Monero)GetBlockChain() map[string]interface{} {
	return m.Base.Request("GET","/xmr/block",nil)
}


func (m *Monero)GetBlock(request  map[string]interface{}) map[string]interface{} {
    _, ok := request["rawtx"]
    if !ok {
		request["rawtx"] = false
    }
    _, ok2 := request["offset"]
    if !ok2 {
		request["offset"] = 0
    }

    _, ok3 := request["limit"]
    if !ok3 {
		request["limit"] = 10
    }

	return m.Base.Request("GET","/xmr/block/" + request["block"].(string) + "",request)
}


func (m *Monero)GetMemPool(request  map[string]interface{}) map[string]interface{} {

        _, ok := request["rawtx"]
    if !ok {
		request["rawtx"] = false
    }
    _, ok2 := request["offset"]
    if !ok2 {
		request["offset"] = 0
    }

    _, ok3 := request["limit"]
    if !ok3 {
		request["limit"] = 10
    }  


	return m.Base.Request("GET","/xmr/mempool",request)
}


func (m *Monero)listAddress(request  map[string]interface{}) map[string]interface{} {
    _, ok := request["offset"]
    if !ok {
            request["offset"] = 0
    }
    _, ok2 := request["limit"]
    if !ok2 {
            request["limit"] = 10
    }
	return m.Base.Request("GET","/xmr/address",request)
}



func (m *Monero)loadAddress(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("POST","/xmr/address/" + request["address_id"].(string) + "/load",request)
}
func (m *Monero)unLoadAddress(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("POST","/xmr/address/" + request["address_id"].(string) + "/unload",request)
}
func (m *Monero)createAddress(request  map[string]interface{}) map[string]interface{} {
    _, ok := request["name"]
    if !ok {
		request["name"] = nil
    }
	
	return m.Base.Request("POST","/xmr/address",request)
}
func (m *Monero)GetAddressInfo(request  map[string]interface{}) map[string]interface{} {
	_, okk := request["reverse"]
    if !okk {
		request["reverse"] = true
    }

    _, ok := request["rawtx"]
    if !ok {
		request["rawtx"] = false
    }

    _, ok2 := request["offset"]
    if !ok2 {
		request["offset"] = 0
    }

    _, ok3 := request["limit"]
    if !ok3 {
		request["limit"] = 10
    }  

	return m.Base.Request("GET","/xmr/address/" + request["address_id"].(string),request)
}

func (m *Monero)GetAddressBalance(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/xmr/address/" + request["address_id"].(string) + "/balance",request)
}








func (m *Monero)sendToAddress(request  map[string]interface{}) map[string]interface{} {

    _, ok := request["kbfee"]
    if !ok {

        blockChain := m.GetBlockChain()
        request["kbfee"] = blockChain["medium_fee_per_kb"].(float64)
    }

    _, ok1 := request["seed_wif"]
    if !ok1 {
            request["seed_wif"] = nil
    }
    _, ok2 := request["password"]
    if !ok2 {
            request["password"] = nil
    }

	return m.Base.Request("POST","/xmr/address/" + request["address_id"].(string) + "/sendtoaddress",request)
}


func (m *Monero)GetTransaction(request  map[string]interface{}) map[string]interface{} {
	return m.Base.Request("GET","/xmr/transaction/" + request["hash"].(string),request)
}







