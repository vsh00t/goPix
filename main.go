//by vsh00t
package main

import (
	"strings"

	"github.com/vsh00t/goPix/database"
	"github.com/vsh00t/goPix/handler"
	"github.com/vsh00t/goPix/ui"
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
		handler.ScanVulns(database.ConnectDB(), host[1], host[2], domain)
	}

}
