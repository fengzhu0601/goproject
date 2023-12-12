package main

import "fmt"

func main() {
	var x *int = nil
	Foo(x)
	var y interface{} = nil
	Foo(y)

	z := interface{}(5)
	Foo(z)
	z = nil
	Foo(z)
}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println(x)
	fmt.Println("non-empty interface")

}

/**
  接口除了有静态类型，还有动态类型和动态值， 当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。 这里的 x 的动态类型是 *int，所以 x 不为 nil
*/
