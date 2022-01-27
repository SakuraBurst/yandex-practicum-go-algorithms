package fibonacciNumbersForAdults_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	fibonacciNumbersForAdults "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/FibonacciNumbersForAdults"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestFibonacciNumbersForAdults(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		fibonacciNumbersForAdults.FibonacciNumbersForAdults(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
