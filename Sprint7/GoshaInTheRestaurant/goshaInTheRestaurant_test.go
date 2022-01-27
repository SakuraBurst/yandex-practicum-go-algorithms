package goshaInTheRestaurant_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	goshaInTheRestaurant "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/GoshaInTheRestaurant"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestGoshaInTheRestaurant(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		goshaInTheRestaurant.GoshaInTheRestaurant(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
