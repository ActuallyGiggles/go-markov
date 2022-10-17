package main

func main() {
	Start(StartInstructions{
		WriteMode:  "counter",
		WriteLimit: 1,
		StartKey:   "!S",
		EndKey:     "!E",
	})

	In("test", "this is a test")
}
