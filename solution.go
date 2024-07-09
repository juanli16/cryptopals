package main

import "fmt"

func SearchString(mem *Memory) (string, error) {
	// Read the first page
	page, err := mem.ReadPage(0)
	if err != nil {
		return "", err
	}

	fmt.Println("page", page)

	for _, p := range page {
		b, err := mem.ReadAddress(int(p))
		if err != nil {
			return "", err
		}

		fmt.Println("address", p, "value", b)
	}

	return "hey", nil
}
