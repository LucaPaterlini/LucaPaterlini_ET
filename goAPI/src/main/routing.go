package main

import (
	"../endpointsHandler"
	"github.com/valyala/fasthttp"
)

// routing
func routingHandler (ctx *fasthttp.RequestCtx){
	ctx.SetContentType("text/json; charset=utf-8")

	switch string(ctx.Path()) {

	//Call n°0
	case "/hello":
		middlewareEndpoint(ctx,endpointsHandler.HandlerHello)

	default:
		ctx.Error(ERRPATH,fasthttp.StatusNotFound)

	}
}



