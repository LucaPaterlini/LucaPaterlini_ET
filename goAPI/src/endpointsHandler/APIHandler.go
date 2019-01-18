package endpointsHandler

import (
	"../coreFunctions"
	"errors"
)

func HandlerHello(d map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(d,[]string{"name"})==true {
		item, e := coreFunctions.Hello(d["name"].(string))
		err = e
		dict = map[string]interface{}{"message": item}
	}else{
		err=errors.New("inputParams: wrong set of input parameters")
	}
	return composeJson(dict,err)
}
