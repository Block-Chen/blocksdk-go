package monero

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

type Monero struct{
    Base base.Base
}

func getBlockChain(request = map[string]interface{}) *Monero {
	return Base.request("GET","/xmr/block")

}


func getBlock(request = map[string]interface{}) *Monero {
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

	return Base.request("GET","/xmr/block/" + strconv.Itoa(request['block']) + "",{
			"rawtx" : request['rawtx'],
			"offset" : request['offset'],
			"limit" : request['limit']
		})

    

}


func getMemPool(request = map[string]interface{}) *Monero {

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


	return Base.request("GET","/xmr/mempool",{
			"rawtx" : request['rawtx'],
			"offset" : request['offset'],
			"limit" : request['limit']
		})

}


func listAddress(request = map[string]interface{}) *Monero {
    value, ok := request['offset']
    if !ok {
            request['offset'] = 0
    }
    value2, ok2 := request['limit']
    if !ok2 {
            request['limit'] = 10
    }
	return Base.request("GET","/xmr/address",{
			"offset" : request['offset'],
			"limit" : request['limit']
		})
    
}



func loadAddress(request = map[string]interface{}) *Monero {
	return Base.request("POST","/xmr/address/" + strconv.Itoa(request['address_id']) + "/load",{
			"private_spend_key" : request['private_spend_key'],
			"password" : request['password']
		})
    
}
func unLoadAddress(request = map[string]interface{}) *Monero {
	return Base.request("POST","/xmr/address/" + strconv.Itoa(request['address_id']) + "/unload")

    
}
func createAddress(request = map[string]interface{}) *Monero {

    value, ok := request['name']
    if !ok {
            request['name'] = nil
    }
	
	return Base.request("POST","/xmr/address",{
			"name" : request['name']
		})
    
}
func getAddressInfo(request = map[string]interface{}) *Monero {

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


	return Base.request("GET","/xmr/address/" + strconv.Itoa(request['address_id']) + "",{
			"offset" : request['offset'],
			"limit" : request['limit'],
			"private_spend_key" : request['private_spend_key'],
		})
    
}

func getAddressBalance(request = map[string]interface{}) *Monero {
	return Base.request("GET","/xmr/address/" + strconv.Itoa(request['address_id']) + "/balance",{
			"private_spend_key" : request['private_spend_key'],
		})
    
}








func sendToAddress(request = map[string]interface{}) *Monero {

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

	return Base.request("POST","/xmr/address/" + strconv.Itoa(request['address_id']) + "/sendtoaddress",{
			"address" : request['address'],
			"amount" : request['amount'],
			"private_spend_key" : request['private_spend_key'],
			"password" : request['password'],
			"kbfee" : request['kbfee']
		})


}


func getTransaction(request = map[string]interface{}) *Monero {
	return Base.request("GET","/xmr/transaction/" + strconv.Itoa(request['hash']) + "")

    
}







