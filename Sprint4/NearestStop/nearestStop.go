package nearestStop

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Cooridnate struct {
	x int
	y int
}

func NearestStop(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	n, _ := strconv.Atoi(scaner.Text())
	metroExistsCoordinates := make([]Cooridnate, 0, n)
	for i := 0; i < n; i++ {
		scaner.Scan()
		metroCoordinate := strings.Fields(scaner.Text())
		x, _ := strconv.Atoi(metroCoordinate[0])
		y, _ := strconv.Atoi(metroCoordinate[1])
		metroExistsCoordinates = append(metroExistsCoordinates, Cooridnate{x: x, y: y})
	}
	scaner.Scan()
	m, _ := strconv.Atoi(scaner.Text())
	busStopsCoordinates := make([]Cooridnate, 0, m)
	for i := 0; i < m; i++ {
		scaner.Scan()
		busStopCoordinate := strings.Fields(scaner.Text())
		x, _ := strconv.Atoi(busStopCoordinate[0])
		y, _ := strconv.Atoi(busStopCoordinate[1])
		busStopsCoordinates = append(busStopsCoordinates, Cooridnate{x: x, y: y})
	}
	max := 0
	maxIndex := -1
	for i, m := range metroExistsCoordinates {
		counter := 0
		for _, b := range busStopsCoordinates {
			difX := m.x - b.x
			difY := m.y - b.y
			sumOfDif := difX*difX + difY*difY
			if sumOfDif <= 400 {
				counter++
			}
		}
		if counter > max {
			max = counter
			maxIndex = i
		}
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(maxIndex + 1))
	writer.Flush()
}

func main() {
	NearestStop(os.Stdin, os.Stdout)
}
