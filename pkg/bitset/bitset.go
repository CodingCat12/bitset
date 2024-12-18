package bitset

type Bitset struct {
	size int
	data []uint64
}

func NewBitset(size int) *Bitset {
	return &Bitset{
		size: size,
		data: make([]uint64, (size+63)/64),
	}
}

func (b *Bitset) Size() int {
	return b.size
}

func (b *Bitset) Data() []uint64 {
	result := make([]uint64, len(b.data))
	copy(result, b.data)
	return result
}
