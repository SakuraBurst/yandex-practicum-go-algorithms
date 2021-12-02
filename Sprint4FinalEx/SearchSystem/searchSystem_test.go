package searchSystem_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	searchSystem "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4FinalEx/SearchSystem"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestSearchSystem(t *testing.T) {
	tests := []Test{{strings.NewReader(`3
i love coffee
coffee with milk and sugar
free tea for everyone
3
i like black coffee without milk
everyone loves new year
mary likes black coffee without milk
`), `1 2
3
2 1
`}, {strings.NewReader(`6
buy flat in moscow
rent flat in moscow
sell flat in moscow
want flat in moscow like crazy
clean flat in moscow on weekends
renovate flat in moscow
1
flat in moscow for crazy weekends
`), `4 5 1 2 3
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		searchSystem.SearchSystem(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
