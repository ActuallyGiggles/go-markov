package markov

import (
	"fmt"
	"strings"
	"sync"
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
		go likelyBeginning(i, c)
		r = <-c
	case "TargetedBeginning":
		go targetedBeginning(i, c)
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
			recursionCounterMx.Unlock()
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
		Output:  "",
		Problem: "",
	}

	chain, exists := jsonToChain("./markov/chains/" + i.Chain + ".json")
	if !exists {
		message.Output = ""
		message.Problem = i.Chain + " does not exist."
		c <- message
		close(c)
		return
	}

	for true {
		if list, ok := chain[nextParent]; !ok {
			message.Output = ""
			message.Problem = i.Chain + " has no messages."
			c <- message
			close(c)
			return
		} else {
			list := list["nextList"]
			child = WeightedRandom(list)
		}

		if child == endKey {
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
	sentence := i.Target + " "
	child := ""
	splitChild := make([]string, 0)
	nextParent := startKey
	message := result{
		Output:  "",
		Problem: "",
	}

	// Check if chain exists and get it back as json
	chain, exists := jsonToChain("./markov/chains/" + i.Chain + ".json")
	if !exists {
		message.Output = ""
		message.Problem = i.Chain + " does not exist."
		c <- message
		close(c)
		return
	}

	options := make(map[string]int)

	// If exists and if the first word of a child matches the word that was chosen, add it to the list
	if list, ok := chain[nextParent]; !ok {
		message.Output = ""
		message.Problem = i.Chain + " has no messages."
		c <- message
		close(c)
		return
	} else {
		for _, combo := range list {
			for child, value := range combo {
				firstWordInChild := strings.Split(child, " ")[0]
				if firstWordInChild == i.Target {
					options[child] = value
				}
			}
		}
	}

	// If no phrase starts with word, ~~recurse~~ or ignore
	if len(options) == 0 {
		message.Output = ""
		message.Problem = "ERROR: no phrase starts with \"" + i.Target + "\" \nSource -> " + i.Chain + " " + nextParent
		c <- message
		close(c)
		return
	}

	// Randomly choose a starting phrase that starts with the word
	nextParent = WeightedRandom(options)

	for true {
		// Look for the nextParent in the chain, if it doesn't exist, return
		if list, ok := chain[nextParent]; !ok {
			message.Output = ""
			message.Problem = fmt.Sprintf("ERROR: %s does not contain nextParent: %s, even though %s was chosen in getSentenceFromAStart", i.Chain, nextParent, nextParent)
			c <- message
			close(c)
			return
		} else {
			list := list["nextList"]
			child = WeightedRandom(list)
		}

		if child == endKey {
			if len(strings.Split(nextParent, " ")) == 1 {
				message.Output = sentence
			} else {
				splitNextParent := strings.Split(nextParent, " ")
				message.Output = sentence + splitNextParent[1]
			}
			c <- message
			close(c)
			return
		}

		splitChild = strings.Split(child, " ")

		// If child is one word, add to sentence
		// Else, take the second word and prepare it to be searched
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
