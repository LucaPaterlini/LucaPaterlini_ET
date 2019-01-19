package endpointsHandler

import (
	"coreFunctions"
	"encoding/json"
	"errors"
)

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



