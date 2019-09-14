package model

type Config struct {
	URL    string `json:"url"`
	Target string `json:"target"`
	Minute int64  `json:"min"`
}
