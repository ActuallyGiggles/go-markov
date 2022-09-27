package markov

import (
	"fmt"
	"time"
)

var (
	startKey string
	endKey   string
	Debug    bool

	nextWriteTime time.Time
)

// Start initializes the Markov  package.
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

	nextWriteTime = time.Now().Add(time.Duration(value) * unit)
	for range time.Tick(time.Duration(value) * unit) {
		nextWriteTime = time.Now().Add(time.Duration(value) * unit)

		writing := 0
		for _, w := range workerMap {
			if Debug {
				fmt.Printf("Worker %d is writing...", w.ID)
				fmt.Println()
			}
			if writing >= 1 {
				w.writeToChain()
				writing = 0
			}
			go w.writeToChain()
			writing += 1
		}
	}
}
