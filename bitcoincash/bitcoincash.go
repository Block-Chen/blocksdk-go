
package bitcoincash

import (  
    "github.com/Block-Chen/blocksdk-go/base"
)

type Bitcoincash struct{
	Base *base.Base
}

func NewBitcoincashClient(api_token string) *Bitcoincash {
	b := base.NewBase(api_token)
	return &Bitcoincash{
		Base : b,
	}
}

func (b *Bitcoincash)GetBlockChain() map[string]interface{}{

	return b.Base.Request("GET","/bch/block",nil)
}


func (b *Bitcoincash)GetBlock(request map[string]interface{})  map[string]interface{} {

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
	
    return b.Base.Request("GET","/bch/block/" + request["block"].(string),request)
}


func (b *Bitcoincash)GetMemPool(request map[string]interface{})  map[string]interface{} {


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

    return b.Base.Request("GET","/bch/mempool/" ,request)
}


func (b *Bitcoincash)GetAddressInfo(request map[string]interface{})  map[string]interface{} {



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

    return b.Base.Request("GET","/bch/address/" + request["address"].(string),request)
    
}
func (b *Bitcoincash)GetAddressBalance(request map[string]interface{})  map[string]interface{} {

    return b.Base.Request("GET","/bch/address/" + request["address"].(string) + "/balance",nil)
    
}
func (b *Bitcoincash)ListWallet(request map[string]interface{})  map[string]interface{} {
    _, ok2 := request["offset"]
    if !ok2 {
            request["offset"] = 0
    }

    _, ok3 := request["limit"]
    if !ok3 {
            request["limit"] = 10
    }  

    return b.Base.Request("GET","/bch/wallet",request)
    
}
func (b *Bitcoincash)CreateWallet(request map[string]interface{})  map[string]interface{} {

    _, ok := request["name"]
    if !ok {
		request["name"] = nil
    }

    return b.Base.Request("POST","/bch/wallet",request)
    
}
func (b *Bitcoincash)LoadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/bch/wallet/" + request["wallet_id"].(string) + "/load",request)
}

func (b *Bitcoincash)UnLoadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/bch/wallet/" + request["wallet_id"].(string) + "/unload",request)
}

func (b *Bitcoincash)ListWalletAddress(request map[string]interface{})  map[string]interface{} {
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
        
    return b.Base.Request("GET","/bch/wallet/" + request["wallet_id"].(string) + "/address",request)
}


func (b *Bitcoincash)CreateWalletAddress(request map[string]interface{})  map[string]interface{} {


    _, ok := request["seed_wif"]
    if !ok {
            request["seed_wif"] = nil
    }

    _, ok1 := request["password"]
    if !ok1 {
            request["password"] = nil
    }

        
    return b.Base.Request("POST","/bch/wallet/" + request["wallet_id"].(string) + "/address",request)
}

func (b *Bitcoincash)GetWalletBalance(request map[string]interface{})  map[string]interface{} {

    return b.Base.Request("GET","/bch/wallet/" + request["wallet_id"].(string) + "/balance",nil)      

}

func (b *Bitcoincash)GetWalletTransaction(request map[string]interface{})  map[string]interface{} {


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


    return b.Base.Request("GET","/bch/wallet/" + request["wallet_id"].(string) + "/transaction",request)
}

func (b *Bitcoincash)SendToAddress(request map[string]interface{})  map[string]interface{} {


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


        
    return b.Base.Request("POST","/bch/wallet/" + request["wallet_id"].(string) + "/sendtoaddress",request)
}

func (b *Bitcoincash)SendMany(request map[string]interface{})  map[string]interface{} {
    _, ok := request["seed_wif"]
    if !ok {

        request["seed_wif"] = nil
    }

    _, ok1 := request["password"]
    if !ok1 {
            request["password"] = nil
    }
            
        
    return b.Base.Request("POST","/bch/wallet/" + request["wallet_id"].(string) + "/sendmany",request)
}

func (b *Bitcoincash)SendTransaction(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/bch/transaction",request)
}

func (b *Bitcoincash)GetTransaction(request map[string]interface{})  map[string]interface{} {
	return b.Base.Request("GET","/bch/transaction/" + request["hash"].(string) + "",nil)
}