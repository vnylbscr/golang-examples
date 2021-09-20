package main

import (
	"fmt"
	"math/rand"
)

var items [5]int

func main() {

	// for index := 0; index < len(items); index++ {
	// 	items[index] = rand.Int()
	// }
	rand.Seed(100)
	assingToArray(items[:])
	fmt.Println(items)
}

func assingToArray(item []int) {
	for i := 0; i < len(item); i++ {
		item[i] = rand.Intn(20)
	}
}
