package markov

import (
	"fmt"
	"time"
)

var (
	startKey string
	endKey   string
	debug    bool = true
)

// Start initializes the Markov package.
//
// Takes:
//		StartInstructions {
// 			Workers       int
// 			WriteInterval int
// 			StartKey      string
// 			EndKey        string
// 		}
func Start(instructions StartInstructions) {
	createChains()

	startKey = instructions.StartKey
	endKey = instructions.EndKey

	toWorker = make(chan input)

	startWorkers(instructions.Workers)

	go writeTicker(instructions.WriteInterval, instructions.IntervalUnit)
}

func writeTicker(value int, intervalUnit string) {
	var unit time.Duration

	switch intervalUnit {
	default:
		unit = time.Minute
	case "seconds":
		unit = time.Second
	case "minutes":
		unit = time.Minute
	case "hours":
		unit = time.Hour
	}

	for range time.Tick(time.Duration(value) * unit) {
		for _, w := range workerMap {
			if debug {
				fmt.Printf("Worker %d is writing...", w.ID)
				fmt.Println()
			}
			w.writeToChain()
			w.Intake = 0
		}
	}
}
