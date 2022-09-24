# go-markov

```go
go get "github.com/ActuallyGiggles/go-markov"
```

Import package
```go
import "go-markov/markov"
```

Initialize Markov
```go
i := markov.StartInstructions{
		Workers:       5,
		WriteInterval: 1,
		IntervalUnit:  "hours",
		StartKey:      "start",
		EndKey:        "end",
}

markov.Start(i)
```

Add to Markov queue
```go
markov.Input("test", "This is a test.")
```

Output a Markov output
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

Get information on workers
```go
ws := markov.WorkersStats()
```
