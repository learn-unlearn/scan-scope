package get_ip

import (
	"net"
	"fmt"
)

func get_local_ip() {}

func GetLocalInternetIp () (net.IP, error) {
	// get the local ip
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, fmt.Errorf("Error getting local IP: %v", err)
	}

	defer conn.Close();
	// fmt.Println("Local Address from udp 8.8.8.8:80:", conn.LocalAddr())
	// fmt.Println("Local IP Address:", strings.Split(conn.LocalAddr().String(), ":")[0])

	local_addr := conn.LocalAddr().(*net.UDPAddr)
	local_ip := local_addr.IP
	fmt.Println("Local Address", local_addr)
	// fmt.Println("Local IP Address:", local_ip)
	// fmt.Println("----------------------------------------------------------------------------")

	return local_ip, nil
}
