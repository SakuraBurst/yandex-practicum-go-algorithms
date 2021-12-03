package hashTable

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type KeyValue struct {
	isExist bool
	key     int
	value   int
}

type HashMap []KeyValue

func (hm HashMap) hash(key int) int {
	return key % len(hm)
}

func (hm HashMap) Get(key int) KeyValue {
	potentialKey := hm.hash(key)
	for i := 0; i < len(hm); i++ {
		if hm[potentialKey].isExist && hm[potentialKey].key != key {
			potentialKey++
			if potentialKey == len(hm) {
				potentialKey = 0
			}
		} else {
			break
		}

	}
	return hm[potentialKey]
}

func (hm HashMap) Put(kv KeyValue) {
	potentialKey := hm.hash(kv.key)
	for i := 0; i < len(hm); i++ {
		if hm[potentialKey].isExist && hm[potentialKey].key != kv.key {
			potentialKey++
			if potentialKey == len(hm) {
				potentialKey = 0
			}
		} else {
			break
		}

	}
	hm[potentialKey] = kv
}

func (hm HashMap) Delete(key int) (int, error) {
	if hm[hm.hash(key)].isExist {
		defer func() {
			hm[hm.hash(key)] = KeyValue{}
		}()

		return hm[hm.hash(key)].value, nil
	} else {
		return 0, errors.New("no such key")
	}

}

func HashTable(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	nCommands, _ := strconv.Atoi(scaner.Text())
	primeForHash := getPrimeForN(nCommands)
	writer := bufio.NewWriter(w)
	hm := make(HashMap, primeForHash)
	for i := 0; i < nCommands; i++ {
		scaner.Scan()
		command := strings.Fields(scaner.Text())
		switch {
		case command[0] == "get":
			getKey, _ := strconv.Atoi(command[1])
			val := hm.Get(getKey)
			if val.isExist {
				writer.WriteString(strconv.Itoa(val.value))
			} else {
				writer.WriteString("None")
			}
			writer.WriteByte('\n')
		case command[0] == "put":
			putKey, _ := strconv.Atoi(command[1])
			putValue, _ := strconv.Atoi(command[2])
			hm.Put(KeyValue{isExist: true, key: putKey, value: putValue})
		case command[0] == "delete":
			deleteKey, _ := strconv.Atoi(command[1])
			val, err := hm.Delete(deleteKey)
			if err != nil {
				writer.WriteString("None")

			} else {
				writer.WriteString(strconv.Itoa(val))
			}
			writer.WriteByte('\n')
		}
	}
	//
	writer.Flush()
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
	case n <= 1000000:
		return 1000003
	}
	return 3
}

func main() {
	HashTable(os.Stdin, os.Stdout)
}
