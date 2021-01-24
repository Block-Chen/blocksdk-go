package webhook

import (  
    "github.com/Block-Chen/blocksdk-go/base"
)

type WebHook struct{
    Base *base.Base
}

func New(api_token string) *WebHook{
	b := base.New(api_token)
	
	return &WebHook{
		Base : b,
	}
}

func (w *WebHook)Create(request  map[string]interface{}) map[string]interface{} {
	return w.Base.Request("POST","/hooks",request)
}


func (w *WebHook)GetWebhooks(request  map[string]interface{}) map[string]interface{} {
	return w.Base.Request("GET","/hooks",request)
}


func (w *WebHook)Get(request  map[string]interface{}) map[string]interface{} {
	return w.Base.Request("GET","/hooks/" + w.Base.ToString(request["hook_id"]),nil)
}


func (w *WebHook)Delete(request  map[string]interface{}) map[string]interface{} {
	return w.Base.Request("DELETE","/hooks/" + w.Base.ToString(request["hook_id"]),nil)
}


func (w *WebHook)ListResponse(request  map[string]interface{}) map[string]interface{} {
	return w.Base.Request("GET","/hooks/responses",request)						
}

func (w *WebHook)GetResponse(request  map[string]interface{}) map[string]interface{} {
	return w.Base.Request("GET","/hooks/" + w.Base.ToString(request["hook_id"]) + "/responses",request)			
}