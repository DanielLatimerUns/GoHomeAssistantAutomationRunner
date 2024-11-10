package automation

import haApi "github.com/DanielLatimerUns/GoHomeAssistantAutomationRunner/HaApi"

func frontDoorOpened() {
	entity := haApi.GetHaEntity("binary_sensor.front_door_sensor_access_control_window_door_is_open")

	if entity.State == "on" {
		haApi.TurnLightOn("light.downstairs_hallway_light")
	}
}
