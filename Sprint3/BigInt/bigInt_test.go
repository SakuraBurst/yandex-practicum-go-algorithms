package bigInt_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	bigInt "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/BigInt"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBigInt(t *testing.T) {
	tests := []Test{{strings.NewReader(`3
15 56 2
`), "56215"}, {strings.NewReader(`3
1 783 2
`), "78321"}, {strings.NewReader(`5
2 4 5 2 10
`), "542210"}, {strings.NewReader(`6
9 10 1 1 1 6
`), "9611110"}, {strings.NewReader(`38
82 58 66 34 64 37 40 97 93 52 28 98 90 64 19 22 21 83 56 70 46 17 31 51 55 41 68 18 98 89 88 74 6 6 31 36 35 8
`), "9898979390898888382747068666664645856555251464140373635343131282221191817"}, {strings.NewReader(`43
66 49 65 1 79 41 24 40 11 64 91 79 30 19 12 16 69 55 20 21 2 12 67 78 83 97 20 95 72 56 95 2 78 67 51 64 77 20 99 89 11 53 36
`), "99979595918983797978787772696767666564645655535149414036302422212020201916121211111"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		bigInt.BigInt(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
