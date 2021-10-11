package listForm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ListForm() {
	stdReader := bufio.NewReader(os.Stdin)
	stdScaner := bufio.NewScanner(stdReader)
	stdScaner.Scan()
	stdScaner.Scan()
	main := stdScaner.Text()
	stdScaner.Scan()
	g := stdScaner.Text()
	plus, _ := strconv.Atoi(strings.TrimSpace(g))
	mainInt, _ := strconv.Atoi(strings.Join(strings.Split(main, " "), ""))
	str := strconv.Itoa(mainInt + plus)
	fmt.Println(strings.Join(strings.Split(str, ""), " "))
}
