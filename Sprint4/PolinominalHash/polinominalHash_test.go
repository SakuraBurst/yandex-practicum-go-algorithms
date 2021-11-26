package polinominalHash_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	polinominalHash "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/PolinominalHash"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestPolinominalHash(t *testing.T) {
	tests := []Test{{strings.NewReader(`123
100003
a`), "97"}, {strings.NewReader(`1000
1000009
abcdefgh`), "6080"}, {strings.NewReader(`123
100003
HaSH`), "56156"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		polinominalHash.PolinominalHash(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}

// func TestBreakPolinominalHash(t *testing.T) { 70 883
// 	tests := []Test{{strings.NewReader(`1000
// 123987123
// ezhgeljkablzwnvuwqvo`), "97"}, {strings.NewReader(`1000
// 123987123
// gbpdcvkumyfxillgnqru`), "97"}}
// 	results := make([]string, 0, 2)
// 	for _, v := range tests {
// 		buf := bytes.NewBuffer([]byte{})
// 		polinominalHash.PolinominalHash(v.inputData, buf)
// 		results = append(results, buf.String())
// 	}
// 	if results[0] == results[1] {
// 		t.Errorf("\nодинаковые:\n%s\ngot:\n%s", results[0], results[1])
// 	} else {
// 		t.Errorf("\nразные:\n%s\ngot:\n%s", results[0], results[1])
// 	}
// }

// ezh

//
