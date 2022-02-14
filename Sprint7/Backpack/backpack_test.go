package backpack_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	backpack "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint7/Backpack"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestBackpack(t *testing.T) {
	tests := []Test{{strings.NewReader(`4 6
2 7
4 2
1 5
2 1
`), `3
4 3 1`}, {strings.NewReader(`92 724
30 52
30 10
25 41
24 69
25 82
7 5
9 21
32 18
17 39
3 48
14 87
22 15
29 82
31 66
25 67
18 18
6 84
29 71
7 94
7 82
34 67
6 73
12 67
11 90
13 67
23 85
3 68
18 25
34 3
3 65
20 55
6 95
18 56
21 21
11 91
14 58
33 90
21 65
14 43
6 59
5 83
7 35
19 73
26 95
30 60
24 43
8 27
4 92
1 38
14 86
5 58
18 25
30 10
28 30
10 65
25 33
15 77
2 58
4 32
22 2
4 37
28 89
28 56
28 47
23 5
21 13
23 95
28 75
3 94
30 59
15 72
33 2
10 25
4 43
15 49
34 9
18 60
5 19
14 2
17 69
21 90
17 61
1 7
26 71
2 36
8 40
8 38
25 23
8 45
14 35
9 39
16 6
`), `58
91 89 87 86 85 84 83 82 81 80 78 77 75 74 73 71 69 68 67 62 61 59 58 57 55 51 50 49 48 47 44 43 42 41 40 39 38 37 36 35 33 32 30 27 26 25 24 23 22 20 19 17 15 13 11 10 5 4`}}
	for _, v := range tests {
		t.Run(v.outputData, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			backpack.Backpack(v.inputData, buf)
			if buf.String() != v.outputData {
				t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
			}
		})

	}
}
