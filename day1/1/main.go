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

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var d int
		_, err = fmt.Sscanf(scanner.Text(), "%d", &d)
		if err != nil {
			log.Fatal(err)
		}

		sum += d
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(sum)
}
