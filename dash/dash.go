
package dash

import (  
    "github.com/Block-Chen/blocksdk-go/base"
)

type Dash struct{
	Base *base.Base
}

func NewDashClient(api_token string) *Dash {
	b := base.NewBase(api_token)
	return &Dash{
		Base : b,
	}
}

func (b *Dash)GetBlockChain() map[string]interface{}{

	return b.Base.Request("GET","/dash/block",nil)
}


func (b *Dash)GetBlock(request map[string]interface{})  map[string]interface{} {

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
	
    return b.Base.Request("GET","/dash/block/" + request["block"].(string),request)
}


func (b *Dash)GetMemPool(request map[string]interface{})  map[string]interface{} {


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

    return b.Base.Request("GET","/dash/mempool/" ,request)
}


func (b *Dash)GetAddressInfo(request map[string]interface{})  map[string]interface{} {



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

    return b.Base.Request("GET","/dash/address/" + request["address"].(string),request)
    
}
func (b *Dash)GetAddressBalance(request map[string]interface{})  map[string]interface{} {

    return b.Base.Request("GET","/dash/address/" + request["address"].(string) + "/balance",nil)
    
}
func (b *Dash)ListWallet(request map[string]interface{})  map[string]interface{} {
    _, ok2 := request["offset"]
    if !ok2 {
            request["offset"] = 0
    }

    _, ok3 := request["limit"]
    if !ok3 {
            request["limit"] = 10
    }  

    return b.Base.Request("GET","/dash/wallet",request)
    
}
func (b *Dash)CreateWallet(request map[string]interface{})  map[string]interface{} {

    _, ok := request["name"]
    if !ok {
		request["name"] = nil
    }

    return b.Base.Request("POST","/dash/wallet",request)
    
}
func (b *Dash)LoadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/dash/wallet/" + request["wallet_id"].(string) + "/load",request)
}

func (b *Dash)UnLoadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/dash/wallet/" + request["wallet_id"].(string) + "/unload",request)
}

func (b *Dash)ListWalletAddress(request map[string]interface{})  map[string]interface{} {
    _, ok := request["address"]
    if !ok {
            request["address"] = nil
    }

    _, ok1 := request["hdkeypath"]
    if !ok1 {
		request["hdkeypath"] = nil
    }
    _, ok2 := request["offset"]
    if !ok2 {
            request["offset"] = 0
    }
    _, ok3 := request["limit"]
    if !ok3 {
            request["limit"] = 10
    }
        
    return b.Base.Request("GET","/dash/wallet/" + request["wallet_id"].(string) + "/address",request)
}


func (b *Dash)CreateWalletAddress(request map[string]interface{})  map[string]interface{} {


    _, ok := request["seed_wif"]
    if !ok {
            request["seed_wif"] = nil
    }

    _, ok1 := request["password"]
    if !ok1 {
            request["password"] = nil
    }

        
    return b.Base.Request("POST","/dash/wallet/" + request["wallet_id"].(string) + "/address",request)
}

func (b *Dash)GetWalletBalance(request map[string]interface{})  map[string]interface{} {

    return b.Base.Request("GET","/dash/wallet/" + request["wallet_id"].(string) + "/balance",nil)      

}

func (b *Dash)GetWalletTransaction(request map[string]interface{})  map[string]interface{} {


    _, ok := request["order"]
    if !ok {
		request["order"] = "desc"
    }

    _, ok1 := request["offset"]
    if !ok1 {
            request["offset"] = 0
    }
    _, ok2 := request["limit"]
    if !ok2 {
            request["limit"] = 10
    }
    _, ok3 := request["category"]
    if !ok3 {
		request["category"] = "all"
    }


    return b.Base.Request("GET","/dash/wallet/" + request["wallet_id"].(string) + "/transaction",request)
}

func (b *Dash)SendToAddress(request map[string]interface{})  map[string]interface{} {


    _, ok := request["kbfee"]
    if !ok {
        blockChain := b.GetBlockChain() 
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


        
    return b.Base.Request("POST","/dash/wallet/" + request["wallet_id"].(string) + "/sendtoaddress",request)
}

func (b *Dash)SendMany(request map[string]interface{})  map[string]interface{} {
    _, ok := request["seed_wif"]
    if !ok {

        request["seed_wif"] = nil
    }

    _, ok1 := request["password"]
    if !ok1 {
            request["password"] = nil
    }
            
        
    return b.Base.Request("POST","/dash/wallet/" + request["wallet_id"].(string) + "/sendmany",request)
}

func (b *Dash)SendTransaction(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/dash/transaction",request)
}

func (b *Dash)GetTransaction(request map[string]interface{})  map[string]interface{} {
	return b.Base.Request("GET","/dash/transaction/" + request["hash"].(string) + "",nil)
}

func (b *Dash)GetTransactionTracking(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/dash/transaction/" + request["hash"].(string) + "/tracking",nil)
    
}