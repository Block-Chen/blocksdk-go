
package bitcoincash

import (  
    "github.com/Block-Chen/blocksdk-go/base"
)

type Bitcoincash struct{
	Base *base.Base
}

func New(api_token string) *Bitcoincash {
	b := base.New(api_token)
	return &Bitcoincash{
		Base : b,
	}
}
func (b *Bitcoincash)GetBlockChain() map[string]interface{}{
	return b.Base.Request("GET","/bch/info",nil)
}


func (b *Bitcoincash)GetBlock(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/bch/blocks/" + b.Base.ToString(request["block"]),request)
}


func (b *Bitcoincash)GetMemPool(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/bch/mempool/" ,request)
}


func (b *Bitcoincash)GetAddressInfo(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/bch/addresses/" +  b.Base.ToString(request["address"]),request)
    
}
func (b *Bitcoincash)GetAddressBalance(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/bch/addresses/" +  b.Base.ToString(request["address"]) + "/balance",nil)
}
func (b *Bitcoincash)GetWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/bch/wallets/"  +  b.Base.ToString(request["wallet_id"]),request)
}
func (b *Bitcoincash)GetWallets(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/bch/wallets",request)
}
func (b *Bitcoincash)CreateHdWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/bch/wallets/hd",request)
}
func (b *Bitcoincash)LoadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/bch/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/load",request)
}

func (b *Bitcoincash)UnloadWallet(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/bch/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/unload",request)
}

func (b *Bitcoincash)GetWalletAddresses(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/bch/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/address",request)
}


func (b *Bitcoincash)CreateWalletAddress(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/bch/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/address",request)
}

func (b *Bitcoincash)GetWalletBalance(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/bch/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/balance",nil)      
}

func (b *Bitcoincash)GetWalletTransactions(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("GET","/bch/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/transactions",request)
}

func (b *Bitcoincash)SendToAddress(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/bch/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/sendtoaddress",request)
}

func (b *Bitcoincash)SendMany(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/bch/wallets/" +  b.Base.ToString(request["wallet_id"]) + "/sendmany",request)
}

func (b *Bitcoincash)SendTransaction(request map[string]interface{})  map[string]interface{} {
    return b.Base.Request("POST","/bch/transactions/send",request)
}

func (b *Bitcoincash)GetTransaction(request map[string]interface{})  map[string]interface{} {
	return b.Base.Request("GET","/bch/transactions/" +  b.Base.ToString(request["hash"]) + "",nil)
}