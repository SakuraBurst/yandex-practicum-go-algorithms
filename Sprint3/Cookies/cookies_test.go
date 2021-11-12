package cookies_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	cookies "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/Cookies"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestCookies(t *testing.T) {
	tests := []Test{{strings.NewReader(`2
1 2
3
2 1 3
`), "2"}, {strings.NewReader(`3
2 1 3
2
1 1
`), "1"}, {strings.NewReader(`10
8 5 5 8 6 9 8 2 4 7
8
9 8 1 1 1 5 10 8
`), "5"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		cookies.Cookies(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
