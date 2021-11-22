package wierdComparison

import (
	"io"
	"bufio"
	"fmt"
)

func WierdComparison(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}