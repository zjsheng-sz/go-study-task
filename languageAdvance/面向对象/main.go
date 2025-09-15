package main

import (
	"fmt"
	"math"
)

func main() {

	//题目一
	rec := Rectangle{
		width:  100.0,
		height: 22.0,
	}

	cir := Circle{
		radius: 11,
	}

	recArea := rec.Area()
	recP := rec.Perimeter()
	fmt.Println("rec area", recArea)
	fmt.Println("rec perimeter", recP)

	cirArea := cir.Area()
	cirP := cir.Perimeter()
	fmt.Println("cir area", cirArea)
	fmt.Println("cir perimeter", cirP)

	//题目二
	emp := Employee{
		Employee: Person{
			Name: "zjs",
			Age:  22,
		},
		EmployeeID: 123,
	}

	emp.PrintInfo()

}

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。*/

type shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (rec Rectangle) Area() float64 {
	return rec.width * rec.height
}

func (rec Rectangle) Perimeter() float64 {
	return 2 * (rec.height + rec.width)
}

type Circle struct {
	radius float64
}

func (cir Circle) Area() float64 {
	return cir.radius * cir.radius * math.Pi
}

func (cir Circle) Perimeter() float64 {
	return 2 * math.Pi * cir.radius
}

/*
使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Employee   Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Println(e)
}
