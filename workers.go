package markov

import (
	"fmt"
	"sync"
)

var (
	workerMap   = make(map[int]*worker)
	workerMapMx sync.Mutex
	toWorker    chan input
)

func startWorkers(workerAmount int) {
	for i := 0; i < workerAmount; i++ {
		go newWorker(i)
	}
}

func newWorker(id int) {
	workerMapMx.Lock()
	workerID := len(workerMap) + 1

	w := &worker{
		ID:    workerID,
		Chain: make(map[string]map[string]map[string]map[string]int),
	}

	workerMap[workerID] = w
	workerMapMx.Unlock()

	w.Status = "Ready"
	w.LastModified = now()

	for in := range toWorker {
		w.addToQueue(in.Chain, in.Content)
	}
}

func (w *worker) addToQueue(chain string, content string) {
	w.Status = "Adding"
	w.LastModified = now()

	w.ChainMx.Lock()
	contentToChain(&w.Chain, chain, content)
	w.Intake += 1
	w.ChainMx.Unlock()

	if debug {
		fmt.Println("added")
	}

	w.Status = "Ready"
	w.LastModified = now()
}

func (w *worker) writeToChain() {
	w.Status = "Writing"
	w.LastModified = now()

	w.ChainMx.Lock()
	for currentChain, currentChainValue := range w.Chain {
		path := "./markov/chains/" + currentChain + ".json"

		existingChain, chainExists := jsonToChain(path)
		if !chainExists {
			existingChain = make(map[string]map[string]map[string]int)
		}

		for parent, parentValue := range currentChainValue {
			if _, ok := existingChain[parent]; !ok {
				existingChain[parent] = make(map[string]map[string]int)
			}
			for list, listValue := range parentValue {
				if _, ok := existingChain[parent][list]; !ok {
					existingChain[parent][list] = make(map[string]int)
				}
				for child, timesUsed := range listValue {
					existingChain[parent][list][child] += timesUsed
				}
			}
		}

		chainToJson(existingChain, path)

		w.Chain = make(map[string]map[string]map[string]map[string]int)
	}
	w.ChainMx.Unlock()

	if debug {
		fmt.Println("written")
	}

	w.Status = "Ready"
	w.LastModified = now()
}

// WorkersStats returns a slice of type WorkerStats
//
//
func WorkersStats() (slice []WorkerStats) {
	workerMapMx.Lock()
	for _, w := range workerMap {
		e := WorkerStats{
			ID:           w.ID,
			Intake:       w.Intake,
			Status:       w.Status,
			LastModified: w.LastModified,
		}

		slice = append(slice, e)
	}
	workerMapMx.Unlock()

	return slice
}
