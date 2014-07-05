package seekerbuf

import (
	"io"
)

type SynchronizedReadWriteSeeker interface {
	io.ReadWriteSeeker
	io.Closer
	Sync() error
}
