package conferenceLovers_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	conferenceLovers "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint3/ConferenceLovers"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestConferenceLovers(t *testing.T) {
	tests := []Test{{strings.NewReader(`7
1 2 3 1 2 3 4
3`), "1 2 3"}, {strings.NewReader(`6
1 1 1 2 2 3
1`), "1"}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		conferenceLovers.ConferenceLovers(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
