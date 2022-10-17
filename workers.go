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
	workerMapMx.Lock()
	defer workerMapMx.Unlock()

	c := jsonToChain(name)

	w := &worker{
		Name:  name,
		Chain: c,
	}

	workerMap[name] = w

	w.Status = "Ready"
	w.LastModified = now()

	return w
}

func (w *worker) addInput(content string) {
	w.ChainMx.Lock()
	defer w.ChainMx.Unlock()

	fmt.Println(w.Chain)
	contentToChain(&w.Chain, w.Name, content)
	fmt.Println(w.Chain)
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
	path := "./markov/chains/" + w.Name + ".json"

	chainToJson(w.Chain, path)

	w.Intake = 0
	w.Status = "Ready"
	w.LastModified = now()
}
