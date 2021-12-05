package returnRange

import (
	"io"
	"bufio"
	"fmt"
)

func ReturnRange(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}