package main

import (
	"fmt"
	"os"
	"runtime"
	"github.com/learn-unlearn/scanscope/scan"
	"github.com/learn-unlearn/scanscope/get_ip"
	"github.com/learn-unlearn/scanscope/subnet"
	"log"
)


func main () {
	fmt.Println("============================================================================")
	fmt.Println("Host operating system:", runtime.GOOS)
	fmt.Println("============================================================================")
	fmt.Println("============================================================================")

	// check for args
	if len(os.Args) < 2 {
		fmt.Println("Usage: %s <target>\n", os.Args[0])
		fmt.Println("Example: ./scanner 192.168.8.1")
		os.Exit(1)
	}

	// get the target
	target := os.Args[1]
	fmt.Println("Scanning target:", target) // fmt.Println("Checking what is my OS:", runtime.GOOS)

	fmt.Println("----------------------------------------------------------------------------")
	router, err := scan.GetGatewayIp()
	if err != nil {
		log.Fatalf("Error finding router IP: %v", err)
	}
	fmt.Printf("Router (Default Gateway) IP: %s\n", router)
	fmt.Println("----------------------------------------------------------------------------")

	local_ip, err := get_ip.GetLocalInternetIp()
	if err != nil {
		log.Fatalf("Error finding local IP: %v", err)
	}
	fmt.Printf("Local IP: %s\n", local_ip)
	fmt.Println("----------------------------------------------------------------------------")

	subnet, err := subnet.GetSubnet(local_ip)
	if err != nil {
		log.Fatalf("Error finding subnet: %v", err)
	}
	fmt.Println("Subnet:", subnet)
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------------")

	scan.DetectDevices(subnet)

	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------------")
}