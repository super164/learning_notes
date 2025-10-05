# tcp-服务器端接受客户端信息

![image-20250807160654069](tcp-服务器监听.assets/image-20250807160654069.png)

## 服务器端代码

```go
package main
import (
	"fmt"
	"io"
	"net"
)
func process(conn net.Conn) {
	//关闭连接
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("close err = %v\n", err)
		}
	}(conn)

	for {
		//每一次都要创建一个新的切片
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//1.等待客户端通过conn
		//2.如果客户端没有write[发送]，那么协程就阻塞在这里
		fmt.Println("服务器等待客户端发送信息", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		//if err != nil {
		//	fmt.Println("服务器端的Read err=", err)
		//	return
		//}
		if err == io.EOF {
			fmt.Println("服务器端的Read err=", err)
			return
		}
		//3.显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}
func main() {
	fmt.Println("服务器开始监听.....")
	//使用tcp协议
	//表示在本地监听，8888端口
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			fmt.Printf("close err =%v\n", err)
		}
	}(listen)

	for {
		//等待客户端来连接
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//在这里起协程，为客服端服务
		go process(conn)
	}
	fmt.Printf("listen suc=%v", listen)
}
```



![image-20250807161157103](tcp-服务器监听.assets/image-20250807161157103.png)

# 客户端代码

```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Println("conn suc=", conn)

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		//从终端读取一行用户输入，发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		line = strings.Trim(line, "\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}
		//将line 发送给服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}

	}
}
```

