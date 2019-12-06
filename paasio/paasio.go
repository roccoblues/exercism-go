package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	reader    io.Reader
	rmu       sync.RWMutex
	bytesRead int64
	readOpts  int
}

type writeCounter struct {
	writer       io.Writer
	wmu          sync.RWMutex
	bytesWritten int64
	writeOpts    int
}

type readWriteCounter struct {
	readWriter io.ReadWriter
	readCounter
	writeCounter
}

// NewReadCounter returns a new ReadCounter
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{reader: r}
}

func (r *readCounter) Read(p []byte) (int, error) {
	r.rmu.Lock()
	defer r.rmu.Unlock()

	r.readOpts++
	read, err := r.reader.Read(p)
	r.bytesRead += int64(read)

	return read, err
}

func (r *readCounter) ReadCount() (int64, int) {
	r.rmu.RLock()
	defer r.rmu.RUnlock()

	return r.bytesRead, r.readOpts
}

// NewWriteCounter returns a new WriteCounter
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{writer: w}
}

func (w *writeCounter) Write(p []byte) (int, error) {
	w.wmu.Lock()
	defer w.wmu.Unlock()

	w.writeOpts++
	written, err := w.writer.Write(p)
	w.bytesWritten += int64(written)

	return written, err
}

func (w *writeCounter) WriteCount() (int64, int) {
	w.wmu.RLock()
	defer w.wmu.RUnlock()

	return w.bytesWritten, w.writeOpts
}

// NewReadWriteCounter returns a new ReadWriteCounter
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	rwc := &readWriteCounter{}
	rwc.reader = rw
	rwc.writer = rw

	return rwc
}
