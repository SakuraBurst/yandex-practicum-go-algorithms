package hashTable

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type KeyValue struct {
	key   int
	value int
}

type HashMap []KeyValue

func (hm HashMap) hash(key int) int {
	return 1
}

func HashTable(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	nCommands, _ := strconv.Atoi(scaner.Text())
	primeForHash := getPrimeForN(nCommands)

	fmt.Println(primeForHash)
	for i := 0; i < nCommands; i++ {
		scaner.Scan()
		command := strings.Fields(scaner.Text())
		switch {
		case command[0] == "get":
			getKey, _ := strconv.Atoi(command[1])
			fmt.Println("get ", getKey)
		case command[0] == "put":
			putKey, _ := strconv.Atoi(command[1])
			putValue, _ := strconv.Atoi(command[2])
			fmt.Println("put ", putKey, putValue)
		case command[0] == "delete":
			deleteKey, _ := strconv.Atoi(command[1])
			fmt.Println("delete ", deleteKey)
		}
	}
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}

func getPrimeForN(n int) int {
	switch {
	case n < 10:
		return 13
	case n < 100:
		return 107
	case n < 1000:
		return 1009
	case n < 10000:
		return 10007
	case n < 100000:
		return 100003
	case n < 1000000:
		return 1000003
	}
	return 3
}
