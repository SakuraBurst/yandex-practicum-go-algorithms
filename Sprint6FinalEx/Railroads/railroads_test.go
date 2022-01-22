package railroads_test

import (
	"bytes"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	railroads "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint6FinalEx/Railroads"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestRailroads(t *testing.T) {
	tests := []Test{{strings.NewReader(`3
RB
R
`), "NO"}, {strings.NewReader(`4
BBB
RB
B
`), "YES"}, {strings.NewReader(`5
RRRB
BRR
BR
R
`), "NO"}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			railroads.Railroads(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})
	}
}

func BenchmarkRailroads(b *testing.B) {
	rand.Seed(time.Now().Unix())
	str := makeRailData(5000)

	for i := 0; i < b.N; i++ {

		test := Test{
			inputData:  strings.NewReader(str),
			outputData: "NO",
		}
		buf := bytes.NewBuffer(nil)
		railroads.Railroads(test.inputData, buf)
		//log.Print(buf.String())
	}
}

func makeRailData(length int) string {
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(length))
	builder.WriteByte('\n')
	for i := length - 1; i >= 0; i-- {
		for g := 0; g < i; g++ {
			if i%2 == 0 {
				if g%2 == 0 {
					builder.WriteString("B")
				} else {
					builder.WriteString("B")
				}
			} else {
				if g%2 == 0 {
					builder.WriteString("B")
				} else {
					builder.WriteString("B")
				}
			}

		}
		builder.WriteByte('\n')
	}
	return builder.String()
}
