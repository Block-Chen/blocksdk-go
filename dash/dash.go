
package dash

import (  
    "github.com/Block-Chen/blocksdk-go/base"
)

type Dash struct{
	Base *base.Base
}

func New(api_token string) *Dash {
	b := base.New(api_token)
	return &Dash{
		Base : b,
	}
}

func (b *Dash)GetBlockChain() map[string]interface{}{
	return b.Base.Request("GET","/dash/info",nil)
}


func (b *Dash)GetBlock(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/dash/blocks/" + b.Base.ToString(request["block"]),request)
}


func (b *Dash)GetMemPool(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/dash/mempool/" ,request)
}


func (b *Dash)GetAddressInfo(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/dash/addresses/" +  b.Base.ToString(request["address"]),request)
    
}
func (b *Dash)GetAddressBalance(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/dash/addresses/" +  b.Base.ToString(request["address"]) + "/balance",nil)
}
func (b *Dash)GetWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/dash/wallets/"  +  b.Base.ToString(request["wallet_id"]),request)
}
func (b *Dash)GetWallets(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/dash/wallets",request)
}
func (b *Dash)CreateHdWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/dash/wallets/hd",request)
}
func (b *Dash)LoadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/dash/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/load",request)
}

func (b *Dash)UnloadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/dash/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/unload",request)
}

func (b *Dash)GetWalletAddresses(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/dash/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/address",request)
}


func (b *Dash)CreateWalletAddress(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/dash/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/address",request)
}

func (b *Dash)GetWalletBalance(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/dash/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/balance",nil)      
}

func (b *Dash)GetWalletTransactions(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/dash/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/transactions",request)
}

func (b *Dash)SendToAddress(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/dash/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/sendtoaddress",request)
}

func (b *Dash)SendMany(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/dash/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/sendmany",request)
}

func (b *Dash)SendTransaction(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/dash/transactions/send",request)
}

func (b *Dash)GetTransaction(request map[string]interface{})  map[string]interface{} {
	return b.Base.Request("GET","/dash/transactions/" +  b.Base.ToString(request["hash"]) + "",nil)
}