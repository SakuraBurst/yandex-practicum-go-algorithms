package exchange

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Price struct {
	day         int
	price       int
	currentPick bool
}

func Exchange(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	nString, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(nString))
	s, _ := reader.ReadString('\n')
	initData := strings.Fields(s)
	priceData := make([]Price, 0, N)
	for i, datum := range initData {
		price, _ := strconv.Atoi(datum)
		priceData = append(priceData, Price{
			day:   i,
			price: price,
		})
	}
	var lowPick Price
	value := 0
	for i := 0; i < N; i++ {
		if !lowPick.currentPick {
			if i+1 < N && priceData[i+1].price < priceData[i].price {
				continue
			}
			lowPick = priceData[i]
			lowPick.currentPick = true
			continue
		}
		if i+1 < N && priceData[i].price < priceData[i+1].price {
			continue
		}
		value += priceData[i].price - lowPick.price
		lowPick.currentPick = false
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(value))
	writer.Flush()
}

func main() {
	Exchange(os.Stdin, os.Stdout)
}
