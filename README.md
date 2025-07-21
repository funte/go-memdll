Windows golang memory module, wrapper of the [MemoryModule][https://github.com/fancycode/MemoryModule/blob/5f83e41].

# Example
* Compile the test dll `cl.exe /LD sum.cpp`.
* Load and test it:
```go
sumDll, _ := NewDLL(data, "sum.dll")
procSum := sumDll.MustFindProc("sum")
r , _, _ := procSum.Call(1, 2) // r should be 3.
```