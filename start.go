package markov

import (
	"fmt"
	"time"
)

var (
	workers       int
	writeInterval int
	intervalUnit  string
	startKey      string
	endKey        string
	Debug         bool

	nextWriteTime time.Time
	peakIntake    struct {
		Amount int
		Time   time.Time
	}
)

// Start initializes the Markov  package.
//
// Takes:
//		StartInstructions {
// 			Workers       int
// 			WriteInterval int
//			IntervalUnit  string
// 			StartKey      string
// 			EndKey        string
// 		}
func Start(instructions StartInstructions) {
	createChains()

	workers = instructions.Workers
	writeInterval = instructions.WriteInterval
	intervalUnit = instructions.IntervalUnit
	startKey = instructions.StartKey
	endKey = instructions.EndKey

	toWorker = make(chan input)

	startWorkers()

	go writeTicker()
}

func writeTicker() {
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

	nextWriteTime = time.Now().Add(time.Duration(writeInterval) * unit)
	for range time.Tick(time.Duration(writeInterval) * unit) {
		nextWriteTime = time.Now().Add(time.Duration(writeInterval) * unit)

		writing := 0
		for _, w := range workerMap {
			if Debug {
				fmt.Printf("Worker %d is writing...", w.ID)
				fmt.Println()
			}

			if w.Intake > peakIntake.Amount {
				peakIntake.Amount = w.Intake
				peakIntake.Time = time.Now()
			}

			if writing >= workers-1 {
				w.writeToChain()
				writing = 0
				continue
			}
			go w.writeToChain()
			writing += 1
		}
	}
}
