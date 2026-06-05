package worker

import (
	"bufio"
	"net"
	"net/http"
	"os"
	"sync"
)

var readerPool = sync.Pool{
	New: func() any {
		return bufio.NewReaderSize(nil, 4096)
	},
}

func handleFD(fd int, handler http.Handler) {
	file := os.NewFile(uintptr(fd), "tcp")
	defer file.Close()

	conn, err := net.FileConn(file)
	if err != nil {
		return
	}
	defer conn.Close()

	br := readerPool.Get().(*bufio.Reader)
	br.Reset(conn)
	defer readerPool.Put(br)

	req, err := http.ReadRequest(br)
	if err != nil {
		return
	}

	rw := newResponseWriter(conn)

	handler.ServeHTTP(rw, req)

	_ = rw.Flush()
}
