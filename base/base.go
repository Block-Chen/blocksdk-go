package base

import (  
    "fmt"
    "time"
    "strconv"
    "net/http"
    "encoding/json"
    "log"
    "strings"
)

const (
    layoutISO = "2006-01-02"
)
type Base struct {  
    api_token string
}

func (b *Base) Init(api_token string){
    b.api_token = api_token
}
func getUsage(request = map[string]interface{}) *Base {
	value, ok := request['start_date']

	if !ok {
		d := time.Now() - 604800
		request['start_date'] = d.Format(layoutISO)
	}

	value2,ok2 :=request['end_date']
	if !=ok2 {
		d := time.Now()
		request['end_date'] = d.Format(layoutISO)
	}

	return request("GET","/usage",{"start_date": request['start_date'],"end_date": request['end_date']})

}

func listPrice(request = map[string]interface{}) *Base{
	return request("GET","/price")
}

func getHashType(request = map[string]interface{}) *Base{
	return request("GET","/auto/" + request ['hash'] + "/type")
}

func request(method string ,path string,data = map[string]interface{}) ) *Base{

	url := "https://api.blocksdk.com/v1" + path
	if method == "GET" && len(data) > 0 {

		for key, element := range data {
			value = element
			if value == true{
				url += key +  "=true&"
			}
			else if value == false {
				url += key + "=false&"
			}
			else {
				url += key+ "=" + strconv.Itoa(value) + "&"
			}
		}


	}




	if method == "POST"{
		jsonReq, err := json.Marshal(api_token)
		response, err := http.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(api_token)) 

	}
	else {
		jsonReq, err := json.Marshal(api_token)
		response, err := http.Get(url, "application/json; charset=utf-8", bytes.NewBuffer(api_token)) 
	}
	if err != nil {
		panic(err)
	}

	body := json.NewDecoder(response.Body)
	
	if method == "POST" {

		body = json.NewDecoder(response.Body)
		converted_json := strings.ReplaceAll(response.Body,':','":')
		converted_json1 := strings.Split(converted_json,'{\n')

		//remaining
		for i := 0; i < len(converted_json1); i++ {
			if converted_json1[i] != ''{

				json2 := strings.Split(converted_json1[i],('":'))
				json2 = strings.Split(json2,' ')
				json2 = json2[len(json2)-1]

				converted_json := strings.ReplaceAll(converted_json, json2,'"'+ json2)

			}

			}

		body = json.loads(converted_json)
	}


			
	if response.Header{
		headers := response.Header
	}
	else{
		headers := {}
	}

	if response.StatusCode{
		status := response.StatusCode
	}
	else{
		status := 0
	}
	
	headers['StatusCode'] = status
	body['Header'] = headers	
	return body
	
	
}












