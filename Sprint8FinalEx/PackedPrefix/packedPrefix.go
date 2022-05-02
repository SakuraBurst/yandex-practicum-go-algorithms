/**
id в контесте 66670661
-- ПРИНЦИП РАБОТЫ --
Привет! Здесь я ничего умного не использовал, у меня даже ощущение, что до прохождения курса я бы решил задачу
точно так же, как и после:) Единственная хитрость, которую у меня получилось сделать кроется в распаковке строк.
Первую строку я распаковываю как обычно, а последующие распаковываю до наибольшего возможного префикса, а дальше останавливаюсь:)
Например, если при распаковке первой строки ее длина оказалась равна 6, а длинна второй оказалась, утрировано, 500,
то мне нет никакого смысла распаковывать все 500 символов. Я распаковываю строку до момента, когда результат >= 6 и возвращаю ровно 6-и символьную строку.
Далее банальным посимвольным сравнением я вычисляю наибольший общий префикс, и в следующей строке уже ищу длину этого префикса. Не знаю, хорошая ли эта оптимизация,
или так и задумано, потому что код с полной распаковкой всех символов я не отправлял:) Единственная проблема со всем этим в том, что у меня есть 2 функции, которые
по сути делают одно и то же, только одна распаковывает полную строку, а другая распаковывает до определенного лимита, как я описал выше. Неприятно, но делать
одну функцию и ставить там какие-то флаги или миллион ифов не хотелось, поэтому 2 функции:)


-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Распаковка идет в 2 этапа: сначала мы ищем начало запакованной строки с помощью проверки руны на число,
потом, с помощью searchForEndOfPackedSequence мы находим конец запакованной строки, и отправляем запакованную строку в следующий шаг рекурсии. Когда вернулась
распакованная строка, мы засовываем ее столько раз сколько нужно в strings.Builder{} (потому что так быстрее)).
Сам алгоритм без распаковки очень простой:
Берем первые две строки, ищем у них общий префикс, берем последующие строки сразу длинны этого префикса,
так-как больше он точно не может быть, повторяем операцию. В самом конце у нас будет наибольший общий префикс


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Я не знаю за сколько распаковывается строка :D Мне кажется за O(n), но меня терзают сомнения:)
В общем, если предположить, что у нас есть все строки в количестве n, и в ней есть минимальная строка длинны m, которая является общим префиксом для всех строк,
то мы имеем на каждом шаге распаковку длинны m и сравнение со старым префиксом тоже длинны m.
Итого O(n * (m + m))


-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Используя описание как во временной сложности, на каждом шаге мы имеем текущий общий префикс длинны m и новую строку для проверки длинны m
Итого O(m + m)

P.S.

Большое спасибо что ревьили мой говнокод эти полгода, жаль что не успел закончить со своим потоком) Вообще, я хотел у вас спросить, если вы
ревьюите на гошке, то наверное и работаете на ней?) Как сложно найти работу джуну? Я сейчас мидл фронт на js, программирую 2 года в общем,
работаю на нормальной работе 1.5 года, но чувствую, что до конца жизни пиксели двигать не хочу)


*/

package packedPrefix

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func PackedPrefix(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	prefix := ""
	currentString := ""
	for i := 0; i < n; i++ {
		scanner.Scan()
		if i == 0 {
			prefix = unpack(scanner.Text())
			continue
		}
		currentString = fixedLengthUnpack(scanner.Text(), len(prefix), 0)
		prefix = currentString[:firstNotEqualsElementIndex(prefix, currentString)]

	}
	io.WriteString(w, prefix)
}

func firstNotEqualsElementIndex(a, b string) int {
	i := 0
	max := 0
	if len(a) > len(b) {
		max = len(a)
	} else {
		max = len(b)
	}
	for ; i < max; i++ {
		if i >= len(a) || i >= len(b) || a[i] != b[i] {
			break
		}
	}
	return i
}

func unpack(s string) string {
	builder := strings.Builder{}
	for i := 0; i < len(s); {
		if unicode.IsDigit(rune(s[i])) {
			howMuch, _ := strconv.Atoi(string(s[i]))
			i += 2
			end := searchForEndOfPackedSequence(s, i)
			unpackedS := unpack(s[i:end])
			for g := 0; g < howMuch; g++ {
				builder.WriteString(unpackedS)
			}
			i = end + 1
		} else {
			builder.WriteByte(s[i])
			i++
		}

	}
	return builder.String()
}

func fixedLengthUnpack(s string, neededLength int, currentLength int) string {
	builder := strings.Builder{}
	for i := 0; i < len(s); {
		length := builder.Len() + currentLength
		if length >= neededLength {
			return cutString(builder.String(), length-neededLength)
		}
		if unicode.IsDigit(rune(s[i])) {
			howMuch, _ := strconv.Atoi(string(s[i]))
			i += 2
			end := searchForEndOfPackedSequence(s, i)
			unpackedS := fixedLengthUnpack(s[i:end], neededLength, builder.Len()+currentLength)
			for g := 0; g < howMuch; g++ {
				length := builder.Len() + currentLength
				if length >= neededLength {
					return cutString(builder.String(), length-neededLength)
				}
				builder.WriteString(unpackedS)
			}

			i = end + 1
		} else {
			builder.WriteByte(s[i])
			i++
		}

	}
	return builder.String()
}

func cutString(s string, cutTo int) string {
	return s[:len(s)-cutTo]
}

func searchForEndOfPackedSequence(s string, start int) int {
	bracketCount := 1
	g := start
	for ; g < len(s); g++ {
		if s[g] == '[' {
			bracketCount++
			continue
		}
		if s[g] == ']' {
			bracketCount--
			if bracketCount == 0 {
				break
			}
		}
	}
	return g
}

func main() {
	PackedPrefix(os.Stdin, os.Stdout)
}
