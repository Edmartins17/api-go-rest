package models

type Personality struct {
	Id      int    `jdon:if`
	Name    string `json:name`
	History string `json:history`
}

var Personalities []Personality
