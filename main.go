package main

import (
	"fmt"
	"time"
)

func main() {
	Start(StartInstructions{
		WriteMode:  "counter",
		WriteLimit: 0,
		StartKey:   "!S",
		EndKey:     "!E",
	})

	//In("test", "i am dead")

	//time.Sleep temporarily here because main exits faster than the goroutine
	time.Sleep(1 * time.Second)

	oi := OutputInstructions{
		Chain:  "test",
		Method: "TargetedBeginning",
		Target: "i",
	}
	output, err := Out(oi)

	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
