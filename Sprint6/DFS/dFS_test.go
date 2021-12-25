package dFS_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	dFS "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6/DFS"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestDFS(t *testing.T) {
	tests := []Test{{strings.NewReader(`4 4
3 2
4 3
1 4
1 2
3
`), "3 2 1 4"}, {strings.NewReader(`2 1
1 2
1
`), "1 2"}, {strings.NewReader(`6 7
3 2
5 4
3 1
1 4
1 6
1 2
1 5
1
`), "1 2 3 4 5 6"}, {strings.NewReader(`3 1
2 3
1
`), "1"}}
	for _, v := range tests {
		buf := bytes.NewBuffer(nil)
		dFS.DFS(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}

func BenchmarkDFC(b *testing.B) {

	for i := 0; i < b.N; i++ {
		m := strings.NewReader(`10 13
9 2
1 3
6 10
5 10
9 10
5 9
3 5
5 2
6 7
5 6
8 3
6 1
4 2
7
`)
		buf := bytes.NewBuffer(nil)
		dFS.DFS(m, buf)
	}
}
