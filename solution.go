package main

import "fmt"

func SearchString(mem *Memory) (string, error) {
	var (
		maxPage = 10
		window  []byte
	)

	for pageN := 0; pageN < maxPages; pageN++ {
		page, err := mem.ReadPage(pageN)
		if err != nil {
			fmt.Println("pageN", pageN)
			return "", err
		}

		for _, p := range page {
			b, err := mem.ReadAddress(int(p))
			if err != nil {
				return "", err
			}

			if len(window) == 5 && string(window) != "gc24{" {
				// reset window
				window = []byte{}
			}

			if len(window) > 5 {
				fmt.Println("window", string(window))
			}

			if len(window) > 5 && string(window[len(window)-1]) == "}" {
				// end of flag
				return string(window), nil
			}

			window = append(window, b)
		}
	}

	return "", fmt.Errorf("flag not found")
}
