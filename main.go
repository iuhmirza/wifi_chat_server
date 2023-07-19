package main

import (
	"fmt"
	"github.com/grandcat/zeroconf"
	"log"
	"net/http"
	"os"
)

func main() {
	host, _ := os.Hostname()
	server, err := zeroconf.Register(
		fmt.Sprintf("Wifi Chat - %v", host),
		"_wifichat._tcp",
		"local.",
		55555,
		[]string{"test"},
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Shutdown()

	http.HandleFunc("/text", textHandler)
	log.Println("Listening and serving on port 55555")
	err = http.ListenAndServe(":55555", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
