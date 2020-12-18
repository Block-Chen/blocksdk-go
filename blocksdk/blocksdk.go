
package Blocksdk

import (  
    "fmt"
    "blocksdk-go/base"
    "blocksdk-go/dash"
    "blocksdk-go/bitcoin"
    "blocksdk-go/ethereum"
    "blocksdk-go/bitcoincash"
    "blocksdk-go/monero"
    "blocksdk-go/litecoin"
    "blocksdk-go/price"
)

type BlockSDK struct {  
    api_token string
}

func (b *BlockSDK) Init(api_token string){
    b.api_token = api_token
}


func createBitcoin() *BlockSDK {
    return Bitcoin(api_token)
}


func createEthereum() *BlockSDK {
    return Ethereum(api_token)
}


func createLitecoin() *BlockSDK {
    return Litecoin(api_token)
}


func createMonero() *BlockSDK {
    return Monero(api_token)
}
func createPrice() *BlockSDK {
    return Price(api_token)
}
func createWebHook() *BlockSDK {
    return WebHook(api_token)
}
func createDash() *BlockSDK {
    return Dash(api_token)
}
func createBitcoinCash() *BlockSDK {
    return Bitcoincash(api_token)

}



