package factorization_test

import (
	"testing"

	factorization "github.com/SakuraBurst/yandex-practicum-go-algorithms/Sprint1/Factorization"
)

func BenchmarkFactor(b *testing.B) {
	factorization.AllSimpleNumbersFromN(b.N)
}
