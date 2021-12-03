package hashTable_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	hashTable "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4FinalEx/HashTable"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestHashTable(t *testing.T) {
	tests := []Test{{strings.NewReader(`10
get 1
put 1 10
put 2 4
get 1
get 2
delete 2
get 2
put 1 5
get 1
delete 2
`), `None
10
4
4
None
5
None
`}, {strings.NewReader(`8
get 9
delete 9
put 9 1
get 9
put 9 2
get 9
put 9 3
get 9
`), `None
None
1
2
3
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		hashTable.HashTable(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
