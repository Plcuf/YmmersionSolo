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
	http.HandleFunc("/creation/treatment", ctrl.Treatment)
	http.HandleFunc("/delete", ctrl.Delete)
	http.HandleFunc("/update", ctrl.Update)
	http.HandleFunc("/update/treatment", ctrl.Treatment2)

	//Pour relier les assets(img/fonts/css) aux templates
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	//Le lien d'ou est lancé le serveur
	fmt.Println("(http://localhost:8081/index) - Server started on port:8081")
	http.ListenAndServe("localhost:8081", nil)
	fmt.Println("Server closed")
}
