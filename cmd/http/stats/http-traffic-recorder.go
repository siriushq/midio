package stats

import (
	"io"
	"net/http"
)

// IncomingTrafficMeter counts the incoming bytes from the underlying request.Body.
type IncomingTrafficMeter struct {
	io.ReadCloser
	countBytes int
}

// Read calls the underlying Read and counts the transferred bytes.
func (r *IncomingTrafficMeter) Read(p []byte) (n int, err error) {
	n, err = r.ReadCloser.Read(p)
	r.countBytes += n
	return n, err
}

// BytesCount returns the number of transferred bytes
func (r IncomingTrafficMeter) BytesCount() int {
	return r.countBytes
}

// OutgoingTrafficMeter counts the outgoing bytes through the responseWriter.
type OutgoingTrafficMeter struct {
	// wrapper for underlying http.ResponseWriter.
	http.ResponseWriter
	countBytes int
}

// Write calls the underlying write and counts the output bytes
func (w *OutgoingTrafficMeter) Write(p []byte) (n int, err error) {
	n, err = w.ResponseWriter.Write(p)
	w.countBytes += n
	return n, err
}

// Flush calls the underlying Flush.
func (w *OutgoingTrafficMeter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

// BytesCount returns the number of transferred bytes
func (w OutgoingTrafficMeter) BytesCount() int {
	return w.countBytes
}
