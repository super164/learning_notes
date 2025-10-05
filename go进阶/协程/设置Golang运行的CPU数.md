# 设置Golang运行的cpu数

为了充分利用多cpu的优势，在Golang程序中，设置运行的CPU数目

```go
package main
import (
	"fmt"
	"runtime"
)
func main() {
	cupNum := runtime.NumCPU()
	fmt.Println("cupNum=", cupNum)

	//自己设置使用多少个cpu
	runtime.GOMAXPROCS(cupNum - 1)
	fmt.Println("ok")
}
```

