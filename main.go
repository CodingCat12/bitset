package main

import (
	"enqr/pkg/bitset"
	"fmt"
)

func main() {
	bits := bitset.NewBitset(256)

	strData := "Hello, World!"

	data := []byte(strData)

	bits.SetBytes(data)
	decodedBytes := bits.GetBytes()

	fmt.Println("Decoded String: ", string(decodedBytes))

	fmt.Print("Decoded Bytes (Hex): ")
	for _, b := range decodedBytes {
		fmt.Printf("%02x ", b)
	}
	fmt.Println()

	fmt.Print("Decoded Bytes (Binary): ")
	for _, b := range decodedBytes {
		fmt.Printf("%08b ", b)
	}
	fmt.Println()
}
