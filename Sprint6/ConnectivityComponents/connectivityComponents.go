package connectivityComponents

import (
	"io"
	"bufio"
	"fmt"
)

func ConnectivityComponents(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}