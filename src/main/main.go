//Package main handle the main of the ET project
package main

import (
	"config"
	"flag"
	"fmt"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
)

var (
	addr = flag.String("addr",config.IPADDR+":"+strconv.Itoa(config.PORT),"TCP address to listen to")
	compress = flag.Bool("compress",false, "Whether to enable transparent response compression ")
)



func main(){

	// Corse Handeler
	withCors := cors.DefaultHandler()

	flag.Parse()
	h := withCors.CorsMiddleware(logPanics(routingHandler))
	if *compress {
		h = fasthttp.CompressHandler(h)
	}
	fmt.Println("Service started")

	if err := fasthttp.ListenAndServe(*addr,h);err != nil{
		log.Fatalf("Errror in ListenAndServer: %s",err)
	}

}


