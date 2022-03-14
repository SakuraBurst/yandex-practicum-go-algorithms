package main

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

//func BenchmarkMain(t *testing.B) {
//	for i := 0; i < t.N; i++ {
//		main()
//	}
//}

var aStr string

var bStr string

func init() {
	bigBigInt := 1 << 30
	aBuilder := strings.Builder{}
	bBuilder := strings.Builder{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < bigBigInt; i++ {
		a := uint8(rand.Intn(122-65) + 65)
		aBuilder.WriteByte(a)
		bBuilder.WriteByte(a)
	}
	aBuilder.WriteString("a")
	bBuilder.WriteString("b")
	aStr = aBuilder.String()
	bStr = bBuilder.String()
}

func BenchmarkStringsCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsCompare(aStr, bStr)
	}
}

func BenchmarkOperatorCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		operatorCompare(aStr, bStr)
	}
}
