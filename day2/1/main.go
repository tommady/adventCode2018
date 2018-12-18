package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var two, three int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counter := make(map[rune]int)
		for _, w := range scanner.Text() {
			counter[w]++
		}

		var twoFound, threeFound bool
	LOOP:
		for _, count := range counter {
			switch count {
			case 2:
				if !twoFound {
					two++
					twoFound = true
				}
			case 3:
				if !threeFound {
					three++
					threeFound = true
				}
			}
			if twoFound && threeFound {
				break LOOP
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(two * three)
}
