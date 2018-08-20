package main

import "log"

func main() {

	machine, err := NewMachine("+++[>+++++<-]>.")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	machine.Run()

}
