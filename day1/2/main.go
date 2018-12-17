package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	nums := make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var d int
		_, err = fmt.Sscanf(scanner.Text(), "%d", &d)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, d)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	seen := map[int]struct{}{0: struct{}{}}
	sum := 0
LOOP:
	for {
		for _, num := range nums {
			sum += num
			if _, exist := seen[sum]; exist {
				break LOOP
			} else {
				seen[sum] = struct{}{}
			}
		}
	}

	log.Println(sum)
}
