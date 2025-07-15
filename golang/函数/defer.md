- 在Golang中，程序遇到defer关键字，不会立即执行defer后的语句，而是即将defer后的语句
压入一个栈中，继续执行函数后的语句
```go
package main
import "fmt"

func main(){
	fmt.Println(add(30,60))
}
func add(num1 int,num2 int) int {
	defer fmt.Println("num1=",num1)
	defer fmt.Println("num2=",num2)

	var sum int = num1 + num2
	fmt.Println("sum=",sum)
	return sum
}
```
- 在defer语句之后的代码部分在压入栈的时候，附带的其变量当前的状态也随之入栈
- defer应用场景： 
    - 比如你想关闭某个使用的资源，在使用的时候直接随手defer，因为defer有延迟执行机制（函数执行完毕再执行defer压入栈的语句）所以你用完随手写了关闭，比较省心，省事 