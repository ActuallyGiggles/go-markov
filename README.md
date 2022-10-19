# go-markov
Simple Markov chain input and output. :)

Get package.
```go
go get "github.com/ActuallyGiggles/go-markov"
```

```go
package main

import "github.com/ActuallyGiggles/go-markov"

func main() {
	i := markov.StartInstructions{
			Workers:       5,
			
			StartKey:      "start",
			EndKey:        "end",
	}

	markov.Start(i)

	markov.Input("test", "This is a test.")

	oi := markov.OutputInstructions{
			Method: "TargetedBeginning",
			Chain:  "test",
			Target: "This",
	}

	output, problem := markov.Output(oi)
	
	fmt.Println(output)
	fmt.Println(problem)
}
