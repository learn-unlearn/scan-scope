package subnet

import (
	"fmt"
	"net"
	"strings"
)

func GetSubnet(local_ip net.IP) (string, error) {
	// Get network subnet
	ifaces, err := net.Interfaces() // returns an array of interfaces -> [{1 65536 lo  up|loopback|running} {2 1500 eth0 32:55:e1:7c:5c:c5 up|broadcast|multicast|running}]
	if err != nil {
		return "", fmt.Errorf("could not run command gateway", err)
	}

	var subnet string
	for _, iface := range ifaces { // fmt.Println("Network Interfaces:", ifaces)
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		fmt.Println("Subnet address range:", addrs)
		for _, addr := range addrs {
			// fmt.Println("Address String:", addr.String())
			if ip_net, ok := addr.(*net.IPNet); ok && !ip_net.IP.IsLoopback() {
				if ip_net.IP.To4() != nil && ip_net.Contains(local_ip) {
					subnet = ip_net.IP.To4().String()[:strings.LastIndex(ip_net.IP.To4().String(), ".")]
					break 
				}
			}
		}

		if subnet != "" {
			break
		}
		// fmt.Println("Interface Addresses:", addrs)
	}

	if subnet == "" {
		return "", fmt.Errorf("could not run command gateway", err)
	}

	// fmt.Println("Local Subnet:", subnet)
	return subnet, nil
}