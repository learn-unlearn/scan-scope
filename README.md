## ScanScope [just a random name to give the project]

### --- Command line tool ---

### Docker

    ```
    docker run -u ${UID}:${UID} -it -v ./:/usr/src/app -w /usr/src/app golang bash
    docker run -it -v ./:/usr/src/app -w /usr/src/app golang bash ====== this runs as root
    docker run -it -v ./:/usr/src/app -w /usr/src/app --network=host golang bash  ====== runs in host mode
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

### System commands to check your cpu and more

    - nproc                            # CPU cores
    - ulimit -n                        # max open files for the shell
    - cat /proc/sys/fs/file-max        # system-wide FD limit

# While running scan, in another terminal:

    - ss -s                            # socket summary
    - ss -tn state established         # established TCP sockets
    - watch -n1 'ss -s; echo; ls /proc/$(pgrep yourbinary)/fd | wc -l'   # monitor sockets + fds