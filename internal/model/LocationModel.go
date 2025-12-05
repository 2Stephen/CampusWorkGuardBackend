package model

type Location struct {
	Adcode    string     `json:"adcode"`
	Name      string     `json:"name"`
	Level     string     `json:"level"`
	Districts []District `json:"districts"`
}
type District struct {
	Adcode string `json:"adcode"`
	Name   string `json:"name"`
	Level  string `json:"level"`
}
