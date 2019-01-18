package coreFunctions

import "testing"

func TestHello(t *testing.T) {
	s, e := Hello("luca")
	expectedAnswer := "Hello luca"
	if e != nil {
		t.Error("Error in TestHello",e.Error())
	}else if s!=expectedAnswer{
		t.Errorf("wrong message; got %s instead of %s",s,expectedAnswer)
	}
}

func TestCreateFillFile(t *testing.T) {
	err :=CreateFillFile("smile.txt","/sunshine/or/rain/","smile always")
	if err!= nil {
		t.Error("Error in TestCreateFillFile", err.Error())
	}
}