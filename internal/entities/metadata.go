package entities

type Metadata struct {
	Duration int64  `json:"duration"`
	Nickname string `json:"nickname"`
	Price    string `json:"price,omitempty"`
	Service  string `json:"service"`
}
