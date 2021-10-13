package sleightofhand

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SleightOfHand() {
	var fingers int
	fmt.Scan(&fingers)
	stdReader := bufio.NewReader(os.Stdin)
	keyMap := map[int]int{}
	for i := 0; i <= 4; i++ {
		s, _ := stdReader.ReadString('\n')
		s = strings.TrimSpace(s)
		for _, v := range s {
			if v == '.' {
				continue
			}
			vNum, _ := strconv.Atoi(string(v))
			keyMap[vNum]++
		}
	}
	score := 0
	for i := 1; i < 10; i++ {
		count, ok := keyMap[i]
		if ok {
			if count <= (fingers * 2) {
				score++
			}
		}
	}
	fmt.Println(score)
}
