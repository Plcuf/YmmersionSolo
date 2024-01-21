package controller

import (
	InitStruct "YmmersionSolo/backend"
	InitTemp "YmmersionSolo/temps"
	"fmt"
	"net/http"
	"os"
)

var err error

// Fonction pour les admins
func Admin(w http.ResponseWriter, r *http.Request) {
	InitStruct.ListeCharacter, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	InitTemp.Temp.ExecuteTemplate(w, "Admin", InitStruct.ListeCharacter)
}

// Fonction pour ajouter un blog
func Add(w http.ResponseWriter, r *http.Request) {
	InitTemp.Temp.ExecuteTemplate(w, "Add", InitStruct.ListeCharacter)
}

// Fonction treatment de l'ajout de blog
