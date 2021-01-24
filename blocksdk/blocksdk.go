
package blocksdk

import (  
    "github.com/Block-Chen/blocksdk-go/bitcoin"
	"github.com/Block-Chen/blocksdk-go/bitcoincash"
    "github.com/Block-Chen/blocksdk-go/litecoin"
    "github.com/Block-Chen/blocksdk-go/dash"
    "github.com/Block-Chen/blocksdk-go/ethereum"
    "github.com/Block-Chen/blocksdk-go/monero"
    "github.com/Block-Chen/blocksdk-go/webhook"
    "github.com/Block-Chen/blocksdk-go/market"
    "github.com/Block-Chen/blocksdk-go/token"
    "github.com/Block-Chen/blocksdk-go/tool"
)

type BlockSDK struct {  
    api_token string
}

func New(api_token string) *BlockSDK{
	return &BlockSDK{
		api_token : api_token,
	}
}

func (b *BlockSDK)CreateBitcoin() *bitcoin.Bitcoin {
    return bitcoin.New(b.api_token)
}
func (b *BlockSDK)CreateBitcoinCash() *bitcoincash.Bitcoincash {
    return bitcoincash.New(b.api_token)
}
func (b *BlockSDK)CreateLitecoin() *litecoin.Litecoin {
    return litecoin.New(b.api_token)
}
func (b *BlockSDK)CreateDash() *dash.Dash {
    return dash.New(b.api_token)
}
func (b *BlockSDK)CreateEthereum() *ethereum.Ethereum {
    return ethereum.New(b.api_token)
}
func (b *BlockSDK)CreateMonero() *monero.Monero {
    return monero.New(b.api_token)
}
func (b *BlockSDK)CreateWebHook() *webhook.WebHook {
    return webhook.New(b.api_token)
}
func (b *BlockSDK)CreateMarket() *market.Market {
    return market.New(b.api_token)
}
func (b *BlockSDK)CreateToken() *token.Token {
    return token.New(b.api_token)
}
func (b *BlockSDK)CreateTool() *tool.Tool {
    return tool.New(b.api_token)
}


func CreateBitcoin(api_token string) *bitcoin.Bitcoin {
    return bitcoin.New(api_token)
}

func CreateBitcoinCash(api_token string) *bitcoincash.Bitcoincash {
    return bitcoincash.New(api_token)
}

func CreateLitecoin(api_token string) *litecoin.Litecoin {
    return litecoin.New(api_token)
}

func CreateDash(api_token string) *dash.Dash {
    return dash.New(api_token)
}

func CreateEthereum(api_token string) *ethereum.Ethereum {
    return ethereum.New(api_token)
}

func CreateMonero(api_token string) *monero.Monero {
    return monero.New(api_token)
}

func CreateWebHook(api_token string) *webhook.WebHook {
    return webhook.New(api_token)
}
func CreateMarket(api_token string) *market.Market {
    return market.New(api_token)
}
func CreateToken(api_token string) *token.Token {
    return token.New(api_token)
}
func CreateTool(api_token string) *tool.Tool {
    return tool.New(api_token)
}