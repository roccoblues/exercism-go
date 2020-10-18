package circular

import (
	"errors"
)

// Buffer implements a circular buffer of bytes.
type Buffer struct {
	storage  []byte
	readPos  int
	writePos int
	count    int
}

// NewBuffer returns a new buffer of the given size.
func NewBuffer(size int) *Buffer {
	buffer := Buffer{
		storage:  make([]byte, size),
		readPos:  0,
		writePos: 0,
	}

	return &buffer
}

// ReadByte returns the next byte from the buffer.
// Returns an error if the buffer is empty.
func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
		return 0, errors.New("buffer is empty")
	}

	if b.readPos >= len(b.storage) {
		b.readPos = 0
	}

	val := b.storage[b.readPos]
	b.readPos++
	b.count--

	return val, nil
}

// WriteByte writes a byte to the buffer.
// Returns an error if the buffer is full.
func (b *Buffer) WriteByte(c byte) error {
	if b.count == len(b.storage) {
		return errors.New("buffer is full")
	}

	if b.writePos >= len(b.storage) {
		b.writePos = 0
	}

	b.storage[b.writePos] = c
	b.writePos++
	b.count++

	return nil
}

// Overwrite writes a byte to the buffer overwriting the oldest byte if the
// buffer is full.
func (b *Buffer) Overwrite(c byte) {
	if b.writePos >= len(b.storage) {
		b.writePos = 0
	}

	if b.count == len(b.storage) {
		b.readPos++
		b.count--
	}

	b.storage[b.writePos] = c
	b.writePos++
	b.count++
}

// Reset the buffer.
func (b *Buffer) Reset() {
	b.readPos = 0
	b.writePos = 0
	b.count = 0
}
