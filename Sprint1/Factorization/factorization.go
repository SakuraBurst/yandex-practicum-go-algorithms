package factorization

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Factorization() {
	var n int
	fmt.Scan(&n)
	g := n
	allEvenNumbers := allSimpleNumbersFrom2ToSqrtN(n)
	result := make([]int, 0, 10)
	for i := 0; i < len(allEvenNumbers); i++ {
		if g == 1 {
			break
		}
		for (g % allEvenNumbers[i]) == 0 {
			result = append(result, allEvenNumbers[i])
			g /= allEvenNumbers[i]
		}
	}
	if len(result) == 0 {
		fmt.Println(n)
		return
	} else {
		if g != 1 {
			result = append(result, g)
		}
		stdWriter := bufio.NewWriter(os.Stdout)
		for _, v := range result {
			fmt.Fprintf(stdWriter, "%d ", v)
		}
		stdWriter.Flush()
	}
}

func allSimpleNumbersFrom2ToSqrtN(n int) []int {
	sq := int(math.Round(math.Sqrt(float64(n))))
	allNubers := make([]int, sq+1)
	simpleNubers := make([]int, 0, sq+1)
	for i := 2; i <= sq; i++ {
		if allNubers[i] == 0 {
			simpleNubers = append(simpleNubers, i)
			for g := i * i; g <= sq; g += i {
				allNubers[g] = -1
			}
		}
	}
	return simpleNubers
}

func factorization(g int, pullSimpleNumbers []int) []int {
	result := make([]int, 0, 10)
	for i := len(pullSimpleNumbers) - 1; i >= 0; {
		if (g % pullSimpleNumbers[i]) == 0 {
			result = append(result, pullSimpleNumbers[i])
			g /= pullSimpleNumbers[i]
			if g == 1 {
				break
			}
		} else {
			i--
		}
	}
	return result
}

func allSimpleNumbersFromNToDived2N(n int) []int {
	sq := n / 2
	allNubers := make([]int, sq+1)
	simpleNubers := make([]int, 0, sq+1)
	for i := 2; i <= sq; i++ {
		if allNubers[i] == 0 {
			simpleNubers = append(simpleNubers, i)
			for g := i * i; g <= sq; g += i {
				allNubers[g] = -1
			}
		}
	}
	return simpleNubers
}

func AllSimpleNumbersFromN(n int) []int {
	allNubers := make([]int, n+1)
	simpleNubers := make([]int, 0, n+1)
	for i := 2; i < n; i++ {
		if allNubers[i] == 0 {
			allNubers[i] = i
			simpleNubers = append(simpleNubers, i)
		}
		for _, v := range simpleNubers {
			z := v * i
			if v > allNubers[i] || z > n {
				break
			}
			allNubers[z] = v
		}
	}
	return simpleNubers
}

func allSimpleNumbersFromNOld(n int) []int {
	allNubers := make([]int, n)
	simpleNubers := make([]int, 0, n)
	for i := 2; i < n; i++ {
		if allNubers[i] == 0 {
			simpleNubers = append(simpleNubers, i)
			for g := i * i; g < n; g += i {
				allNubers[g] = -1
			}
		}
	}
	return simpleNubers
}

// allNubers := make([]int, n)
// for i := 2; i < n; {
// 	if g == 1 {
// 		break
// 	}
// 	if allNubers[i] == 0 {
// 		allNubers[i] = 1
// 		for g := i * i; g < n; g += i {
// 			allNubers[g] = -1
// 		}
// 		if (g % i) == 0 {
// 			result = append(result, i)
// 			g /= i
// 		} else {
// 			i++
// 		}
// 	}
// 	if allNubers[i] == 1 {
// 		if (g % i) == 0 {
// 			result = append(result, i)
// 			g /= i
// 		} else {
// 			i++
// 		}
// 	}
// 	if allNubers[i] == -1 {
// 		i++
// 	}
// }

// result := make([]int, 0, 10)
// 	for i := len(pullSimpleNumbers) - 1; i >= 0; {
// 		if (g % pullSimpleNumbers[i]) == 0 {
// 			result = append(result, pullSimpleNumbers[i])
// 			g /= pullSimpleNumbers[i]
// 			if g == 1 {
// 				break
// 			}
// 		} else {
// 			i--
// 		}
// 	}
// for i := 0; i < len(pullSimpleNumbers); {
// 	if (g % pullSimpleNumbers[i]) == 0 {
// 		result = append(result, pullSimpleNumbers[i])
// 		g /= pullSimpleNumbers[i]
// 		if g == 1 {
// 			break
// 		}
// 	} else {
// 		i++
// 	}
// }

// var pullSimpleNumbers = AllSimpleNumbersFromN(n)
// fmt.Println("done", len(pullSimpleNumbers))
// if n <= 100 {
// 	pullSimpleNumbers = allSimpleNumbersFromNToDived2N(n)
// } else {
// 	pullSimpleNumbers = allSimpleNumbersFrom2ToSqrtN(n)
// }
// fmt.Println(pullSimpleNumbers)
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"
// )

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	g := n
// 	var allSimpleNumbers []int
// 	if n <= 100 {
// 		allSimpleNumbers = allSimpleNumbersFromNToDived2N(n)
// 	} else {
// 		allSimpleNumbers = allSimpleNumbersFrom2ToSqrtN(n)
// 	}
// 	result := make([]int, 0, 10)
// 	for i := len(allSimpleNumbers) - 1; i >= 0; {
// 		if (g % allSimpleNumbers[i]) == 0 {
// 			result = append(result, allSimpleNumbers[i])
// 			g /= allSimpleNumbers[i]
// 			if g == 1 {
// 				break
// 			}
// 		} else {
// 			i--
// 		}
// 	}
// 	if len(result) == 0 {
// 		fmt.Println(g)
// 		return
// 	} else {
// 		stdWriter := bufio.NewWriter(os.Stdout)
// 		reverse(result)
// 		for _, v := range result {
// 			fmt.Fprintf(stdWriter, "%d ", v)
// 		}
// 		stdWriter.Flush()
// 	}
// }

// func reverse(v []int) {
// 	for i, g := 0, len(v)-1; i < g; i, g = i+1, g-1 {
// 		v[i], v[g] = v[g], v[i]
// 	}
// }

// func allSimpleNumbersFrom2ToSqrtN(n int) []int {
// 	sq := int(math.Round(math.Sqrt(float64(n))))
// 	allNubers := make([]int, sq+1)
// 	simpleNubers := make([]int, 0, sq+1)
// 	for i := 2; i <= sq; i++ {
// 		if allNubers[i] == 0 {
// 			simpleNubers = append(simpleNubers, i)
// 			for g := i * i; g <= sq; g += i {
// 				allNubers[g] = -1
// 			}
// 		}
// 	}
// 	return simpleNubers
// }

// func allSimpleNumbersFromNToDived2N(n int) []int {
// 	sq := n / 2
// 	allNubers := make([]int, sq+1)
// 	simpleNubers := make([]int, 0, sq+1)
// 	for i := 2; i <= sq; i++ {
// 		if allNubers[i] == 0 {
// 			simpleNubers = append(simpleNubers, i)
// 			for g := i * i; g <= sq; g += i {
// 				allNubers[g] = -1
// 			}
// 		}
// 	}
// 	return simpleNubers
// }
// allNumbers := make([]int, n+1)
// primes := make([]int, 0, n+1)
// for i := 2; i < n; i++ {
// 	if g == 1 {
// 		break
// 	}
// 	if allNumbers[i] == 0 {
// 		allNumbers[i] = i
// 		primes = append(primes, i)
// 		for g%i == 0 {
// 			g /= i
// 			result = append(result, i)
// 		}
// 	}
// 	for _, v := range primes {
// 		z := v * i
// 		if v > allNumbers[i] || z > n {
// 			break
// 		}
// 		allNumbers[z] = v
// 	}
// }
