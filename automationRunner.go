package main

import (
	"github.com/DanielLatimerUns/GoHomeAssistantAutomationRunner/automation"
	"time"
)

func main() {
	// start main automation loop
	for {
		// loop executes every 2 seconds
		time.Sleep(1 * time.Second)
		run()
	}
}

func run() {
	automationsToRun := automation.ConfigureAutomations().Automations

	for _, a := range automationsToRun {
		// run each automation in its own go routine
		go a.Run()
	}
}
