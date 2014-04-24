package seekerbuf

import (
	"errors"
	"fmt"
	"io"
)

type bufferImpl struct {
	data []byte
	offset int64
}

func len64(arr []byte) int64 {
	return int64(len(arr))
}

// Reads a len(p) bytes out of the buffer from it's current position. Returns
// the number of bytes read, as well as an error (potentially). If a partial
// read occurs, n will be less than len(p) and p may not be filled fully.
func (self *bufferImpl) Read(p []byte) (n int, err error) {
	end := self.offset + len64(p)

	// Are we about to do a partial read?
	if end > len64(self.data) {
		end = len64(self.data)
	}

	copy(p, self.data[self.offset:end])

	// if it's read past the end, we want n < len(p), otherwise n = len(p) as
	// len(p) bytes were read.
	if len64(self.data) < (self.offset + len64(p)) {
		n = len(self.data) - int(self.offset)
		self.offset = len64(self.data)
	} else {
		n = len(self.data)
		self.offset = len64(p)
	}

	return
}

// Writes len(p) bytes to the buffer. If there isn't enough space left
// internally for the bytes in p, the buffer will expand. Returns the number of
// bytes written (always len(p)) unless an error occurs.
func (self *bufferImpl) Write(p []byte) (n int, err error) {
	if (self.offset + len64(p)) > len64(self.data) {
		// grow it up!
		cp := make([]byte, len(self.data) + len(p))
		copy(cp, self.data)
		self.data = cp
	}

	copy(self.data[self.offset:], p)
	n = len(p)
	self.offset += len64(p)

	return
}

// Seeks resets the offset internal to the buffer. whence can be set to 0
// indicating offset is relative to the beginning of the buffer, 1 indicating
// offset is relative to the current position in the buffer, or 2 indicating
// offset is relative to the end of the file.
// 
// This buffer does not support setting offset beyond the end of the file. An
// attempt to put the current position beyond the end of the file will set the
// current position to the end of the file.
//
// If whence is not in 0, 1, or 2 an error will be returned.
func (self *bufferImpl) Seek(offset int64, whence int) (int64, error) {
	if whence == 0 {
		self.offset = offset
	} else if whence == 1 {
		self.offset += offset
	} else if whence == 2 {
		self.offset = len64(self.data) - offset

		// In case they passed in negative offset.
		if self.offset > len64(self.data) {
			self.offset = len64(self.data)
		}
	} else {
		return -1, errors.New(fmt.Sprintf("invalid whence: %d", whence))
	}

	return self.offset, nil
}

// Allocates a new ReadWriteSeekerBuffer with the supplied capacity.
func NewReadWriteSeekerBuffer(capacity int) io.ReadWriteSeeker {
	buf := new(bufferImpl)
	buf.data = make([]byte, capacity)
	return buf
}
