package jumpingUpTheStairs_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	jumpingUpTheStairs "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/JumpingUpTheStairs"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestJumpingUpTheStairs(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		jumpingUpTheStairs.JumpingUpTheStairs(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
