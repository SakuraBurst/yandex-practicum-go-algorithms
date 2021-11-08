package combinations

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var phoneKeys = map[string][]string{
	"1": nil,
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
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
