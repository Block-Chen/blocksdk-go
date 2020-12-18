package Webhook

import (  
    "fmt"
    "blocksdk-go/base"
)

type WebHook struct{
    base.Base
}

func create(request = map[string]interface{}) *WebHook {
	return Base.request("POST","/hook",{"callback" : request['callback'],"category" : request['category'],"address" : request['address']})

}


func list(request = map[string]interface{}) *WebHook {



    value, ok := request['offset']
    if !ok {
        request['offset'] = 0
    }

    value1, ok1 := request['limit']
    if !ok1 {
            request['limit'] = 10
    }
		
	return Base.request("GET","/hook",{"offset" : request['offset'],"limit" : request['limit']})
		
}


func get(request = map[string]interface{}) *WebHook {
		return Base.request("GET","/hook/" + strconv.Itoa(request['hook_id']))

    
}


func delete(request = map[string]interface{}) *WebHook {
	return Base.request("POST","/hook/" + strconv.Itoa(request['hook_id']) + "/delete")

    
}


func listResponse(request = map[string]interface{}) *WebHook {

    value, ok := request['offset']
    if !ok {
        request['offset'] = 0
    }

    value1, ok1 := request['limit']
    if !ok1 {
            request['limit'] = 10
    }
		
	return Base.request("GET","/hook/response",{"offset" : request['offset'],"limit" : request['limit']})			
			
}

func getResponse(request = map[string]interface{}) *WebHook {

    value, ok := request['offset']
    if !ok {
        request['offset'] = 0
    }

    value1, ok1 := request['limit']
    if !ok1 {
            request['limit'] = 10
    }
	
	return Base.request("GET","/hook/" + strconv.Itoa(request['hook_id']) + "/response",{"offset" : request['offset'],"limit" : request['limit']})			


}





