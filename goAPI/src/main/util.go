package main

import (
	"../calc"
	"fmt"
	gocache "github.com/pmylund/go-cache"
	"github.com/valyala/fasthttp"
	"log"
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
	d := calc.ParseQueryStringToDict(ctx.QueryArgs().String())
	key:=string(ctx.Path())+ctx.QueryArgs().String()
	if _, found := cache.Get(key); !found {
		cache.Set(key, f(d), gocache.DefaultExpiration)
	}
	response,_:= cache.Get(key)
	_,err :=fmt.Fprint(ctx, fmt.Sprintf("%v", response))
	if err== nil {ctx.SetStatusCode(fasthttp.StatusOK)
	} else {ctx.SetStatusCode(fasthttp.StatusInternalServerError)}
}

