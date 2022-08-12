package bench

import (
	"fmt"
	"testing"
	"bufio"
	"os"
)

var wtr = bufio.NewWriter(os.Stdout)

const (
	FIZZ = 3
	BUZZ = 5
)

func Fizzbuzz(n int) {

	for i := 1; i <= n; i++ {
		if i%FIZZ == 0 && i%BUZZ == 0 {
			fmt.Println("FizzBuzz")

		} else if i%FIZZ == 0 {
			fmt.Println("Fizz")

		} else if i%BUZZ == 0 {
			fmt.Println("Buzz")

		} else {
			fmt.Printf("%d\n", i)
		}
	}
}

func answerFizzbuzz(n int) string {
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

func printWithBuffer(count int) {
	for i := 1; i <= count; i++ {
		fmt.Fprintln(wtr, answerFizzbuzz(i))
	}
	wtr.Flush()
}

func fzbzM (count int) {
	dst := make([]chan string, count)

	for i := 1; i <= count; i++ {
		dst[i-1] = make(chan string)

		go func(src chan string, req int) {
			src <- answerFizzbuzz(req)
		}(dst[i-1], i)
	}
	for i := 0; i < count; i++ {
		//fmt.Fprintln(wtr, <-dst[i])
		fmt.Println(<-dst[i])
	}
	//wtr.Flush()
}

//func TestFizzbuzzBasic(t *testing.T) {
//	Fizzbuzz(100000)
//}

//func TestFizzbuzzBuffer(t *testing.T) {
//	printWithBuffer(15)
//}

func ExampleFzbzM() {
	fzbzM(15)
	// Output:
	// 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// 8
	// Fizz
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
}

func BenchmarkFizzbuzz(b *testing.B) {
	printWithBuffer(b.N)
}

func BenchmarkFzbzM(b *testing.B) {
	fzbzM(b.N)
}
