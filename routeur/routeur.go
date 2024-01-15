package routeur

import (
	ctrl "YmmersionSolo/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServe() {
	/*Initialisation des routes
	http.HandleFunc("Route actuel, fonction activé")
	Lorsque on se situe sur une route la fonction associé va s'activé */
	http.HandleFunc("/index", ctrl.Accueil)
	http.HandleFunc("/creation", ctrl.Detail)

	//Pour relier les assets(img/fonts/css) aux templates
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	//Le lien d'ou est lancé le serveur
	fmt.Println("(http://localhost:8080/index) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)
	fmt.Println("Server closed")
}
