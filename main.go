package main

import (
	"awsce"
	"log"
)

func main() {
	costOutputs, err := awsce.FetchTotalCost()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(costOutputs)
}
