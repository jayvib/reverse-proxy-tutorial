package main

import (
	"log"
	"net/http"

	proxy "github.com/jayvib/reverse-proxy-tutorial"
)

func main() {
	proxy.LogSetup()
	http.HandleFunc("/", proxy.HandleRequestAndRedirect)
	if err := http.ListenAndServe(proxy.GetListenAddress(), nil); err != nil {
		log.Fatal(err)
	}
}
