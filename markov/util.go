package markov

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func jsonToChain(path string) (chain map[string]map[string]map[string]int, exists bool) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Failed reading file: %s", err)
		return nil, false
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("jsonToChain error: ", path, "\n", err)
		return nil, false
	}

	err = json.Unmarshal(content, &chain)
	if err != nil {
		log.Println("Error when unmarshalling file:", path, "\n", err)
		return nil, false
	}

	return chain, true
}

func chainToJson(chain map[string]map[string]map[string]int, path string) {
	file, _ := json.MarshalIndent(chain, "", " ") // Convert data to JSON
	_ = ioutil.WriteFile(path, file, 0644)        // Save data as JSON
}

// PrettyPrint prints out an object in a pretty format
func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
}

// createChains creates a markovdb folder if it doesn't exist.
func createChains() {
	// Create or check if main markov db folder exists
	_, dberr := os.Stat("markov/chains")
	if os.IsNotExist(dberr) {
		err := os.MkdirAll("markov/chains", 0755)
		if err != nil {
			log.Println(err)
		}
	}
}

func now() string {
	return time.Now().Format("15:04:05")
}
