package main

import (
	"config"
	"endpointsHandler"
	"github.com/valyala/fasthttp"
)

// routing
func routingHandler (ctx *fasthttp.RequestCtx){
	ctx.SetContentType("text/json; charset=utf-8")

	switch string(ctx.Path()) {

	case "/createfile":
		middlewareEndpoint(ctx,endpointsHandler.HandlerCreateFillFile)
	case "/getcontentfile":
		middlewareEndpoint(ctx,endpointsHandler.HandlerGetContentFile)
	case "/deletefoldersub":
		middlewareEndpoint(ctx,endpointsHandler.HandlerDeleteFolderSubs)
	case "/getanaliticspath":
		middlewareEndpoint(ctx,endpointsHandler.HandlerGetAnaliticsPath)
	default:
		ctx.Error(config.ERRPATH,fasthttp.StatusNotFound)

	}
}



