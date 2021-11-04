/*
A. Ближайший ноль


Привет) Не знаю где писать, поэтому напишу здесь и в ридми файле, если можно, то напишите где лучше писать)
+ я не знаю в какой форме писать, если вам нужны сухие цифры, то напишите, а то ниже будет решение, + то как я к нему пришел +
мои мысли о нем

В общем id в контесте 54508476. На самом деле я не верил, что это решение пройдет и очень удивился, потому что на мой взгляд
оно наивное, но так-как люди там уже начали второй спринт пару дней назад, а я только первый заканчиваю, и контест сказал, что все ок,
высылаю это решение.

В целом оно неплохое, у меня +600 мс в запасе, с другой стороны всего 20 мегабайт осталось, это не есть хорошо)

Суть алгоритма в том, что мы 2 раза пройдемся по массиву. В первом цикле, мы идем от начала до встречи первого нуля мы ставим -1, чтобы помочь второму циклу.
Когда встрели первый нуль, отсчитываем от него с помощью каунтера. Когда нашли новый нуль - отсчитываем заново и т.д.
В следующем цикле мы уже идем с конца и после встречи первого нуля отсчитываем пока каунтер меньше тех чисел, которые ему попадаются, и это число не -1
Когда каунтер становится равен или больше, останавливаеся и идем дальше до встречи новых нулей. -1 нужен для тех элементов из самого начала массива,
расстояние которых от нуля мы не узнаем за первый цикл. Вроде все объяснил)

Сложность, я считаю, O(n2), потому что мы в любом случае продйемся по массиву 2 раза


*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main11() {
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
