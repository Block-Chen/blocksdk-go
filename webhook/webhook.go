package webhook

import (  
    "github.com/Block-Chen/blocksdk-go/base"
)

type WebHook struct{
    Base *base.Base
}

func NewWebHook(api_token string) *WebHook{
	b := base.NewBase(api_token)
	
	return &WebHook{
		Base : b,
	}
}

func (w *WebHook)create(request  map[string]interface{}) map[string]interface{} {
	return w.Base.Request("POST","/hook",request)
}


func (w *WebHook)list(request  map[string]interface{}) map[string]interface{} {
    _, ok := request["offset"]
    if !ok {
        request["offset"] = 0
    }

    _, ok1 := request["limit"]
    if !ok1 {
            request["limit"] = 10
    }
		
	return w.Base.Request("GET","/hook",request)
		
}


func (w *WebHook)get(request  map[string]interface{}) map[string]interface{} {
		return w.Base.Request("GET","/hook/" + request["hook_id"].(string),nil)

    
}


func (w *WebHook)delete(request  map[string]interface{}) map[string]interface{} {
	return w.Base.Request("POST","/hook/" + request["hook_id"].(string) + "/delete",nil)
}


func (w *WebHook)listResponse(request  map[string]interface{}) map[string]interface{} {

    _, ok := request["offset"]
    if !ok {
        request["offset"] = 0
    }

    _, ok1 := request["limit"]
    if !ok1 {
            request["limit"] = 10
    }
		
	return w.Base.Request("GET","/hook/response",request)						
}

func (w *WebHook)getResponse(request  map[string]interface{}) map[string]interface{} {

    _, ok := request["offset"]
    if !ok {
		request["offset"] = 0
    }

    _, ok1 := request["limit"]
    if !ok1 {
       request["limit"] = 10
    }
	
	return w.Base.Request("GET","/hook/" + request["hook_id"].(string) + "/response",request)			
}