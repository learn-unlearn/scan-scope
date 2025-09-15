## ScanScope [just a random name to give the project]

### --- Command line tool ---

### Docker

    ```
    docker run -u ${UID}:${UID} -it -v ./:/usr/src/app -w /usr/src/app golang bash
    docker run -it -v ./:/usr/src/app -w /usr/src/app golang bash ====== this runs as root
    ```

## Run in container

    ```
    - export GOCACHE=/tmp/go-build // run this in docker container
    - apt-get update && apt-get install -y iproute2
    - apt-get install -y iputils-ping
    - go run main.go
    - go build -o scanner || go build
    ```

### Objectives

    [] Discover active hosts on a local network (ICMP "ping sweep").
    [] Scan these hosts (or a specific target) for open TCP ports.
    [] Be efficient by scanning multiple hosts/ports concurrently using goroutines.

### Core Features

    [] Input: Accept a target IP address, hostname, or a CIDR range (e.g., 192.168.1.0/24).
    [] Ping Sweep: Identify which hosts in a range are online.
    [] Port Scanning: Check a list of common ports or a user-defined range of ports.
    [] Output: Clearly display which hosts are online and which ports are open on them.
    [] Concurrency: Use goroutines to make the scan fast.

### Additional [if possible]

    [] Service banner grabbing (connecting to an open port and reading the first response).
    [] UDP port scanning (much more complex and unreliable).
    [] Adjustable timeout and number of concurrent workers.