// Package calc provides utility functions that requires only cpu and ram
package calc

import (
	"github.com/asaskevich/govalidator"
	"strconv"
	"strings"
)


// ParseQueryStringToDict parse the query string (key=val&key1=val1)
// and returns a dictionary 
func ParseQueryStringToDict(a string) map[string]interface{}{
	d := make(map[string]interface{})
	if len(a)<1 {return d}
	for _,t := range  strings.Split(a,"&"){
		g:= strings.Split(t,"=")
		if len(g)<2{continue}
		k,v :=g[0],g[1]
		if govalidator.IsInt(v){
			d[k], _ =strconv.ParseInt(v,10,64)
		}else if govalidator.IsFloat(v){
			d[k], _ =strconv.ParseFloat(v,64)
		}else {
			d[k]=v
		}
	}
	return d
}


