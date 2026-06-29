package main

/*

#include <stdio.h>

void sayHello() {
    printf("Hello from C!\n");
}
*/
import "C" // 和上面的注释要紧贴，不能有空行

/*
# method 1
go run main.go

# method 2
go run .

# method 3
go build -o myprogram main.go
./myprogram
*/

func main() {
	C.sayHello() // 直接调用C函数
}
