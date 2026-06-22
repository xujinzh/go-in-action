package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}
	fmt.Printf("%d,%d\n", cap(slice), len(slice))
	newSlice := slice[1:3]
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("newSlice: %v\n", newSlice)
	// 如果新append后,切片的数据长度还没有超过原来slice的容量,那么直接在底层数组上进行覆盖
	// 原底层数组将被修改,任何引用底层数组的切片也会发生变化
	newSlice = append(newSlice, 60)
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("newSlice: %v\n", newSlice)

	newSlice2 := slice[1:3]
	// 如果新append后，切片数据长度超过原来slice的容量，那么底层创建全新数组
	// 原底层数组保持不变
	newSlice2 = append(newSlice2, 100, 101, 102, 103, 104, 105)
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("newSlice2: %v\n", newSlice2)
	fmt.Printf("newSlice: %v\n", newSlice)

	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 创建长度是1（=3-2），容量是2（=4-2）的新切片
	newSource := source[2:3:4]
	fmt.Printf("newSource: %v\n", newSource)
	fmt.Printf("%d-%d\n", len(newSource), cap(newSource))
	// 设置容量大于已有容量会发生运行时错误
	// newSource1 := source[2:3:(2+5)]
	// fmt.Printf("newSource1: %v\n", newSource1)

	// 设置长度和容量一样，append时会自动创建新的底层数组，不影响原来的切片和底层数组
	newSource2 := source[2:3:3]
	// 但是直接修改新切片中的值会影响原来的切片和底层数组
	newSource2[0] = "haha"
	fmt.Printf("source: %v\n", source)
	fmt.Printf("newSource2: %v\n", newSource2)
	// append 后就会创建新的数组
	newSource2 = append(newSource2, "Peach")
	fmt.Printf("source: %v\n", source)
	fmt.Printf("newSource2: %v\n", newSource2)

	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, s2...))
}
