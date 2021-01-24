package bitcoin

import (  
    "github.com/Block-Chen/blocksdk-go/base"
)

type Bitcoin struct{
	Base *base.Base
}

func New(api_token string) *Bitcoin {
	b := base.New(api_token)
	return &Bitcoin{
		Base : b,
	}
}

func (b *Bitcoin)GetBlockChain() map[string]interface{}{
	return b.Base.Request("GET","/btc/info",nil)
}


func (b *Bitcoin)GetBlock(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/btc/blocks/" + b.Base.ToString(request["block"]),request)
}


func (b *Bitcoin)GetMemPool(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/btc/mempool/" ,request)
}


func (b *Bitcoin)GetAddressInfo(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/btc/addresses/" +  b.Base.ToString(request["address"]),request)
    
}
func (b *Bitcoin)GetAddressBalance(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/btc/addresses/" +  b.Base.ToString(request["address"]) + "/balance",nil)
}
func (b *Bitcoin)GetWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/btc/wallets/"  +  b.Base.ToString(request["wallet_id"]),request)
}
func (b *Bitcoin)GetWallets(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/btc/wallets",request)
}
func (b *Bitcoin)CreateHdWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/btc/wallets/hd",request)
}
func (b *Bitcoin)LoadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/btc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/load",request)
}

func (b *Bitcoin)UnloadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/btc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/unload",request)
}

func (b *Bitcoin)GetWalletAddresses(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/btc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/address",request)
}


func (b *Bitcoin)CreateWalletAddress(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/btc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/address",request)
}

func (b *Bitcoin)GetWalletBalance(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/btc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/balance",nil)      
}

func (b *Bitcoin)GetWalletTransactions(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/btc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/transactions",request)
}

func (b *Bitcoin)SendToAddress(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/btc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/sendtoaddress",request)
}

func (b *Bitcoin)SendMany(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/btc/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/sendmany",request)
}

func (b *Bitcoin)SendTransaction(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/btc/transactions/send",request)
}

func (b *Bitcoin)GetTransaction(request map[string]interface{})  map[string]interface{} {
	return b.Base.Request("GET","/btc/transactions/" +  b.Base.ToString(request["hash"]) + "",nil)
}


