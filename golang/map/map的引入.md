1. 映射(map) , Go语言中内置的一种类型,它将键值对相关联,我们可以通过键key来获取对应的值value,类似其它语言的集合   
>键值对：一对匹配的信息 例如：
>学生学号- 学生姓名 
>20095452 - 赵珊珊 I 
>20095459 - 张三 

2. 基本语法: var map 变量名 map[keytype]valuetype 
    1. PS: key、value的类型： bool、数字、string、指针、channel、还可以是只包含前面几个类型的接口、结构体、数组
    2. PS： key通常为int、string类型，value通常为数字（整数、浮点数）、string、map、结构体
    3. PS:  对于key部分slice、 map、 function不可以
3. map的特点：
    1. map集合在使用前一定要make
    2. map的key-value,是按键的大小来排列的
    3. key键不会重复的，如果重复的话，后者会替换前一个value
    4. make函数的第二个参数size 可以省略，默认就分配一个内存