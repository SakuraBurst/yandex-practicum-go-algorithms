/*
-- ПРИНЦИП РАБОТЫ --
id в контесте 61657768

Привет! Ну тут все просто, потому что как и с предыдущей задачей решение было описано в теории почти полностью)

К нам приходит корень двоичное дерево поиска (далее BST) и мы начинаем рекурсивный поиск по нему в поисках нужной нам ноды.
Если текущаяя нода нам не подходит, и ее значение больше, чем данный нам ключ, то ищем в левом поддереве, если меньше - то в правом.
В результате нам вернется измененная (или не изменная) левое или правое поддерево, в зависимости от того, в какое поддерево мы пошли.
Если мы нашли нужную нам ноду, то у нас три варианта того что произойдет:
1. Нужная нам нода оказалась листом. В данном случае мы просто возваращаем nil тем самым удалаяя ноду
2. Нужная нам нода имеет только одно поддерево. В данном случае мы просто возвращаем поддерево, тем самым удаляя ноду
3. Нужная нам нода оказалась полноценным родителем. В данном случае я выбрал поиск наименьшего значения в правом поддереве

--поиск наименьшего значения в правом поддереве (функция getAndDeleteSmallestValInTree)--
Поиск наименьшего значения работает почти также как и основаня функция, за исключением того, что она идет только влево
и возвращает нам значение удаленной ноды:
Принимая правое поддерево функция идет рекурсивно влево, пока не найдет ноду, где левого поддерева нет, и когда найдет
у нас будет два варианта что произойдет:
1. Нужная нам нода оказалась листом. В данном случае мы просто возваращаем nil тем самым удалаяя ноду и value ноды, чтобы вернуть значение
в нашу основную функцию
2. Нужная нам нода имеет правое поддерево. В данном случае мы возвращам поддерево и значение

Когда функция что-то вернула, в одном из рекурсивных циклов, то значение удаленной ноды передается наверх, а изменное поддерево
встает на место левого поддерева

--поиск наименьшего значения в правом поддереве (функция getAndDeleteSmallestValInTree)--

Возвращаемся к основной функции, так вот, если у нас выпал 3 вариант, то в резльтате выполнения getAndDeleteSmallestValInTree нам вернется
изменное правое поддерево и value удаленной ноды. В value текущий ноды мы записывам value удаленной ноды, тем самым удаляя текущее value
и на место правого поддерева мы записывам изменное правое поддерево (это нужно, например, если правое поддерево оказалось листом и изменное правое
поддерево возвращается как nil)

Если нужного нам значения не оказалось в BST и мы дошли до несуществующей ноды (самый первый базовый случай, node == nil) то просто возвращаем nil
В результате нам вернется тоже самое дерево благодря тому, что на место правых и левых поддеревьев верунтся теже самые поддеревья

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Главный фактор корректности, на мой взгляд в том, что если мы нашли нужную ноду, то мы используем корректный алгорим по замене этой ноды,
на наименьшее значение в правом поддереве.
А если не нашли, то так как мы возвращаем изменное поддерево для того, чтобы поместить его в родителя, то нам ничего не мешает возвращать
теже самые значения и оставить дерево в исконном виде


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

сложность всегда составляет O(h) где h - высота дерева:
Если искомый ключ находится в листе, то нам нужно преодолеть h нод чтобы до него добраться.
Если, предположим, искомый ключ находится где-то в начале или середине то чтобы найти потребуется, конечно,
меньше времени (O(1) если в начале, O(h / 2) если в середине) но чтобы найти наименьший элемент потребуется все оставшееся время.

Кстати,если дерево сбалансированно, то удаление будет занимать O(log n)

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Мы не создам доп переменных, поэтому O(1)
*/

package deleteNode

type TestNode struct {
	Value int
	Left  *TestNode
	Right *TestNode
}

func Remove(node *TestNode, key int) *TestNode {
	if node == nil {
		return node
	}
	if node.Value == key {
		// если нужная нам нода это лист, то просто возвращаем nil
		if node.Left == nil && node.Right == nil {
			return nil
		}
		// если есть только правое или левое поддерево, то просто возвращаем эту поддерево
		if node.Left == nil {
			return node.Right
		}

		if node.Right == nil {
			return node.Left
		}
		// если есть и правое и левое поддерево, то ищем в правой самое маленькое значение
		// попутно удаляя его и ставя на место правого поддерева новое значение
		val, replaceNode := getAndDeleteSmallestValInTree(node.Right)
		node.Right = replaceNode
		node.Value = val
		return node
	}
	if node.Value > key {
		newNode := Remove(node.Left, key)
		node.Left = newNode
	} else {
		newNode := Remove(node.Right, key)
		node.Right = newNode
	}

	return node
}

func getAndDeleteSmallestValInTree(n *TestNode) (int, *TestNode) {
	if n.Left == nil {
		if n.Right == nil {
			return n.Value, nil
		}
		return n.Value, n.Left
	}
	val, replaceNode := getAndDeleteSmallestValInTree(n.Left)
	n.Left = replaceNode
	return val, n
}
