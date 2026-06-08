package loadBalancer

import (
	"log"
	"net"
)

func (l *LoadBalancer) ListenTcp() error {
	listener, err := net.Listen("tcp", ":"+l.httpPort)
	if err != nil {
		return err
	}

	l.listener = listener

	for {
		conn, err := listener.Accept()

		if err != nil {
			if ne, ok := err.(*net.OpError); ok && !ne.Temporary() {
				return nil
			}

			continue
		}

		log.Printf("[LB] accepted connection from=%s", conn.RemoteAddr())

		l.handleTCP(conn)
	}
}

func (l *LoadBalancer) handleTCP(conn net.Conn) {
	defer conn.Close()

	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		return
	}

	file, err := tcpConn.File()
	if err != nil {
		return
	}
	defer file.Close()

	fd := int(file.Fd())

	worker := l.pickWorker()
	if worker == nil {
		log.Printf("[LB] no worker available")
		return
	}

	log.Printf("[LB] forwarding connection to worker=%d fd=%d", worker.ID, fd)

	if err := sendFD(worker.socketFD, fd); err != nil {
		log.Printf("[LB] failed forwarding connection worker=%d err=%v", worker.ID, err)
		return
	}

	log.Printf("[LB] connection dispatched worker=%d", worker.ID)
}
