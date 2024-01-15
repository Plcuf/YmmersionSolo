package controller

import (
	InitStruct "YmmersionSolo/backend"
	InitTemp "YmmersionSolo/temps"
	"fmt"
	"net/http"
	"os"
)

// Fonction de la page index pour avoir les recommandations aléatoires
func Accueil(w http.ResponseWriter, r *http.Request) {
	InitStruct.ListeCharacter, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	//execution du templates index.html
	InitTemp.Temp.ExecuteTemplate(w, "index", InitStruct.ListeCharacter)
}

// Fonction de la page du blogs
func Detail(w http.ResponseWriter, r *http.Request) {
	InitStruct.ListeCharacter, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}
	//execution du template Detail.html
	InitTemp.Temp.ExecuteTemplate(w, "Create", InitStruct.ListeCharacter)
}
