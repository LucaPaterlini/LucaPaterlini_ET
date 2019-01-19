package coreFunctions

import (
	"config"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)


func CreateFillFile(filename,path,payload string,replace bool) (err error){
	targetPath := filepath.Join(config.DefaultPathStorage,path)
	targetPathAndFile := filepath.Join(targetPath,filename)
	// check if the file path and the file exists
	if !replace {
		if _, err = os.Stat(targetPath); !os.IsNotExist(err) {
			return errors.New("File already Inserted: " + targetPathAndFile)
		}
	}
	err = os.MkdirAll(targetPath, 0700)
	if err!=nil{return }
	f, err := os.Create(targetPathAndFile)
	if err!=nil {return }
	_,err = f.Write([]byte(payload))
	if err!=nil{return }
	err = f.Close()
	return
}

func GetContentFile(filename,path string)(payload string, err error){
	targetPathAndFile := filepath.Join(config.DefaultPathStorage,path,filename)
	b, err := ioutil.ReadFile(targetPathAndFile)
	if err!=nil {return }
	payload=string(b)
	return
}

func DeleteFolderSubs(path string)(err error){
	targetPath := filepath.Join(config.DefaultPathStorage,path)
	if _, err = os.Stat(targetPath); os.IsNotExist(err) {
		return errors.New("Directory already Deleted: " + targetPath)
	}
	err = os.RemoveAll(targetPath)
	return
}
func GetAnaliticsPath(path string)(json string,err error){
	targetPath := filepath.Join(config.DefaultPathStorage,path)
	b, err := exec.Command("python3",config.DefaultPathPythonScript,targetPath).CombinedOutput()
	json=string(b)
	return
}