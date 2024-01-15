package backend

type Character struct {
	Name        string `json:"name"`
	Id          int    `json:"id"`
	Description string `json:"description"`
	Age         string `json:"age"`
	Genre       string `json:"genre"`
	Image       string `json:"image"`
}

var Char Character
var ListeCharacter []Character
var LstIDSuppr []int
