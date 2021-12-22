package timeToGetOut_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	timeToGetOut "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/TimeToGetOut"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestTimeToGetOut(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		timeToGetOut.TimeToGetOut(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}