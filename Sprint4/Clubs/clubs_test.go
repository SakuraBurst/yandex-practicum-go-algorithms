package clubs_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	clubs "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/Clubs"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestClubs(t *testing.T) {
	tests := []Test{{strings.NewReader(`8
вышивание крестиком
рисование мелками на парте
настольный керлинг
настольный керлинг
кухня африканского племени ужасмай
тяжелая атлетика
таракановедение
таракановедение`), `вышивание крестиком
рисование мелками на парте
настольный керлинг
кухня африканского племени ужасмай
тяжелая атлетика
таракановедение
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		clubs.Clubs(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
