package markov

import (
	"errors"
	"fmt"
	"strings"
)

func Out(oi OutputInstructions) (output string, err error) {
	name := oi.Chain
	method := oi.Method
	target := oi.Target

	switch method {
	case "LikelyBeginning":
		output, err = LikelyBeginning(name)
	case "TargetedBeginning":
		output, err = TargetedBeginning(name, target)
	}

	return output, err
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

					if len(parentSplit) == 1 {
						output = output + parent
						return output, nil
					}

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

func TargetedBeginning(name string, target string) (output string, err error) {
	output = ""
	parent := ""
	child := ""

	c, err := jsonToChain(name)
	if err != nil {
		return "", err
	}

	initial := true
	var initialList []string

	for true {
		parentExists := false
		for parentNumber, cParent := range c.Parents {
			if initial {
				if parentNumber >= len(c.Parents)-1 {
					initial = false
					parentExists = true
					parent = pickRandomParent(initialList)
					parentSplit := strings.Split(parent, " ")
					output = parentSplit[0] + " "
					break
				}

				potentialParentSplit := strings.Split(cParent.Word, " ")
				if potentialParentSplit[0] == target {
					initialList = append(initialList, cParent.Word)
					continue
				} else {
					continue
				}
			}

			if cParent.Word == parent {
				parentExists = true
				child = getNextWord(cParent)

				if child == endKey {
					parentSplit := strings.Split(parent, " ")

					if len(parentSplit) == 1 {
						output = output + parent
						return output, nil
					}

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
			return output, errors.New(fmt.Sprintf("parent(%s) does not exist in chain(%s)", parent, name))
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

func pickRandomParent(parents []string) (parent string) {
	parent = PickRandomFromSlice(parents)

	return parent
}