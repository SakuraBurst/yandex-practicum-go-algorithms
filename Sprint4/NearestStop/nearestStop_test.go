package nearestStop_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	nearestStop "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/NearestStop"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestNearestStop(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		nearestStop.NearestStop(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}