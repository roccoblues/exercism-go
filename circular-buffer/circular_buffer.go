package circular

import (
	"errors"
)

// Buffer implements a circular buffer of bytes.
type Buffer struct {
	storage  []byte
	size     int
	readPos  int
	writePos int
	count    int
}

// NewBuffer returns a new buffer of the given size.
func NewBuffer(size int) *Buffer {
	return &Buffer{
		storage:  make([]byte, size),
		size:     size,
		readPos:  0,
		writePos: 0,
		count:    0,
	}
}

// ReadByte returns the next byte from the buffer.
// Returns an error if the buffer is empty.
func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
		return 0, errors.New("buffer is empty")
	}

	val := b.storage[b.readPos]
	b.advanceReader()
	b.count--

	return val, nil
}

// WriteByte writes a byte to the buffer.
// Returns an error if the buffer is full.
func (b *Buffer) WriteByte(c byte) error {
	if b.count == b.size {
		return errors.New("buffer is full")
	}

	b.storage[b.writePos] = c
	b.advanceWriter()
	b.count++

	return nil
}

// Overwrite writes a byte to the buffer overwriting the oldest byte if the
// buffer is full.
func (b *Buffer) Overwrite(c byte) {
	b.storage[b.writePos] = c
	b.advanceWriter()

	if b.count == b.size {
		b.advanceReader()
	} else {
		b.count++
	}
}

// Reset the buffer.
func (b *Buffer) Reset() {
	b.readPos = 0
	b.writePos = 0
	b.count = 0
}

func (b *Buffer) advanceReader() {
	b.readPos++
	if b.readPos == b.size {
		b.readPos = 0
	}
}

func (b *Buffer) advanceWriter() {
	b.writePos++
	if b.writePos == b.size {
		b.writePos = 0
	}
}
