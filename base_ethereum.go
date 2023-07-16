package blocksdk_go

import (
	"fmt"
)

type BaseEthereum struct {
	*Base
}

func NewBaseEthereum(chainName string, apiToken string, endpoint string, version string) *BaseEthereum {
	return &BaseEthereum{
		Base: NewBase(chainName, apiToken, endpoint, version),
	}
}

func (e *BaseEthereum) GetBlockChainInfo() (map[string]interface{}, error) {
	return e.Request("GET", "/info", nil)
}

func (e *BaseEthereum) GetBlock(request map[string]interface{}) (map[string]interface{}, error) {
	block := request["block"].(string)
	return e.Request("GET", fmt.Sprintf("/block/%s", block), request)
}

func (e *BaseEthereum) GetAddresses(request map[string]interface{}) (map[string]interface{}, error) {
	return e.Request("GET", "/address", request)
}

func (e *BaseEthereum) CreateAddress(request map[string]interface{}) (map[string]interface{}, error) {
	return e.Request("POST", "/address", request)
}

func (e *BaseEthereum) GetAddressInfo(request map[string]interface{}) (map[string]interface{}, error) {
	address := request["address"].(string)
	return e.Request("GET", fmt.Sprintf("/address/%s", address), request)
}

func (e *BaseEthereum) GetAddressBalance(request map[string]interface{}) (map[string]interface{}, error) {
	address := request["address"].(string)
	return e.Request("GET", fmt.Sprintf("/address/%s/balance", address), nil)
}

func (e *BaseEthereum) Send(request map[string]interface{}) (map[string]interface{}, error) {
	fromAddress := request["from_address"].(string)
	return e.Request("POST", fmt.Sprintf("/address/%s/send", fromAddress), request)
}

func (e *BaseEthereum) SendTransaction(request map[string]interface{}) (map[string]interface{}, error) {
	return e.Request("POST", "/transaction/send", request)
}

func (e *BaseEthereum) GetTransaction(request map[string]interface{}) (map[string]interface{}, error) {
	hash := request["hash"].(string)
	return e.Request("GET", fmt.Sprintf("/transaction/%s", hash), nil)
}

func (e *BaseEthereum) GetTokenInfo(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	return e.Request("GET", fmt.Sprintf("/token/%s/info", contractAddress), nil)
}

func (e *BaseEthereum) SendToken(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	fromAddress := request["from_address"].(string)
	return e.Request("POST", fmt.Sprintf("/token/%s/%s/transfer", contractAddress, fromAddress), request)
}

func (e *BaseEthereum) GetTokenBalance(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	fromAddress := request["from_address"].(string)
	return e.Request("GET", fmt.Sprintf("/token/%s/%s/balance", contractAddress, fromAddress), nil)
}

func (e *BaseEthereum) GetTokenTxs(request map[string]interface{}) (map[string]interface{}, error) {
	fromAddress := request["from_address"].(string)
	return e.Request("GET", fmt.Sprintf("/token/%s/transactions", fromAddress), request)
}

func (e *BaseEthereum) GetTokenContractTxs(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	fromAddress := request["from_address"].(string)
	return e.Request("GET", fmt.Sprintf("/token/%s/%s/transactions", contractAddress, fromAddress), request)
}

func (e *BaseEthereum) GetTokenAllBalance(request map[string]interface{}) (map[string]interface{}, error) {
	fromAddress := request["from_address"].(string)
	return e.Request("GET", fmt.Sprintf("/token/%s/all-balance", fromAddress), request)
}

func (e *BaseEthereum) GetSingleNfts(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	return e.Request("GET", fmt.Sprintf("/single-nft/%s/nfts", contractAddress), request)
}

func (e *BaseEthereum) GetSingleOwnerNfts(request map[string]interface{}) (map[string]interface{}, error) {
	ownerAddress := request["owner_address"].(string)
	return e.Request("GET", fmt.Sprintf("/single-nft/%s/owner-nfts", ownerAddress), request)
}

func (e *BaseEthereum) GetSingleCreatorNfts(request map[string]interface{}) (map[string]interface{}, error) {
	creatorAddress := request["creator_address"].(string)
	return e.Request("GET", fmt.Sprintf("/single-nft/%s/creator-nfts", creatorAddress), request)
}

func (e *BaseEthereum) GetSingleTxs(request map[string]interface{}) (map[string]interface{}, error) {
	fromAddress := request["from_address"].(string)
	return e.Request("GET", fmt.Sprintf("/single-nft/%s/transactions", fromAddress), request)
}

func (e *BaseEthereum) GetSingleNftOwnerNfts(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	ownerAddress := request["owner_address"].(string)
	return e.Request("GET", fmt.Sprintf("/single-nft/%s/%s/owner-nfts", contractAddress, ownerAddress), request)
}

func (e *BaseEthereum) GetSingleNftCreatorNfts(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	creatorAddress := request["creator_address"].(string)
	return e.Request("GET", fmt.Sprintf("/single-nft/%s/%s/creator-nfts", contractAddress, creatorAddress), request)
}

func (e *BaseEthereum) GetSingleNftTxs(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	fromAddress := request["from_address"].(string)
	return e.Request("GET", fmt.Sprintf("/single-nft/%s/%s/from-transactions", contractAddress, fromAddress), request)
}

func (e *BaseEthereum) GetSingleNftInfo(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	tokenId := request["token_id"].(string)
	return e.Request("GET", fmt.Sprintf("/single-nft/%s/%s/info", contractAddress, tokenId), request)
}

func (e *BaseEthereum) GetSingleNftTokenTxs(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	tokenId := request["token_id"].(string)
	return e.Request("GET", fmt.Sprintf("/single-nft/%s/%s/nft-transactions", contractAddress, tokenId), request)
}

func (e *BaseEthereum) GetSingleNftAuctionNfts(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	return e.Request("GET", fmt.Sprintf("/single-nft/%s/auction-nfts", contractAddress), request)
}

func (e *BaseEthereum) GetSingleNftSellerNfts(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	sellerAddress := request["seller_address"].(string)
	return e.Request("GET", fmt.Sprintf("/single-nft/%s/%s/seller-nfts", contractAddress, sellerAddress), request)
}

func (e *BaseEthereum) GetSingleNftTokenBids(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	tokenId := request["token_id"].(string)
	return e.Request("GET", fmt.Sprintf("/single-nft/%s/%s/nft-bids", contractAddress, tokenId), request)
}

func (e *BaseEthereum) GetMultiNfts(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	return e.Request("GET", fmt.Sprintf("/multi-nft/%s/nfts", contractAddress), request)
}

func (e *BaseEthereum) GetMultiOwnerNfts(request map[string]interface{}) (map[string]interface{}, error) {
	ownerAddress := request["owner_address"].(string)
	return e.Request("GET", fmt.Sprintf("/multi-nft/%s/owner-nfts", ownerAddress), request)
}

func (e *BaseEthereum) GetMultiCreatorNfts(request map[string]interface{}) (map[string]interface{}, error) {
	creatorAddress := request["creator_address"].(string)
	return e.Request("GET", fmt.Sprintf("/multi-nft/%s/creator-nfts", creatorAddress), request)
}

func (e *BaseEthereum) GetMultiTxs(request map[string]interface{}) (map[string]interface{}, error) {
	fromAddress := request["from_address"].(string)
	return e.Request("GET", fmt.Sprintf("/multi-nft/%s/transactions", fromAddress), request)
}

func (e *BaseEthereum) GetMultiNftOwnerNfts(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	ownerAddress := request["owner_address"].(string)
	return e.Request("GET", fmt.Sprintf("/multi-nft/%s/%s/owner-nfts", contractAddress, ownerAddress), request)
}

func (e *BaseEthereum) GetMultiNftCreatorNfts(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	creatorAddress := request["creator_address"].(string)
	return e.Request("GET", fmt.Sprintf("/multi-nft/%s/%s/creator-nfts", contractAddress, creatorAddress), request)
}

func (e *BaseEthereum) GetMultiNftTxs(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	fromAddress := request["from_address"].(string)
	return e.Request("GET", fmt.Sprintf("/multi-nft/%s/%s/from-transactions", contractAddress, fromAddress), request)
}

func (e *BaseEthereum) GetMultiNftInfo(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	tokenId := request["token_id"].(string)
	return e.Request("GET", fmt.Sprintf("/multi-nft/%s/%s/info", contractAddress, tokenId), request)
}

func (e *BaseEthereum) GetMultiNftTokenTxs(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	tokenId := request["token_id"].(string)
	return e.Request("GET", fmt.Sprintf("/multi-nft/%s/%s/nft-transactions", contractAddress, tokenId), request)
}

func (e *BaseEthereum) GetMultiNftSellerNfts(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	sellerAddress := request["seller_address"].(string)
	return e.Request("GET", fmt.Sprintf("/multi-nft/%s/%s/seller-nfts", contractAddress, sellerAddress), request)
}

func (e *BaseEthereum) ReadContract(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	return e.Request("POST", fmt.Sprintf("/contract/%s/read", contractAddress), request)
}

func (e *BaseEthereum) WriteContract(request map[string]interface{}) (map[string]interface{}, error) {
	contractAddress := request["contract_address"].(string)
	return e.Request("POST", fmt.Sprintf("/contract/%s/write", contractAddress), request)
}
