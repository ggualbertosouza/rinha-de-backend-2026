package worker

import (
	"log"
	"net"
	"net/http"
	"time"
)

func Run(unixSocket string, handler http.Handler) {
	for {
		conn, err := net.DialUnix(
			"unix",
			nil,
			mustUnixAddr(unixSocket),
		)

		if err != nil {
			log.Printf("unable to connect to lb: %v", err)

			time.Sleep(time.Second)

			continue
		}

		log.Println("connected to lb")

		runWorker(conn, handler)

		log.Println("connection lost, reconnecting")
	}
}

func runWorker(conn *net.UnixConn, handler http.Handler) {
	defer conn.Close()

	for {
		fd, err := recvFD(conn)
		if err != nil {
			return
		}

		if fd == 0 {
			continue
		}

		go handleFD(fd, handler)
	}
}
