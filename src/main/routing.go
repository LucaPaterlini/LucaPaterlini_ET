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
	case "/getanaliticspath":
		middlewareEndpoint(ctx,endpointsHandler.HandlerGetAnaliticsPath)
	case "/replaceconentfile":
		middlewareEndpoint(ctx,endpointsHandler.HandlerReplaceContentFile)
	case "/deletefoldersub":
		middlewareEndpoint(ctx,endpointsHandler.HandlerDeleteFolderSubs)


	default:
		ctx.Error(config.ERRPATH,fasthttp.StatusNotFound)

	}
}



