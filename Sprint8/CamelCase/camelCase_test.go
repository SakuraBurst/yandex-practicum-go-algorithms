package camelCase_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	camelCase "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint8/CamelCase"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestCamelCase(t *testing.T) {
	tests := []Test{{strings.NewReader(`3
MamaMilaRamu
MamaMia
MonAmi
2
MM
MA
`), `MamaMia
MamaMilaRamu
MonAmi
`}, {strings.NewReader(`2
AlphaBetaGgamma
AbcdBcdGggg
2
ABGG
ABG
`), `AbcdBcdGggg
AlphaBetaGgamma
`}, {strings.NewReader(`5
WudHnagkbhfwrbci
WCUkvoxboxufsdap
jdrxomezzrpuhbgi
ZcGHdrPplfoldemu
cylbtqwuxhiveznc
3
WGHV
NKVDT
ZGHU

`), ``}, {strings.NewReader(`10
bFHdRdJaXN
MhqzvipuzL
DPSVGjUfuy
EpayAivGPR
Mqlmchycpc
JKLDUIbhzh
AHyiFHMMZQ
EJuoAvxGPR
hDqCNojsGl
DsPSVGUApr
10
RWOVOVAU
JKLD
DPSAGUAA
ABOBF
D
DJTLGKRMV
TVEGDUMZ
EAGPRWSCX
UTCNGVHY
DPSVGUAAJ

`), `JKLDUIbhzh
DPSVGjUfuy
DsPSVGUApr
hDqCNojsGl
`}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			camelCase.CamelCase(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
