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

	words := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minDiff := 2 << 10
	var minDiffA, minDiffB string
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			diff := 0
			for k := 0; k < len(words[j]); k++ {
				if words[i][k] != words[j][k] {
					diff++
				}
			}
			if diff < minDiff {
				minDiffA = words[i]
				minDiffB = words[j]
				minDiff = diff
			}
		}
	}

	ans := ""
	for i := 0; i < len(minDiffA); i++ {
		if minDiffA[i] == minDiffB[i] {
			ans += string(minDiffA[i])
		}
	}
	log.Println(ans)
}
