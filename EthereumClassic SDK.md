### 이더리움 클래식 V3 REST API 문서
[이더리움 클래식 개발자 문서 바로가기](https://documenter.getpostman.com/view/20292093/Uz5FKwxw#6287a2d3-2386-495a-98bc-27498f724ca2)
함수 호출에 필요한 매개변수 또는 반환되는 데이터에 대해서는 REST API 개발자 문서를 참고해 주시길 바랍니다.

### 클라이언트 생성 (테스트넷)
```go
package main

import (
	"fmt"
	blocksdk "github.com/Block-Chen/blocksdk-go"
)

func main() {
    etcClient := blocksdk.NewEthereumClassic("YOU_TOKEN", "https://testnet-api.blocksdk.com/")
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
    etcClient := blocksdk.NewEthereumClassic("YOU_TOKEN", "https://mainnet-api.blocksdk.com/")
}
```

### 블록체인 정보
```
GET /v3/etc/info
```
```go
result, err := etcClient.GetBlockChainInfo()
```

### 블록 정보
```
GET /v3/etc/block/<block>
```
```go
result, err := etcClient.GetBlock(map[string]interface{}{
    "block" : "blockNumber 또는 blockHash",
});
```

### 주소 목록
```
GET /v3/etc/address
```
```go
result, err := etcClient.GetAddresses(map[string]interface{}{
    "offset" : "0",
    "limit" : "10",
});
```

### 주소 정보
```
GET /v3/etc/address/<address>/info
```
```go
result, err := etcClient.GetAddressInfo(map[string]interface{}{
    "address" : "주소",
    "offset" : "0",
    "limit" : "10",
});
```

### 주소 생성
```
POST /v3/etc/address
```
```go
result, err := etcClient.CreateAddress(map[string]interface{}{
    "name" : "test",
});
```

### 주소 잔액
```
GET /v3/etc/address/<address>/balance
```
```go
result, err := etcClient.GetAddressBalance(map[string]interface{}{
    "address" : "주소",
});
```

### 주소 거래 전송
```
POST /v3/etc/address/<from_address>/send
```
```go
result, err := etcClient.Send(map[string]interface{}{
    "from" : "주소",
    "to" : "주소",
    "amount" : "보낼 양",
    "private_key" : "보내는 주소 키",
});
```

### 거래 전송
```
POST /v3/etc/transaction/send
```
```go
result, err := etcClient.SendTransaction(map[string]interface{}{
    "hex" : "서명된 트랜잭션 hex",
});
```

### 거래 조회
```
GET /v3/etc/transaction/<tx_hash>
```
```go
result, err := etcClient.GetTransaction(map[string]interface{}{
    "hash" : "트랜잭션 해쉬",
});
```

### ERC20 토큰 정보
```
GET /v3/etc/token/<contract_address>/info
```
```go
result, err := etcClient.GetTokenInfo(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
});
```

### ERC20 토큰 잔액
```
GET /v3/etc/token/<contract_address>/<from_address>/balance
```
```go
result, err := etcClient.GetTokenBalance(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "잔액을 조회할 주소",
});
```

### ERC20 토큰 전송
```
POST /v3/etc/token/<contract_address>/<from_address>/transfer
```
```go
result, err := etcClient.SendToken(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "토큰을 전송할 주소",
    "to" : "주소",
    "amount" : "보낼 양",
    "private_key" : "보내는 주소 키",
});
```

### ERC20 특정 주소 거래 조회
```
GET /v3/etc/token/<from_address>/transactions
```
```go
result, err := etcClient.GetTokenTxs(map[string]interface{}{
    "from_address" : "거래 내역을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC20 특정 컨트렉트 거래 조회
```
GET /v3/etc/token/<contract_address>/<from_address>/transactions
```
```go
result, err := etcClient.GetTokenContractTxs(map[string]interface{}{
    "contract_address" : "ERC20 토큰 컨트렉트 주소",
    "from_address" : "거래 내역을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC20 소유내역 조회
```
GET /v3/etc/token/<from_address>/all-balance
```
```go
result, err := etcClient.GetTokenAllBalance(map[string]interface{}{
    "from_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC721(NFT) 특정 컨트렉트 NFT 목록
```
GET /v3/etc/single-nft/<contract_address>/nfts
```
```go
result, err := etcClient.GetSingleNfts(map[string]interface{}{
    "contract_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 월렛 소유중인 NFT 조회
```
GET /v3/etc/single-nft/<owner_address>/owner-nfts
```
```go
result, err := etcClient.GetSingleOwnerNfts(map[string]interface{}{
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 월렛 생성한 NFT 조회
```
GET /v3/etc/single-nft/<creator_address>/creator-nfts
```
```go
result, err := etcClient.GetSingleCreatorNfts(map[string]interface{}{
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```

### ERC721(NFT) 특정 월렛 NFT 거래 조회
```
GET /v3/etc/single-nft/<from_address>/transactions
```
```go
result, err := etcClient.GetSingleTxs(map[string]interface{}{
    "from_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛 의 소유중인 NFT 조회
```
GET /v3/etc/single-nft/<contract_address>/<owner_address>/owner-nfts
```
```go
result, err := etcClient.GetSingleNftOwnerNfts(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "owner_address" : "월렛 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛의 생성한 NFT 조회
```
GET /v3/etc/single-nft/<contract_address>/<creator_address>/creator-nfts
```
```go
result, err := etcClient.GetSingleNftCreatorNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "creator_address" : "토큰 목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 컨트렉트,월렛의 NFT 거래 조회
```
GET /v3/etc/single-nft/<contract_address>/<from_address>/from-transactions
```
```go
result, err := etcClient.GetSingleNftTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 NFT 거래내역 조회
```
GET /v3/etc/single-nft/<contract_address>/<token_id>/nft-transactions
```
```go
result, err := etcClient.GetSingleNftTokenTxs(map[string]interface{}{
    "contract_address" :  "NFT 컨트렉트 주소",
    "token_id" :  "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC721(NFT) 특정 NFT 정보 조회
```
GET /v3/etc/single-nft/<contract_address>/<token_id>/info
```
```go
result, err := etcClient.GetSingleNftInfo(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" :  "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트 NFT 목록 조회
```
GET /v3/etc/multi-nft/<contract_address>/nfts
```
```go
result, err := etcClient.GetMultiNfts(map[string]interface{}{
    "contract_address" :"NFT 컨트렉트 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 소유하고 있는 NFT 목록 조회
```
GET /v3/etc/multi-nft/<owner_address>/owner-nfts
```
```go
result, err := etcClient.GetMultiOwnerNfts(map[string]interface{}{
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 생성한 NFT 목록 조회
```
GET /v3/etc/multi-nft/<creator_address>/creator-nfts
```
```go
result, err := etcClient.GetMultiCreatorNfts(map[string]interface{}{
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 월렛 NFT 거래 내역 조회
```
GET /v3/etc/multi-nft/<from_address>/transactions
```
```go
result, err := etcClient.GetMultiTxs(map[string]interface{}{
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT)  특정 컨트렉트,월렛이 소유한 NFT 목록 조회
```
GET /v3/etc/multi-nft/<contract_address>/<owner_address>/owner-nfts
```
```go
result, err := etcClient.GetMultiNftOwnerNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "owner_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트,월렛이 생성한 NFT 목록 조회
```
GET /v3/etc/multi-nft/<contract_address>/<creator_address>/creator-nfts
```
```go
result, err := etcClient.GetMultiNftCreatorNfts(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "creator_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 컨트렉트,월렛의 NFT 거래 내역 조회
```
GET /v3/etc/multi-nft/<contract_address>/<from_address>/from-transactions
```
```go
result, err := etcClient.GetMultiNftTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "from_address" : "목록을 조회할 주소",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 NFT 정보&소유자 조회
```
GET /v3/etc/multi-nft/<contract_address>/<token_id>/info
```
```go
result, err := etcClient.GetMultiNftInfo(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" : "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### ERC1155(NFT) 특정 NFT 거래내역 조회
```
GET /v3/etc/multi-nft/<contract_address>/<token_id>/nft-transactions
```
```go
result, err := etcClient.GetMultiNftTokenTxs(map[string]interface{}{
    "contract_address" : "NFT 컨트렉트 주소",
    "token_id" : "NFT 토큰 ID",
    "offset" : "0",
    "limit" : "10",
});
```


### 스마트 계약 함수호출(읽기)
```
POST /v3/etc/contract/<contract_address>/read
```
```go
result, err := etcClient.ReadContract(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "method" : "실행할 함수 명",
    "return_type" : "반환 데이터 타입",
    "parameter_type" : []interface{}{"인풋 파라미터 타입"},
    "parameter_data" : []interface{}{"인풋 파라미터 데이터"},
});
```


### 스마트 계약 함수호출(쓰기)
```
POST /v3/etc/contract/<contract_address>/write
```
```go
result, err := etcClient.WriteContract(map[string]interface{}{
    "contract_address" : "컨트렉트 주소",
    "from" : "트랜잭션을 생성할 주소",
    "private_key" : "from 의 프라이빗키",
    "method" : "실행할 함수 명",
    "parameter_type" : []interface{}{"인풋 파라미터 타입"},
    "parameter_data" : []interface{}{"인풋 파라미터 데이터"},
});
```