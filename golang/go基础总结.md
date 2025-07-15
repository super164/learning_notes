# go语言基础部分的知识梳理

## 目录
- [编写时的注意事项](#编写时的注意事项)
- [变量的创建和使用](#变量的创建和使用)
- [常量的创建和使用](#常量的创建和使用)
- [常用数据类型](#常用数据类型)
- [转义字符](#转义字符)
- [运算符](#运算符)
- [条件控制语句](#条件控制语句)
- [循环语句](#循环语句)
- [函数](#函数)
- [指针](#指针)
- [数组](#数组)
- [切片](#切片) 
- [Map](#Map)
- [结构体](#结构体)

## 编写时的注意事项
    1. 源文件以`go`为扩展名
    2. 程序的执行入口是main()函数
    3. 严格分区大小写
    4. 方法是由一条条语句构成，每个语句后不需要加分号，体现出go语言的简洁性
    5. go编译器是一行行进行编译的，所以不能子啊同一行写多条语句，会报错
    6. 代码中定义的变量或者导的包没有用到，子啊编译的时候会报错
    7. 在代码中括号的使用必须是成对出现的  

## 变量的创建和使用
    在go语言中，变两个创建一共有四种创建方式,如下  
   1. `var a int = 1`，这种方式创建的变量名，类别以及赋值都有
   2. `var a int`,第二种的话只定义不赋值
   3. `var a = 1`,第三种方式，省略了变量的类别，根据后面赋值来确定
   4. `a := 1`,第四种方式更加的简便，只有变量名和值，注意以此方式创建的话，不能省略`:=`  

    创建完变量之后，就可以在程序种使用和改变其的值。
    >对于变量的声明，当变量首字母大写的话，则是全局变量，首字母小写的话则是局部变量
   
## 常用数据类型
在go语言中常用的数据类型，可以理解为整数型，浮点型，布尔型，字符串类型，在整数类型中还可以分为`int8` `int16` `int32` `int64` `int`，整数中可分为`float32` `float64`可以总结为以下的表格：  
|数据类型|占用储存空间|默认值|
|:---:|:---:|:---:|
|int8|1字节|0|
|int16|2字节|0|
|int32|4字节|0|
|int64|8字节|0|
|float32|4字节|0|
|float64|8字节|0|
|bool|1字节|false|
|string|\ |" "|
>字符类型的数据，其本质上就是一个整数，可以直接参与运算，输出字符的时候，直接输出了对应字符的UTF-8编码值，Unicode是其中的一个字符集
如果需要输出字符的话，则采用格式化输出`fmt.Printf("c1对应的具体字符为：%c",c1)`
## 转义字符
简单介绍以下在go语言中比较常用的几个转义字符，可以对数据进行一系列操作，例如下面代码所展示的几个功能：
```go
fmt.Println("aaaa\nbbb")
	//\b 退格将光标前移一位
	fmt.Println("aaaa\bbbb") //aaabbb
	//\r 将光标移至最前边
	fmt.Println("aaaa\rbbb") //bbba
	//\t 补全一个制表符
	fmt.Println("aaaa\tbbb") //aaaa    bbb
	//\" 输出双引号，单引号类似
	fmt.Printf("\"GoLang\"") //"GoLang"
```
## 运算符
        在运算符这部分，比较常用到计算运算符的有`+` `-` `*` `/` `%` `!` `=`分别为加、减、乘、除、取模、取反、等号。
        计算运算符，除此之外还有`++` `--` `+=` `-=` `*=` `/=` `%=`分别为加一，减一，加等,减等，乘等，除等，模等
        以及条件判断运算符有`>` `<` `==` `<=` `>=` `||` `&&` `!=`,分别为大于、小于、等于、小于等于、大于等于、或、且、不等于
>上述我介绍了一些比较常用的运算符，运算符包括但不限于上述介绍，对于运算符的执行先后顺序，我有一个简单粗暴的操作，就是你想让谁先运行就直接给谁加上`()`
## 条件控制语句
1. ### if条件句（单分支）
    1. 在go语言中对于每一个if条件句都要有`{}`，不管代码有多少行都不能省略，这一点和其他的函数有所不同。
    2. 条件表达式左右的()可以不写,也建议不写
    3. if和表达式中间,一定要有空格  
    代码示例如下：
    ```go
    package main
    import "fmt"
    func main(){
	    var num int = 29
	    if num < 30{
		    fmt.Println("对不起，口罩的库存不足")
	    }

	    if num1 := 100;num1 >= 100{
		    fmt.Println("口罩的数量充足")
	    }
    }   
    ```
2. ### if条件句(双分支)
    1. 基本语法
    ```go
    if 条件表达式{  
        逻辑代码1  
    }else {  
        逻辑代码2  
    }
    ```
    2. 当条件表达式成立，即执行逻辑代码1，否则执行逻辑代码2。{}也是必须有的。  
    代码示例如下:
    ```go
    package main
    import "fmt"
    func main(){
	    //定义口罩的数量小于三十，提示库存不足，否则库存充足
	    var num int
	    fmt.Println("请输入口罩的数量：") 
	    fmt.Scanln(&num)
	    if num < 30{
		    fmt.Println("对不起，口罩的库存不足")
	    } else {
		    fmt.Println("口罩库存充足")
	    }
    }
    ```
3. ### if条件句（多分支）
    1. 基本语法  
    ```go
    if 条件表达式{  
        逻辑代码1  
    }else if 条件表达式{  
        逻辑代码2  
    }else {  
        逻辑代码3  
    }
    ```
    2. 当其中if条件句中某个条件符合的时候，将会直接跳出if条件句，不会截止执行下面的代码
    ```go
    package main
    import "fmt"
    func main(){
    	var score int 
    	fmt.Println("请输入分数：") 
    	fmt.Scanln(&score)
    	if score >= 90 {
    		fmt.Println("你的成绩为A级别")
    	} else if score >= 80 {
    		fmt.Println("你的成绩为B级别")
    	}else if score >= 70 {
    		fmt.Println("你的成绩为C级别")
    	}else if score >= 60 {
    		fmt.Println("你的成绩为D级别")
    	}else {
    		fmt.Println("你的成绩为E级别")
    	}
    }
    ```
4. ### Switch 条件语句
    1. 基础语句：  
    ```go
    switch表达式{  
        case 值1,值2,..  
        语句块1  
        case 值3,值4,..  
        语句块2  
        default:  
        语句块  
    }
    ```
    2. Switch语句中，每一个case块后不需要写`break`语句，以及`default`语句的位置可以在任意的地方，也可以不写
    3. switch后面是一个表达式，这个表达式的结果依次跟case进行比较，满足结果的话就执行冒号后面的代码。
>在Switch语句中要实现分支实现穿透可以使用fallthrough关键字，在case 块后增加fallthrough,会继续指向下**一个**case块
## 循环语句
>在go语言中的循环语句只有for循环和for-range循环，没有while循环
1. ### for循环
    1. 基本语句：
    ```go
    for (初始表达式;布尔表达式;迭代因子){
        循环体;
    }
    ```
    >for循环的初始表达式不能使用var来定义变量，可以使用`:=`的方式来定义  
    可将条件初始化放置在for循环之外，迭代式放在循环体中
    ```go
    func main(){
	    var sum int = 0
	    for i := 0;i <= 5;i++{
		    sum += i
	    } 
	    fmt.Println(sum)
    }
    ```
2. ### for-range循环
    1. 基本结构：
    ```go
    for i , value := range str {
	    循环体
    }
    ```
    上述结构中以字符串为例，`i`指的是字符串中每个字符的索引，`value`是`i`所指的字符
    ```go
    func main(){
	    var str string = "hello golang"
	    for i := 0;i < len(str); i++ {
		    fmt.Printf("%c ",str[i])
	    }
	    //使用for range来编写循环
	    for i , value := range str {
		    fmt.Printf("索引是: %d,具体的值为：%c\n",i,value)
	    }
    }
    ```

## 函数
>对于函数部分，其就是将代码中的实现某一个功能的代码给封装起来，独立于main函数之外，如果要使用该功能可以调用该函数。

1. 基础结构：
    ```go
    func 函数名(形参)(返回值类型){
        逻辑代码
        return 返回值
    }
    ``` 
    >其中形参部分**名称在前,类型在后**
2. 对于函数部分，有些函数没有返回值，则返回值类型部分可以省略`return`可以不写
3. >有多个返回值的情况如下
    ```go
    func Sum1 (num1 int,num2 int)(int,int){
	    var numble int = 0
	    numble = num1 + num2
	    result := num1-num2
	    return numble,result
    }
    ```
    >有多个返回值的情况下，若只想得到其中的部分返回值，则可以``使用 `_` 代替``用于忽略那个值

4. 当调用函数的时候，传入的形参，为直接传入值的情况是，在函数里面改变的变量的值，不会改变main函数中的值。如果传入的是指针变量的形参的话，则会改变。
5. **函数可以作为一种数据类型赋值给变量，赋值后的变量可以直接当做该函数使用**
    ```go
    func a (num1 int,){
        fmt.Println(num1)
    }
    func main(){
        c := 1
        b := a
        b(c)
    }
    ```
6. ### 匿名函数使用方式： 
    1. 在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次 
    ```go
    package main
    import "fmt"
    func main(){
    	result := func (num1 int,num2 int) int {
		    return num1 + num2
	    }(2,6)
	    fmt.Println(result)
    }
    ```
7. ### init函数
    1. init函数:初始化函数,可以用来进行一些初始化的操作 每一个源文件都可以包含一个init函数,该函数会在main函数执行前,被Go运行框架调用。 
    2. 全局变量定义, init函数, main函数的执行流程?
    - 全局,init,main
    3. 多个源文件都有init函数的时候,如何执行:
    - 导入包中的先执行，接着main包中的全局变量，init,main函数
8. ### 包
>#### 包就是go语言中函数的所处在的文件，不管是系统函数还是自定义函数，如果想在一个程序中使用这些函数，都要提前导入该包，才能使用其中的函数。
> main包是程序入口包，一般main函数会放在这个包下，要求是main函数一定要放在main包下  
> 包的声明和所在的文件夹的名字可以不一致  
> 一个目录下的同一级的文件必须归属同一个包，即声明的时候为同一个包
9. ### 闭包
1. 闭包的本质： 
    闭包本质依旧是一个匿名函数,只是这个函数引入外界的变量/参数匿名函数+引用的变量/参数=闭包
```go
package main
import "fmt"
func getSum() func (int) int {
	var sum int = 0
	return func (num int) int {
		sum = sum + num
		return sum
	}
}
//闭包：返回的匿名函数+匿名函数以外的变量num
//匿名函数中引用的那个变量会一直保存在内存中,可以一直使用
func main(){
	g := getSum()
	fmt.Println(g(1))
	fmt.Println(g(2))
	fmt.Println(g(3))

}
```
>闭包应用场景：闭包可以保留上次引用的某个值，我们传入一次就可以反复使用了

## 指针
go语言中的指针和C语言的指针类似，指针中就搞明白两个符号就行，`&`指针变量所指向的地址，`*`这个地址所指向的值
```go
package main
import "fmt"

func main(){
	var age int = 20
	//&取当前数据的存放的地址
	fmt.Println(&age)//0xc00000a108
	//定义一个指针变量指向age
	var prt *int = &age
	fmt.Println("prt这个指针指向的那个数据为：",*prt)
}
```
#### 指针变量改变指向值
```go
package main
import "fmt"

func main(){
	var age int =18
	fmt.Println(age)
	//定义一个指针变量指向age,更改值，再次看age的值变不变
	var prt *int = &age
	*prt = 20
	fmt.Println(age)
}
``` 
>指针变量接受的一定是一个地址值  
>指针变量的地址不可以不匹配

## 数组
1. ### 一维数组：
    >数组的var 数组名 [数组大小]数据类型
    - 数组的初始化方式:(四种)
        ```go
        func main(){
        //第一种直接定义数组的类型
        var arr1 int[3] =int[3]{1,2,3}
        //第二种等号前面的数据类型以及大小可以省略
        var arr2 = int[3]{1,2,3}
        //第三种，动态容量的数组，长度可变
        var arr3 = [...]int{1,2,3,4,5,6}
        //第四种，长度可变并且可以指定每个索引对应那个值
        var arr4 = [...]int{2:32,1:44,0:12}
        }
        ```
    - 数组的遍历：
        ```go
        for i := 0;i < len(arr); i++ {
            fmt.Println(arr[i])
        }
        for k,v = range arr {
            fmt.Pritnf("数组中第%d个元素的值为：%d",k.v)
        }
        ```
2. ### 二维数组
    >二维数组的定义格式:var 数组名 [数组大小] [数组大小] 数据类型
    - **数组的初始化方式:(四种):**
        ```go
        func main(){
            //第一种，直接定义
            var arr1 int [2][2] = [2][2]int{{2,2},{2,2}}
            //第二种，可省略int[][]
            var arr2 = [2][2]int{{2,2},{2,2}}
            //第三种，不固定大小的定义动态数组
            var arr3 = [][]{{3,5,8,4},{,1,2,3}}
            //第四种，定义不固定大小的数组，并且每一个索引对应一个值
            var arr4 = [...][]int{{2: 1, 0: 2, 1: 3}, {1: 4, 2: 5, 0: 6}}
        }
        ```
    - **数组的遍历：**
        ```go
        //for循环
        func main(){
            for i := 0;i < len(arr);i++ {
                for j := 0;j < len(arr);j++ {
                    fmt.Print(arr[i][j],"\t")
                }
                fmt.Println()
            }
        }
        //for-range循环
        for k1,v1 = range arr {
            for k2,v2 = range v1 {
                fmt.Printf("arr[%v][%v]=%v\t", i, j, v2)
            }
            fmt.Println()
        }
        ```
## 切片
### 切片的定义：
1. 方式1:定义一个切片,然后让切片去引用一个已经创建好的数组。 
   切片是建立在数组的基础上的
    - 基本结构：var slice []int = 数组[起始索引:结束索引]（**不包含结束索引**）  
    slice := 数组[起始索引:结束索引]
2. 方式2:通过make内置函数来创建切片,基本语法:   
    var 切片名 []type = make([], len,[cap]) 
    ```go
    func main() {
	    //定义切片：make函数的三个参数，1.切片类型2.切片长度3.切片的容量
	    slice := make([]int, 4, 20)
	    fmt.Println(slice)
	    fmt.Printf("切片的长度：%v\n", len(slice))
	    fmt.Printf("切片的容量：%v\n", cap(slice))
	    slice[0] = 43
	    slice[1] = 33
	    fmt.Println(slice)
    }
    ```
    >ps:这样定义的切片，数组是在内部进行创建的，不能直接调用，只能通过slice简介访问  
3. 方式3:定一个切片，直接就指定具体数组，使用原理类似make的方式。
    >`slice2 := []int{1, 4, 7}`

### 切片的遍历：
两种遍历方式：`for`循环实现遍历,`for range`循环实现遍历
```go
package main

import "fmt"

func main() {
	slice := make([]int, 4, 20)
	slice[0] = 66
	slice[1] = 77
	slice[2] = 88
	slice[3] = 99
	//for循环遍历
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v\t", i, slice[i])
	}
	fmt.Println("\n---------------------------------------------")
	//for-range循环遍历
	for i, v := range slice {
		fmt.Printf("slice[%v] = %v\t", i, v)
	}
}

```
4. 切片可以动态增长
    >底层逻辑：
    >1. 底层追加元素的时候对数组进行扩容，老数组扩容为新数组
    >2. 创建一个新数组，将老数组中的数复制到新数组中，在新数组中加上新数
    >3. 如果想要得到追加后的slice，可以直接将结果赋值给原来的切片
    ```go
    func main() {
	    var intarr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	    //定义切片
	    var slice []int = intarr[1:4]
	    fmt.Println(slice)
	    //增长
	    slice1 := append(slice, 34, 23)
	    fmt.Println(slice1)

	    slice = append(slice, 34, 23)
	    fmt.Println(slice)
    }
    ```
    >通过append函数进行追加，也可以将切片追加给切片  
    slice3 := []int{99,11}  
    slice = append(slice,slice3...)//其中的`...`是必写的  

5. 切片的拷贝
    运用copy()函数，a,b两个切片，copy(a,b)这是将b赋值给a
## Map
1. Map的创建：
    1. 方式一:var 变量名 map[keytype]valuetype 
    ```go
    //方式一
	    //定义map变量
	    var a map[int]string
	    //只声明map内存是没有分配空间的
	    //通过make函数进行初始化，才会分配空间
	    a = make(map[int]string, 10) //可存放十个键值对
	    //存入键值对
	    a[2303] = "张三"
	    a[2301] = "李四"
	    a[2302] = "王五"
	    fmt.Println(a)
    ```

    2. 方式二:
    ```go 
    //方式二：
	    b := make(map[int]string)
	    b[2303] = "张三"
	    b[2301] = "李四"
	    fmt.Println(b)
    ```

    3. 方式三
    ```go
    //方式三
	    c := map[int]string{
		    2303: "张三",
		    2301: "李四",
	    }
	    c[2305] = "王五"
	    fmt.Println(c)
    ```
2. 关于Map的一些操作：
    1. ### 增加和更新操作：
        map["key"]= value 如果key还没有，就是增加，如果key存在就是修改。 
    2. ### 删除操作： 
        delete(map, "key"), delete是一个内置函数，如果key存在，就删除该key-value，如果k的y不存在，不操作，但是也不会报错 
    3. ### 清空操作: 
        1. 如果我们要删除map的所有key,没有一个专门的方法一次删除，可以遍历一下key,逐个删除 
        2. 或者map = make(.), make一个新的,让原来的成为垃圾,被gc回收 
    4. ### 查找操作：
        value ,bool = map[key] 
        value为返回的value，bool为是否返回，要么true 要么false

    ```go
    func main() {
	    a := make(map[int]string)
	    //增加操作
	    a[202301] = "张三"
	    a[202302] = "王二麻子"
	    //修改操作
	    a[202302] = "李四"
	    fmt.Println(a)
	    //查找操作
	    value, flag := a[202301]
	    fmt.Println(value)
	    fmt.Println(flag)
	    value1, flag1 := a[202307]
	    fmt.Println(value1)
	    fmt.Println(flag1)
	    //删除操作
	    delete(a, 202301)
	    fmt.Println(a)
    }

    ```

    5. ### map的值为新的map
    >类似于二维的，map里面套map,第一层索引对应了一个新的map
    ```go
    /加深
	    b := make(map[string]map[int]string)
	    b["班1"] = make(map[int]string, 3)
	    b["班1"][202301] = "张三"
	    b["班1"][202302] = "李四"
	    b["班1"][202303] = "王二麻子"

	    b["班2"] = make(map[int]string, 3)
	    b["班2"][202401] = "小小"
	    b["班2"][202402] = "晓晓"
	    b["班2"][202403] = "潇潇"

	    for k1, v1 := range b {
		    fmt.Println(k1)
		    for k2, v2 := range v1 {
			    fmt.Printf("学生学号为：%v，学生的姓名为：%v\n", k2, v2)
		    }
		fmt.Println()
	}
    ```
3. Map的复制：
    ```go
    package main
    func main() {
        // 示例源映射
        a := map[int]string{
            1: "apple",
            2: "banana",
        }
        // 创建一个新的目标映射
        c := make(map[int]string)
        // 遍历源映射，逐个复制键值对
        for k, v := range a {
            c[k] = v
        }
        // 现在c包含a的所有键值对
    }
    ```
    >创建目标映射：使用 make(map[int]string) 创建一个新的映射 c。  
    遍历源映射：通过 for k, v := range a 遍历源映射 a 的所有键值对。  
    复制键值对：在每次循环中，将键值对 k, v 赋值给目标映射 c。
## 结构体
>用变量创建对对象的时候，有如下缺点：  
    1. 不利于数据的管理，维护  
    2. 人物的很多属性都属于一个对象，用变量管理太分散了  
    因此就引出了结构体，通过抽取出每一个对象的共同点，进行一个封装，构成结构体，每次创建对象的时候直接通过结构体创建，这样便于管理

```go
//定义一个结构体
type Teacher struct {
	//变量大写表示外界可以进行访问
	Name   string
	Age    int
	School string
}
```
1. 结构体的创建：
```go
func main() {
	//方法一：
	var t1 Teacher
	t1.Name = "牛鬼蛇蛇"
	t1.Age = 11
	t1.School = "牛马大学"
    //方法二
    var t Teacher = Teacher{"牛鬼蛇蛇",11,"牛马大学"}
    //方法三：返回的是结构体的指针
    var t *Teacher = new(Teacher)
	(*t).Name = "风鹰侠"
	(*t).Age = 18
	t.School = "铠甲大学"
	//为了符合程序员的编程习惯，go提供了简化的赋值方式：

    //方法四：用取地址的符号在创建结构体的时候赋初始值
    var t *Teacher = &Teacher{"牛鬼蛇蛇", 11, "牛马大学"}
}
```
2. 结构体之间的转换：
    1. 结构体是用户单独定义的类型，和其他类型进行转换的时候需要完全相同的字段(名字，个数和类型)
    ```go
    package main
    import "fmt"
    type Student struct {
	    Age int
    }
    type Preson struct {
	    Age int
    }
    func main() {
	    var s Student = Student{10}
	    var p Preson = Preson{10}
	    s = Student(p)
	    fmt.Println(s)
	    fmt.Println(p)
    }
    ```

    2. 结构体进行type重新定义（相当于取别名），Golang认为是新的数据类型，但是相互间可以强转
    ```go
    package main
    import "fmt"
    type Student struct {
	    Age int
    }
    type Stu Student

    func main() {
	    var S1 Student = Student{10}
	    var S2 Stu = Stu{3}
	    fmt.Println(S1)
	    S1 = Student(S2)
	    fmt.Println(S1)
	    fmt.Println(S2)
    }
    ```
