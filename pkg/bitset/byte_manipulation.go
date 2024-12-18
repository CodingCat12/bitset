package bitset

import "errors"

func (b *Bitset) GetByte(index int) (byte, error) {
	if index < 0 || index+7 >= b.size {
		return 0, errors.New("index out of bounds")
	}

	wordIndex := index / 64
	offset := index % 64

	mask := uint64(0b11111111) << offset
	result := (b.data[wordIndex] & mask) >> offset

	return byte(result), nil
}

func (b *Bitset) SetByte(index int, value byte) error {
	if index < 0 || index+7 >= b.size {
		return errors.New("index out of bounds")
	}

	wordIndex := index / 64
	offset := index % 64

	mask := uint64(value) << offset

	b.data[wordIndex] &^= (uint64(0b11111111) << offset)
	b.data[wordIndex] |= mask

	return nil
}

func (b *Bitset) GetBytes() []byte {
	byteCount := (b.size + 7) / 8
	result := make([]byte, byteCount)

	for i := range byteCount {
		index := i * 8

		wordIndex := index / 64
		offset := index % 64

		mask := uint64(0b11111111) << offset
		byteValue := (b.data[wordIndex] & mask) >> offset

		result[i] = byte(byteValue)
	}

	return result
}

func (b *Bitset) SetBytes(data []byte) error {
	if len(data)*8 > b.size {
		return errors.New("data exceeds bitset size")
	}

	for i, byteVal := range data {
		index := i * 8

		wordIndex := index / 64
		offset := index % 64

		mask := uint64(0b11111111) << offset

		b.data[wordIndex] &^= mask
		b.data[wordIndex] |= uint64(byteVal) << offset
	}
	return nil
}
