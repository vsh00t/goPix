//by vsh00t
package main

import (
	"main/database"
	"main/handler"
	"main/ui"
	"strings"
)

func main() {
	ui.Colorize()
	domain := handler.CreateDomain()
	all := database.CuentaAllSubdomains(database.ConnectDB(), domain)
	ui.NumAllSubdomains(all)
	ui.InicioScaneo()
	hosts := database.IsActiveScan(database.ConnectDB(), domain)
	for _, host := range hosts {
		host := strings.Split(host, ":")
		go handler.IsWeb(database.ConnectDB(), domain, host[1], host[2])
		handler.ScanVulns(database.ConnectDB(), host[1], host[2], domain)
	}

}
