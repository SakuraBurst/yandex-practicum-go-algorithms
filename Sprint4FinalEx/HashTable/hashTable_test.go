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

func TestHashTable(t *testing.T){
	tests := []Test{{strings.NewReader(""), ""}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		hashTable.HashTable(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}