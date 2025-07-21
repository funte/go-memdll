Windows golang memory module, wrapper of the [MemoryModule][https://github.com/fancycode/MemoryModule/blob/5f83e41].

# Example
* Compile the test dll `cl.exe /LD sum.cpp`.
* Load and test sum.dll:
```go
package main

import (
	_ "embed"
	"fmt"

	"github.com/funte/go-memdll"
)

//go:embed sum.dll
var data []byte

func main() {
	sumDll, _ := memdll.NewDLL(data, "sum.dll")
	procSum := sumDll.MustFindProc("sum")
	r, _, _ := procSum.Call(1, 2) // r should be 3.
	fmt.Println("sum(1, 2) =", int(r))
}

```