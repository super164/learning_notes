# 读文件的操作应用实例

1. 读取文件的内容并显示在终端（带缓冲区的方式），使用`os.Open、file.Close、bufio.NewReader()、reader.ReadString`函数和方法

   代码如下

   ```go
   package main
   import (
   	"bufio"
   	"fmt"
   	"io"
   	"os"
   )
   func main() {
   	//打开文件
   	//file 1,file对象 2,file指针 3,文件句柄
   	file, err := os.Open("e:/test.txt")
   	if err != nil {
   		fmt.Println("open file err=", err)
   	}
   	//函数退出时及时的关闭file
   	defer file.Close()
   
   	//创建一个*Reader，带缓冲
   	reader := bufio.NewReader(file)
   	for {
   		str, err := reader.ReadString('\n')
   		if err == io.EOF {
   			break
   		}
   		fmt.Print(str)
   	}
   	fmt.Println("文件读取结束")
   }
   ```

2. 读取文件内容并显示在终端（使用`os`词义将整个文件读入内存中），这种方式适用于文件不大的情况，相关方法和函数（`os.ReadFile`）

   ```go
   package main
   import (
   	"fmt"
   	"os"
   )
   func main() {
   	file := "e:/test.txt"
   	content, err := os.ReadFile(file)
   	if err != nil {
   		fmt.Printf("read file err = %v", err)
   
   	}
   	fmt.Printf("%v", string(content))
   }
   ```
   
   