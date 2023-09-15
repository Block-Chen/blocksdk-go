### 이더리움 V3 REST API 문서
[이더리움 개발자 문서 바로가기](https://documenter.getpostman.com/view/20292093/Uz5FKwxw#3dc0d320-c1b1-449f-b950-daca8249d474)
함수 호출에 필요한 매개변수 또는 반환되는 데이터에 대해서는 REST API 개발자 문서를 참고해 주시길 바랍니다.

### 클라이언트 생성 (테스트넷)
```go
package main

import (
	"fmt"
	blocksdk "github.com/Block-Chen/blocksdk-go"
)

func main() {
    ethClient := blocksdk.NewEthereum("YOU_TOKEN", "https://testnet-api.blocksdk.com/")
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
    ethClient := blocksdk.NewEthereum("YOU_TOKEN", "https://mainnet-api.blocksdk.com/")
}
```

### 블록체인 정보
```
GET /v3/eth/info
```
```go
result, err := ethClient.GetBlockChainInfo()
```

### 블록 정보
```
GET /v3/eth/block/<block>
```
```go
result, err := ethClient.GetBlock(map[string]interface{}{
    "block" : "blockNumber 또는 blockHash",
});
```

### 주소 목록
```
GET /v3/eth/address
```
```go
result, err := ethClient.GetAddresses(map[string]interface{}{
    "offset" : "0",
    "limit" : "10",
});
```

### 주소 정보
```
GET /v3/eth/address/<address>/info
```
```go
result, err := ethClient.GetAddressInfo(map[string]interface{}{
    "address" : "주소",
    "offset" : "0",
    "limit" : "10",
});
```

### 주소 생성
```
POST /v3/eth/address
```
```go
result, err := ethClient.CreateAddress(map[string]interface{}{
    "name" : "test",
});
```

### 주소 잔액
```
GET /v3/eth/address/<address>/balance
```
```go
result, err := ethClient.GetAddressBalance(map[string]interface{}{
    "address" : "주소",
});
```

### 주소 거래 전송
```
POST /v3/eth/address/<from_address>/send
```
```go
result, err := ethClient.Send(map[string]interface{}{
    "from" : "주소",
    "to" : "주소",
    "amount" : "보낼 양",
    "private_key" : "보내는 주소 키",
});
```

### 거래 전송
```
POST /v3/eth/transaction/send
```
```go
result, err := ethClient.SendTransaction(map[string]interface{}{
    "hex" : "서명된 트랜잭션 hex",
});
```

### 거래 조회
```
GET /v3/eth/transaction/<tx_hash>
```
```go
result, err := ethClient.GetTransaction(map[string]interface{}{
    "hash" : "트랜잭션 해쉬",
});
```

### ERC20 토큰 정보
```
GET /v3/eth/token/<contract_address>/info
```
```go
result, err := ethClient.GetTokenInfo(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
});
```

### ERC20 토큰 잔액
```
GET /v3/eth/token/<contract_address>/<from_address>/balance
```
```go
result, err := ethClient.GetTokenBalance(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "잔액을 조회할 주소",
});
```

### ERC20 토큰 전송
```
POST /v3/eth/token/<contract_address>/<from_address>/transfer
```
```go
result, err := ethClient.SendToken(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "토큰을 전송할 주소",
    "to" : "주소",
    "amount" : "보낼 양",
    "private_key" : "보내는 주소 키",
});
```

### ERC20 특정 주소 거래 조회
```
GET /v3/eth/token/<from_address>/transactions
```
```go
result, err := ethClient.GetTokenTxs(map[string]interface{}{
    "from_address" : "거래 내역을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC20 특정 컨트렉트 거래 조회
```
GET /v3/eth/token/<contract_address>/<from_address>/transactions
```
```go
result, err := ethClient.GetTokenContractTxs(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "거래 내역을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC20 소유내역 조회
```
GET /v3/eth/token/<from_address>/all-balance
```
```go
result, err := ethClient.GetTokenAllBalance(map[string]interface{}{
    "from_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC721(NFT) 특정 컨트렉트 NFT 목록
```
GET /v3/eth/single-nft/<contract_address>/nfts
```
```go
result, err := ethClient.GetSingleNfts(map[string]interface{}{
    "contract_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 월렛 소유중인 NFT 조회
```
GET /v3/eth/single-nft/<owner_address>/owner-nfts
```
```go
result, err := ethClient.GetSingleOwnerNfts(map[string]interface{}{
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 월렛 생성한 NFT 조회
```
GET /v3/eth/single-nft/<creator_address>/creator-nfts
```
```go
result, err := ethClient.GetSingleCreatorNfts(map[string]interface{}{
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC721(NFT) 특정 월렛 NFT 거래 조회
```
GET /v3/eth/single-nft/<from_address>/transactions
```
```go
result, err := ethClient.GetSingleTxs(map[string]interface{}{
    "from_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛 의 소유중인 NFT 조회
```
GET /v3/eth/single-nft/<contract_address>/<owner_address>/owner-nfts
```
```go
result, err := ethClient.GetSingleNftOwnerNfts(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "owner_address" : "월렛 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛의 생성한 NFT 조회
```
GET /v3/eth/single-nft/<contract_address>/<creator_address>/creator-nfts
```
```go
result, err := ethClient.GetSingleNftCreatorNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "creator_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛의 NFT 거래 조회
```
GET /v3/eth/single-nft/<contract_address>/<from_address>/from-transactions
```
```go
result, err := ethClient.GetSingleNftTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 NFT 거래내역 조회
```
GET /v3/eth/single-nft/<contract_address>/<token_id>/nft-transactions
```
```go
result, err := ethClient.GetSingleNftTokenTxs(map[string]interface{}{
    "contract_address" :  "NFT 컨트렉트 주소",
    "token_id" :  "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 NFT 정보 조회
```
GET /v3/eth/single-nft/<contract_address>/<token_id>/info
```
```go
result, err := ethClient.GetSingleNftInfo(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" :  "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트 NFT 목록 조회
```
GET /v3/eth/multi-nft/<contract_address>/nfts
```
```go
result, err := ethClient.GetMultiNfts(map[string]interface{}{
    "contract_address" :"NFT 컨트렉트 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 소유하고 있는 NFT 목록 조회
```
GET /v3/eth/multi-nft/<owner_address>/owner-nfts
```
```go
result, err := ethClient.GetMultiOwnerNfts(map[string]interface{}{
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 생성한 NFT 목록 조회
```
GET /v3/eth/multi-nft/<creator_address>/creator-nfts
```
```go
result, err := ethClient.GetMultiCreatorNfts(map[string]interface{}{
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 NFT 거래 내역 조회
```
GET /v3/eth/multi-nft/<from_address>/transactions
```
```go
result, err := ethClient.GetMultiTxs(map[string]interface{}{
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT)  특정 컨트렉트,월렛이 소유한 NFT 목록 조회
```
GET /v3/eth/multi-nft/<contract_address>/<owner_address>/owner-nfts
```
```go
result, err := ethClient.GetMultiNftOwnerNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트,월렛이 생성한 NFT 목록 조회
```
GET /v3/eth/multi-nft/<contract_address>/<creator_address>/creator-nfts
```
```go
result, err := ethClient.GetMultiNftCreatorNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트,월렛의 NFT 거래 내역 조회
```
GET /v3/eth/multi-nft/<contract_address>/<from_address>/from-transactions
```
```go
result, err := ethClient.GetMultiNftTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 NFT 정보&소유자 조회
```
GET /v3/eth/multi-nft/<contract_address>/<token_id>/info
```
```go
result, err := ethClient.GetMultiNftInfo(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" : "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 NFT 거래내역 조회
```
GET /v3/eth/multi-nft/<contract_address>/<token_id>/nft-transactions
```
```go
result, err := ethClient.GetMultiNftTokenTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" : "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### 스마트 계약 함수호출(읽기)
```
POST /v3/eth/contract/<contract_address>/read
```
```go
result, err := ethClient.ReadContract(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "method" : "실행할 함수 명",
    "return_type" : "반환 데이터 타입",
    "parameter_type" : []interface{}{"인풋 파라미터 타입"},
    "parameter_data" : []interface{}{"인풋 파라미터 데이터"},
});
```


### 스마트 계약 함수호출(쓰기)
```
POST /v3/eth/contract/<contract_address>/write
```
```go
result, err := ethClient.WriteContract(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "from" : "트랜잭션을 생성할 주소",
    "private_key" : "from 의 프라이빗키",
    "method" : "실행할 함수 명",
    "parameter_type" : []interface{}{"인풋 파라미터 타입"},
    "parameter_data" : []interface{}{"인풋 파라미터 데이터"},
});
```