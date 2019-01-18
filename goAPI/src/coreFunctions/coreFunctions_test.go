package coreFunctions

import (
	"testing"
)

func TestCreateFillFile(t *testing.T) {
	err :=CreateFillFile("smile.txt","/sunshine/or/rain/","smile always",false)
	if err!=nil {
		t.Error(err.Error())
	}
}

func TestGetContentFile(t *testing.T) {
	msg,err := GetContentFile("smile.txt","/sunshine/or/rain/")
	expectedResponse := "smile always"
	if err!=nil {t.Error(err.Error())
	}else if msg!=expectedResponse{
		t.Errorf("Expected: %s Got: %s",expectedResponse,msg)
	}
}

func TestReplaceFile(t *testing.T) {
	err :=CreateFillFile("smile.txt","/sunshine/or/rain/","smile always2",true)
	if err!=nil {
		t.Error(err.Error())
	}
}

func TestDeleteFolderFiles(t *testing.T) {
	err :=DeleteFolderSubs("/sunshine/or/rain/")
	if err!=nil {t.Error(err.Error())
	}
}

func TestGetAnaliticsPath(t *testing.T) {
	_,err := GetAnaliticsPath("/sunshine")
	if err!=nil {t.Error(err.Error())}
}