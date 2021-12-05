package filterDown

import (
	"io"
	"bufio"
	"fmt"
)

func FilterDown(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}