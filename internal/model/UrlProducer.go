package model

type UrlProducer struct {
	Id      int64  `json:"id"`
	Enabled bool   `json:"enabled"`
	Key     string `json:"key"`
}
