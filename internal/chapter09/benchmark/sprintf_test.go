package benchmark

import (
	"fmt"
	"strconv"
	"testing"
)

// // BenchmarkSprintf 对 fmt.Sprintf 函数进行基准测试
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", number)
	}

}

// BenchmarkFormat 对 strconv.FormatInt 函数进行基准测试
func BenchmarkFormat(b *testing.B) {
	number := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa 对 strconv.Itoa 函数进行基准测试
func BenchmarkItoa(b *testing.B) {
	number := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(number)
	}
}
