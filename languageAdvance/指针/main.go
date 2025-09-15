package main

import "fmt"

/*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/

func addTen(num *int) {

	*num += 10

}

/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/

func doubleSlice(ptr *[]int) {

	for i := 0; i < len(*ptr); i++ {
		(*ptr)[i] = (*ptr)[i] * 2
	}
}


func main() {
	
	
	n := 100
	addTen(&n)
	fmt.Println(n)


	nums := []int{1,2,3,4,5}
	doubleSlice(&nums)
	fmt.Println(nums)

}
