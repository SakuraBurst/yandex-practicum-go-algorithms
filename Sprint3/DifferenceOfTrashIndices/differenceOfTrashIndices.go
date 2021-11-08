package differenceOfTrashIndices

import (
	"io"
	"bufio"
	"fmt"
)

func DifferenceOfTrashIndices(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	fmt.Println(reader, writer)
}