package price

import (  
    "fmt"
    "blocksdk-go/base"
)

type Price struct{
    Base base.Base
}


func listPrice(request = map[string]interface{}) *Price {
	return Base.request("GET","/price")
}