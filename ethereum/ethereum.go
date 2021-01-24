package ethereum

import ( 
    "github.com/Block-Chen/blocksdk-go/base"
)

type Ethereum struct{
    Base *base.Base
}

func New(api_token string) *Ethereum {
	b := base.New(api_token)
	return &Ethereum{
		Base : b,
	}
}

func (e *Ethereum) GetBlockChain() map[string]interface{}{
	return e.Base.Request("GET","/eth/info",nil)		
}


func (e *Ethereum) GetBlock(request map[string]interface{}) map[string]interface{} {   
	return e.Base.Request("GET","/eth/blocks/" + e.Base.ToString(request["block"]),request)
}


func (e *Ethereum) GetMemPool(request map[string]interface{}) map[string]interface{} {
	return e.Base.Request("GET","/eth/mempool",request)
}


func (e *Ethereum)GetAddresses(request map[string]interface{}) map[string]interface{}{
	return e.Base.Request("GET","/eth/addresses",request)   
}



func (e *Ethereum)LoadAddress(request map[string]interface{})  map[string]interface{} {
	return e.Base.Request("POST","/eth/addresses/" + e.Base.ToString(request["address"]) + "/load",request)

}


func (e *Ethereum)UnloadAddress(request  map[string]interface{})  map[string]interface{} {
	return e.Base.Request("POST","/eth/addresses/" + e.Base.ToString(request["address"]) + "/unload",nil)
}


func (e *Ethereum)CreateAddress(request  map[string]interface{})  map[string]interface{} {
	return e.Base.Request("POST","/eth/addresses",request)
}


func (e *Ethereum)GetAddressInfo(request  map[string]interface{})  map[string]interface{} {
	return e.Base.Request("GET","/eth/addresses/" + e.Base.ToString(request["address"]) ,request)
}

func (e *Ethereum)GetAddressBalance(request  map[string]interface{})  map[string]interface{} {

	return e.Base.Request("GET","/eth/addresses/" + e.Base.ToString(request["address"]) + "/balance",nil)

}
func (e *Ethereum)SendToAddress(request  map[string]interface{})  map[string]interface{} {
	return e.Base.Request("POST","/eth/addresses/" + e.Base.ToString(request["from"]) + "/sendtoaddress",request)

}

func (e *Ethereum)SendTransaction(request map[string]interface{})  map[string]interface{} {
	return e.Base.Request("POST","/eth/transactions/send",request)
}
func (e *Ethereum)GetTransaction(request map[string]interface{})  map[string]interface{} {
	return e.Base.Request("GET","/eth/transactions/" + e.Base.ToString(request["hash"]),nil)
}







