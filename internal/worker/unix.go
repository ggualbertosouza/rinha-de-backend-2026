package worker

import (
	"log"
	"net"
)

func mustUnixAddr(path string) *net.UnixAddr {
	addr, err := net.ResolveUnixAddr("unix", path)
	if err != nil {
		log.Fatal(err)
	}
	return addr
}
