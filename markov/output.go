package markov

import "sync"

var (
	recursionCounter   = make(map[string]int)
	recursionCounterMx sync.Mutex
)

func Output(instructions OutputInstructions) (output string, problem string) {
	c := make(chan result)
	go outputController(instructions, c)
	r := <-c

	return r.Output, r.Problem
}

func outputController(i OutputInstructions, outputC chan result) {
	var r result
	c := make(chan result)

	switch i.Method {
	case "LikelyBeginning":
		//go likelyBeginning(i, c)
		r = <-c
	case "TargetedBeginning":
		//go targetedBeginning(i, c)
		r = <-c
	case "LikelyEnd":
	case "TargetedEnd":
	case "TargetedMiddle":
	}

	if r.Problem == "" {
		outputC <- r
		return
	} else {
		recursionCounterMx.Lock()
		recursionCounter[i.Chain] += 1

		if recursionCounter[i.Chain] > 5 {
			recursionCounter[i.Chain] = 0
			recursionCounterMx.Unlock()
			outputC <- r
			return
		} else {
			go outputController(i, outputC)
		}
	}

	return
}
