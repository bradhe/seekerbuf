package seekerbuf

import (
	"io"
)

type SynchronizedReadWriteSeeker interface {
	io.ReadWriteSeeker
	Sync() error
}
