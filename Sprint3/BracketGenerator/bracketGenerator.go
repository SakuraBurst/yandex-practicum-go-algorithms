package bracketGenerator

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

func BracketGenerator(r io.Reader, w io.Writer) {
	stdReader := bufio.NewReader(r)
	nS, _ := stdReader.ReadString('\n')
	n, _ := strconv.Atoi(nS)
	bufferWriter := bufio.NewWriter(w)
	if n == 1 {
		bufferWriter.WriteString("()")
		bufferWriter.Flush()
		return
	}
	bS := make([]string, 0, n*2)
	recursionBracketAppend(&bS, n*2, "")

	sort.Strings(bS)

	for _, s := range bS {
		bufferWriter.WriteString(s)
		bufferWriter.WriteByte('\n')
	}
	bufferWriter.Flush()
}

func isCorrectBracketSequense(s string) bool {
	openedBrackets := 0
	for _, v := range s {
		if v == 40 {
			openedBrackets++
		} else {
			openedBrackets--
			if openedBrackets < 0 {
				return false
			}
		}
	}
	return openedBrackets == 0
}

func recursionBracketAppend(dc *[]string, n int, currentSquense string) {
	if n == 0 {
		if isCorrectBracketSequense(currentSquense) {
			*dc = append(*dc, currentSquense)
		}

	} else {
		recursionBracketAppend(dc, n-1, currentSquense+"(")
		recursionBracketAppend(dc, n-1, currentSquense+")")
	}
}

// func reverseWithBracketReverse(s []byte) string {
// 	for i, g := 0, len(s)-1; i < g; i, g = i+1, g-1 {
// 		s[i], s[g] = reverseBracket(s[g]), reverseBracket(s[i])
// 	}
// 	return string(s)
// }

// func reverseBracket(b byte) byte {
// 	if b == '(' {
// 		return ')'
// 	}
// 	return '('
// }

// type Deck struct {
// 	front                int
// 	back                 int
// 	deck                 []string
// 	currentLenght        int
// 	maxLength            int
// 	openedBracketCounter int
// }

// func (d *Deck) PushFront(s string) error {
// 	if d.isFull() {
// 		return errors.New("deck is full")
// 	}
// 	if isOpenedBracket(s) {
// 		d.openedBracketCounter++
// 	} else {
// 		d.openedBracketCounter--
// 	}
// 	d.deck[d.front] = s
// 	d.front = incrementToTheLimit(d.front, d.maxLength)
// 	d.currentLenght++
// 	return nil
// }

// func (d *Deck) PushBack(s string) error {
// 	if d.isFull() {
// 		return errors.New("deck is full")
// 	}
// 	if isOpenedBracket(s) {
// 		d.openedBracketCounter++
// 	} else {
// 		d.openedBracketCounter--
// 	}
// 	d.deck[d.back] = s
// 	d.back = decrementToTheZero(d.front, d.maxLength)
// 	d.currentLenght++
// 	return nil
// }
// func (d *Deck) PopFront() (string, error) {
// 	if d.isEmpty() {
// 		return "", errors.New("deck is empty")
// 	}
// 	d.front = decrementToTheZero(d.front, d.maxLength)
// 	if isOpenedBracket(d.deck[d.front]) {
// 		d.openedBracketCounter--
// 	} else {
// 		d.openedBracketCounter++
// 	}
// 	d.currentLenght--
// 	return d.deck[d.front], nil
// }
// func (d *Deck) PopBack() (string, error) {
// 	if d.isEmpty() {
// 		return "", errors.New("deck is empty")
// 	}
// 	d.back = incrementToTheLimit(d.front, d.maxLength)
// 	if isOpenedBracket(d.deck[d.back]) {
// 		d.openedBracketCounter--
// 	} else {
// 		d.openedBracketCounter++
// 	}
// 	d.currentLenght--
// 	return d.deck[d.back], nil
// }

// func (d Deck) FullDeckString() (string, error) {
// 	if !d.isFull() {
// 		return "", errors.New("deck is not full")
// 	}
// 	return strings.Join(d.deck, ""), nil
// }
// func (d Deck) isFull() bool {
// 	return d.currentLenght == d.maxLength
// }
// func (d Deck) isEmpty() bool {
// 	return d.currentLenght == 0
// }
// func (d Deck) HaveOpenedBracket() bool {
// 	return d.openedBracketCounter > 0
// }

// func incrementToTheLimit(value, limit int) int {
// 	value++
// 	if value == limit {
// 		return 0
// 	}
// 	return value
// }

// func decrementToTheZero(value, fallBack int) int {
// 	value--
// 	if value < 0 {
// 		return fallBack
// 	}
// 	return value
// }

// func isOpenedBracket(s string) bool {
// 	return s == "("
// }

// func BracketGenerator(r io.Reader, w io.Writer) {
// 	stdReader := bufio.NewReader(r)
// 	nS, _ := stdReader.ReadString('\n')
// 	n, _ := strconv.Atoi(nS)
// 	bufferWriter := bufio.NewWriter(w)
// 	dc := DeckContainer{make([]Deck, 0, n*2), n}

// 	recursionAppend(&dc)

// 	resultSlice := make([]string, 0, len(dc.decks))
// 	for _, v := range dc.decks {
// 		s, err := v.FullDeckString()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		resultSlice = append(resultSlice, s)
// 	}
// 	sort.Strings(resultSlice)

// 	for _, s := range resultSlice {
// 		bufferWriter.WriteString(s)
// 		bufferWriter.WriteByte('\n')
// 	}

// 	bufferWriter.Flush()
// }

// type DeckContainer struct {
// 	decks []Deck
// 	deep  int
// }

// func recursionAppend(dc *DeckContainer) {
// 	if dc.deep == 2 {
// 		for g := 0; g < dc.deep; g++ {
// 			deck := Deck{0, dc.deep - 1, make([]string, dc.deep*2), 0, dc.deep * 2, 0}
// 			for m := 0; m < g+1; m++ {
// 				deck.PushFront("(")
// 			}
// 			for i := 0; i < (dc.deep*2 - g); i++ {
// 				if deck.HaveOpenedBracket() {
// 					deck.PushFront(")")
// 				} else {
// 					deck.PushFront("(")
// 				}
// 			}
// 			dc.decks = append(dc.decks, deck)
// 		}
// 	} else {
// 		currentDeep := dc.deep
// 		dc.deep--
// 		recursionAppend(dc)
// 		newDecks := make([]Deck, 0, len(dc.decks)*3)
// 		for _, d := range dc.decks {
// 			for g := 0; g < 2; g++ {
// 				deck := Deck{0, dc.deep - 1, make([]string, currentDeep*2), 0, currentDeep * 2, 0}
// 				s, err := d.FullDeckString()
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				if g%2 == 0 {
// 					for _, r := range s {
// 						deck.PushFront(string(r))
// 					}
// 					deck.PushFront("(")
// 				} else {
// 					deck.PushFront("(")
// 					for _, r := range s {
// 						deck.PushFront(string(r))
// 					}
// 				}
// 				if deck.HaveOpenedBracket() {
// 					deck.PushFront(")")
// 				} else {
// 					deck.PushFront("(")
// 				}
// 				newDecks = append(newDecks, deck)
// 			}
// 		}
// 		reversedDecks := []Deck{}
// 		for _, nd := range newDecks {
// 			nds, err := nd.FullDeckString()
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			reverseNds := reverseWithBracketReverse([]byte(nds))
// 			if nds != reverseNds {
// 				deck := Deck{0, dc.deep - 1, make([]string, currentDeep*2), 0, currentDeep * 2, 0}
// 				for _, r := range reverseNds {
// 					deck.PushFront(string(r))
// 				}
// 				reversedDecks = append(reversedDecks, deck)
// 			}
// 		}
// 		dc.decks = append(newDecks, reversedDecks...)
// 	}
// }

// func reverseWithBracketReverse(s []byte) string {
// 	for i, g := 0, len(s)-1; i < g; i, g = i+1, g-1 {
// 		s[i], s[g] = reverseBracket(s[g]), reverseBracket(s[i])
// 	}
// 	return string(s)
// }

// func reverseBracket(b byte) byte {
// 	if b == '(' {
// 		return ')'
// 	}
// 	return '('
// }
