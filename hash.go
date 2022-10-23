package bloomfilter

type (
	SimpleHash struct {
		cap  uint
		seed uint
	}
)

var (
	seeds = []uint{7, 11, 13, 31, 37}
)

func (h SimpleHash) hash(value string) uint {
	var result uint = 0
	for i := 0; i < len(value); i++ {
		result = result*h.seed + uint(value[i])
	}
	return (h.cap - 1) & result
}
