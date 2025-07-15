package main

import "fmt"

func a(num1 int) {
	fmt.Println(num1)
}
func main() {
	// //转义字符
	// //\n 换行
	// fmt.Println("aaaa\nbbb")
	// //\b 退格将光标前移一位
	// fmt.Println("aaaa\bbbb") //aaabbb
	// //\r 将光标移至最前边
	// fmt.Println("aaaa\rbbb") //bbba
	// //\t 补全一个制表符
	// fmt.Println("aaaa\tbbb") //aaaa    bbb
	// //\" 输出双引号，单引号类似
	// fmt.Printf("\"GoLang\"") //"GoLang"

	// //定义一个二维数组
	// var arr1 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(arr1)
	// //
	// var arr2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(arr2)

	// var arr3 = [...][]int{{1, 2, 5, 3}, {4, 5, 6}}
	// fmt.Println(arr3)

	// var arr4 = [...][]int{{2: 1, 0: 2, 1: 3}, {1: 4, 2: 5, 0: 6}}
	// fmt.Println(arr4)

	// var arr = [3][3]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	// for i := 0; i < len(arr); i++ {
	// 	for j := 0; j < len(arr); j++ {
	// 		fmt.Print(arr[i][j])
	// 	}
	// }

	// a := make(map[int]string)
	// c := map[int]string{
	// 	2303: "张三",
	// 	2301: "李四",
	// }
	// c[2305] = "王五"
	// copy(a,c)
	// fmt.Println(a)
	// fmt.Println(c)
}
