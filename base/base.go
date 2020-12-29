package base

import (  
    "bytes"
    "time"
    "net/http"
    "encoding/json"
)

const (
    layoutISO = "2006-01-02"
)

type Base struct {  
    api_token string
}

func NewBase(api_token string) *Base{
	return &Base{
		api_token : api_token,
	}
}

func (b *Base)getUsage(request map[string]interface{}) interface{}{
	_, ok := request["start_date"]
	if !ok {
		d := time.Now().AddDate(0,0,-7)
		request["start_date"] = d.Format(layoutISO)
	}

	_,ok2 := request["end_date"]
	if !ok2 {
		d := time.Now()
		request["end_date"] = d.Format(layoutISO)
	}
	
	return b.Request("GET","/usage",request)
}


func (b *Base)getHashType(request map[string]interface{}) interface{}{
	return b.Request("GET","/auto/" + request["hash"].(string) + "/type",request)
}

func (b *Base)Request(method string ,path string,data map[string]interface{}) map[string]interface{}{

	url := "https://api.blocksdk.com/v1" + path
	if method == "GET" && data != nil && len(data) > 0 {

		for key, element := range data {
		
			if element.(bool) == true{
				url += key +  "=true&"
			}else if element.(bool) == false {
				url += key + "=false&"
			}else {
				url += key+ "=" + element.(string) + "&"
			}
		}


	}




	client := &http.Client{}
	buff := bytes.NewBuffer(nil)
	if method == "POST"{
		pbytes, _ := json.Marshal(data)
		buff = bytes.NewBuffer(pbytes)
	}
	
	req, err := http.NewRequest(method,url,buff)
	req.Header.Add("X-API-KEY",b.api_token)
	req.Header.Add("Content-Type","application/json")
 
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
	defer resp.Body.Close()
 
	body := json.NewDecoder(resp.Body)

	var res  map[string]interface{}
	body.Decode(&res)
	
	res["Header"] = resp.Header
	res["StatusCode"] = resp.StatusCode
	
	return res
}












