# udpflow

**udpflow** is a Go program that forwards UDP datagrams between two sockets.
It is primarily useful for establishing a proxy between IPv4 and IPv6.

## Installation

This program is written in Go.
You can compile and install this program with:

```bash
env CGO_ENABLED=0 go install github.com/yoursunny/udpflow@main
```

## Usage

The program accepts four positional arguments:

1. Local endpoint A.
2. Remote endpoint A.
3. Local endpoint B.
4. Remote endpoint B.

Each endpoint is written as `IPv4:port` or `[IPv6]:port`.
The program listens on the two local endpoints.
For each UDP datagram received from a remote endpoint, it is sent to the other remote endpoint.

Sample command:

```bash
./udpflow 192.0.2.1:4000 192.0.2.2:4000 [2001:db8:8ce8:70ef::1]:4000 [2001:db8:8ce8:70ef::2]:4000
```
