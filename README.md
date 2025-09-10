### ScanScope [just a random name to give the project]

#### --- Command line tool ---

##### Solor docker run
docker run -u ${UID}:${UID} -it -v ./:/usr/src/app -w /usr/src/app golang bash

###### Run in container
go run main.go

####
[] Discover active hosts on a local network (ICMP "ping sweep").
[] Scan these hosts (or a specific target) for open TCP ports.
[] Be efficient by scanning multiple hosts/ports concurrently using goroutines.