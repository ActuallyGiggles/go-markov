package markov

import (
	"sync"
	wr "github.com/mroth/weightedrand"
)

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

func likelyBeginning(i OutputInstructions, c chan result) {
	sentence := ""
	child := ""
	splitChild := make([]string, 0)
	nextParent := startKey
	message := result{
		Output: "",
		Problem:   "",
	}

	chain, exists := jsonToChain("./markovdb/" + channel + ".json")
	if !exists {
		message.Output = "We ran into a problem!"
		message.Problem = "ERROR: " + channel + " does not exist."
		c <- message
		close(c)
		return
	}

	for true {
		if list, ok := chain[nextParent]; !ok {
			message.Output = channel + " has no messages!"
			message.Problem = "ERROR: " + channel + " has no messages."
			c <- message
			close(c)
			return
		} else {
			list := list["nextList"]
			child = WeightedRandom(list)
		}

		if child == endPhrase {
			if len(strings.Split(nextParent, " ")) == 1 {
				message.Output = sentence + nextParent
			} else {
				message.Output = sentence + splitChild[1]
			}
			c <- message
			close(c)
			return
		}

		splitChild = strings.Split(child, " ")

		if !DoesSliceContainIndex(splitChild, 1) {
			if sentence == "" {
				sentence = child
			}
			message.Output = sentence
			c <- message
			close(c)
			return
		} else {
			nextParent = child
			sentence = sentence + splitChild[0] + " "
		}
	}
	message.Output = sentence
	c <- message
	close(c)
	return
}

func targetedBeginning(i OutputInstructions, c chan result) {
	
}
