package main

import "fmt"

func main() {
	slice := [][]int{{10}, {100, 200}}
	slice[0] = append(slice[0], 20)
	fmt.Printf("slice: %v\n", slice)

	a1 := average(90, 93, 98, 92)
	a2 := average(89, 99, 88, 98, 98)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)

	average2("xiaoming", 99, 89, 98, 97, 97)
	average2("xiaohua", 88, 89, 98, 99, 99)

}

func average(scores ...int) int {
	if len(scores) == 0 {
		return 0
	}
	sum := 0
	for _, score := range scores {
		sum += score
	}
	ave := sum / len(scores)
	return ave
}

func average2(name string, scores ...int) {
	if len(scores) == 0 {
		fmt.Printf("name = %s, average score: %d\n", name, 0)
	}
	sum := 0
	for _, score := range scores {
		sum += score
	}
	ave := sum / len(scores)
	fmt.Printf("name = %s, average score: %d\n", name, ave)
}
