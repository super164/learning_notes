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
```go
//加深
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