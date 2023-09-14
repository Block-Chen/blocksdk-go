# GO REST API SDK for BlockSDK
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Block-Chen/blocksdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Block-Chen/blocksdk-go)](https://goreportcard.com/report/github.com/Block-Chen/blocksdk-go)

BlockSDK GO에 오신 것을 환영합니다. 이 저장소에는 BlockSDK의 GO SDK와 REST API용 샘플이 포함되어 있습니다.

## 지원중인 블록체인 네트워크
비트코인 , 라이트코인 , 비트코인 캐시 , 웹후크 는 V2버전 에서 지원되고 있습니다.
```
1.이더리움
2.클레이튼  
3.바이낸스 스마트 체인
4.폴리곤
5.아발란체
6.이더리움 클래식
```
## 개발자 문서
* [BlockSDK REST API V3 문서](https://documenter.getpostman.com/view/20292093/Uz5FKwxw)
* [BlockSDK REST API V2 문서](https://docs-v2.blocksdk.com/ko/#fa255f0ccc)
* [BLOCKSDK GO SDK V3 문서](https://github.com/Block-Chen/blocksdk-go/wiki)

## 시작하기
### go-get 을 사용하여 Install
```sh
go get -u github.com/Block-Chen/blocksdk-go
```

## 코드 샘플
### 이더리움 테스트넷 클라이언트 생성
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
### 이더리움 메인넷 클라이언트 생성
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
### 이더리움 블록체인 정보 가져오기
```go
result, err := ethClient.GetBlockChainInfo()

if err != nil {
    return
}

fmt.Println(result)
```

### 이더리움 테스트넷 특정 컨트렉트 NFT 목록 가져오기
```go
request := map[string]interface{}{
    "contract_address" : "0xf5de760f2e916647fd766b4ad9e85ff943ce3a2b",
    "includeMetadata" : "false",
    "offset" : "0",
    "limit" : "10",
}

result, err := ethClient.GetSingleNfts(request)

if err != nil {
    return
}

fmt.Println(result)
```
