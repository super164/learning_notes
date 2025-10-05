# JSON的序列化

**json的序列化是指，将有key-value结构的数据类型（比如结构体，map,切片）序列化成json字符串的操作**

实现结构体、map、和切片三种数据的序列化，如下



1. 结构体的序列化:

   ```go
   func testStruct() {
   	monster := Monster{
   		Name:     "牛魔王",
   		Age:      1000,
   		Birthday: "23-6-30",
   		Sal:      10000.0,
   		Skill:    "蛮牛冲撞",
   	}
   	//将monster序列化
   	data, err := json.Marshal(&monster)
   	if err != nil {
   		fmt.Printf("序列化失败，err=%v", err)
   	}
   	//输出序列化后的结果
   	fmt.Printf("monster序列化后=%v", string(data))
   }
   ```

   

2. map的序列化:

   ```go
   func testMap() {
   	var m map[string]interface{}
   	m = make(map[string]interface{})
   	m["name"] = "红孩子"
   	m["age"] = 19
   	m["address"] = "火云洞"
   	//序列化
   	data, err := json.Marshal(m)
   	if err != nil {
   		fmt.Printf("序列化失败，err=%v\n", err)
   	}
   	//输出序列化后的结果
   	fmt.Printf("m序列化后=%v\n", string(data))
   }
   ```

   

3. 切片的序列化：

   ```go
   func testSlice() {
   	var slice []map[string]interface{}
   	var m1 map[string]interface{}
   	m1 = make(map[string]interface{})
   	m1["name"] = "tom"
   	m1["age"] = 19
   	m1["address"] = [2]string{"北京", "上海"}
   
   	slice = append(slice, m1)
   	var m2 map[string]interface{}
   	m2 = make(map[string]interface{})
   	m2["name"] = "mary"
   	m2["age"] = 18
   	m2["address"] = [2]string{"天津", "廊坊"}
   	slice = append(slice, m2)
   
   	data, err := json.Marshal(slice)
   	if err != nil {
   		fmt.Printf("序列化失败,err=%v", err)
   	}
   	//输出序列化后的结果
   	fmt.Printf("slice序列化后=%v\n", string(data))
   }
   ```

对基本数据类型进行序列化没有什么实际的意义



## 对结构体进行tag操作

```go
type Monster struct {
	Name     string `json:"monster_name"`//反射机制
	Age      int	`json:"monster_age"`
	Birthday string
	Sal      float64
	Skill    string
}
func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      1000,
		Birthday: "23-6-30",
		Sal:      10000.0,
		Skill:    "蛮牛冲撞",
	}
	//将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化失败，err=%v", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v", string(data))
}
```

tag操作类似于对于结构体中数据序列化之后，取个别名，想要展示的名字，上述代码展示序列化之后，牛魔王的键为`monster_name`

对于结构体的序列化,如果我们希望序列化后的key的名字,又我们自己重新制定,那么可以给struct指定一个tag 标签
