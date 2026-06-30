package main

import "fmt"

func add(x, y int) int {
	fmt.Printf("add %v and %v\n", x, y)
	return x + y
}

func main() {
	// defer 虽然整个语句不是先执行，但是里面的函数 add 需要先执行
	// 结果压入栈中，在 main 函数结束后执行 defer
	defer fmt.Println("add two number:", add(1, 2))
	fmt.Println("main done!")
}
