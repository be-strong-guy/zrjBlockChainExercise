package task2

/**
定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，
创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/
import "fmt"

type Shape interface {
	Area()
	Perimeter()
}
type Rectangle struct {
}

func (r Rectangle) Area() {
	fmt.Println("调用了Rectangle的接口Area（）")
}
func (r Rectangle) Perimeter() {
	fmt.Println("调用了Rectangle的接口Perimeter（）")
}

type Circle struct {
}

func (c Circle) Area() {
	fmt.Println("调用了Circle的接口Area（）")
}
func (c Circle) Perimeter() {
	fmt.Println("调用了Rectangle的接口Perimeter（）")
}
