package main

import (
	"fmt"
	"bufio"
	"os"
//	"time"
)

var wtr = bufio.NewWriter(os.Stdout)

const (
	FIZZ = 3
	BUZZ = 5
	N = 15
)

func answerFizzBuzz(n int) string {
	if n%FIZZ == 0 && n%BUZZ == 0 {
		return "FizzBuzz"

	} else if n%FIZZ == 0 {
		return "Fizz"

	} else if n%BUZZ == 0 {
		return "Buzz"

	} else {
		return fmt.Sprintf("%d", n)
	}
}

func concurrentFizzBuzz (count int) {
	ch := make([]chan string, count)

	for i := 1; i <= count; i++ {
		ch[i-1] = make(chan string)

		go func(src chan string, req int) {
			// time.Sleep(100*time.Millisecond)
			// ↑を挟んでも100ms程度でプログラムが終了する
			src <- answerFizzBuzz(req)
		}(ch[i-1], i)
	}
	for i := 0; i < count; i++ {
		fmt.Fprintln(wtr, <-ch[i])
	}
	wtr.Flush()
}

func main() {
	concurrentFizzBuzz(N)
}
