package main

import "fmt"

func SearchString(mem *Memory) (string, error) {
	// Read the first page
	page, err := mem.ReadPage(0)
	if err != nil {
		return "", err
	}

	fmt.Println("page", string(page))

	// Search for the string "hey"
	for i, p := range page {
	}

	fmt.Println("here")

	return "hey", nil
}
