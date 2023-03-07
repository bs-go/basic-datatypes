package main

import "fmt"

func byteSlices() {
	b := make([]byte, 12)
	fmt.Println("Bytes slice:", b)

	b = []byte("Byte slice €")
	fmt.Println("Byte slice:", b)

	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))

	fmt.Println("Length of b:", len(b))
}
