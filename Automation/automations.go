package automations

import (
	haApi "github.com/DanielLatimerUns/GoHomeAssistantAutomationRunner/HaApi"
	"time"
)

func frontDoorOpened() {
	entity := haApi.GetHaEntity("binary_sensor.front_door_sensor_access_control_window_door_is_open")

	if entity.State == "on" {
		haApi.TurnLightOn("light.downstairs_hallway_light")
	}
}

func frontDoorMotion() {
	entity := haApi.GetHaEntity("event.front_door_motion")

	lastMotion, err := time.Parse("2006-01-02T15:04:05.000-07:00", entity.State)
	handleErr(err)

	lastMotionWithBuffer := lastMotion.Add(time.Second * 2)

	if time.Now().Before(lastMotionWithBuffer) {
		haApi.TurnLightOff("light.downstairs_hallway_light")
		time.Sleep(time.Second * 2)
		haApi.TurnLightOn("light.downstairs_hallway_light")
	}
}
