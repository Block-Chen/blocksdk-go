### 폴리곤 V3 REST API 문서
[폴리곤 개발자 문서 바로가기](https://documenter.getpostman.com/view/20292093/Uz5FKwxw#90d3a851-df92-4c4e-b3f7-1e1566bc891e)
함수 호출에 필요한 매개변수 또는 반환되는 데이터에 대해서는 REST API 개발자 문서를 참고해 주시길 바랍니다.

### 클라이언트 생성 (테스트넷)
```go
package main

import (
	"fmt"
	blocksdk "github.com/Block-Chen/blocksdk-go"
)

func main() {
    maticClient := blocksdk.NewPolygon("YOU_TOKEN", "https://testnet-api.blocksdk.com/")
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
    maticClient := blocksdk.NewPolygon("YOU_TOKEN", "https://mainnet-api.blocksdk.com/")
}
```

### 블록체인 정보
```
GET /v3/matic/info
```
```go
result, err := maticClient.GetBlockChainInfo()
```

### 블록 정보
```
GET /v3/matic/block/<block>
```
```go
result, err := maticClient.GetBlock(map[string]interface{}{
    "block" : "blockNumber 또는 blockHash",
});
```

### 주소 목록
```
GET /v3/matic/address
```
```go
result, err := maticClient.GetAddresses(map[string]interface{}{
    "offset" : "0",
    "limit" : "10",
});
```

### 주소 정보
```
GET /v3/matic/address/<address>/info
```
```go
result, err := maticClient.GetAddressInfo(map[string]interface{}{
    "address" : "주소",
    "offset" : "0",
    "limit" : "10",
});
```

### 주소 생성
```
POST /v3/matic/address
```
```go
result, err := maticClient.CreateAddress(map[string]interface{}{
    "name" : "test",
});
```

### 주소 잔액
```
GET /v3/matic/address/<address>/balance
```
```go
result, err := maticClient.GetAddressBalance(map[string]interface{}{
    "address" : "주소",
});
```

### 주소 거래 전송
```
POST /v3/matic/address/<from_address>/send
```
```go
result, err := maticClient.Send(map[string]interface{}{
    "from" : "주소",
    "to" : "주소",
    "amount" : "보낼 양",
    "private_key" : "보내는 주소 키",
});
```

### 거래 전송
```
POST /v3/matic/transaction/send
```
```go
result, err := maticClient.SendTransaction(map[string]interface{}{
    "hex" : "서명된 트랜잭션 hex",
});
```

### 거래 조회
```
GET /v3/matic/transaction/<tx_hash>
```
```go
result, err := maticClient.GetTransaction(map[string]interface{}{
    "hash" : "트랜잭션 해쉬",
});
```

### ERC20 토큰 정보
```
GET /v3/matic/token/<contract_address>/info
```
```go
result, err := maticClient.GetTokenInfo(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
});
```

### ERC20 토큰 잔액
```
GET /v3/matic/token/<contract_address>/<from_address>/balance
```
```go
result, err := maticClient.GetTokenBalance(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "잔액을 조회할 주소",
});
```

### ERC20 토큰 전송
```
POST /v3/matic/token/<contract_address>/<from_address>/transfer
```
```go
result, err := maticClient.SendToken(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "토큰을 전송할 주소",
    "to" : "주소",
    "amount" : "보낼 양",
    "private_key" : "보내는 주소 키",
});
```

### ERC20 특정 주소 거래 조회
```
GET /v3/matic/token/<from_address>/transactions
```
```go
result, err := maticClient.GetTokenTxs(map[string]interface{}{
    "from_address" : "거래 내역을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC20 특정 컨트렉트 거래 조회
```
GET /v3/matic/token/<contract_address>/<from_address>/transactions
```
```go
result, err := maticClient.GetTokenContractTxs(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "거래 내역을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC20 소유내역 조회
```
GET /v3/matic/token/<from_address>/all-balance
```
```go
result, err := maticClient.GetTokenAllBalance(map[string]interface{}{
    "from_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC721(NFT) 특정 컨트렉트 NFT 목록
```
GET /v3/matic/single-nft/<contract_address>/nfts
```
```go
result, err := maticClient.GetSingleNfts(map[string]interface{}{
    "contract_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 월렛 소유중인 NFT 조회
```
GET /v3/matic/single-nft/<owner_address>/owner-nfts
```
```go
result, err := maticClient.GetSingleOwnerNfts(map[string]interface{}{
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 월렛 생성한 NFT 조회
```
GET /v3/matic/single-nft/<creator_address>/creator-nfts
```
```go
result, err := maticClient.GetSingleCreatorNfts(map[string]interface{}{
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC721(NFT) 특정 월렛 NFT 거래 조회
```
GET /v3/matic/single-nft/<from_address>/transactions
```
```go
result, err := maticClient.GetSingleTxs(map[string]interface{}{
    "from_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛 의 소유중인 NFT 조회
```
GET /v3/matic/single-nft/<contract_address>/<owner_address>/owner-nfts
```
```go
result, err := maticClient.GetSingleNftOwnerNfts(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "owner_address" : "월렛 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛의 생성한 NFT 조회
```
GET /v3/matic/single-nft/<contract_address>/<creator_address>/creator-nfts
```
```go
result, err := maticClient.GetSingleNftCreatorNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "creator_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛의 NFT 거래 조회
```
GET /v3/matic/single-nft/<contract_address>/<from_address>/from-transactions
```
```go
result, err := maticClient.GetSingleNftTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 NFT 거래내역 조회
```
GET /v3/matic/single-nft/<contract_address>/<token_id>/nft-transactions
```
```go
result, err := maticClient.GetSingleNftTokenTxs(map[string]interface{}{
    "contract_address" :  "NFT 컨트렉트 주소",
    "token_id" :  "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 NFT 정보 조회
```
GET /v3/matic/single-nft/<contract_address>/<token_id>/info
```
```go
result, err := maticClient.GetSingleNftInfo(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" :  "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트 NFT 목록 조회
```
GET /v3/matic/multi-nft/<contract_address>/nfts
```
```go
result, err := maticClient.GetMultiNfts(map[string]interface{}{
    "contract_address" :"NFT 컨트렉트 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 소유하고 있는 NFT 목록 조회
```
GET /v3/matic/multi-nft/<owner_address>/owner-nfts
```
```go
result, err := maticClient.GetMultiOwnerNfts(map[string]interface{}{
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 생성한 NFT 목록 조회
```
GET /v3/matic/multi-nft/<creator_address>/creator-nfts
```
```go
result, err := maticClient.GetMultiCreatorNfts(map[string]interface{}{
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 NFT 거래 내역 조회
```
GET /v3/matic/multi-nft/<from_address>/transactions
```
```go
result, err := maticClient.GetMultiTxs(map[string]interface{}{
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT)  특정 컨트렉트,월렛이 소유한 NFT 목록 조회
```
GET /v3/matic/multi-nft/<contract_address>/<owner_address>/owner-nfts
```
```go
result, err := maticClient.GetMultiNftOwnerNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트,월렛이 생성한 NFT 목록 조회
```
GET /v3/matic/multi-nft/<contract_address>/<creator_address>/creator-nfts
```
```go
result, err := maticClient.GetMultiNftCreatorNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트,월렛의 NFT 거래 내역 조회
```
GET /v3/matic/multi-nft/<contract_address>/<from_address>/from-transactions
```
```go
result, err := maticClient.GetMultiNftTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 NFT 정보&소유자 조회
```
GET /v3/matic/multi-nft/<contract_address>/<token_id>/info
```
```go
result, err := maticClient.GetMultiNftInfo(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" : "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 NFT 거래내역 조회
```
GET /v3/matic/multi-nft/<contract_address>/<token_id>/nft-transactions
```
```go
result, err := maticClient.GetMultiNftTokenTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" : "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### 스마트 계약 함수호출(읽기)
```
POST /v3/matic/contract/<contract_address>/read
```
```go
result, err := maticClient.ReadContract(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "method" : "실행할 함수 명",
    "return_type" : "반환 데이터 타입",
    "parameter_type" : []interface{}{"인풋 파라미터 타입"},
    "parameter_data" : []interface{}{"인풋 파라미터 데이터"},
});
```


### 스마트 계약 함수호출(쓰기)
```
POST /v3/matic/contract/<contract_address>/write
```
```go
result, err := maticClient.WriteContract(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "from" : "트랜잭션을 생성할 주소",
    "private_key" : "from 의 프라이빗키",
    "method" : "실행할 함수 명",
    "parameter_type" : []interface{}{"인풋 파라미터 타입"},
    "parameter_data" : []interface{}{"인풋 파라미터 데이터"},
});
```