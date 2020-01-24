package main

import (
	"bufio"
	"os"
	"log"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		log.Fatalln(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

