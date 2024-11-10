package haApi

import "encoding/json"

type HaService struct {
	Name       string `json:"-"`
	ServiceUrl string `json:"-"`
	EntityID   string `json:"entity_id"`
}

func TurnLightOn(entityID string) {
	callService(HaService{EntityID: entityID, ServiceUrl: "light/turn_on"})
}

func TurnLightOff(entityID string) {
	callService(HaService{EntityID: entityID, ServiceUrl: "light/turn_off"})
}

func TurnSwitchOn(entityID string) {
	callService(HaService{EntityID: entityID, ServiceUrl: "switch/turn_on"})
}

func TurnSwitchOff(entityID string) {
	callService(HaService{EntityID: entityID, ServiceUrl: "switch/turn_off"})
}

func callService(service HaService) {
	var url = haUrl + "services/" + service.ServiceUrl

	payload, err := json.Marshal(service)
	handleErr(err)

	resp, err := httpPost(url, payload)
	handleErr(err)
	handleErr(resp.Body.Close())
}
