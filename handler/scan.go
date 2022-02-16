package handler

import (
	"database/sql"
	"fmt"
	"log"
	"main/database"
	"main/ui"
	"os/exec"
	"strings"
)

var nucleiTemplates string = "/home/jorge/nuclei-templates/"

func ScanVulns(db *sql.DB, subdomain string, port string, domain string) {
	fmt.Println("")
	fmt.Println("Escaneando", subdomain+":"+port)
	fmt.Println("")
	cmd := "echo " + subdomain + "| httpx --silent -ports " + port + " | nuclei --silent -c 800 -rl 500 -t " + nucleiTemplates + " -json resultado.json"
	ui.ProgBar()
	out, err := exec.Command("/bin/bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Println(err, "Error al escanear", subdomain+":"+port)
		log.Fatal(err)
	}
	stringout := string(out)
	var vulns []string = strings.Split(stringout, "\n")
	//fmt.Println(vulns)
	vulnstring := ParserVuln(vulns)
	fmt.Printf("%s\n", vulnstring)
	database.UpdateVuln(database.ConnectDB(), subdomain, port, vulnstring)

}
