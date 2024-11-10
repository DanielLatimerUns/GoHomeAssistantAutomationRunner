package haApi

import (
	"encoding/json"
	"io"
)

type HaEntity struct {
	Attributes  map[string]interface{} `json:"attributes"`
	EntityID    string                 `json:"entity_id"`
	LastChanged string                 `json:"last_changed"`
	LastUpdated string                 `json:"last_updated"`
	State       string                 `json:"state"`
}

func GetHaEntity(entityId string) HaEntity {
	var url = haUrl + "states/" + entityId

	resp, err := httpGet(url)
	handleErr(err)

	body, err := io.ReadAll(resp.Body)
	handleErr(err)

	var entity HaEntity
	err = json.Unmarshal(body, &entity)
	handleErr(err)

	return entity
}

func SetHaEntityState(entity HaEntity) {
	payload, err := json.Marshal(entity)
	handleErr(err)

	resp, err := httpPost("states"+entity.EntityID, payload)
	handleErr(err)
	handleErr(resp.Body.Close())
}
