package brokenarray

import (
	"math"
)

func BrokenSearch(arr []int, k int) int {
	// главная проверка. Если искомое число больше, чем последний элемент массива, то, условно, мы будем искать в правой части, если нет - в левой
	isKHigerThenLastElement := k > arr[len(arr)-1]
	result := -1
	start := 0
	end := len(arr)
	// если результат не найден за логарифмическое вермя, значит его нет и result остается -1
	// почему-то ингода он не успевает найти, поэтому я добавли + 1 шаг, с чем связано без понятия
	logArrLenght := int(math.Log2(float64(len(arr))))
	for i := 0; i <= logArrLenght; i++ {
		// часто замечал в исходниках го, что там деление на 2 производится с помощью шифта вправо
		// может это бысрее чем обычное деление на 2? Я не знаю :D Просто копирую)
		middle := (start + end) >> 1
		if arr[middle] == k {
			result = middle
			break
		}
		if arr[middle] > k {
			if isKHigerThenLastElement {
				end = middle
			} else {
				if arr[middle] > arr[len(arr)-1] {
					start = middle
				} else {
					end = middle
				}
			}
		} else {
			if isKHigerThenLastElement {
				if arr[middle] < arr[len(arr)-1] {
					end = middle
				} else {
					start = middle
				}
			} else {
				start = middle
			}
		}
	}

	return result
}

/*
для примера возьмем массив
[7,8,9,10,1,2,3,4,5,6]

Если выбранное число > k && k > последний элемент массива {
	искомое число не может быть правее, потому что самое большое число справа меньше чем k
	то тогда идем левее
	пример:
	k = 9
	guess = 10
	мы теперь ищем в [7,8,9]
}

Если выбранное число > k && k < последний элемент массива {
	плохой исход, так как если k меньше поледнего элемента, то числа больше могут быть как слева, так и справа
	придется делать проверку
	Если выбранное число > последний элемент массива {
		ищем в правой части
		пример:
		k = 5
		guess = 10
		мы теперь ищем в [1,2,3,4,5,6]
	} В ином случае {
		ищем в левой части
		пример:
		k = 5
		guess 6
		мы теперь ищем в [7,8,9,10,1,2,3,4,5]
	}
}

Если выбранное число < k && k > последний элемент массива {
	плохой исход, так как если k больше поледнего элемента, то числа меньше могут быть как слева, так и справа
	плохой исход, придется делать две проверки
	Если выбранное число < последний элемент массива {
		ищем в левой части
		пример:
		k = 10
		guess = 2
		мы теперь ищем в [7,8,9,10,1]
	} В ином случае {
		ищем в правой части
		пример:
		k = 10
		guess 7
		мы теперь ищем в [8,9,10,1,2,3,4,5,6]
	}

	то тогда мы ищем правее
	пример:
	k = 9
	guess = 8
	мы теперь ищем в [9,10,1,2,3,4,5,6]
}



Если выбранное число < k && k < последний элемент массива {
	если k < последний элемент массива, то числа меньше k могут быть только справа
	то тогда ищем слева
	пример k = 4
	guess = 3
	мы теперь ищем в [,4,5,6]
}



*/

func BrokenSearch2(arr []int, k int) int {
	// Your code
	// “ヽ(´▽｀)ノ”
	arrStart := 0
	if len(arr) == 1 {
		if arr[0] == k {
			return 0
		} else {
			return -1
		}
	}
	// проверяем на то, что начало массива в самом конце потому что это O(1) и мы можем сэкономить время
	if arr[len(arr)-1] < arr[len(arr)-2] {
		arrStart = len(arr) - 1
		// потом проверяем, что массив сбился, если не сиблся то arrStart остается 0
	} else if arr[len(arr)-1] < arr[0] {
		start := 0
		end := len(arr)
		for arrStart == 0 {
			guess := (start + end) >> 1
			if arr[guess] < arr[len(arr)-1] {
				if arr[guess] < arr[guess-1] {
					arrStart = guess
				} else {
					end = guess
				}
			} else {
				start = guess
			}
		}
	}

	// fmt.Println("nachalo massiva nahoditsya po indexsu = ", arrStart, "i ravno = ", arr[arrStart])
	result := -1
	start := 0
	end := len(arr)
	// если результат не найден за логарифмическое вермя, значит его нет
	// почему-то ингода он не успевает найти, поэтому я добавли + 1 шаг, с чем связано без понятия
	logArrLenght := int(math.Log2(float64(len(arr))))
	for i := 0; i <= logArrLenght; i++ {
		guess := (start + end) >> 1
		fixedGuess := ((guess + arrStart) % len(arr))
		if arr[fixedGuess] == k {
			result = fixedGuess
			break
		}
		if arr[fixedGuess] < k {
			start = guess
		} else {
			end = guess
		}

	}

	return result
}
