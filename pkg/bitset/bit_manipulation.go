package bitset

import (
	"errors"
)

func (b *Bitset) Set(index int) error {
	if index < 0 || index >= b.size {
		return errors.New("index out of bounds")
	}

	wordIndex := index / 64
	offset := index % 64

	b.data[wordIndex] |= 1 << offset
	return nil
}

func (b *Bitset) Clear(index int) error {
	if index < 0 || index >= b.size {
		return errors.New("index out of bounds")
	}

	wordIndex := index / 64
	offset := index % 64

	b.data[wordIndex] &^= 1 << offset
	return nil
}

func (b *Bitset) Flip(index int) error {
	if index < 0 || index >= b.size {
		return errors.New("index out of bounds")
	}

	wordIndex := index / 64
	offset := index % 64

	b.data[wordIndex] ^= 1 << offset
	return nil
}

func (b *Bitset) Get(index int) (bool, error) {
	if index < 0 || index >= b.size {
		return false, errors.New("index out of bounds")
	}

	wordIndex := index / 64
	offset := index % 64

	return (b.data[wordIndex] & (1 << offset)) != 0, nil
}

func (b *Bitset) SetAll() {
	for i := range b.data {
		b.data[i] = ^uint64(0)
	}
}

func (b *Bitset) ClearAll() {
	b.data = make([]uint64, len(b.data))
}

func (b *Bitset) GetAll() []bool {
	result := make([]bool, b.size)
	for i := 0; i < b.size; i++ {
		wordIndex := i / 64
		offset := i % 64
		result[i] = (b.data[wordIndex] & (1 << offset)) != 0
	}
	return result
}

func (b *Bitset) FlipAll() {
	for i := range b.data {
		b.data[i] = ^b.data[i]
	}
}
