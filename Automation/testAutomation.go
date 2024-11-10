package automation

import haApi "github.com/DanielLatimerUns/GoHomeAssistantAutomationRunner/HaApi"

func testAutomation() {
	entity := haApi.GetHaEntity("light.sofa_light")

	if entity.State == "on" {
		haApi.TurnLightOff(entity.EntityID)
	} else {
		haApi.TurnLightOn(entity.EntityID)
	}
}
