package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // スライスは可変長。配列と違い長さを指定する必要がない。
	// prims[1:4]は最初の要素を含むが、最後の要素は含まない。半開区間
	fmt.Println(s)
}
