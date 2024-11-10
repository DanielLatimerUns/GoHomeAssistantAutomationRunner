package automation

import (
	"github.com/google/uuid"
	"log"
)

type Automation struct {
	id   uuid.UUID
	name string
	fn   func()
}

type Collection struct {
	Automations []Automation
}

func ConfigureAutomations() Collection {
	return Collection{
		Automations: []Automation{
			{id: newUUID(), fn: testAutomation, name: "Test Automation"},
		},
	}
}

func (a Automation) Run() {
	log.Println("Starting Automation: " + a.name + " (" + a.id.String() + ")")
	a.fn()
	log.Println("Completed Automation: " + a.name + " (" + a.id.String() + ")")
}

func newUUID() uuid.UUID {
	u := uuid.New()
	return u
}
