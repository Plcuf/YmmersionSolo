package backend

type Character struct {
	Name        string `json:"Name"`
	Id          string `json:"id"`
	Description string `json:"description"`
	Age         string `json:"Age"`
	Race        string `json:"race"`
	Classe      string `json:"classe"`
}

var Char Character
var ListeCharacter []Character
var LstIDSuppr []int
