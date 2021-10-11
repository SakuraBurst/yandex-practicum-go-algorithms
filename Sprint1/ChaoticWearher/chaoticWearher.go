package chaoticWeather

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChaoticWeather() {
	var n int
	fmt.Scan(&n)
	stdReader := bufio.NewReader(os.Stdin)
	// после скана остается байт переноса строки, я его просто считываю
	stdReader.ReadByte()
	line, _ := stdReader.ReadString('\n')
	mainWeatherInfo := strings.Split(strings.TrimSpace(line), " ")
	stdWriter := bufio.NewWriter(os.Stdout)
	if len(mainWeatherInfo) == 1 {
		stdWriter.WriteString("1")
		stdWriter.Flush()
		return
	}
	var counter = 0
	for i := 0; i < n; i++ {
		switch {
		case i == 0:
			if isHiger(mainWeatherInfo[i], mainWeatherInfo[i+1]) {
				counter++
				i++
			} else {
				continue
			}
		case i == (n - 1):
			if isHiger(mainWeatherInfo[i], mainWeatherInfo[i-1]) {
				counter++
				i++
			} else {
				continue
			}
		default:
			if isHiger(mainWeatherInfo[i], mainWeatherInfo[i-1]) && isHiger(mainWeatherInfo[i], mainWeatherInfo[i+1]) {
				counter++
				i++
			} else {
				continue
			}
		}
	}
	fmt.Fprint(stdWriter, counter)
	stdWriter.Flush()
}

func isHiger(i, j string) bool {
	m, _ := strconv.Atoi(i)
	l, _ := strconv.Atoi(j)
	return m > l
}
