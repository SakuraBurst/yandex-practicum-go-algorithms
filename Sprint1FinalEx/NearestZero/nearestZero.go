package nearestzero

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func NearestZero() {
	stdReader := bufio.NewReader(os.Stdin)
	nString, _ := stdReader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nString))
	housesString, _ := stdReader.ReadString('\n')
	houses := strings.Split(strings.TrimSpace(housesString), " ")
	isInCountFromZero := false
	awayFromZeroCounter := 0
	for i := 0; i < n; i++ {
		if houses[i] == "0" {
			isInCountFromZero = true
			awayFromZeroCounter = 0
		}
		if isInCountFromZero {
			count := strconv.Itoa(awayFromZeroCounter)
			houses[i] = count
			awayFromZeroCounter++
			continue
		}
		houses[i] = "-1"
	}
	isInCountFromZero = false
	awayFromZeroCounter = 1
	for i := len(houses) - 1; i >= 0; i-- {
		if houses[i] == "0" {
			isInCountFromZero = true
			awayFromZeroCounter = 1
			continue
		}
		if isInCountFromZero {
			count := strconv.Itoa(awayFromZeroCounter)
			currentNumber, _ := strconv.Atoi(houses[i])
			if awayFromZeroCounter >= currentNumber && currentNumber != -1 {
				isInCountFromZero = false
				awayFromZeroCounter = 1
				continue
			}
			houses[i] = count
			awayFromZeroCounter++
		}
	}
	stdWiter := bufio.NewWriter(os.Stdout)
	stdWiter.WriteString(strings.Join(houses, " "))
	stdWiter.Flush()
}
