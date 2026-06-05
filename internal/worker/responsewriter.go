package worker

import (
	"bufio"
	"bytes"
	"net"
	"net/http"
	"strconv"
	"sync"
)

var bodyPool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}

var writerPool = sync.Pool{
	New: func() any {
		return bufio.NewWriterSize(nil, 4096)
	},
}

type responseWriter struct {
	conn        net.Conn
	header      http.Header
	statusCode  int
	wroteHeader bool

	body *bytes.Buffer
}

func newResponseWriter(conn net.Conn) *responseWriter {
	buf := bodyPool.Get().(*bytes.Buffer)
	buf.Reset()

	return &responseWriter{
		conn:       conn,
		header:     make(http.Header, 4),
		statusCode: http.StatusOK,
		body:       buf,
	}
}

func (w *responseWriter) Header() http.Header {
	return w.header
}

func (w *responseWriter) WriteHeader(status int) {
	if w.wroteHeader {
		return
	}

	w.statusCode = status
	w.wroteHeader = true
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}

	return w.body.Write(b)
}

func (w *responseWriter) Flush() error {
	defer bodyPool.Put(w.body)

	contentType := w.header.Get("Content-Type")
	if contentType == "" {
		contentType = "text/plain; charset=utf-8"
	}

	bw := writerPool.Get().(*bufio.Writer)
	bw.Reset(w.conn)

	defer func() {
		bw.Reset(nil)
		writerPool.Put(bw)
	}()

	statusText := http.StatusText(w.statusCode)
	if statusText == "" {
		statusText = "Unknown"
	}

	bw.WriteString("HTTP/1.1 ")
	bw.WriteString(strconv.Itoa(w.statusCode))
	bw.WriteByte(' ')
	bw.WriteString(statusText)
	bw.WriteString("\r\n")

	bw.WriteString("Content-Type: ")
	bw.WriteString(contentType)
	bw.WriteString("\r\n")

	bw.WriteString("Content-Length: ")
	bw.WriteString(strconv.Itoa(w.body.Len()))
	bw.WriteString("\r\n")

	bw.WriteString("Connection: close\r\n")

	for k, values := range w.header {
		switch k {
		case "Content-Type":
			continue
		case "Content-Length":
			continue
		case "Connection":
			continue
		}

		for _, v := range values {
			bw.WriteString(k)
			bw.WriteString(": ")
			bw.WriteString(v)
			bw.WriteString("\r\n")
		}
	}

	bw.WriteString("\r\n")

	if _, err := bw.Write(w.body.Bytes()); err != nil {
		return err
	}

	return bw.Flush()
}
