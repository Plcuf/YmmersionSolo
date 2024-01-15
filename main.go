package main

import (
	r "YmmersionSolo/routeur" //Route vers mes routes
	t "YmmersionSolo/temps"   //Route vers mes templates
)

func main() {
	t.InitTemplate() //Initialise mes templates
	r.InitServe()    //Initialise mes routes / assets et lance le serveur
}
