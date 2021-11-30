package anagramGrouping_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	anagramGrouping "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint4/AnagramGrouping"
)

type Test struct {
	inputData  io.Reader
	outputData string
}

func TestAnagramGrouping(t *testing.T) {
	tests := []Test{{strings.NewReader(`6
tan eat tea ate nat bat`), `0 4
1 2 3
5
`}, {strings.NewReader(`81
rjvew ivbmo kwfem zeboo pcdkw tfsbz yuwtm wmxjw nrqus nnqpp gcwpw ornbs pthyj yjolz rvnjg ybqwr hwpes mgzas juliq dyhhz njxgb bwzoh tosuk wobdr gnalw cuorl vukpl xkumi mkzph kkgqj uspch jluhu iljup fjuvk krwsd whpdw eixvf wjxmp jbldf fcxpk htcae vomzc fosvw hoplx fbqzl sgrnk bvywp crlmx twzih wlett bbvvc vbaej ljckp ohidr jzxng emspk dsjpq irvzq ysvlv wcity djyqw ehtqq kpvyp khucp bkulg wzbna fqmri nsaqs tohsd qrdxv bwlty bvpyr cxtds cqhbe nyzak mkebf jadyu sbrpk zykgr eqjns gfhob`), `0
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
45
46
47
48
49
50
51
52
53
54
55
56
57
58
59
60
61
62
63
64
65
66
67
68
69
70
71
72
73
74
75
76
77
78
79
80
`}}
	for _, v := range tests {
		buf := bytes.NewBuffer([]byte{})
		anagramGrouping.AnagramGrouping(v.inputData, buf)
		if buf.String() != v.outputData {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", v.outputData, buf.String())
		}
	}
}
