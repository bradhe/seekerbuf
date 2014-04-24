# A ReadWriteSeeker Buffer

This package mimics the functionality of ReadWriteSeeker in memory. This is
useful for testing systems that would otherwise write to a file.

## Usage

Pretty straight forward. I haven't actually tried building this but...you'll
get it.

```go
package main

import (
  "fmt"
  "github.com/bradhe/seekerbuf"
)

func main() {
  buf := seekerbuf.NewReadWriteSeekerBuffer(1024)
  buf.Write([]byte("This is only a test, do not worry!"))

  arr := make([]byte, len("This is only a test, do not worry!"))
  buf.Seek(0, 0)
  buf.Read(arr)

  fmt.Printf("%v", buf)
}
```
