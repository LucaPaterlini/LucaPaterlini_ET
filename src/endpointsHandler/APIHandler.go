//Package endpointsHandler incapsulate the handler of the endpoint handlers of the ET project
package endpointsHandler

import (
	"coreFunctions"
	"encoding/json"
	"errors"
)

//HandlerCreateFillFile  check the parameters and then execute the core function CreateFillFile
//managing the responses in case of error as result of the function or in case the input is incorrect
func HandlerCreateFillFile(d map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(d,[]string{"filename","path","payload"})==true {
		e := coreFunctions.CreateFillFile(d["filename"].(string),d["path"].(string),d["payload"].(string),false)
		err = e
	}else{
		err=errors.New("inputParams: wrong set of input parameters")
	}
	return composeJson(dict,err)
}

//HandlerGetContentFile  check the parameters and then execute the core function GetContentFile
//managing the responses in case of error as result of the function or in case the input is incorrect
func HandlerGetContentFile(d map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(d,[]string{"filename","path",})==true {
		msg,e := coreFunctions.GetContentFile(d["filename"].(string),d["path"].(string))
		err = e
		dict = map[string]interface{}{"content": msg}
	}else{
		err=errors.New("inputParams: wrong set of input parameters")
	}
	return composeJson(dict,err)
}

//HandlerReplaceContentFile  check the parameters and then execute the core function ReplaceContentFile
//managing the responses in case of error as result of the function or in case the input is incorrect
func HandlerReplaceContentFile(d map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(d,[]string{"filename","path","payload"})==true {
		e := coreFunctions.CreateFillFile(d["filename"].(string),d["path"].(string),d["payload"].(string),true)
		err = e
	}else{
		err=errors.New("inputParams: wrong set of input parameters")
	}
	return composeJson(dict,err)
}

//HandlerDeleteFolderSubs  check the parameters and then execute the core function DeleteFolderSubs
//managing the responses in case of error as result of the function or in case the input is incorrect
func HandlerDeleteFolderSubs(d map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(d,[]string{"path"})==true {
		e := coreFunctions.DeleteFolderSubs(d["path"].(string))
		err = e
	}else{
		err=errors.New("inputParams: wrong set of input parameters")
	}
	return composeJson(dict,err)
}


//HandlerDeleteFolderSubs  check the parameters and then execute the core function DeleteFolderSubs
//managing the responses in case of error as result of the function or in case the input is incorrect
func HandlerGetAnaliticsPath(d map[string]interface{})string{
	dict :=make(map[string]interface{})
	var raw []map[string]map[string]interface{}
	var err error
	if checkKeys(d,[]string{"path"})==true {
		msg,e := coreFunctions.GetAnaliticsPath(d["path"].(string))
		err = e
		if err==nil{
			err =json.Unmarshal([]byte(msg), &raw)
			dict = map[string]interface{}{"data": raw}
		}

	}else{
		err=errors.New("inputParams: wrong set of input parameters")
	}
	return composeJson(dict,err)
}



