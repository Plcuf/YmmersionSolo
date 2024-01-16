package backend

import (
	"encoding/json"
	"fmt"
	"os"
)

// Fonction pour modifié le JSON
func EditJSON(ModifiedChar []Character) {

	modifiedJSON, errMarshal := json.Marshal(ModifiedChar)
	if errMarshal != nil {
		fmt.Println("Error encodage ", errMarshal.Error())
		return
	}
	// Écrire le JSON modifié dans le fichier
	err := os.WriteFile("JSON/bdd.json", modifiedJSON, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier JSON modifié:", err)
		return
	}
}

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
