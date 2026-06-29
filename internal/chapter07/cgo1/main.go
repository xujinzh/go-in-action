package main

/*
#include <stdio.h>
#include <stdlib.h> // 使用 free() 函数必须引入这个头文件

void sayHello(const char *name, int age) {
	printf("Hello %s from C! You are %d years old.\n", name, age);
}
*/
import "C"
import "unsafe" // 使用 unsafe.Pointer 是否内存

func main() {
	goName := "Alice"
	goAge := 25

	// 1. 将 Go 字符串转为 C 字符串（C 会在堆区分配内存）
	cName := C.CString(goName)

	// 2. 必须使用 defer 在函数退出时释放 C 字符串内存
	defer C.free(unsafe.Pointer(cName))

	// 3. 将 Go 的 int 强转为 C 的 int
	cAge := C.int(goAge)

	// 4. 调用 C 函数并传参
	C.sayHello(cName, cAge)
}
