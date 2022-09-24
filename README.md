# go-markov

```go
go get "github.com/ActuallyGiggles/go-markov"
```

Import package.
```go
import "go-markov/markov"
```

Initialize Markov.
```go
i := markov.StartInstructions{
		Workers:       5,
		WriteInterval: 10,
		IntervalUnit:  "minutes",
		StartKey:      "start",
		EndKey:        "end",
}

markov.Start(i)
```

Add to Markov queue. Important: markov will only write to a chain after the write interval has ticked.
```go
markov.Input("test", "This is a test.")
```

Output a Markov output.
```go
oi := markov.OutputInstructions{
		Method: "TargetedBeginning",
		Chain:  "test",
		Target: "This",
}

output, problem := markov.Output(oi)

fmt.Println(output)
fmt.Println(problem)
```

Get information on workers.
```go
ws := markov.WorkersStats()
```
