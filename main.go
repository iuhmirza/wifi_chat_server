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
	addresses, err := net.InterfaceAddrs()
	ips := []net.IP{}
	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.To4())
			}
		}
	}

	host, _ := os.Hostname()
	service, err := mdns.NewMDNSService(
		fmt.Sprintf("Wifi Chat - %v", host),
		"_wifichat._tcp",
		"",
		"",
		55555,
		ips,
		[]string{"Wifi Chat Server"},
	)

	fmt.Println(service.IPs)
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
