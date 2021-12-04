/*
-- ПРИНЦИП РАБОТЫ --
id в контесте 59995888

Привет! Тут, судя по контесту я справился гораздо лучше, чем в первой задаче, оставив в запасе еще 4 секунды) Но решение не лишено недостатков

Суть проста, есть тип HashMap который под собой имеет простой слайс, и реализует нужные методы. Я очень долго думал, как же выбирать размер исконного слайса
потому что в теории было написано просто "это должно быть простое число, которое выбирается исходя из ситуации", минут 20 поискав в гугле (может есть какой-то
алгоритм выбора размера хеш мапы, кто знает), я так и не нашел ничего, потому что, я так понял, рехеширование с увеличением размера исконного массива
это стандарт, поэтому никто не заморачивается с выбором конкретнго размера. Я заниматься рехешированием из-за ограниченного времени не хотел, поэтому
просто выбрал простое число которое больше, чем количество запросов к хеш мапе (функция getPrimeForN) (наверное это читерство, я не знаю, ведь если ты реализуешь реальную хеш мапу
тебе никто не напишет сколько будет конкретно запросов к ней).

Дальше создали хеш мапу с нужным нам размером. Хеш мапа содержит структуру типа KeyValue с полями, собственно key value и двумя хелпер полями:
isDeleted - было ли поле удалено
isWrited - записано ли поле
Второе мне не нравится, я думаю мб можно было бы хранить не сами структуры, а ссылки на них, а если поле не записано, то можно хранить nil

Так как как-то само получилось, что я справляюсь с коллизиями методом открытой адресации, в каждом моем методе есть цикл от 0 до N (тут тоже кроется проблема)
Каждый метод берет первоначальный индекс с помощью метода hash
Запускатеся цикл
Для метдов Get и Delete если они встретили удаленное поле ИЛИ они встретили поле, которое записано и у него не тот ключ, который мы ищем, то изначальный индекс
инкрементится на 1.
Для метода Put есть только пропуск индексов, которые записаны, но имею не тот ключ, который мы хотим вставить.
Если Put нашел свой ключ, то он записывает новое значение в это поле. Если не нашел, то записывает в первое попавшееся пустое или удаленное поле

Если Get и Delete результатом своих поисков нашли записанное поле, то они возвращают его значение (ну Delete еще и удаляет его), если не нашли, то erorr

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Реализовывая массив с простым количеством ячеек, и получая хеш от целочисленного ключа по модулю от количества ячеек мы получаем хеш мапу, как и просит задание
Коллизии разрешаются методом открытой адресации


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

В среднем время получения значение, вставки и удаления O(1), но, если представить, то что мы сначала заполнили весю хешмапу, а потом удаллии все занчения,
то удаление и получение будут искать пока не найдут либо пустое либо исконное поле, а удаленные поля будут пропускать, что приведет к тому, что операции
удаления и получения будут занимать O(N).
Помог бы совет из теории: каунтер удаленных значений, при превышении лимита за O(N) создать новый массив уже без удаленных полей.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
мы сразу создаем массив уже заполненный пустыми значениями, так что от начала и до конца O(N)


-- PS --
Я думаю, что в это раз вы все таки отправите меня дорабатывать решение) Мне вообще эта тема с трудом далась, + неприятно отсылать недоделанное решение,
но что поделать. Я так старался догнать мою группу, и мне дали еще шанс догнать в каникулы, и я все равно ничего не успел, сложно, конечно, совмещать работу,
где сейчас делайн по проекту и алгоритмы, но я так старался, что я не могу не отправить это решение, хоть и понимаю головой, что это не лучшее решение, которое я мог бы сделать
*/

package hashTable

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type KeyValue struct {
	isWrited  bool
	isDeleted bool
	key       int
	value     int
}

type HashMap []KeyValue

func (hm HashMap) hash(key int) int {
	return key % len(hm)
}

func (hm HashMap) Get(key int) (int, error) {
	potentialKey := hm.hash(key)
	for i := 0; i < len(hm); i++ {
		if hm[potentialKey].isDeleted || (hm[potentialKey].isWrited && hm[potentialKey].key != key) {
			potentialKey++
			if potentialKey == len(hm) {
				potentialKey = 0
			}
		} else {
			break
		}

	}
	if hm[potentialKey].isWrited {
		return hm[potentialKey].value, nil
	} else {
		return 0, errors.New("no such key")
	}
}

func (hm HashMap) Put(kv KeyValue) {
	potentialKey := hm.hash(kv.key)
	for i := 0; i < len(hm); i++ {
		if hm[potentialKey].isWrited && hm[potentialKey].key != kv.key {
			potentialKey++
			if potentialKey == len(hm) {
				potentialKey = 0
			}
		} else {
			break
		}

	}
	hm[potentialKey] = kv
}

func (hm HashMap) Delete(key int) (int, error) {
	potentialKey := hm.hash(key)
	for i := 0; i < len(hm); i++ {
		if hm[potentialKey].isDeleted || (hm[potentialKey].isWrited && hm[potentialKey].key != key) {
			potentialKey++
			if potentialKey == len(hm) {
				potentialKey = 0
			}
		} else {
			break
		}

	}
	if hm[potentialKey].isWrited {
		defer func() {
			hm[potentialKey] = KeyValue{isDeleted: true}
		}()

		return hm[potentialKey].value, nil
	} else {
		return 0, errors.New("no such key")
	}

}

func HashTable(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	nCommands, _ := strconv.Atoi(scaner.Text())
	primeForHash := getPrimeForN(nCommands)
	writer := bufio.NewWriter(w)
	hm := make(HashMap, primeForHash)
	for i := 0; i < nCommands; i++ {
		scaner.Scan()
		command := strings.Fields(scaner.Text())
		switch {
		case command[0] == "get":
			getKey, _ := strconv.Atoi(command[1])
			val, err := hm.Get(getKey)
			if err != nil {
				writer.WriteString("None")

			} else {
				writer.WriteString(strconv.Itoa(val))
			}
			writer.WriteByte('\n')
		case command[0] == "put":
			putKey, _ := strconv.Atoi(command[1])
			putValue, _ := strconv.Atoi(command[2])
			hm.Put(KeyValue{isWrited: true, key: putKey, value: putValue})
		case command[0] == "delete":
			deleteKey, _ := strconv.Atoi(command[1])
			val, err := hm.Delete(deleteKey)
			if err != nil {
				writer.WriteString("None")

			} else {
				writer.WriteString(strconv.Itoa(val))
			}
			writer.WriteByte('\n')
		}
	}
	//
	writer.Flush()
}

func getPrimeForN(n int) int {
	switch {
	case n <= 10:
		return 13
	case n <= 100:
		return 107
	case n <= 1000:
		return 1009
	case n <= 10000:
		return 10007
	case n <= 100000:
		return 100003
	default:
		return 1000003
	}
}

func main() {
	HashTable(os.Stdin, os.Stdout)
}
