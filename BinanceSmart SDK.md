### 바이낸스 스마트 체인 V3 REST API 문서
[바이낸스 스마트 체인 개발자 문서 바로가기](https://documenter.getpostman.com/view/20292093/Uz5FKwxw#f6ddeb8a-8e34-4900-9706-9aea7c18fbe9)
함수 호출에 필요한 매개변수 또는 반환되는 데이터에 대해서는 REST API 개발자 문서를 참고해 주시길 바랍니다.

### 클라이언트 생성 (테스트넷)
```go
package main

import (
	"fmt"
	blocksdk "github.com/Block-Chen/blocksdk-go"
)

func main() {
    bscClient := blocksdk.NewBinanceSmart("YOU_TOKEN", "https://testnet-api.blocksdk.com/")
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
    bscClient := blocksdk.NewBinanceSmart("YOU_TOKEN", "https://mainnet-api.blocksdk.com/")
}
```

### 블록체인 정보
```
GET /v3/bsc/info
```
```go
result, err := bscClient.GetBlockChainInfo()
```

### 블록 정보
```
GET /v3/bsc/block/<block>
```
```go
result, err := bscClient.GetBlock(map[string]interface{}{
    "block" : "blockNumber 또는 blockHash",
});
```

### 주소 목록
```
GET /v3/bsc/address
```
```go
result, err := bscClient.GetAddresses(map[string]interface{}{
    "offset" : "0",
    "limit" : "10",
});
```

### 주소 정보
```
GET /v3/bsc/address/<address>/info
```
```go
result, err := bscClient.GetAddressInfo(map[string]interface{}{
    "address" : "주소",
    "offset" : "0",
    "limit" : "10",
});
```

### 주소 생성
```
POST /v3/bsc/address
```
```go
result, err := bscClient.CreateAddress(map[string]interface{}{
    "name" : "test",
});
```

### 주소 잔액
```
GET /v3/bsc/address/<address>/balance
```
```go
result, err := bscClient.GetAddressBalance(map[string]interface{}{
    "address" : "주소",
});
```

### 주소 거래 전송
```
POST /v3/bsc/address/<from_address>/send
```
```go
result, err := bscClient.Send(map[string]interface{}{
    "from" : "주소",
    "to" : "주소",
    "amount" : "보낼 양",
    "private_key" : "보내는 주소 키",
});
```

### 거래 전송
```
POST /v3/bsc/transaction/send
```
```go
result, err := bscClient.SendTransaction(map[string]interface{}{
    "hex" : "서명된 트랜잭션 hex",
});
```

### 거래 조회
```
GET /v3/bsc/transaction/<tx_hash>
```
```go
result, err := bscClient.GetTransaction(map[string]interface{}{
    "hash" : "트랜잭션 해쉬",
});
```

### ERC20 토큰 정보
```
GET /v3/bsc/token/<contract_address>/info
```
```go
result, err := bscClient.GetTokenInfo(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
});
```

### ERC20 토큰 잔액
```
GET /v3/bsc/token/<contract_address>/<from_address>/balance
```
```go
result, err := bscClient.GetTokenBalance(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "잔액을 조회할 주소",
});
```

### ERC20 토큰 전송
```
POST /v3/bsc/token/<contract_address>/<from_address>/transfer
```
```go
result, err := bscClient.SendToken(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "토큰을 전송할 주소",
    "to" : "주소",
    "amount" : "보낼 양",
    "private_key" : "보내는 주소 키",
});
```

### ERC20 특정 주소 거래 조회
```
GET /v3/bsc/token/<from_address>/transactions
```
```go
result, err := bscClient.GetTokenTxs(map[string]interface{}{
    "from_address" : "거래 내역을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC20 특정 컨트렉트 거래 조회
```
GET /v3/bsc/token/<contract_address>/<from_address>/transactions
```
```go
result, err := bscClient.GetTokenContractTxs(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "거래 내역을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC20 소유내역 조회
```
GET /v3/bsc/token/<from_address>/all-balance
```
```go
result, err := bscClient.GetTokenAllBalance(map[string]interface{}{
    "from_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC721(NFT) 특정 컨트렉트 NFT 목록
```
GET /v3/bsc/single-nft/<contract_address>/nfts
```
```go
result, err := bscClient.GetSingleNfts(map[string]interface{}{
    "contract_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 월렛 소유중인 NFT 조회
```
GET /v3/bsc/single-nft/<owner_address>/owner-nfts
```
```go
result, err := bscClient.GetSingleOwnerNfts(map[string]interface{}{
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 월렛 생성한 NFT 조회
```
GET /v3/bsc/single-nft/<creator_address>/creator-nfts
```
```go
result, err := bscClient.GetSingleCreatorNfts(map[string]interface{}{
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC721(NFT) 특정 월렛 NFT 거래 조회
```
GET /v3/bsc/single-nft/<from_address>/transactions
```
```go
result, err := bscClient.GetSingleTxs(map[string]interface{}{
    "from_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛 의 소유중인 NFT 조회
```
GET /v3/bsc/single-nft/<contract_address>/<owner_address>/owner-nfts
```
```go
result, err := bscClient.GetSingleNftOwnerNfts(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "owner_address" : "월렛 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛의 생성한 NFT 조회
```
GET /v3/bsc/single-nft/<contract_address>/<creator_address>/creator-nfts
```
```go
result, err := bscClient.GetSingleNftCreatorNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "creator_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛의 NFT 거래 조회
```
GET /v3/bsc/single-nft/<contract_address>/<from_address>/from-transactions
```
```go
result, err := bscClient.GetSingleNftTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 NFT 거래내역 조회
```
GET /v3/bsc/single-nft/<contract_address>/<token_id>/nft-transactions
```
```go
result, err := bscClient.GetSingleNftTokenTxs(map[string]interface{}{
    "contract_address" :  "NFT 컨트렉트 주소",
    "token_id" :  "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 NFT 정보 조회
```
GET /v3/bsc/single-nft/<contract_address>/<token_id>/info
```
```go
result, err := bscClient.GetSingleNftInfo(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" :  "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트 NFT 목록 조회
```
GET /v3/bsc/multi-nft/<contract_address>/nfts
```
```go
result, err := bscClient.GetMultiNfts(map[string]interface{}{
    "contract_address" :"NFT 컨트렉트 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 소유하고 있는 NFT 목록 조회
```
GET /v3/bsc/multi-nft/<owner_address>/owner-nfts
```
```go
result, err := bscClient.GetMultiOwnerNfts(map[string]interface{}{
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 생성한 NFT 목록 조회
```
GET /v3/bsc/multi-nft/<creator_address>/creator-nfts
```
```go
result, err := bscClient.GetMultiCreatorNfts(map[string]interface{}{
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 NFT 거래 내역 조회
```
GET /v3/bsc/multi-nft/<from_address>/transactions
```
```go
result, err := bscClient.GetMultiTxs(map[string]interface{}{
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT)  특정 컨트렉트,월렛이 소유한 NFT 목록 조회
```
GET /v3/bsc/multi-nft/<contract_address>/<owner_address>/owner-nfts
```
```go
result, err := bscClient.GetMultiNftOwnerNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트,월렛이 생성한 NFT 목록 조회
```
GET /v3/bsc/multi-nft/<contract_address>/<creator_address>/creator-nfts
```
```go
result, err := bscClient.GetMultiNftCreatorNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트,월렛의 NFT 거래 내역 조회
```
GET /v3/bsc/multi-nft/<contract_address>/<from_address>/from-transactions
```
```go
result, err := bscClient.GetMultiNftTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 NFT 정보&소유자 조회
```
GET /v3/bsc/multi-nft/<contract_address>/<token_id>/info
```
```go
result, err := bscClient.GetMultiNftInfo(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" : "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 NFT 거래내역 조회
```
GET /v3/bsc/multi-nft/<contract_address>/<token_id>/nft-transactions
```
```go
result, err := bscClient.GetMultiNftTokenTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" : "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### 스마트 계약 함수호출(읽기)
```
POST /v3/bsc/contract/<contract_address>/read
```
```go
result, err := bscClient.ReadContract(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "method" : "실행할 함수 명",
    "return_type" : "반환 데이터 타입",
    "parameter_type" : []interface{}{"인풋 파라미터 타입"},
    "parameter_data" : []interface{}{"인풋 파라미터 데이터"},
});
```


### 스마트 계약 함수호출(쓰기)
```
POST /v3/bsc/contract/<contract_address>/write
```
```go
result, err := bscClient.WriteContract(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "from" : "트랜잭션을 생성할 주소",
    "private_key" : "from 의 프라이빗키",
    "method" : "실행할 함수 명",
    "parameter_type" : []interface{}{"인풋 파라미터 타입"},
    "parameter_data" : []interface{}{"인풋 파라미터 데이터"},
});
```