package main

import (
	"github.com/DanielLatimerUns/GoHomeAssistantAutomationRunner/Automation"
	"github.com/google/uuid"
	"log"
	"time"
)

func main() {
	// start main automation loop
	for {
		// loop executes every 1 seconds
		time.Sleep(1 * time.Second)
		// execute each automation run in its own go routine to stop slow automations blocking new runs
		go run()
	}
}

func run() {
	automationsToRun := automations.GetAutomations()
	runId := uuid.New().String()

	log.Print("Starting Run: " + "(" + runId + ")")
	for _, a := range automationsToRun {
		// execute each automation in its own go routine
		go a.Execute(runId)
	}
}
