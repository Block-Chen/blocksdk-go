package bitcoin

import (  
    "github.com/Block-Chen/blocksdk-go/base"
)

type Bitcoin struct{
	Base *base.Base
}

func NewBitcoinClient(api_token string) *Bitcoin {
	b := base.NewBase(api_token)
	return &Bitcoin{
		Base : b,
	}
}

func (b *Bitcoin)GetBlockChain() map[string]interface{}{

	return b.Base.Request("GET","/btc/block",nil)
}


func (b *Bitcoin)GetBlock(request map[string]interface{})  map[string]interface{} {

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
	
    return b.Base.Request("GET","/btc/block/" + request["block"].(string),request)
}


func (b *Bitcoin)GetMemPool(request map[string]interface{})  map[string]interface{} {


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

    return b.Base.Request("GET","/btc/mempool/" ,request)
}


func (b *Bitcoin)GetAddressInfo(request map[string]interface{})  map[string]interface{} {



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

    return b.Base.Request("GET","/btc/address/" + request["address"].(string),request)
    
}
func (b *Bitcoin)GetAddressBalance(request map[string]interface{})  map[string]interface{} {

    return b.Base.Request("GET","/btc/address/" + request["address"].(string) + "/balance",nil)
    
}
func (b *Bitcoin)ListWallet(request map[string]interface{})  map[string]interface{} {
    _, ok2 := request["offset"]
    if !ok2 {
            request["offset"] = 0
    }

    _, ok3 := request["limit"]
    if !ok3 {
            request["limit"] = 10
    }  

    return b.Base.Request("GET","/btc/wallet",request)
    
}
func (b *Bitcoin)CreateWallet(request map[string]interface{})  map[string]interface{} {

    _, ok := request["name"]
    if !ok {
		request["name"] = nil
    }

    return b.Base.Request("POST","/btc/wallet",request)
    
}
func (b *Bitcoin)LoadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/btc/wallet/" + request["wallet_id"].(string) + "/load",request)
}

func (b *Bitcoin)UnLoadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/btc/wallet/" + request["wallet_id"].(string) + "/unload",request)
}

func (b *Bitcoin)ListWalletAddress(request map[string]interface{})  map[string]interface{} {
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
        
    return b.Base.Request("GET","/btc/wallet/" + request["wallet_id"].(string) + "/address",request)
}


func (b *Bitcoin)CreateWalletAddress(request map[string]interface{})  map[string]interface{} {


    _, ok := request["seed_wif"]
    if !ok {
            request["seed_wif"] = nil
    }

    _, ok1 := request["password"]
    if !ok1 {
            request["password"] = nil
    }

        
    return b.Base.Request("POST","/btc/wallet/" + request["wallet_id"].(string) + "/address",request)
}

func (b *Bitcoin)GetWalletBalance(request map[string]interface{})  map[string]interface{} {

    return b.Base.Request("GET","/btc/wallet/" + request["wallet_id"].(string) + "/balance",nil)      

}

func (b *Bitcoin)GetWalletTransaction(request map[string]interface{})  map[string]interface{} {


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


    return b.Base.Request("GET","/btc/wallet/" + request["wallet_id"].(string) + "/transaction",request)
}

func (b *Bitcoin)SendToAddress(request map[string]interface{})  map[string]interface{} {


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


        
    return b.Base.Request("POST","/btc/wallet/" + request["wallet_id"].(string) + "/sendtoaddress",request)
}

func (b *Bitcoin)SendMany(request map[string]interface{})  map[string]interface{} {
    _, ok := request["seed_wif"]
    if !ok {

        request["seed_wif"] = nil
    }

    _, ok1 := request["password"]
    if !ok1 {
            request["password"] = nil
    }
            
        
    return b.Base.Request("POST","/btc/wallet/" + request["wallet_id"].(string) + "/sendmany",request)
}

func (b *Bitcoin)SendTransaction(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/btc/transaction",request)
}

func (b *Bitcoin)GetTransaction(request map[string]interface{})  map[string]interface{} {
	return b.Base.Request("GET","/btc/transaction/" + request["hash"].(string) + "",nil)
}

func (b *Bitcoin)GetTransactionTracking(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/btc/transaction/" + request["hash"].(string) + "/tracking",nil)
    
}

