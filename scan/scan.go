package scan

import (
	"fmt"
	"sync"
	"os/exec"
	"runtime"
	"strings"
	"net"
	// "log"
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

func DetectDevices(subnet string, batch_size int) {
	timeout := 100 * time.Millisecond
	fmt.Println("Scaning network started: %v/255", subnet)
	fmt.Println("Scaning network default timeout: %v", timeout)

	ips := make([]string, 0, 254)
	for i := 1; i <= 254; i++ {
		ips = append(ips, fmt.Sprintf("%s.%d", subnet, i))
	}
	// fmt.Println("Generated IP addresses:", ips)

	// generate port
	ports := make([]int, 0, 9000)
	for p := 1000; p <= 9999; p++ {
		ports = append(ports, p)
	}
	// fmt.Println("Generated ports:", ports)

	type job struct {
		ip   string
		port int
	}

	worker := func(id int, jobs <-chan job, results chan<- string, wg *sync.WaitGroup) {
		defer wg.Done()
		for j := range jobs {
			addr := fmt.Sprintf("%s:%d", j.ip, j.port)
			conn, err := net.DialTimeout("tcp", addr, timeout)
			if err == nil {
				conn.Close()
				results <- fmt.Sprintf("Device found at %s:%d \n", j.ip, j.port)
			}
		}
	}

	// result channel & collector goroutine
	results := make(chan string, 1000)
	var resWg sync.WaitGroup
	resWg.Add(1)
	go func() {
		defer resWg.Done()
		for r := range results {
			fmt.Println(r)
		}
	}()

	for start := 0; start < len(ips); start += batch_size {
		end := start + batch_size
		if end > len(ips) {
			end = len(ips)
		}

		batch := ips[start:end]
		fmt.Println("starting batch scan for %d..%d (size=%d)", start + 1, end, len(batch))

		// create job channel for this batch
		jobs := make(chan job, len(batch) * 10)
		var wg sync.WaitGroup

		// start worker for this batch
		for w := 0; w < runtime.NumCPU(); w++ {
			wg.Add(1)
			go worker(w, jobs, results, &wg)
		}

		// Enqueue jobs for this batch
		for _, ip := range batch {
			for _, port := range ports {
				jobs <- job{ip: ip, port: port}
			}
		}

		close(jobs)
		wg.Wait()
		fmt.Println("batch scan complete for %d..%d", start + 1, end)
	}

	close(results)
	resWg.Wait()
	fmt.Println("Network Scan complete")

	// var wg sync.WaitGroup


	// for i := 1; i <= 255; i++ {
	// 	ip := fmt.Sprintf("%s.%d", subnet, i)
	// 	// fmt.Println("Pinging IP:", ip)
	// 	for _, port := range 9999 {
	// 		wg.Add(1)
	// 		go func (ip string, port int) {
	// 			defer wg.Done()
	// 			addr = fmt.Sprintf("%s:%d", ip, port)
	// 			conn, err := net.DialTimeout("tcp", addr, timeout)
	// 			if err == nil {
	// 				conn.Close()
	// 				fmt.Println("Device found at %s \n", ip)
	// 			}
	// 		}(ip, port)
	// 	}
	// 	wg.Add(1)
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
	// }
	// wg.Wait()
	// log.Println("Network Scan complete")
}