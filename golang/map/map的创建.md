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