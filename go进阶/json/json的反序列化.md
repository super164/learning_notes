# `json`反序列化

## 基本概念：

`json`反序列化是指，将`json`字符串反序列化成对应的数据类型（如结构体、map、切片）的操作



## 将`json`字符串分别反序列为结构体，map，切片

1. 结构体：

```go
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func unmarshalStruct() {
	//真正的是在项目中，是通过网络传输获取到的
	str := "{\"Name\":\"牛魔王\",\"Age\":1000,\"Birthday\":\"23-6-30\",\"Sal\":10000,\"Skill\":\"蛮牛冲撞\"}"
	//定义Monster
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("反序列化失败,err=%v", err)
	}
	fmt.Println("反序列化后monster= ", monster)
}
```

2. map

   ```go
   func unmarshalMap() {
   	str := "{\"address\":\"火云洞\",\"age\":19,\"name\":\"红孩子\"}"
   	var m map[string]interface{}
   	//反序列化map不需要make,底层会自动的make
   	//m = make(map[string]interface{})
   	err := json.Unmarshal([]byte(str), &m)
   	if err != nil {
   		fmt.Printf("反序列化失败,err=%v", err)
   	}
   	fmt.Println("反序列化后m= ", m)
   }
   ```

3. 切片

   ```go
   func unmarshalSlice() {
   	str := "[{\"address\":[\"北京\",\"上海\"],\"age\":19,\"name\":\"tom\"}," +
   		"{\"address\":[\"天津\",\"廊坊\"],\"age\":18,\"name\":\"mary\"}]\n"
   	var slice []map[string]interface{}
   
   	err := json.Unmarshal([]byte(str), &slice)
   	if err != nil {
   		fmt.Printf("反序列化失败,err=%v", err)
   	}
   	fmt.Println("反序列化后slice= ", slice)
   }
   ```

   

细节：

1. 在反序列化一个`json`字符串时，要确保反序列化后的数据类型和原来序列化数据类型一致
2. 如果`json`字符串是通过程序获取到的，则不需要在对双引号进行转义。