```

$ cd cgo/cgo
$ make
$ mv cgo libcgo.so

$ cd cgo/main
$ LD_LIBRARY_PATH=../cgo/  go run .
$ go build
$ LD_LIBRARY_PATH=../cgo/  ./main

// https://golang.google.cn/cmd/cgo/
These two call malloc to make copies free them after:
C.CString, C.CBytes,

```
