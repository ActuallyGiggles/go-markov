package main

import (
	"fmt"
	"sync"
)

var (
	workerMap   = make(map[string]*worker)
	workerMapMx sync.Mutex
)

func newWorker(name string) *worker {
	c, err := jsonToChain(name)
	var w *worker
	if err != nil {
		w = &worker{
			Name:  name,
			Chain: chain{},
		}
	} else {
		w = &worker{
			Name:  name,
			Chain: c,
		}
	}

	workerMapMx.Lock()
	workerMap[name] = w
	workerMapMx.Unlock()

	w.Status = "Ready"
	w.LastModified = now()

	return w
}

func (w *worker) addInput(content string) {
	w.ChainMx.Lock()
	defer w.ChainMx.Unlock()

	contentToChain(&w.Chain, w.Name, content)
	w.Intake += 1

	// for _, parent := range w.Chain.Parents {
	// 	for _, child := range parent.Next {
	// 		fmt.Println(parent.Word, "->", child.Word)
	// 	}
	// 	for _, grandparent := range parent.Previous {
	// 		fmt.Println(parent.Word, "<-", grandparent.Word)
	// 	}
	// }
}

func (w *worker) writeToFile() {
	defer duration(track(w.Name))

	w.ChainMx.Lock()
	defer w.ChainMx.Unlock()

	w.Status = "Writing"
	w.LastModified = now()

	fmt.Println("writing")
	chainToJson(w.Chain, w.Name)

	w.Intake = 0
	w.Status = "Ready"
	w.LastModified = now()
}
