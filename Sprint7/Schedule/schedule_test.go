package schedule_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	schedule "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/Schedule"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestSchedule(t *testing.T) {
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		schedule.Schedule(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
