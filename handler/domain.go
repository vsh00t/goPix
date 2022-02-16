package handler

import (
	"fmt"
	"main/database"
	"main/ui"
)

func CreateDomain() (domain string) {
	database.InitDB(database.ConnectDB())
	dom := database.GetDomain(database.ConnectDB())
	for dom == "" {
		ui.Colorize()
		fmt.Println("No existen dominios en la Base de datos.")
		fmt.Println("Ingrese el dominio a escanear: ")
		fmt.Scanln(&dom)
		if dom != "" {
			database.InsertDomain(database.ConnectDB(), dom)
			dom = database.GetDomain(database.ConnectDB())
		}
	}
	ui.Inicio(dom)
	ListSubdomains(dom)
	return dom
}
