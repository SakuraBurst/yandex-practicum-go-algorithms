package filterUp

import (
	"io"
	"bufio"
	"fmt"
)

func FilterUp(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}