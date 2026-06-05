package loadBalancer

import (
	"net"
	"os"
	"sync/atomic"
)

type LoadBalancer struct {
	httpPort string
	unixPath string

	listener     net.Listener
	unixListener *net.UnixListener

	workers atomic.Pointer[[]*Worker]

	nextWorker atomic.Uint64
}

type Worker struct {
	ID int

	Conn *net.UnixConn

	socketFD int
}

func NewLoadBalancer(port, unixPath string) *LoadBalancer {
	lb := &LoadBalancer{
		httpPort: port,
		unixPath: unixPath,
	}

	workers := make([]*Worker, 0, 8)
	lb.workers.Store(&workers)

	return lb
}

func (l *LoadBalancer) Close() error {
	if l.listener != nil {
		_ = l.listener.Close()
	}

	if l.unixListener != nil {
		_ = l.unixListener.Close()
	}

	if l.unixPath != "" {
		_ = os.Remove(l.unixPath)
	}

	return nil
}

func (l *LoadBalancer) pickWorker() *Worker {
	workersPtr := l.workers.Load()

	if workersPtr == nil {
		return nil
	}

	workers := *workersPtr

	if len(workers) == 0 {
		return nil
	}

	idx := l.nextWorker.Add(1)

	return workers[idx%uint64(len(workers))]
}
