package _struct

type Metrics struct {
	CreatedAt    int64          `json:"created_at"`
	Config       map[string]interface{}          `json:"config"`
	Id           string          `json:"id"`
	Service      interface{}          `json:"service"`
	Name         string          `json:"name"`
	Protocols    interface{}     `json:"protocols"`
	Enabled      bool            `json:"enabled"`
	RunOn        interface{}          `json:"run_on"`
	Consumer     map[string]string          `json:"consumer"`
	Route        interface{}          `json:"route"`
	Tags         interface{}          `json:"tags"`
}

type AllData struct {
	Next      string          `json:"next"`
	Data      []*Metrics          `json:"data"`
}