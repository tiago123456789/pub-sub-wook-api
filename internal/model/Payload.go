package model

type Payload struct {
	Event string                 `json:"event"`
	Token string                 `json:"token"`
	Data  map[string]interface{} `json:"data"`
}
