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

func TestCookies(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		cookies.Cookies(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}