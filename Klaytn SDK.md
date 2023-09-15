### 클레이튼 V3 REST API 문서
[클레이튼 개발자 문서 바로가기](https://documenter.getpostman.com/view/20292093/Uz5FKwxw#7b85cc54-fd97-4a30-9a6b-06808938c1bc)
함수 호출에 필요한 매개변수 또는 반환되는 데이터에 대해서는 REST API 개발자 문서를 참고해 주시길 바랍니다.

### 클라이언트 생성 (테스트넷)
```go
package main

import (
	"fmt"
	blocksdk "github.com/Block-Chen/blocksdk-go"
)

func main() {
    klayClient := blocksdk.NewKlaytn("YOU_TOKEN", "https://testnet-api.blocksdk.com/")
}
```

### 클라이언트 생성 (메인넷)
```go
package main

import (
	"fmt"
	blocksdk "github.com/Block-Chen/blocksdk-go"
)

func main() {
    klayClient := blocksdk.NewKlaytn("YOU_TOKEN", "https://mainnet-api.blocksdk.com/")
}
```

### 블록체인 정보
```
GET /v3/klay/info
```
```go
result, err := klayClient.GetBlockChainInfo()
```

### 블록 정보
```
GET /v3/klay/block/<block>
```
```go
result, err := klayClient.GetBlock(map[string]interface{}{
    "block" : "blockNumber 또는 blockHash",
});
```

### 주소 목록
```
GET /v3/klay/address
```
```go
result, err := klayClient.GetAddresses(map[string]interface{}{
    "offset" : "0",
    "limit" : "10",
});
```

### 주소 정보
```
GET /v3/klay/address/<address>/info
```
```go
result, err := klayClient.GetAddressInfo(map[string]interface{}{
    "address" : "주소",
    "offset" : "0",
    "limit" : "10",
});
```

### 주소 생성
```
POST /v3/klay/address
```
```go
result, err := klayClient.CreateAddress(map[string]interface{}{
    "name" : "test",
});
```

### 주소 잔액
```
GET /v3/klay/address/<address>/balance
```
```go
result, err := klayClient.GetAddressBalance(map[string]interface{}{
    "address" : "주소",
});
```

### 주소 거래 전송
```
POST /v3/klay/address/<from_address>/send
```
```go
result, err := klayClient.Send(map[string]interface{}{
    "from" : "주소",
    "to" : "주소",
    "amount" : "보낼 양",
    "private_key" : "보내는 주소 키",
});
```

### 거래 전송
```
POST /v3/klay/transaction/send
```
```go
result, err := klayClient.SendTransaction(map[string]interface{}{
    "hex" : "서명된 트랜잭션 hex",
});
```

### 거래 조회
```
GET /v3/klay/transaction/<tx_hash>
```
```go
result, err := klayClient.GetTransaction(map[string]interface{}{
    "hash" : "트랜잭션 해쉬",
});
```

### ERC20 토큰 정보
```
GET /v3/klay/token/<contract_address>/info
```
```go
result, err := klayClient.GetTokenInfo(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
});
```

### ERC20 토큰 잔액
```
GET /v3/klay/token/<contract_address>/<from_address>/balance
```
```go
result, err := klayClient.GetTokenBalance(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "잔액을 조회할 주소",
});
```

### ERC20 토큰 전송
```
POST /v3/klay/token/<contract_address>/<from_address>/transfer
```
```go
result, err := klayClient.SendToken(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "토큰을 전송할 주소",
    "to" : "주소",
    "amount" : "보낼 양",
    "private_key" : "보내는 주소 키",
});
```

### ERC20 특정 주소 거래 조회
```
GET /v3/klay/token/<from_address>/transactions
```
```go
result, err := klayClient.GetTokenTxs(map[string]interface{}{
    "from_address" : "거래 내역을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC20 특정 컨트렉트 거래 조회
```
GET /v3/klay/token/<contract_address>/<from_address>/transactions
```
```go
result, err := klayClient.GetTokenContractTxs(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "거래 내역을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC20 소유내역 조회
```
GET /v3/klay/token/<from_address>/all-balance
```
```go
result, err := klayClient.GetTokenAllBalance(map[string]interface{}{
    "from_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC721(NFT) 특정 컨트렉트 NFT 목록
```
GET /v3/klay/single-nft/<contract_address>/nfts
```
```go
result, err := klayClient.GetSingleNfts(map[string]interface{}{
    "contract_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 월렛 소유중인 NFT 조회
```
GET /v3/klay/single-nft/<owner_address>/owner-nfts
```
```go
result, err := klayClient.GetSingleOwnerNfts(map[string]interface{}{
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 월렛 생성한 NFT 조회
```
GET /v3/klay/single-nft/<creator_address>/creator-nfts
```
```go
result, err := klayClient.GetSingleCreatorNfts(map[string]interface{}{
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC721(NFT) 특정 월렛 NFT 거래 조회
```
GET /v3/klay/single-nft/<from_address>/transactions
```
```go
result, err := klayClient.GetSingleTxs(map[string]interface{}{
    "from_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛 의 소유중인 NFT 조회
```
GET /v3/klay/single-nft/<contract_address>/<owner_address>/owner-nfts
```
```go
result, err := klayClient.GetSingleNftOwnerNfts(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "owner_address" : "월렛 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛의 생성한 NFT 조회
```
GET /v3/klay/single-nft/<contract_address>/<creator_address>/creator-nfts
```
```go
result, err := klayClient.GetSingleNftCreatorNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "creator_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛의 NFT 거래 조회
```
GET /v3/klay/single-nft/<contract_address>/<from_address>/from-transactions
```
```go
result, err := klayClient.GetSingleNftTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 NFT 거래내역 조회
```
GET /v3/klay/single-nft/<contract_address>/<token_id>/nft-transactions
```
```go
result, err := klayClient.GetSingleNftTokenTxs(map[string]interface{}{
    "contract_address" :  "NFT 컨트렉트 주소",
    "token_id" :  "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 NFT 정보 조회
```
GET /v3/klay/single-nft/<contract_address>/<token_id>/info
```
```go
result, err := klayClient.GetSingleNftInfo(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" :  "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트 NFT 목록 조회
```
GET /v3/klay/multi-nft/<contract_address>/nfts
```
```go
result, err := klayClient.GetMultiNfts(map[string]interface{}{
    "contract_address" :"NFT 컨트렉트 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 소유하고 있는 NFT 목록 조회
```
GET /v3/klay/multi-nft/<owner_address>/owner-nfts
```
```go
result, err := klayClient.GetMultiOwnerNfts(map[string]interface{}{
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 생성한 NFT 목록 조회
```
GET /v3/klay/multi-nft/<creator_address>/creator-nfts
```
```go
result, err := klayClient.GetMultiCreatorNfts(map[string]interface{}{
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 NFT 거래 내역 조회
```
GET /v3/klay/multi-nft/<from_address>/transactions
```
```go
result, err := klayClient.GetMultiTxs(map[string]interface{}{
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT)  특정 컨트렉트,월렛이 소유한 NFT 목록 조회
```
GET /v3/klay/multi-nft/<contract_address>/<owner_address>/owner-nfts
```
```go
result, err := klayClient.GetMultiNftOwnerNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트,월렛이 생성한 NFT 목록 조회
```
GET /v3/klay/multi-nft/<contract_address>/<creator_address>/creator-nfts
```
```go
result, err := klayClient.GetMultiNftCreatorNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트,월렛의 NFT 거래 내역 조회
```
GET /v3/klay/multi-nft/<contract_address>/<from_address>/from-transactions
```
```go
result, err := klayClient.GetMultiNftTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 NFT 정보&소유자 조회
```
GET /v3/klay/multi-nft/<contract_address>/<token_id>/info
```
```go
result, err := klayClient.GetMultiNftInfo(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" : "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 NFT 거래내역 조회
```
GET /v3/klay/multi-nft/<contract_address>/<token_id>/nft-transactions
```
```go
result, err := klayClient.GetMultiNftTokenTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" : "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### 스마트 계약 함수호출(읽기)
```
POST /v3/klay/contract/<contract_address>/read
```
```go
result, err := klayClient.ReadContract(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "method" : "실행할 함수 명",
    "return_type" : "반환 데이터 타입",
    "parameter_type" : []interface{}{"인풋 파라미터 타입"},
    "parameter_data" : []interface{}{"인풋 파라미터 데이터"},
});
```


### 스마트 계약 함수호출(쓰기)
```
POST /v3/klay/contract/<contract_address>/write
```
```go
result, err := klayClient.WriteContract(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "from" : "트랜잭션을 생성할 주소",
    "private_key" : "from 의 프라이빗키",
    "method" : "실행할 함수 명",
    "parameter_type" : []interface{}{"인풋 파라미터 타입"},
    "parameter_data" : []interface{}{"인풋 파라미터 데이터"},
});
```