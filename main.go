package main

import (
	"bufio"
	"os"
	"go-finder/finder"
	"log"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

    f := finder.Finder {
        MaxCountWorkers: 5
        SearchWord: "Go",
    }

    f.Render()
	for scanner.Scan() {
        f.Start(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

