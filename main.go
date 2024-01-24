package main

import (
	"log"
	"os"

	"github.com/1garo/futur3/cmd"
)


func main() {
	yfile, err := os.ReadFile("example.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.CalculateInterest(yfile)
	if err != nil {
		log.Fatal(err)
	}
}
