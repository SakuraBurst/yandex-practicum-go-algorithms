package functionvalue

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func FunctionValue() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	var x int64
	var ax2 int64
	var bx int64
	var c int64
Loop:
	for i := 0; scanner.Scan(); i++ {
		// мне супер не нравится этот метод, но у меня нет никакой друго инфы о том как взять числа для функции
		// кроме того, что а идет первым, х вторым, b третим и c четвертым
		currentInt, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		switch i {
		case 0:
			ax2 += currentInt
		case 1:
			x = currentInt
			ax2 *= x * x
		case 2:
			bx = currentInt * x
		case 3:
			c = currentInt
			break Loop
		}

	}
	fmt.Fprint(writer, ax2+bx+c)
	writer.Flush()
}
