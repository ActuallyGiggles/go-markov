package markov

import (
	"strings"
)

// Input adds content to a chain writing queue
//
// 	Takes:
// 		Chain string
//		Content string
func Input(chain string, content string) {
	i := input{
		Chain:   chain,
		Content: content,
	}

	toWorker <- i
}

func contentToChain(chainMap *map[string]map[string]map[string]map[string]int, chain string, content string) {
	// Prepare the message to be more easily processed as a slice and add the startKey and endKey
	slice := prepareContentForChainProcessing(content)

	// Grab the current, previous, and next words in the slice and save it to the chain in proper order
	extractStartAndSaveToChain(*chainMap, chain, slice)
	extractEndAndSaveToChain(*chainMap, chain, slice)
	extractAndSaveToChain(*chainMap, chain, slice)
}

/*
prepareContentForChainProcessing processes content into a slice that
should be easier to manipulate when adding into an existing chain.
It accounts for any number of items and add start/end keys.

Example:
	Content: This is a test.
	Result: [startKey], This is, is a, a test., [endKey]]
*/
func prepareContentForChainProcessing(content string) []string {
	var returnSlice []string
	returnSlice = append(returnSlice, startKey)
	slice := strings.Split(content, " ")
	if len(slice) > 1 {
		for i := 0; i < len(slice)-1; i++ {
			firstWord := slice[i]
			secondWord := slice[i+1]
			returnSlice = append(returnSlice, firstWord+" "+secondWord)
		}
	} else {
		returnSlice = append(returnSlice, slice[0])
	}
	returnSlice = append(returnSlice, endKey)
	return returnSlice
}

// extractStartAndSaveToChain extracts the starts from slices and saves them to a Chain.
func extractStartAndSaveToChain(c map[string]map[string]map[string]map[string]int, channel string, slice []string) {
	start := slice[0]
	next := slice[1]

	_, channelOK := c[channel] // Does channel exist in c?
	if !channelOK {            // If no...
		c[channel] = make(map[string]map[string]map[string]int) // Create channel
	}
	_, startOK := c[channel][start] // Does start word exist in c?
	if !startOK {                   // If no...
		c[channel][start] = make(map[string]map[string]int) // Create start
	}
	_, nextListOK := c[channel][start]["nextList"] // Does the next list exist?
	if !nextListOK {                               // If no...
		c[channel][start]["nextList"] = make(map[string]int) // Create nextList
	}
	_, nextOK := c[channel][start]["nextList"][next] // Does next word exist in next list?
	if !nextOK {                                     // If no...
		c[channel][start]["nextList"][next] = 1 // Create next word and set counter to 1
	} else {
		c[channel][start]["nextList"][next] = c[channel][start]["nextList"][next] + 1 // Add 1 to existing next word counter
	}
}

// extractEndAndSaveToChain extracts the ends from slices and saves them to a chain.
func extractEndAndSaveToChain(c map[string]map[string]map[string]map[string]int, channel string, slice []string) {
	end := slice[len(slice)-1]
	previous := slice[len(slice)-2]

	_, channelOK := c[channel] // Does channel exist in c?
	if !channelOK {            // If no...
		c[channel] = make(map[string]map[string]map[string]int) // Create channel
	}
	_, endOK := c[channel][end] // Does start word exist in c?
	if !endOK {                 // If no...
		c[channel][end] = make(map[string]map[string]int) // Create end
	}
	_, previousListOK := c[channel][end]["previousList"] // Does the previous list exist?
	if !previousListOK {                                 // If no...
		c[channel][end]["previousList"] = make(map[string]int) // Create previousList
	}
	_, previousOK := c[channel][end]["previousList"][previous] // Does previous word exist in previous list?
	if !previousOK {                                           // If no...
		c[channel][end]["previousList"][previous] = 1 // Create previous word and set counter to 1
	} else {
		c[channel][end]["previousList"][previous] = c[channel][end]["previousList"][previous] + 1 // Add 1 to existing previous word counter
	}
}

// extractAndSaveToChain extracts from slices and saves them to a chain.
func extractAndSaveToChain(c map[string]map[string]map[string]map[string]int, channel string, slice []string) {
	for i := 0; i < len(slice)-2; i++ {
		previous := slice[i]
		current := slice[i+1]
		next := slice[i+2]

		_, channelOK := c[channel] // Does channel exist in c?
		if !channelOK {            // If no...
			c[channel] = make(map[string]map[string]map[string]int) // Create channel
		}
		_, currentOK := c[channel][current] // Does current word exist in c?
		if !currentOK {                     // If no...
			c[channel][current] = make(map[string]map[string]int) // Create current
		}
		_, previousListOK := c[channel][current]["previousList"] // Does the previous list exist?
		if !previousListOK {                                     // If no...
			c[channel][current]["previousList"] = make(map[string]int) // Create previousList
		}
		_, previousOK := c[channel][current]["previousList"][previous] // Does previous word exist in previous list?
		if !previousOK {                                               // If no...
			c[channel][current]["previousList"][previous] = 1 // Create previous word and set counter to 1
		} else {
			c[channel][current]["previousList"][previous] = c[channel][current]["previousList"][previous] + 1 // Add 1 to existing previous word counter
		}
		_, nextListOK := c[channel][current]["nextList"] // Does the next list exist?
		if !nextListOK {                                 // If no...
			c[channel][current]["nextList"] = make(map[string]int) // Create nextList
		}
		_, nextOK := c[channel][current]["nextList"][next] // Does next word exist in next list?
		if !nextOK {                                       // If no...
			c[channel][current]["nextList"][next] = 1 // Create next word and set counter to 1
		} else {
			c[channel][current]["nextList"][next] = c[channel][current]["nextList"][next] + 1 // Add 1 to existing next word counter
		}
	}
}