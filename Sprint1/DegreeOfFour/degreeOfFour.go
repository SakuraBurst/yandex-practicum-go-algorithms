package degreeOfFour

import (
	"fmt"
	"math"
)

func DegreeOfFour() {
	var n float64
	fmt.Scan(&n)
	if n == 1 {
		fmt.Println("True")
		return
	}
	if n < 4 {
		fmt.Println("False")
		return
	}
	possibleAnswer := n / 2
	for i := math.Round(math.Log2(n)); i >= 0; i-- {
		currentResult := math.Pow(4, possibleAnswer)
		// fmt.Println(currentResult)
		if currentResult == n {
			fmt.Println("True")
			return
		}
		if currentResult > n {

			possibleAnswer = math.Round(possibleAnswer / 2)
		}
		if currentResult < n {
			possibleAnswer += math.Round(possibleAnswer / 2)
		}
	}
	fmt.Println("False")
}
