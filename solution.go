package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/semaphore"
)

func SearchString(mem *Memory) (string, error) {
	var (
		window []byte
		sem    = semaphore.NewWeighted(4)
		ctx    = context.Background()
		res    string
	)

	for pageN := 0; pageN < 1000; pageN++ {
		if err := sem.Acquire(ctx, 1); err != nil {
			break
		}

		go func(pageN int) {
			defer sem.Release(1)
			// Read the first page
			page, err := mem.ReadPage(pageN)
			if err != nil {
				return
			}

			for _, p := range page {
				b, err := mem.ReadAddress(int(p))
				if err != nil {
					return
				}

				if len(window) == 5 && string(window) != "gc24{" {
					// reset window
					window = []byte{}
				}

				if len(window) > 5 && string(window[len(window)-1]) == "}" {
					// end of flag
					res = string(window)
					return
				}

				window = append(window, b)
			}

			return
		}(pageN)
	}

	// We block here until done.
	if err := sem.Acquire(ctx, 4); err != nil {
		fmt.Printf("Failed to acquire semaphore: %v", err)
	}

	if res != "" {
		return res, nil
	}

	return "", fmt.Errorf("flag not found")
}
