package fizzbuzz

import (
	"fmt"
	"io"
	"os"
)

const (
	fizzNumber    = 3
	buzzNumber    = 5
	fizzValue     = "Fizz"
	buzzValue     = "Buzz"
	fizzBuzzValue = "FizzBuzz"
)

var fizzBuzzNumber = fizzNumber * buzzNumber

type FizzBuzzer interface {
	List() []interface{}
	Print(io.Writer)
}

type FizzBuzzStruct struct {
	n int
}

func NewFizzBuzzStruct(n int) FizzBuzzStruct {
	return FizzBuzzStruct{n}
}

func (fb FizzBuzzStruct) List() []interface{} {
	list := make([]interface{}, fb.n)
	for i := 1; i <= fb.n; i++ {
		var dat interface{}
		if i%fizzBuzzNumber == 0 {
			dat = fizzBuzzValue
		} else if i%buzzNumber == 0 {
			dat = buzzValue
		} else if i%fizzNumber == 0 {
			dat = fizzValue
		} else {
			dat = i
		}
		list[i-1] = dat
	}
	return list
}

func (fb FizzBuzzStruct) Print(w io.Writer) {
	for _, val := range fb.List() {
		fmt.Fprintln(w, val)
	}
}

func PrintFizzBuzz(n int) {
	NewFizzBuzzStruct(n).Print(os.Stdout)
}

func FizzBuzz(n int) []interface{} {
	return NewFizzBuzzStruct(n).List()
}
