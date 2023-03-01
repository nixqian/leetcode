package main

import "fmt"

func alloc8191() {
	nums := make([]int, 8191)
	for i := 0; i < 8191; i++ {
		nums[i] = i
	}
}

func alloc8192() {
	nums := make([]int, 8192)
	for i := 0; i < 8192; i++ {
		nums[i] = i
	}
}

func alloc8193() {
	nums := make([]int, 8193)
	for i := 0; i < 8193; i++ {
		nums[i] = i
	}
}
func allocN(n int) {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
}

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func testAccocc(n int) {
	a := make([]int, 0)
	for i := 0; i < n; i++ {
		a = append(a, i)
		fmt.Println(len(a), cap(a))
	}
}

func main() {
	testAccocc(1024)
}
