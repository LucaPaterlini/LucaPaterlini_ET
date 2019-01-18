package coreFunctions

import (
	"../config"
	"fmt"
	"os"

)

func Hello(name string)(s string, e error){
	s=fmt.Sprintf("Hello %s",name)
	return
}


func CreateFillFile(filename,path,payload string) (e error){
	f, err := os.Create(config.DefaultPath+path+filename)
	if err!=nil {return }
	_,err = f.Write([]byte(payload))
	return
}