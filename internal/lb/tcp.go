package loadBalancer

import (
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
		return
	}

	_ = sendFD(worker.socketFD, fd)
}
