package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/hashicorp/mdns"
)

func main() {
	host, _ := os.Hostname()
	service, err := mdns.NewMDNSService(
		fmt.Sprintf("Wifi Chat - %v", host),
		"_wifichat._tcp",
		"",
		"",
		55555,
		[]net.IP{},
		[]string{"Wifi Chat Server"},
	)
	if err != nil {
		log.Println(err)
		return
	}
	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Shutdown()

	http.HandleFunc("/text", textHandler)
	log.Println("Listening and serving on port 55555")
	http.ListenAndServe(":55555", nil)
}
