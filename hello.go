package main

import (
	"fmt"
)

const Pi = 3.14

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("add function ", add(10, 11))

	a, b := swap("hello", "world")
	fmt.Println(a, b)      //world hello
	fmt.Println(split(17)) //7 10

	/*sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)*/

	/*sum := 0
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)*/

}
