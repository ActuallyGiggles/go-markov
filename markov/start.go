package markov

import (
	"log"
	"time"
)

var (
	startKey string
	endKey   string
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

	go writeTicker(instructions.WriteInterval)
}

func writeTicker(workerWriteInterval int) {
	for range time.Tick(time.Duration(workerWriteInterval) * time.Minute) {
		for _, w := range workerMap {
			log.Printf("Worker %d is writing...", w.ID)
			w.writeToChain()
			w.Intake = 0
		}
	}
}
