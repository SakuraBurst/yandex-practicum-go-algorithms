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
	tests := []Test{{strings.NewReader(`5
9 10
9.3 10.3
10 11
10.3 11.3
11 12
`), `3
9 10
10 11
11 12
`}, {strings.NewReader(`3
9 10
11 12.25
12.15 13.3
`), `2
9 10
11 12.25
`}, {strings.NewReader(`7
19 19
7 14
12 14
8 22
22 23
5 21
9 23
`), `3
7 14
19 19
22 23
`}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			schedule.Schedule(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
