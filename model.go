package markov

import "sync"

// StartInstructions are the instructions on how Markov should be started
//
// 		Workers: How many workers to create
// 		WriteInterval: How often to write to chains (in minutes)
//		IntervalUnit: What unit to use for WriteInterval (default minutes)
//			"seconds"
//			"minutes"
//			"hours"
//		StartKey: A string that is used to signify the natural beginning to the content
// 		EndKey: A string that is used to signify the natural end to the content
//
// (StartKey and EndKey should be strings that are unlikely to be natural content)
type StartInstructions struct {
	Workers       int
	WriteInterval int
	IntervalUnit  string
	StartKey      string
	EndKey        string
}

// WorkerStats contains the current statistics on a worker
//
// 		ID: ID of worker
//		Intake: How many inputs this worker has gotten (clears after every Instructions.WriteInterval)
//		Status: Current status of worker
//			"Ready": Ready to work
//			"Adding": Adding content to queue
//			"Writing": Writing queue to chains
//		LastModified: The last time the status was updated
type WorkerStats struct {
	ID           int
	Intake       int
	Status       string
	LastModified string
}

// OutputInstructions are the instructions to give Markov when asking for an output
//
// 		Method: Which method to use when constructing the output
//			"LikelyBeginning": Build from a likely beginning
//			"TargetedBeginning": Build from a targeted beginning
//			"LikelyEnding": Build from a likely Ending (to be added)
//			"TargetedEnding": Build from a targeted ending (to be added)
//			"TargetedMiddle": Build from a targeted middle (to be added)
//		Chain: Which chain to use
// 		Target: Optional target (leave blank if N/A)
type OutputInstructions struct {
	Method string
	Chain  string
	Target string
}

type input struct {
	Chain   string
	Content string
}

type worker struct {
	ID           int
	Chain        map[string]map[string]map[string]map[string]int
	ChainMx      sync.Mutex
	Intake       int
	Status       string
	LastModified string
}

type result struct {
	Output  string
	Problem string
}
