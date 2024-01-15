package controller

import (
	InitStruct "YmmersionSolo/backend"
	InitTemp "YmmersionSolo/temps"
	"fmt"
	"io"
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
func InitAdd(w http.ResponseWriter, r *http.Request) {
	InitStruct.ListeCharacter, err = InitStruct.ReadJSON() //Met le fichier JSON dans ma struct
	if err != nil {
		fmt.Println("Error encodage ", err.Error())
		os.Exit(1)
	}

	//Prend les valeurs demandés
	InitStruct.Char.Name = r.FormValue("name")
	InitStruct.Char.Id = InitStruct.GenerateID() //Je génére un id pas utilisé
	InitStruct.Char.Description = r.FormValue("description")
	InitStruct.Char.Age = r.FormValue("age")
	InitStruct.Char.Genre = r.FormValue("genre")
	//Prend les données ne dépassant cette taille (pout l'image)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, errFile := r.FormFile("Image") //Récupère le fichier image
	if errFile != nil {
		http.Error(w, errFile.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	InitStruct.Char.Image = handler.Filename
	filepath := "./assets/img/" + handler.Filename //Chemin où mettre le fichier
	f, _ := os.Create(filepath)
	defer f.Close()
	io.Copy(f, file) //Met l'image au chemin donnée
	//je ne met pas dans ma struct

	InitStruct.ListeCharacter = append(InitStruct.ListeCharacter, InitStruct.Char)
	InitStruct.EditJSON(InitStruct.ListeCharacter) //Met les données dans le JSON
	http.Redirect(w, r, "/admin", http.StatusMovedPermanently)
}
