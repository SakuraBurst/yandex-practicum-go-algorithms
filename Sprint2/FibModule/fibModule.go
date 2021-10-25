package fibmodule

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var SQRT_5 = math.Sqrt(5)

var GOLDEN_RATIO = (1 + SQRT_5) / 2

var INVERSE_GOLDEN_RATIO = (1 - SQRT_5) / 2

func FibModule() {
	stdReader := bufio.NewReader(os.Stdin)
	str, _ := stdReader.ReadString('\n')
	data := strings.Fields(str)
	n, _ := strconv.Atoi(data[0])
	k, _ := strconv.Atoi(data[1])
	mod := int(math.Pow(10, float64(k)))
	result := nFibForTest(n+1, mod)
	fmt.Println(result)
}

func nFibFor(n int) *big.Int {
	results := [...]*big.Int{big.NewInt(1), big.NewInt(1)}
	for i := 0; i < n-2; i++ {
		results[0], results[1] = results[1], results[0].Add(results[0], results[1])
	}
	return results[1]
}

func nFibForTest(n, limit int) int {
	results := [...]int{1, 1}
	for i := 0; i < n-2; i++ {
		results[0], results[1] = results[1], (results[0]+results[1])%limit
	}
	return results[1]
}

var BIG_SQRT_5 = big.NewFloat(SQRT_5)
var BIG_GOLDEN_RATIO = big.NewFloat(GOLDEN_RATIO)
var BIG_INVERSE_GOLDEN_RATIO = big.NewFloat(INVERSE_GOLDEN_RATIO)

// func BineMethod(n int) *big.Int {
// 	var BIG_N = big.NewInt(n)
// 	var RESULT = big.NewFloat(0)
// 	var BIG_SQRT_5 = big.NewFloat(SQRT_5)
// 	var BIG_GOLDEN_RATIO = big.NewFloat(GOLDEN_RATIO)
// 	var BIG_INVERSE_GOLDEN_RATIO = big.NewFloat(INVERSE_GOLDEN_RATIO)
// 	return BIG_N.Exp()
// }

func BineMethod(n int) *big.Float {
	return big.NewFloat((math.Pow(GOLDEN_RATIO, float64(n)) - math.Pow(INVERSE_GOLDEN_RATIO, float64(n))) / SQRT_5)
}
