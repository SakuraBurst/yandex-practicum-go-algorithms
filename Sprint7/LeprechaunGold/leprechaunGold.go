package leprechaunGold

import (
	"bufio"
	"fmt"
	"io"
)

func LeprechaunGold(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}