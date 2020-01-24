package main

import (
	"bufio"
	"os"
	"go-finder/finder"
	"log"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    f := finder.NewFinder("Go")
    f.SetMaxCountWorkers(5)

    f.Render()
    for scanner.Scan() {
        f.Start(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatalln(err)
    }

    f.StopWait()
}
