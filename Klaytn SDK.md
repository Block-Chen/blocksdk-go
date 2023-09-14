### 클레이튼 V3 REST API 문서
[클레이튼 개발자 문서 바로가기](https://documenter.getpostman.com/view/20292093/Uz5FKwxw#7b85cc54-fd97-4a30-9a6b-06808938c1bc)
함수 호출에 필요한 매개변수 또는 반환되는 데이터에 대해서는 REST API 개발자 문서를 참고해 주시길 바랍니다.

### 클라이언트 생성 (테스트넷)
```js
const BLOCKSDK = require('blocksdk-js');
const client = new BLOCKSDK("YOU_TOKEN");
```

### 클라이언트 생성 (메인넷)
엔드 포인트를 지정해주지 않는경우 테스트넷으로 기본 설정되어 호출 됩니다 메인넷은 아래 예시와 같이 클라이언트 생성시 두번째 매개변수를 메인넷으로 지정해 주시길 바랍니다.
```js
const BLOCKSDK = require('blocksdk-js');
const client = new BLOCKSDK("YOU_TOKEN","https://mainnet-api.blocksdk.com/klay");
```

### 블록체인 정보
```
GET /v3/klay/info
```
```js
const result = client.klaytn.GetBlockChainInfo();
```

### 블록 정보
```
GET /v3/klay/block/<block>
```
```js
result = client.klaytn.GetBlock({
    'block' : "blockNumber 또는 blockHash"
});
```

### 주소 목록
```
GET /v3/klay/address
```
```js
result = client.klaytn.GetAddresses({
    'offset' : 0,
    'limit' : 10
});
```

### 주소 정보
```
GET /v3/klay/address/<address>/info
```
```js
result = client.klaytn.GetAddressInfo({
    'address' : "주소",
    'offset' : 0,
    'limit' : 10
});
```

### 주소 생성
```
POST /v3/klay/address
```
```js
result = client.klaytn.CreateAddress({
    'name' : "test"
});
```

### 주소 잔액
```
GET /v3/klay/address/<address>/balance
```
```js
result = client.klaytn.GetAddressBalance({
    'address' : "주소"
});
```

### 주소 거래 전송
```
POST /v3/klay/address/<from_address>/send
```
```js
result = client.klaytn.Send({
    'from' : "주소",
    'to' : "주소",
    'amount' : "보낼 양",
    'private_key' : "보내는 주소 키"
});
```

### 거래 전송
```
POST /v3/klay/transaction/send
```
```js
result = client.klaytn.SendTransaction({
    'hex' : "서명된 트랜잭션 hex"
});
```

### 거래 조회
```
GET /v3/klay/transaction/<tx_hash>
```
```js
result = client.klaytn.GetTransaction({
    'hash' : "트랜잭션 해쉬"
});
```

### ERC20 토큰 정보
```
GET /v3/klay/token/<contract_address>/info
```
```js
result = client.klaytn.GetTokenInfo({
    'contract_address' : "ERC20 토큰 컨트렉트 주소"
});
```

### ERC20 토큰 잔액
```
GET /v3/klay/token/<contract_address>/<from_address>/balance
```
```js
result = client.klaytn.GetTokenBalance({
    'contract_address' : "ERC20 토큰 컨트렉트 주소",
    'from_address' : "잔액을 조회할 주소"
});
```

### ERC20 토큰 전송
```
POST /v3/klay/token/<contract_address>/<from_address>/transfer
```
```js
result = client.klaytn.SendToken({
    'contract_address' : "ERC20 토큰 컨트렉트 주소",
    'from_address' : "토큰을 전송할 주소",
    'to' : "주소",
    'amount' : "보낼 양",
    'private_key' : "보내는 주소 키"
});
```

### ERC20 특정 주소 거래 조회
```
GET /v3/klay/token/<from_address>/transactions
```
```js
result = client.klaytn.GetTokenTxs({
    'from_address' : "거래 내역을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```

### ERC20 특정 컨트렉트 거래 조회
```
GET /v3/klay/token/<contract_address>/<from_address>/transactions
```
```js
result = client.klaytn.GetTokenContractTxs({
    'contract_address' : "ERC20 토큰 컨트렉트 주소",
    'from_address' : "거래 내역을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```

### ERC20 소유내역 조회
```
GET /v3/klay/token/<from_address>/all-balance
```
```js
result = client.klaytn.GetTokenAllBalance({
    'from_address' : "토큰 목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```

### ERC721(NFT) 특정 컨트렉트 NFT 목록
```
GET /v3/klay/single-nft/<contract_address>/nfts
```
```js
result = client.klaytn.GetSingleNfts({
    'contract_address' : "목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC721(NFT) 특정 월렛 소유중인 NFT 조회
```
GET /v3/klay/single-nft/<owner_address>/owner-nfts
```
```js
result = client.klaytn.GetSingleOwnerNfts({
    'owner_address' : "목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC721(NFT) 특정 월렛 생성한 NFT 조회
```
GET /v3/klay/single-nft/<creator_address>/creator-nfts
```
```js
result = client.klaytn.GetSingleCreatorNfts({
    'creator_address' : "목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```

### ERC721(NFT) 특정 월렛 NFT 거래 조회
```
GET /v3/klay/single-nft/<from_address>/transactions
```
```js
result = client.klaytn.GetSingleTxs({
    'from_address' : "토큰 목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC721(NFT) 특정 컨트렉트,월렛 의 소유중인 NFT 조회
```
GET /v3/klay/single-nft/<contract_address>/<owner_address>/owner-nfts
```
```js
result = client.klaytn.GetSingleNftOwnerNfts({
    'contract_address' : "컨트렉트 주소",
    'owner_address' : "월렛 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC721(NFT) 특정 컨트렉트,월렛의 생성한 NFT 조회
```
GET /v3/klay/single-nft/<contract_address>/<creator_address>/creator-nfts
```
```js
result = client.klaytn.GetSingleNftCreatorNfts({
    'contract_address' : "NFT 컨트렉트 주소",
    'creator_address' : "토큰 목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC721(NFT) 특정 컨트렉트,월렛의 NFT 거래 조회
```
GET /v3/klay/single-nft/<contract_address>/<from_address>/from-transactions
```
```js
result = client.klaytn.GetSingleNftTxs({
    'contract_address' : "NFT 컨트렉트 주소",
    'from_address' : "목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC721(NFT) 특정 NFT 거래내역 조회
```
GET /v3/klay/single-nft/<contract_address>/<token_id>/nft-transactions
```
```js
result = client.klaytn.GetSingleNftTokenTxs({
    'contract_address' :  "NFT 컨트렉트 주소",
    'token_id' :  "NFT 토큰 ID",
    'offset' : 0,
    'limit' : 10
});
```


### ERC721(NFT) 특정 NFT 정보 조회
```
GET /v3/klay/single-nft/<contract_address>/<token_id>/info
```
```js
result = client.klaytn.GetSingleNftInfo({
    'contract_address' : "NFT 컨트렉트 주소",
    'token_id' :  "NFT 토큰 ID",
    'offset' : 0,
    'limit' : 10
});
```


### ERC1155(NFT) 특정 컨트렉트 NFT 목록 조회
```
GET /v3/klay/multi-nft/<contract_address>/nfts
```
```js
result = client.klaytn.GetMultiNfts({
    'contract_address' :"NFT 컨트렉트 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC1155(NFT) 특정 월렛 소유하고 있는 NFT 목록 조회
```
GET /v3/klay/multi-nft/<owner_address>/owner-nfts
```
```js
result = client.klaytn.GetMultiOwnerNfts({
    'owner_address' : "목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC1155(NFT) 특정 월렛 생성한 NFT 목록 조회
```
GET /v3/klay/multi-nft/<creator_address>/creator-nfts
```
```js
result = client.klaytn.GetMultiCreatorNfts({
    'creator_address' : "목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC1155(NFT) 특정 월렛 NFT 거래 내역 조회
```
GET /v3/klay/multi-nft/<from_address>/transactions
```
```js
result = client.klaytn.GetMultiTxs({
    'from_address' : "목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC1155(NFT)  특정 컨트렉트,월렛이 소유한 NFT 목록 조회
```
GET /v3/klay/multi-nft/<contract_address>/<owner_address>/owner-nfts
```
```js
result = client.klaytn.GetMultiNftOwnerNfts({
    'contract_address' : "NFT 컨트렉트 주소",
    'owner_address' : "목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC1155(NFT) 특정 컨트렉트,월렛이 생성한 NFT 목록 조회
```
GET /v3/klay/multi-nft/<contract_address>/<creator_address>/creator-nfts
```
```js
result = client.klaytn.GetMultiNftCreatorNfts({
    'contract_address' : "NFT 컨트렉트 주소",
    'creator_address' : "목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC1155(NFT) 특정 컨트렉트,월렛의 NFT 거래 내역 조회
```
GET /v3/klay/multi-nft/<contract_address>/<from_address>/from-transactions
```
```js
result = client.klaytn.GetMultiNftTxs({
    'contract_address' : "NFT 컨트렉트 주소",
    'from_address' : "목록을 조회할 주소",
    'offset' : 0,
    'limit' : 10
});
```


### ERC1155(NFT) 특정 NFT 정보&소유자 조회
```
GET /v3/klay/multi-nft/<contract_address>/<token_id>/info
```
```js
result = client.klaytn.GetMultiNftInfo({
    'contract_address' : "NFT 컨트렉트 주소",
    'token_id' : "NFT 토큰 ID",
    'offset' : 0,
    'limit' : 10
});
```


### ERC1155(NFT) 특정 NFT 거래내역 조회
```
GET /v3/klay/multi-nft/<contract_address>/<token_id>/nft-transactions
```
```js
result = client.klaytn.GetMultiNftTokenTxs({
    'contract_address' : "NFT 컨트렉트 주소",
    'token_id' : "NFT 토큰 ID",
    'offset' : 0,
    'limit' : 10
});
```


### 스마트 계약 함수호출(읽기)
```
POST /v3/klay/contract/<contract_address>/read
```
```js
result = client.klaytn.ReadContract({
    'contract_address' : "컨트렉트 주소",
    'method' : "실행할 함수 명",
    'return_type' : "반환 데이터 타입",
    'parameter_type' : ["인풋 파라미터 타입"],
    'parameter_data' : ["인풋 파라미터 데이터"]
});
```


### 스마트 계약 함수호출(쓰기)
```
POST /v3/klay/contract/<contract_address>/write
```
```js
result = client.klaytn.WriteContract({
    'contract_address' : "컨트렉트 주소",
    'from' : "트랜잭션을 생성할 주소",
    'private_key' : "from 의 프라이빗키",
    'method' : "실행할 함수 명",
    'parameter_type' : ["인풋 파라미터 타입"],
    'parameter_data' : ["인풋 파라미터 데이터"]
});
```