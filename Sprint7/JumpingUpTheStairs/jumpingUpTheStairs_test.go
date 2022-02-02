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
	tests := []Test{
		{strings.NewReader("6 3"), "13"},
		{strings.NewReader("7 7"), "32"},
		{strings.NewReader("2 2"), "1"},
		{strings.NewReader("62 44"), "535806680"},
	}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			jumpingUpTheStairs.JumpingUpTheStairs(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
