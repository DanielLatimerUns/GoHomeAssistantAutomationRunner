package automations

import (
	"github.com/google/uuid"
	"log"
)

type AutomationEntity struct {
	id   string
	name string
	fn   func()
}

func GetAutomations() []AutomationEntity {
	return []AutomationEntity{
		{id: "3e5bf3ef-6621-4026-8ce3-7a3f35cc6631", fn: frontDoorOpened, name: "Front Door Opened"},
		{id: "e0b2e6b1-e875-40f4-9cbc-ce00f153560f", fn: frontDoorMotion, name: "Front Door Motion"},
	}
}

func (a AutomationEntity) Execute(runId string) {
	log.Println("Starting Automation: " + a.name + " (" + a.id + ") (" + runId + ")")
	a.fn()
	log.Println("Completed Automation: " + a.name + " (" + a.id + ") (" + runId + ")")
}

func newUUID() uuid.UUID {
	u := uuid.New()
	return u
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
