package main

import (
	"calc"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net/url"
)

func logPanics(function fasthttp.RequestHandler)fasthttp.RequestHandler{
	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if x := recover(); x!=nil{
				log.Printf("[%v] caught panic: %v",ctx.RemoteAddr(),x)
			}
		}()
		function(ctx)
	}
}

func middlewareEndpoint(ctx *fasthttp.RequestCtx,f func(map[string]interface{})string){
	query,err := url.QueryUnescape(ctx.QueryArgs().String())
	if err==nil {
		d := calc.ParseQueryStringToDict(query)
		_, err = fmt.Fprint(ctx, fmt.Sprintf("%v", f(d)))
	}
	if err== nil {ctx.SetStatusCode(fasthttp.StatusOK)
	} else {ctx.SetStatusCode(fasthttp.StatusInternalServerError)}
}

