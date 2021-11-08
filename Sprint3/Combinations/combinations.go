package combinations

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var phoneKeys = map[string][]string{
	"1": nil,
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
	"0": nil,
}

func Combinations(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	nStr, _ := reader.ReadString('\n')
	buttons := strings.Split(nStr, "")
	sliceLength := 1
	for i := 0; i < len(buttons); i++ {
		sliceLength *= len(phoneKeys[buttons[i]])
	}
	fmt.Println(sliceLength)
	result := make([]string, 0, sliceLength)
	combinationsMaker(&result, buttons, "")
	writer := bufio.NewWriter(w)
	writer.WriteString(strings.Join(result, " "))
	writer.Flush()
}

func combinationsMaker(results *[]string, buttons []string, result string) {
	if len(buttons) == 0 {
		*results = append(*results, result)
	} else {
		currentButton := buttons[0]
		buttons = buttons[1:]
		for _, v := range phoneKeys[currentButton] {
			combinationsMaker(results, buttons, result+v)
		}
	}

}
