Go build-id extractor
=====================

Build ID is some sequence of bytes which is embedded into the executable file.
Currently, Go supports build ID only on Linux out of the box; MacOS X can be handled
as well (read below how). This package defines GetBuildID function which returns
[]byte embedded into the calling assembly on Linux and MacOS X.

Example
-------
```sh
go get github.com/vmarkovtsev/gobid
```

```go
package main

import (
    "fmt"
	"github.com/vmarkovtsev/gobid"
)

func main() {
  bid, _ := gobid.GetBuildID()
  fmt.Printf("%x\n", bid)
}
```

```sh
# Linux
go build -ldflags "-B 0xdeadbeef" test.go 

# MacOSX
printf "\xde\xad\xbe\xef" > bid.bin
go build -ldflags "-linkmode external -extldflags \"-Wl,-sectcreate,note,build-id,bid.bin\"" test.go 
```

License: MIT