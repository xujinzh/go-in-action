// 因为不是总能获取一个值的地址，所以值的方法集只包括了使用值接收者实现的方法。

/*
Values                     Methods Receivers
-----------------------------------------------
T                          (t T)
*T                         (t T) and (t *T)

Methods Receivers           Values
-----------------------------------------------
(t T)                       T and *T
(t *T)                      *T
*/
package main

import "fmt"

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {

	d := duration(42)
	result := (&d).pretty()
	fmt.Printf("result: %v\n", result)
}
