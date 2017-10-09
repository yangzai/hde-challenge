package main

import (
	"fmt"
	"os"
	"math/rand"
)
func main() {
	r := rand.New(rand.NewSource(99))
	f, err := os.Create("input-test.txt")
	fmt.Println(err)
	fmt.Fprintf(f, "%v\n", 1000)
	for i := 0; i < 1000; i++ {
		fmt.Fprintf(f, "%v\n", 1000)
		for i := 0; i < 999; i++ {
			fmt.Fprintf(f, "%v ", r.Intn(100))
		}
		fmt.Fprintf(f, "%v\n", r.Intn(100))
	}

}
