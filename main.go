package main

import "os"

// go run main.go run [cmd] [args]

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Please help!")
	}
}

