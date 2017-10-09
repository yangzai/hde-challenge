package main

import (
	"fmt"
	"os"
)

var f *os.File


func sqr(x int) int {
	var temp int
	temp = 0
	if (x > 0) {
		temp = x * x
	} else {
		temp = 0
	}
	return temp
}

func calc(x int) int {
	var ret int
	if (x > 0){
		var n int
		fmt.Fscanf(f, "%d",&n)
		ret =  sqr(n) + calc(x-1)
	} else {
		ret = 0
	}
	return ret
}

func exec(x int) int {
	if (x > 0){
		var n int
		fmt.Fscanf(f, "%d",&n);
		fmt.Println(calc(n));
		exec(x - 1);
	}
	return 0
}

func main() {
	var err error
	f, err = os.Open("input-test.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var t int
	fmt.Fscanf(f,"%d",&t)
	exec(t)
}