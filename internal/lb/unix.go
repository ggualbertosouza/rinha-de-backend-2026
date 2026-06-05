package loadBalancer

import (
	"net"
	"os"
)

func (l *LoadBalancer) ListenUnix() error {
	_ = os.Remove(l.unixPath)

	addr := &net.UnixAddr{
		Name: l.unixPath,
		Net:  "unix",
	}

	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		return err
	}

	l.unixListener = listener

	for {
		conn, err := listener.AcceptUnix()
		if err != nil {
			if ne, ok := err.(*net.OpError); ok && !ne.Temporary() {
				return nil
			}

			continue
		}

		l.registerWorker(conn)
	}
}

func (l *LoadBalancer) registerWorker(conn *net.UnixConn) {
	raw, err := conn.SyscallConn()
	if err != nil {
		_ = conn.Close()
		return
	}

	var socketFD int

	err = raw.Control(func(fd uintptr) {
		socketFD = int(fd)
	})

	if err != nil {
		_ = conn.Close()
		return
	}

	for {
		currentPtr := l.workers.Load()

		current := *currentPtr

		next := make([]*Worker, len(current)+1)

		copy(next, current)

		next[len(current)] = &Worker{
			ID:       len(current) + 1,
			Conn:     conn,
			socketFD: socketFD,
		}

		if l.workers.CompareAndSwap(currentPtr, &next) {
			return
		}
	}
}
