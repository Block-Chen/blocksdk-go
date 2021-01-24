package litecoin

import (  
    "github.com/Block-Chen/blocksdk-go/base"
)

type Litecoin struct{
	Base *base.Base
}

func New(api_token string) *Litecoin {
	b := base.New(api_token)
	return &Litecoin{
		Base : b,
	}
}

func (b *Litecoin)GetBlockChain() map[string]interface{}{
	return b.Base.Request("GET","/ltc/info",nil)
}


func (b *Litecoin)GetBlock(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/ltc/blocks/" + b.Base.ToString(request["block"]),request)
}


func (b *Litecoin)GetMemPool(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/ltc/mempool/" ,request)
}


func (b *Litecoin)GetAddressInfo(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/ltc/addresses/" +  b.Base.ToString(request["address"]),request)
    
}
func (b *Litecoin)GetAddressBalance(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/ltc/addresses/" +  b.Base.ToString(request["address"]) + "/balance",nil)
}
func (b *Litecoin)GetWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/ltc/wallets/"  +  b.Base.ToString(request["wallet_id"]),request)
}
func (b *Litecoin)GetWallets(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/ltc/wallets",request)
}
func (b *Litecoin)CreateHdWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/ltc/wallets/hd",request)
}
func (b *Litecoin)LoadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/ltc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/load",request)
}

func (b *Litecoin)UnloadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/ltc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/unload",request)
}

func (b *Litecoin)GetWalletAddresses(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/ltc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/address",request)
}


func (b *Litecoin)CreateWalletAddress(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/ltc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/address",request)
}

func (b *Litecoin)GetWalletBalance(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/ltc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/balance",nil)      
}

func (b *Litecoin)GetWalletTransactions(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/ltc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/transactions",request)
}

func (b *Litecoin)SendToAddress(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/ltc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/sendtoaddress",request)
}

func (b *Litecoin)SendMany(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/ltc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/sendmany",request)
}

func (b *Litecoin)SendTransaction(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/ltc/transactions/send",request)
}

func (b *Litecoin)GetTransaction(request map[string]interface{})  map[string]interface{} {
	return b.Base.Request("GET","/ltc/transactions/" +  b.Base.ToString(request["hash"]) + "",nil)
}