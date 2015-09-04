package main

import (
	"crypto/sha256"
	"encoding/binary"
	"math"
)

type BloomFilter struct {
	bitmap     [65536]int
	num_hashes int
}

func NewBloomFilter(num_hashes int) *BloomFilter {
	var filter = BloomFilter{num_hashes: num_hashes}
	return &filter
}

func OptimalNumHashes(num_elems int) (int, float64) {
	var previous float64 = 1
	k := 1

	for k <= 1000 {
		falsePositive := falsePositiveRate(num_elems, 65536, k)

		if previous < falsePositive {
			break
		}

		previous = falsePositive
		k++
	}

	return k - 1, previous
}

func falsePositiveRate(n, m, k int) float64 {
	exp := math.Exp(float64(-k*n) / float64(m))
	return math.Pow(1-exp, float64(k))
}

func (filter *BloomFilter) AddElement(str string) {
	indices := filter.getIndices([]byte(str))

	for _, n := range indices {
		filter.bitmap[n] += 1
	}
}

func (filter *BloomFilter) RemoveElement(str string) {
	indices := filter.getIndices([]byte(str))

	for _, n := range indices {
		if filter.bitmap[n] > 0 {
			filter.bitmap[n] -= 1
		}
	}
}

func (filter *BloomFilter) ContainsElement(str string) bool {
	indices := filter.getIndices([]byte(str))

	for _, n := range indices {
		if filter.bitmap[n] == 0 {
			return false
		}
	}

	return true
}

func (filter *BloomFilter) CountElements() int {
	sum := 0

	for _, n := range filter.bitmap {
		sum += n
	}

	return sum / filter.num_hashes
}

func (filter *BloomFilter) GetEmpty() int {
	count := 0

	for _, bit := range filter.bitmap {
		if bit == 0 {
			count += 1
		}
	}

	return count
}

func (filter *BloomFilter) FalsePositiveRate() float64 {
	n, m, k := filter.CountElements(), len(filter.bitmap), filter.num_hashes
	return falsePositiveRate(n, m, k)
}

func (filter *BloomFilter) getIndices(arr []byte) []uint16 {
	result := sha256.Sum256(arr)
	indices := make([]uint16, filter.num_hashes)

	for i, j := 0, 0; i < filter.num_hashes; i, j = i+1, j+2 {
		sum := binary.BigEndian.Uint16(result[j:(j + 2)])
		indices[i] = sum
	}

	return indices
}
