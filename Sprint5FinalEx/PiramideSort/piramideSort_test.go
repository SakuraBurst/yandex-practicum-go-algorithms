package piramideSort_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	piramideSort "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint5FinalEx/PiramideSort"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestPiramideSort(t *testing.T) {
	tests := []Test{{strings.NewReader(`5
alla 4 100
gena 6 1000
gosha 2 90
rita 2 90
timofey 4 80`), `gena
timofey
alla
gosha
rita
`}, {strings.NewReader(`5
alla 0 0
gena 0 0
gosha 0 0
rita 0 0
timofey 0 0`), `alla
gena
gosha
rita
timofey
`}, {strings.NewReader(`5
alla 0 10
gena 0 9
gosha 0 8
rita 0 7
timofey 0 0
debilniy_comment`), `timofey
rita
gosha
gena
alla
`}, {strings.NewReader(`5
a 0 0
ab 0 0
abc 0 0
abcd 0 0
abcde 0 0
Проверка лексикографической сортировки`), `a
ab
abc
abcd
abcde
`}, {strings.NewReader(`13
tufhdbi 76 58
rqyoazgbmv 59 78
qvgtrlkmyrm 35 27
tgcytmfpj 70 27
xvf 84 19
jzpnpgpcqbsmczrgvsu 30 3
evjphqnevjqakze 92 15
wwzwv 87 8
tfpiqpwmkkduhcupp 1 82
tzamkyqadmybky 5 81
amotrxgba 0 6
easfsifbzkfezn 100 28
kivdiy 70 47
Просто тест`), `easfsifbzkfezn
evjphqnevjqakze
wwzwv
xvf
tufhdbi
tgcytmfpj
kivdiy
rqyoazgbmv
qvgtrlkmyrm
jzpnpgpcqbsmczrgvsu
tzamkyqadmybky
tfpiqpwmkkduhcupp
amotrxgba
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		piramideSort.PiramideSort(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
