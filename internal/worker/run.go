package worker

import (
	"log"
	"net"
	"net/http"
)

func Run(unixSocket string, handler http.Handler) {
	conn, err := net.DialUnix(
		"unix",
		nil,
		mustUnixAddr(unixSocket),
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fd, err := recvFD(conn)
		if err != nil {
			continue
		}

		if fd == 0 {
			continue
		}

		go handleFD(fd, handler)
	}
}
