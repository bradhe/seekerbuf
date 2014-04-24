package seekerbuf

import (
	"testing"
)

func TestWritesSeeksAndReads(t *testing.T) {
	original := []byte("Hello, World!")

	buf := NewReadWriteSeekerBuffer(0)
	buf.Write(original)
	buf.Seek(0, 0)

	bytes := make([]byte, len(original))
	buf.Read(bytes)

	for i := range(original) {
		if original[i] != bytes[i] {
			t.Logf(`Expected "%v" to equal "%v"`, bytes, original)
			t.Fail()
			return
		}
	}
}

func TestReturnsPartialReads(t *testing.T) {
	original := []byte("Hello, World!")

	buf := NewReadWriteSeekerBuffer(0)
	buf.Write(original)
	buf.Seek(0, 0)

	bytes := make([]byte, len(original) * 10)
	n, _ := buf.Read(bytes)

	if n != len(original) {
		t.Logf("Expected to read %d bytes, read %d", len(original), n)
		t.Fail()
	}
}

func TestReturnsWriteBytes(t *testing.T) {
	original := []byte("Hello, World!")

	buf := NewReadWriteSeekerBuffer(0)
	n, _ := buf.Write(original)

	if n != len(original) {
		t.Logf("Expected to read %d bytes, read %d", len(original), n)
		t.Fail()
	}
}
