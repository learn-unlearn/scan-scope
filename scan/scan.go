package scan

import (
	"fmt"
	"sync"
	"os/exec"
	"runtime"
	"strings"
	"net"
	"log"
	"time"
)

func GetGatewayIp() (string, error) {
	var cmd *exec.Cmd

	if runtime.GOOS == "linux" {
		cmd = exec.Command("ip", "route", "show", "default")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("could not run command gateway", err)
	}

	lines := strings.Split(string(output), "\n")
	for _,line := range lines {
		trimmed_lines := strings.TrimSpace(line) // why does this lines of split array return string here?
		if runtime.GOOS == "linux" && strings.HasPrefix(trimmed_lines, "default") {
			fields := strings.Fields(trimmed_lines)
			if len(fields) >= 2 {
				// fmt.Println(fields)
				return fields[2], nil
			}
		}
	}

	return "", fmt.Errorf("could not parse gateway IP from command output")
}

func DetectDevices(subnet string) {
	fmt.Println("Scaning network started: %v/255", subnet)
	var wg sync.WaitGroup

	timeout := 100 * time.Millisecond

	for i := 1; i <= 255; i++ {
		ip := fmt.Sprintf("%s.%d", subnet, i)
		// fmt.Println("Pinging IP:", ip)
		for _, port := range 9999 {
			wg.Add(1)
			go func (ip string, port int) {
				defer wg.Done()
				addr = fmt.Sprintf("%s:%d", ip, port)
				conn, err := net.DialTimeout("tcp", addr, timeout)
				if err == nil {
					conn.Close()
					fmt.Println("Device found at %s \n", ip)
				}
			}(ip, port)
		}
		wg.Add(1)
		// go func (ip string) {
		// 	defer wg.Done()
		// 	// Attempt to connect to common port
		// 	conn, err := net.DialTimeout("tcp", ip+":80", timeout)
		// 	if err == nil {
		// 		conn.Close()
		// 		fmt.Println("Device found at %s \n", ip)
		// 	}
		// 	// fmt.Printf("Error:", err) // print a dot for each ping attempt
		// }(ip)
	}
	wg.Wait()
	log.Println("Network Scan complete")
}