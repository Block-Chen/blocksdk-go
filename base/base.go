package base

import (  
    "strconv"
    "bytes"
    "net/http"
    "encoding/json"
)


type Base struct {  
    api_token string
}

func New(api_token string) *Base{
	return &Base{
		api_token : api_token,
	}
}

func (b *Base)ToString(data interface{}) string{
	switch v := data.(type) { 
		case bool: 
			if v == true {
				return "true"
			}else{
				return "false"
			}
		case int: 
			return strconv.Itoa(v)
		case string: 
			return v
	} 
	
	return ""
}


func (b *Base)Request(method string ,path string,data map[string]interface{}) map[string]interface{}{

	url := "https://api.blocksdk.com/v2" + path
	if method == "GET" && data != nil && len(data) > 0 {

		for key, element := range data {
			url += key+ "=" + b.ToString(element) + "&"
		}
	}


	client := &http.Client{}
	buff := bytes.NewBuffer(nil)
	if method == "POST"{
		pbytes, _ := json.Marshal(data)
		buff = bytes.NewBuffer(pbytes)
	}
	
	req, err := http.NewRequest(method,url,buff)
	req.Header.Add("X-API-TOKEN",b.api_token)
	req.Header.Add("Content-Type","application/json")
 
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
	defer resp.Body.Close()
 
	body := json.NewDecoder(resp.Body)

	var res  map[string]interface{}
	body.Decode(&res)
	
	return res
}












