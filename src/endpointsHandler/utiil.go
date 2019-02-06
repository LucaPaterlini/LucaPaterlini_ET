package endpointsHandler

import "encoding/json"

type ResponseJson struct {
	Err bool `json:"err"`
	Data map[string]interface{} `json:"data"`
}
//composeJson compose the answer of each endpoint using a struct ResponseJson
//handling the add of the error message in case it's required
func composeJson(params map[string]interface{},err error)(s string){
	objR := &ResponseJson{}
	objR.Err=err!=nil
	objR.Data=make(map[string]interface{})
	if err==nil{
		for k,v := range params{objR.Data[k]=v}
	}else {
		objR.Data["errMsg"]=err.Error()
	}
	jsonString, _ := json.MarshalIndent(objR,"","\t")
	return string(jsonString)
}
//checkKeys checks the presence of all the keys in the map in the keys array
func checkKeys(d map[string]interface{},key[]string)bool{
	for _,v := range key {
		_,ok:=d[v]
		if ok==false{return false}
	}
	return true
}

