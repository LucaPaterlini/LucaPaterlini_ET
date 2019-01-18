package calc

import (
	"testing"
)

func TestParseQueryStringToDict(t *testing.T){
	str := "text=hello&fun=true"
	dict:=ParseQueryStringToDict(str)
	if dict["text"]!="hello" {t.Error("Wrong dictionary read")}
}
