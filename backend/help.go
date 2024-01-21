package backend

import (
	"encoding/json"
	"fmt"
	"os"
)

// Fonction pour mettre le JSON dans une struct
func ReadJSON() ([]Character, error) {
	jsonFile, err := os.ReadFile("JSON/bdd.json")
	if err != nil {
		fmt.Println("Error reading", err.Error())
	}
	var jsonData []Character
	err = json.Unmarshal(jsonFile, &jsonData)
	return jsonData, err
}

// Fonction pour savoir si l'id existe déjà
func IdAlreadyExists(nb string) bool {
	for i := 0; i < len(ListeCharacter); i++ {
		if ListeCharacter[i].Id == nb {
			return true
		}
	}
	return false
}

// Fonction pour générer un Id disponible
func GenerateID() string {
	if !IdAlreadyExists(string(len(ListeCharacter) + 1)) {
		return string(len(ListeCharacter) + 1)
	} else {
		t := LstIDSuppr[0]
		if len(LstIDSuppr) > 1 {
			LstIDSuppr = LstIDSuppr[1:]
		} else {
			LstIDSuppr = []int{}
		}
		return string(t)
	}
}

func GetCharById(Id string) {
	filepath := "JSON/bdd.json"
	contenu, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("pas cool : ", err)
	}
	var Personnage []Character
	err = json.Unmarshal(contenu, &Personnage)
	if err != nil {
		fmt.Println("pas cool : ", err)
	}
	var index = -1
	for i, char := range Personnage {
		fmt.Println(char.Id, Id, i)
		if char.Id == Id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("pa cool :(")
	}
	Personnage = append(Personnage[:index], Personnage[index+1:]...)
	updatedData, err := json.Marshal(Personnage)
	if err != nil {
		fmt.Println("pas cool : ", err)
	}
	err = os.WriteFile("JSON/bdd.json", updatedData, 0644)
	if err != nil {
		fmt.Println("pas cool : ", err)
	}
}
