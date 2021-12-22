package topologicalSort

import (
	"io"
	"bufio"
	"fmt"
)

func TopologicalSort(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}