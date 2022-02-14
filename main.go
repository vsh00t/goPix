package main

import (
	"fmt"
	"main/database"
	"main/handler"
	"strings"
)

func main() {
	database.InitDB(database.ConnectDB())
	domain := database.GetDomain(database.ConnectDB())
	if domain == "" {
		panic("No existen dominios activos, agregue uno.")
	} else {
		fmt.Println("Iniciando el descubrimiento de subdominios y escaneo de puertos del dominio: ", domain)
		handler.ListSubdomains(domain)
	}
	fmt.Println("Iniciando el escaneo de vulnerabilidades")
	hosts := database.IsActiveScan(database.ConnectDB(), domain)

	for _, host := range hosts {
		//fmt.Println(host)
		//split hots by :
		host := strings.Split(host, ":")
		fmt.Println("Escaneando: " + host[1] + ":" + host[2])
		handler.ScanVulns(database.ConnectDB(), host[1], host[2])
	}

}
