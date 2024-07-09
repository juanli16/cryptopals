package main

import "fmt"

func SearchString(mem *Memory) (string, error) {
	// Read the first page
	page, err := mem.ReadPage(0)
	if err != nil {
		return "", err
	}

	var window []byte
	for _, p := range page {
		b, err := mem.ReadAddress(int(p))
		if err != nil {
			return "", err
		}

		if len(window) == 5 && string(window) != "gc24{" {
			// reset window
			window = []byte{}
		}

		if len(window) > 5 && string(window[len(window)-1]) == "}" {
			// end of flag
			fmt.Println("found ", string(window))
			break
		}

		window = append(window, b)
		fmt.Println(string(window))
	}

	return "hey", nil
}
