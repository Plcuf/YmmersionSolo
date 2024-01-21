package controller

import (
	InitStruct "YmmersionSolo/backend"
	InitTemp "YmmersionSolo/temps"
	"encoding/json"
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

func Treatment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newChar := InitStruct.Character{
		Name:        r.FormValue("Name"),
		Id:          InitStruct.GenerateID(), //Je génére un id pas utilisé
		Description: r.FormValue("description"),
		Age:         r.FormValue("Age"),
		Classe:      r.FormValue("classe"),
		Race:        r.FormValue("race"),
	}
	filename := "JSON/bdd.json"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("pas cool : ", err)
	}

	err = json.Unmarshal(data, &InitStruct.ListeCharacter)
	if err != nil {
		fmt.Println("pas cool : ", err)
	}

	InitStruct.ListeCharacter = append(InitStruct.ListeCharacter, newChar)

	updatedData, err := json.MarshalIndent(InitStruct.ListeCharacter, "", " ")
	if err != nil {
		fmt.Println("pas cool : ", err)
	}

	err = os.WriteFile(filename, updatedData, 0644)
	if err != nil {
		fmt.Println("pas cool : ", err)
	}

	fmt.Println("séboon")

	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	InitStruct.GetCharById(Id)
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

func Update(w http.ResponseWriter, r *http.Request) {
	date, err := os.ReadFile("JSON/bdd.json")
	if err != nil {
		fmt.Println("pas cool : ", err)
	}
	var personnage []InitStruct.Character
	if err := json.Unmarshal(date, &personnage); err != nil {
		fmt.Println("pa cool : ", err)
	}
	Id := r.URL.Query().Get("id")
	var selectedPers InitStruct.Character
	for _, char := range personnage {
		if char.Id == Id {
			selectedPers = char
			break
		}
	}
	if selectedPers.Id == "" {
		fmt.Println("pa cool :(")
	}
	InitTemp.Temp.ExecuteTemplate(w, "update", selectedPers)
}

func Treatment2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newChar := InitStruct.Character{
		Name:        r.FormValue("Name"),
		Id:          r.URL.Query().Get("Id"),
		Description: r.FormValue("description"),
		Age:         r.FormValue("Age"),
		Classe:      r.FormValue("classe"),
		Race:        r.FormValue("race"),
	}
	filename := "JSON/bdd.json"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("pas cool : ", err)
	}

	err = json.Unmarshal(data, &InitStruct.ListeCharacter)
	if err != nil {
		fmt.Println("pas cool : ", err)
	}

	var index = -1
	for i, char := range InitStruct.ListeCharacter {
		if char.Id == newChar.Id {
			index = i
			break
		}
	}
	if index != -1 {
		InitStruct.ListeCharacter[index] = newChar
		updatedData, err := json.MarshalIndent(InitStruct.ListeCharacter, "", " ")
		if err != nil {
			fmt.Println("pas cool : ", err)
		}

		err = os.WriteFile(filename, updatedData, 0644)
		if err != nil {
			fmt.Println("pas cool : ", err)
		}
		fmt.Println("séboon")
	} else {
		fmt.Println("sépaboon")
	}

	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
