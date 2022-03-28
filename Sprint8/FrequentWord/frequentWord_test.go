package frequentWord_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	frequentWord "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/FrequentWord"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestFrequentWord(t *testing.T) {
	tests := []Test{
		{strings.NewReader(`5
caba
aba
caba
abac
aba
`), "aba"},
		{strings.NewReader(`3
b
bc
bcd
`), "b"},
		{strings.NewReader(`10
ciwlaxtnhhrnenw
ciwnvsuni
ciwaxeujmsmvpojqjkxk
ciwnvsuni
ciwnvsuni
ciwuxlkecnofovq
ciwuxlkecnofovq
ciwodramivid
ciwlaxtnhhrnenw
ciwnvsuni`), "ciwnvsuni"},
		{strings.NewReader(`10
vuoikjyxxynwfxc
vuojgssmntswtw
vuohonbwjkjakrc
vuohonbwjkjakrc
vuo
vuodzixbzuha
vuo
vuodzixbzuha
vuoikjyxxynwfxc
vuoikjyxxynwfxc`), "vuoikjyxxynwfxc"},
	}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			frequentWord.FrequentWord(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
