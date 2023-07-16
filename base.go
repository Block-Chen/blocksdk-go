package blocksdk_go

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Base struct {
	ChainName string
	ApiToken  string
	Endpoint  string
	Version   string
}

func NewBase(chainName string, apiToken string, endpoint string, version string) *Base {
	return &Base{
		ChainName: chainName,
		ApiToken:  apiToken,
		Endpoint:  endpoint,
		Version:   version,
	}
}

func (b *Base) Request(method string, path string, data map[string]interface{}) (map[string]interface{}, error) {
	baseURL, _ := url.Parse(b.Endpoint)
	baseURL.Path = b.Version + "/" + b.ChainName + path

	if method == "GET" && len(data) > 0 {
		params := url.Values{}
		for key, value := range data {
			strValue := ""
			if value == true {
				strValue = "true"
			} else if value == false {
				strValue = "false"
			} else {
				strValue = value.(string)
			}
			params.Add(key, strValue)
		}
		baseURL.RawQuery = params.Encode()
	}

	var jsonStr []byte
	if method == "POST" || method == "DELETE" {
		jsonStr, _ = json.Marshal(data)
	}

	req, _ := http.NewRequest(method, baseURL.String(), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-token", b.ApiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	return result, nil
}
