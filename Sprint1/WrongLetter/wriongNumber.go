package wrongLetter

import (
	"bufio"
	"fmt"
	"os"
)

func WrongLetter() {
	stdReader := bufio.NewReader(os.Stdin)
	stdScanner := bufio.NewScanner(stdReader)
	stdScanner.Scan()
	firstStr := stdScanner.Text()
	stdScanner.Scan()
	secondStr := stdScanner.Text()
	str1Map := map[rune]int{}
	for _, v := range firstStr {
		str1Map[v]++
	}
	for _, v := range secondStr {
		_, ok := str1Map[v]
		if !ok {
			fmt.Println(string(v))
			return
		}
		str1Map[v]--
		if str1Map[v] < 0 {
			fmt.Println(string(v))
			return
		}
	}
}

// func sortTwoStrings(s1, s2 string){
// 	for i := 0; i < len(s2); i++ {

// 	}
// }
