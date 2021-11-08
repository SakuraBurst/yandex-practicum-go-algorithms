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

func TestConferenceLovers(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		conferenceLovers.ConferenceLovers(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}