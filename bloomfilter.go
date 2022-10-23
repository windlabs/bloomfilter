package bloomfilter

import (
	"github.com/bits-and-blooms/bitset"
)

const (
	DEFAULT_SIZE = 1 << 31
)

type (
	BloomFilter struct {
		set   *bitset.BitSet
		funcs [5]SimpleHash
	}
)

func NewBloomFilter() (bf *BloomFilter) {
	bf = new(BloomFilter)
	for i := 0; i < len(bf.funcs); i++ {
		bf.funcs[i] = SimpleHash{DEFAULT_SIZE, seeds[i]}
	}
	bf.set = bitset.New(DEFAULT_SIZE)

	return bf
}

func (bf BloomFilter) Add(value string) {
	for _, f := range bf.funcs {
		bf.set.Set(f.hash(value))
	}
}

func (bf BloomFilter) Contains(value string) (ret bool) {
	if value == "" {
		return false
	}
	ret = true
	for _, f := range bf.funcs {
		ret = ret && bf.set.Test(f.hash(value))
	}
	return ret
}
