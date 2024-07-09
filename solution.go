package solution

import "fmt"

type Memory interface {
	// Reads a byte from the specified address.
	ReadAddress(address int) (byte, error)

	// Reads an entire page by its ID.
	ReadPage(pageID int) ([]byte, error)
}

func SearchString(mem *Memory) (string, error) {
	fmt.Println("here")

	return "hey", nil
}
