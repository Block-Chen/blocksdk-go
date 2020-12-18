package ethereum

import (  
    "fmt"
    "time"
    "strconv"
    "net/http"
    "encoding/json"
    "log"
    "strings"
    "blocksdk-go/base"
)

type Ethereum struct{
    Base base.Base
}

func getBlockChain(request = map[string]interface{}) *Ethereum {
	return Base.request("GET","/eth/block")		
}


func getBlock(request = map[string]interface{}) *Ethereum {
    value, ok := request['rawtx']
    if !ok {
            request['rawtx'] = false
    }
    value2, ok2 := request['offset']
    if !ok2 {
            request['offset'] = 0
    }

    value3, ok3 := request['limit']

    if !ok2 {
            request['limit'] = 10
    }        
	return Base.request("GET","/eth/block/" + strconv.Itoa(request['block']),{"rawtx" : request['rawtx'],"offset": request['offset'],"limit" : request['limit']})


}


func getMemPool(request = map[string]interface{}) *Ethereum {

        value, ok := request['rawtx']
    if !ok {
            request['rawtx'] = false
    }
    value2, ok2 := request['offset']
    if !ok2 {
            request['offset'] = 0
    }

    value3, ok3 := request['limit']

    if !ok2 {
            request['limit'] = 10
    }  
	return self.request("GET","/eth/mempool",{"rawtx" : request['rawtx'],"offset" : request['offset'],"limit" : request['limit']})

}


func listAddress(request = map[string]interface{}) *Ethereum {
    value, ok := request['offset']
    if !ok {
            request['offset'] = 0
    }
    value2, ok2 := request['limit']
    if !ok2 {
            request['limit'] = 10
    }

		
	return Base.request("GET","/eth/address",{
			"offset" : request['offset'],
			"limit" : request['limit']
		})   
}



func loadAddress(request = map[string]interface{}) *Ethereum {
	return Base.request("POST","/eth/address/" + request['address'] + "/load",{"seed_wif" : request['seed_wif'],"password" : request['password']})

}


func unLoadAddress(request = map[string]interface{}) *Ethereum {
	return Base.request("POST","/eth/address/" + request['address'] + "/unload")
}


func createAddress(request = map[string]interface{}) *Ethereum {
   valuee, okk := request['name']
    if !okk {
            request['name'] = nil
    }	
	return Base.request("POST","/eth/address",{"name" : request['name']})
}


func getAddressInfo(request = map[string]interface{}) *Ethereum {
   valuee, okk := request['reverse']
    if !okk {
            request['reverse'] = true
    }
    value, ok := request['rawtx']
    if !ok {
            request['rawtx'] = false
    }

    value2, ok2 := request['offset']
    if !ok2 {
            request['offset'] = 0
    }

    value3, ok3 := request['limit']

    if !ok2 {
            request['limit'] = 10
    }  
	
	return Base.request("GET","/eth/address/" + request['address'],{"reverse" : request['reverse'],"rawtx"  : request['rawtx'],"offset" : request['offset'],"limit"  : request['limit']})

    
}

func getAddressBalance(request = map[string]interface{}) *Ethereum {

	return Base.request("GET","/eth/address/" + request['address'] + "/balance")

}
func sendToAddress(request = map[string]interface{}) *Ethereum {

        value, ok := request['kbfee']
    if !ok {

        blockChain = *getBlockChain()
        request['kbfee'] = blockChain['medium_fee_per_kb']
    }

    value1, ok1 := request['seed_wif']
    if !ok1 {
            request['seed_wif'] = nil
    }
    value2, ok2 := request['password']
    if !ok2 {
            request['password'] = nil
    }

	return Base.request("POST","/eth/address/" + request['from'] + "/sendtoaddress",{"address" : request['to'],"amount" : request['amount'],"private_key" : request['private_key'],"password" : request['password'],"gwei" : request['gwei']})

}

func sendTransaction(request = map[string]interface{}) *Ethereum {
	return Base.request("POST","/eth/transaction",{"sign_hex" : request['sign_hex']})

}
func getTransaction(request = map[string]interface{}) *Ethereum {
	return Base.request("GET","/eth/transaction/" + request['hash'])

    
}







