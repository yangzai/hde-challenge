package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

// tail call recursion, read from channel instead of array
func sumSquaredPositives(c <-chan int, acc int) int {
	val, ok := <-c

	if !ok { //channel closed
		return acc
	}

	if val > 0 {
		acc += val * val
	}

	return sumSquaredPositives(c, acc)
}

func fillOuts(outs []chan int) {
	outs[0] = make(chan int, 1)

	if len(outs) > 1 {
		fillOuts(outs[1:])
	}
}

func processTestCases(scanner *bufio.Scanner, outs []chan int) {
	scanner.Scan()
	l, _ := strconv.Atoi(scanner.Text())

	in := make(chan int, l)

	go func(in <-chan int, out chan int) {
		defer close(out)
		out <- sumSquaredPositives(in, 0)
	}(in, outs[0])

	streamLine(scanner, in, l)
	close(in)
	in = nil //gc

	if len(outs) > 1 {
		processTestCases(scanner, outs[1:])
	}
}

func streamLine(scanner *bufio.Scanner, in chan int, l int) {
	if l == 0 {
		return
	}

	scanner.Scan()
	elem, _ := strconv.Atoi(scanner.Text())
	in <- elem

	streamLine(scanner, in, l - 1)
}

// print output in order
func processResults(outs []chan int, done chan struct{}) {
	if done != nil {
		defer close(done)
	}

	fmt.Println(<-outs[0])
	outs[0] = nil //gc

	if len(outs) > 1 {
		processResults(outs[1:], nil)
	}
}

func main() {
	f, err := os.Open("input-test.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(f)
	//scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// set up output pipeline
	outs := make([]chan int, n)
	fillOuts(outs)
	done := make(chan struct{})
	go processResults(outs, done)

	processTestCases(scanner, outs)

	<-done
}
