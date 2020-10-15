
package dash

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

type Dash struct{
    Base base.Base
}

func getBlockChain(request = map[string]interface{}) *Dash {
	return Base.request("GET","/dash/block")
}


func getBlock(request = map[string]interface{}) *Dash {

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

	return Base.request("GET","/dash/block/" + strconv.Itoa(request['block']),{"rawtx" : request['rawtx'],"offset": request['offset'],"limit" : request['limit']})
    
}


func getMemPool(request = map[string]interface{}) *Dash {
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
	return Base.request("GET","/dash/mempool",{"rawtx" : request['rawtx'],"offset" : request['offset'],"limit" : request['limit']})

    
}


func getAddressInfo(request = map[string]interface{}) *Dash {
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
	return Base.request("GET","/dash/address/" + strconv.Itoa(request['address']),{"rawtx"  : request['rawtx'],"offset" : request['offset'],"limit"  : request['limit']})
    
}
func getAddressBalance(request = map[string]interface{}) *Dash {
	return Base.request("GET","/dash/address/" + strconv.Itoa(request['address']) + "/balance")


    
}
func listWallet(request = map[string]interface{}) *Dash {
    value2, ok2 := request['offset']
    if !ok2 {
            request['offset'] = 0
    }

    value3, ok3 := request['limit']

    if !ok2 {
            request['limit'] = 10
    }  

    return Base.request("GET","/dash/wallet",{
            "offset" : request['offset'],
            "limit" : request['limit']
        })
    
}
func createWallet(request = map[string]interface{}) *Dash {

    
    value, ok := request['name']
    if !ok {
            request['name'] = nil
    }

    return Base.request("POST","/dash/wallet",{
                 "name" : request['name']
    })
}
func loadWallet(request = map[string]interface{}) *Dash {
		return Base.request("POST","/dash/wallet/" + strconv.Itoa(request['wallet_id']) + "/load",{"seed_wif" : request['seed_wif'],"password" : request['password']})

    
}
func unLoadWallet(request = map[string]interface{}) *Dash {
	return Base.request("POST","/dash/wallet/" + strconv.Itoa(request['wallet_id']) + "/unload")    
}
func listWalletAddress(request = map[string]interface{}) *Dash {

    value, ok := request['address']
    if !ok {
            request['address'] = nil
    }

    value1, ok1 := request['hdkeypath']
    if !ok1 {
            request['hdkeypath'] = nil
    }
    value2, ok2 := request['offset']
    if !ok2 {
            request['offset'] = 0
    }
    value3, ok3 := request['limit']
    if !ok3 {
            request['limit'] = 10
    }
        
    return self.request("GET","/dash/wallet/" + strconv.Itoa(request['wallet_id']) + "/address",{
            "address" : request['address'],
            "hdkeypath" : request['hdkeypath'],
            "offset" : request['offset'],
            "limit" : request['limit']
        })
    
}

func createWalletAddress(request = map[string]interface{}) *Dash {

    
    value, ok := request['seed_wif']
    if !ok {
            request['seed_wif'] = nil
    }

    value1, ok1 := request['password']
    if !ok1 {
            request['password'] = nil
    }

        
    return Base.request("POST","/dash/wallet/" + strconv.Itoa(request['wallet_id']) + "/address",{
            "seed_wif" : request['seed_wif'],
            "password" : request['password']
        })
}

func getWalletBalance(request = map[string]interface{}) *Dash {

 		return Base.request("GET","/dash/wallet/" + strconv.Itoa(request['wallet_id']) + "/balance")
    
}
func getWalletTransaction(request = map[string]interface{}) *Dash {

	    value, ok := request['order']
    if !ok {
            request['order'] = 'desc'
    }

    value1, ok1 := request['offset']
    if !ok1 {
            request['offset'] = 0
    }
    value2, ok2 := request['limit']
    if !ok2 {
            request['limit'] = 10
    }
    value3, ok3 := request['category']
    if !ok3 {
            request['category'] = 'all'
    }


    return Base.request("GET","/dash/wallet/" + strconv.Itoa(request['wallet_id']) + "/transaction",{
            "category" : request['category'],
            "order" : request['order'],
            "offset" : request['offset'],
            "limit" : request['limit']
        })
    
}

func sendToAddress(request = map[string]interface{}) *Dash {


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


        
	return self.request("POST","/dash/wallet/" + strconv.Itoa(request['wallet_id']) + "/sendtoaddress",{"address" : request['address'],"amount" : request['amount'],"seed_wif" : request['seed_wif'],"password" : request['password'],"kbfee" : request['kbfee']})

    
}
func sendMany(request = map[string]interface{}) *Dash {
    value, ok := request['seed_wif']
    if !ok {

        request['seed_wif'] = nil
    }



    value1, ok1 := request['password']
    if !ok1 {
            request['password'] = nil
    }
            
        
	return Base.request("POST","/dash/wallet/" + strconv.Itoa(request['wallet_id']) + "/sendmany",{"to" : request['to'],"seed_wif" : request['seed_wif'],"password" : request['password']})	

    
}
func sendTransaction(request = map[string]interface{}) *Dash {

	return Base.request("POST","/dash/transaction",{"sign_hex" : request['sign_hex']})

}
func getTransaction(request = map[string]interface{}) *Dash {

	return Base.request("GET","/dash/transaction/" + request['hash'])		

}







