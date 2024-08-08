package entities

type KafkaMcMessage struct {
	Duration int64  `json:"duration"`
	Nickname string `json:"Nickname"`
	Service  string `json:"service"`
}
