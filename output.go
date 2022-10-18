package main

import (
	"errors"
	"fmt"
	"strings"
)

func Out(oi OutputInstructions) (output string, err error) {
	name := oi.Chain
	// method := oi.Method
	// target := oi.Target
	output, err = LikelyBeginning(name)

	return output, nil
}

func LikelyBeginning(name string) (output string, err error) {
	output = ""
	parent := startKey
	child := ""

	c, err := jsonToChain(name)
	if err != nil {
		return "", err
	}

	for true {
		parentExists := false
		for _, cParent := range c.Parents {
			if cParent.Word == parent {
				parentExists = true
				child = getNextWord(cParent)

				if child == endKey {
					parentSplit := strings.Split(parent, " ")
					output = output + parentSplit[1]
					return output, nil
				} else {
					childSplit := strings.Split(child, " ")
					output = output + childSplit[0] + " "

					parent = child
					continue
				}
			}
		}

		if !parentExists {
			return output, errors.New(fmt.Sprintf("parent %s does not exist in chain %s", parent, name))
		}
	}

	return output, nil
}

func getNextWord(parent parent) (child string) {
	var wrS []wRand
	for _, word := range parent.Next {
		w := word.Word
		v := word.Value
		item := wRand{
			Word:  w,
			Value: v,
		}
		wrS = append(wrS, item)
	}
	child = weightedRandom(wrS)

	return child
}
