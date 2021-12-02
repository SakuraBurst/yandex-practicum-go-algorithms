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
`}, {strings.NewReader(`10
ing ioizzlux knjklikxxq
lntw qinx iltxhmrmfp
kmsjgmfyu gic xivg pu
nlbudxrvnv vrbs spdeuoqr
eyew sosx iecoujab xtokvidtwf
spdeuoqr qpcmtqzs ioizzlux
jeqoseryf cvoafqw hnovlrgzp
uanxvs gic ljymhdwgw
spdeuoqr pu xivg mjotkrq
rg uqblrnrgwr pu aeiv
5
uqblrnrgwr
rg cvoafqw
vrbs xivg pu
kqweuyzgkt
rg gic xivg
`), `10 
7 10 
3 9 4 10 

3 8 9 10 
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		searchSystem.SearchSystem(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
