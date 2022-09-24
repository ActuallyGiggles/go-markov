# go-markov

```go
package main

import (
	"fmt"
	"go-markov/markov"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	instructions := markov.StartInstructions{
		Workers:       5,
		WriteInterval: 1,
		IntervalUnit:  "hours",
		StartKey:      "start",
		EndKey:        "end",
	}

	markov.Start(instructions)

	outputI := markov.OutputInstructions{
		Method: "TargetedBeginning",
		Chain:  "test",
		Target: "This",
	}

	output, problem := markov.Output(outputI)

	fmt.Println(output)
	fmt.Println(problem)

	<-sc
	log.Println("Stopping...")
}```
