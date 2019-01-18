package endpointsHandler

import (
"fmt"
"testing"
)

func TestHandlerHello(t *testing.T) {
	d:=map[string]interface{}{"name":"luca"}
	fmt.Print(HandlerHello(d))
}
