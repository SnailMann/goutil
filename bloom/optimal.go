package bloom

import (
	"githut.com/SnailMann/goutil/number"
	"math"
)

// OptimalNumOfHashFunctions Computes the number of hash functions (k)
func OptimalNumOfHashFunctions(n uint32, m uint32) uint32 {
	return uint32(math.Max(1, math.Round(float64(m/n)*math.Ln2)))
}

// OptimalNumOfBits Compute number of bits (m)
func OptimalNumOfBits(n uint32, p float64) uint32 {
	if p == 0 {
		p = number.Float64MinValue
	}
	return uint32((-float64(n) * math.Log(p)) / (math.Ln2 * math.Ln2))
}

// OptimalFpp Computes the false positive probability
// p = (1-e^(-kn/m))^k
func OptimalFpp(n uint32, m uint32, k uint32) float64 {
	var point float64 = 100000
	return math.Round(point*math.Pow(float64(1)-math.Exp(float64((-k*n)/m)), float64(k))) / point
}

// BitsOfElement Compute the bits each element (c)
// c = m / n
func BitsOfElement(n uint32, m uint32) float64 {
	return float64(m / n)
}

func BitCount(bytes []byte) uint32 {
	var c uint32 = 0
	for _, b := range bytes {
		a := b
		for a != 0 {
			a &= (a - 1)
			c++
		}
	}
	return c
}
