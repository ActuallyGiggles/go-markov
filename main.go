package main

import (
	"fmt"
	"time"

	"go-markov/markov"
)

func main() {
	markov.Start(markov.StartInstructions{
		WriteMode:  "counter",
		WriteLimit: 0,
		StartKey:   "!S",
		EndKey:     "!E",
	})

	//In("test", "i am dead")

	//time.Sleep temporarily here because main exits faster than the goroutine
	time.Sleep(1 * time.Second)

	oi := markov.OutputInstructions{
		Chain:  "test",
		Method: "TargetedBeginning",
		Target: "i",
	}
	output, err := markov.Out(oi)

	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
