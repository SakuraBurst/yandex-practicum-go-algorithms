package hashTable

import (
	"io"
	"bufio"
	"fmt"
)

func HashTable(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}