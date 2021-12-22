package bipartitionCheck

import (
	"io"
	"bufio"
	"fmt"
)

func BipartitionCheck(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}