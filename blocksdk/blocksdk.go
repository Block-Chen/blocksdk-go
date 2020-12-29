
package blocksdk

import (  
    "github.com/Block-Chen/blocksdk-go/bitcoin"
	"github.com/Block-Chen/blocksdk-go/bitcoincash"
    "github.com/Block-Chen/blocksdk-go/litecoin"
    "github.com/Block-Chen/blocksdk-go/dash"
    "github.com/Block-Chen/blocksdk-go/ethereum"
    "github.com/Block-Chen/blocksdk-go/monero"
    "github.com/Block-Chen/blocksdk-go/webhook"
)

type BlockSDK struct {  
    api_token string
}

func NewBlockSDK(api_token string) *BlockSDK{
	return &BlockSDK{
		api_token : api_token,
	}
}

func (b *BlockSDK)CreateBitcoin() *bitcoin.Bitcoin {
    return bitcoin.NewBitcoinClient(b.api_token)
}
func (b *BlockSDK)CreateBitcoinCash() *bitcoincash.Bitcoincash {
    return bitcoincash.NewBitcoincashClient(b.api_token)
}
func (b *BlockSDK)CreateLitecoin() *litecoin.Litecoin {
    return litecoin.NewLitecoinClient(b.api_token)
}
func (b *BlockSDK)CreateDash() *dash.Dash {
    return dash.NewDashClient(b.api_token)
}
func (b *BlockSDK)CreateEthereum() *ethereum.Ethereum {
    return ethereum.NewEthereumClient(b.api_token)
}
func (b *BlockSDK)CreateMonero() *monero.Monero {
    return monero.NewMoneroClient(b.api_token)
}
func (b *BlockSDK)CreateWebHook() *webhook.WebHook {
    return webhook.NewWebHook(b.api_token)
}


func CreateBitcoin(api_token string) *bitcoin.Bitcoin {
    return bitcoin.NewBitcoinClient(api_token)
}

func CreateBitcoinCash(api_token string) *bitcoincash.Bitcoincash {
    return bitcoincash.NewBitcoincashClient(api_token)
}

func CreateLitecoin(api_token string) *litecoin.Litecoin {
    return litecoin.NewLitecoinClient(api_token)
}

func CreateDash(api_token string) *dash.Dash {
    return dash.NewDashClient(api_token)
}

func CreateEthereum(api_token string) *ethereum.Ethereum {
    return ethereum.NewEthereumClient(api_token)
}

func CreateMonero(api_token string) *monero.Monero {
    return monero.NewMoneroClient(api_token)
}

func CreateWebHook(api_token string) *webhook.WebHook {
    return webhook.NewWebHook(api_token)
}