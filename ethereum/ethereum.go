package ethereum

import ( 
    "github.com/Block-Chen/blocksdk-go/base"
)

type Ethereum struct{
    Base *base.Base
}

func NewEthereumClient(api_token string) *Ethereum {
	b := base.NewBase(api_token)
	return &Ethereum{
		Base : b,
	}
}

func (e *Ethereum) GetBlockChain() map[string]interface{}{
	return e.Base.Request("GET","/eth/block",nil)		
}


func (e *Ethereum) GetBlock(request map[string]interface{}) map[string]interface{} {
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
	return e.Base.Request("GET","/eth/block/" + request["block"].(string),request)
}


func (e *Ethereum) GetMemPool(request map[string]interface{}) map[string]interface{} {

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
	
	return e.Base.Request("GET","/eth/mempool",request)
}


func (e *Ethereum)ListAddress(request map[string]interface{}) map[string]interface{}{
    _, ok := request["offset"]
    if !ok {
		request["offset"] = 0
    }
    _, ok2 := request["limit"]
    if !ok2 {
		request["limit"] = 10
    }

		
	return e.Base.Request("GET","/eth/address",request)   
}



func (e *Ethereum)LoadAddress(request map[string]interface{})  map[string]interface{} {
	return e.Base.Request("POST","/eth/address/" + request["address"].(string) + "/load",request)

}


func (e *Ethereum)UnLoadAddress(request  map[string]interface{})  map[string]interface{} {
	return e.Base.Request("POST","/eth/address/" + request["address"].(string) + "/unload",nil)
}


func (e *Ethereum)CreateAddress(request  map[string]interface{})  map[string]interface{} {
	_, okk := request["name"]
    if !okk {
		request["name"] = nil
    }	
	return e.Base.Request("POST","/eth/address",request)
}


func (e *Ethereum)getAddressInfo(request  map[string]interface{})  map[string]interface{} {
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
	
	return e.Base.Request("GET","/eth/address/" + request["address"].(string) ,request)

    
}

func (e *Ethereum)GetAddressBalance(request  map[string]interface{})  map[string]interface{} {

	return e.Base.Request("GET","/eth/address/" + request["address"].(string) + "/balance",nil)

}
func (e *Ethereum)SendToAddress(request  map[string]interface{})  map[string]interface{} {

	_, ok := request["kbfee"]
    if !ok {
        blockChain := e.GetBlockChain()
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

	return e.Base.Request("POST","/eth/address/" + request["from"].(string) + "/sendtoaddress",request)

}

func (e *Ethereum)SendTransaction(request map[string]interface{})  map[string]interface{} {
	return e.Base.Request("POST","/eth/transaction",request)

}
func (e *Ethereum)GetTransaction(request map[string]interface{})  map[string]interface{} {
	return e.Base.Request("GET","/eth/transaction/" + request["hash"].(string),nil)
}







