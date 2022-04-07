package crib_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	crib "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8FinalEx/Crib"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestCrib(t *testing.T) {
	tests := []Test{
		{strings.NewReader(`examiwillpasstheexam
5
will
pass
the
exam
i
`), "YES"},
		{strings.NewReader(`abacaba
2
abac
caba
`), "NO"},
		{strings.NewReader(`abacaba
3
abac
caba
aba
`), "YES"},
		{strings.NewReader(`sscevscescescscsscevscevscesscsc
4
sce
s
scev
sc
`), "YES"},
		{strings.NewReader(`syuwbawbayuwba
4
wba
yuwb
syu
syuw
`), "YES"},
		{strings.NewReader(`scbscbscba
4
scb
sc
b
ba
`), "YES"},
		{strings.NewReader(`bwvfbtrjqpbwvfbbwvbwbbwbbwvbwvf
4
bwvf
b
bw
bwv
`), "NO"},
		{strings.NewReader(`jnvsrjgdyvgdgdgdgdyvggdggdgdygd
4
g
gdyv
gdy
gd
`), "NO"},
	}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			crib.Crib(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
