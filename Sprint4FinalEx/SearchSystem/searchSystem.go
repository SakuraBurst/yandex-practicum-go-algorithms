package searchSystem

import (
	"io"
	"bufio"
	"fmt"
)

func SearchSystem(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}