package main

func main() {
	Start(StartInstructions{
		WriteMode:  "counter",
		WriteLimit: 0,
		StartKey:   "!S",
		EndKey:     "!E",
	})

	In("test", "this is a pie")
}
