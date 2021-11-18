package brokenarray

import (
	"fmt"
	"math"
)

func BrokenSearch(arr []int, k int) int {
	// Your code
	// “ヽ(´▽｀)ノ”
	arrStart := 0
	guess := len(arr) >> 1
	// проверяем, что массив сбился, если не сиблся то arrStart остается 0
	if arr[len(arr)-1] < arr[0] {
		counter := 0
		for guess == 0 {
			counter++
			if arr[guess] < arr[len(arr)-1] {
				if arr[guess] < arr[guess-1] {
					arrStart = guess
				} else {
					guess >>= 1
				}
			} else {
				guess += guess >> 1
			}
		}
		fmt.Println(counter)
	}

	guess = len(arr) >> 1
	result := -1
	for i := 0; i < int(math.Log2(float64(len(arr)))); i++ {
		fixedGuess := (guess + arrStart) % len(arr)
		if arr[fixedGuess] == k {
			result = fixedGuess
			break
		}
		if arr[fixedGuess] < k {
			guess += guess >> 1
		} else {
			guess >>= 1
		}

	}

	return result + 1
}
