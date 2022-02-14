package handler

import (
	"database/sql"
	"fmt"
	"log"
	"main/database"
	"os/exec"
	"strings"
)

var nucleiTemplates string = "/home/jorge/nuclei-templates/"

func ScanVulns(db *sql.DB, subdomain string, port string) {
	cmd := exec.Command("nuclei", "-u", subdomain+":"+port, "--silent", "-c", "800", "-rl", "500", "-t", nucleiTemplates, "-json", "resultado.json") //nuclei --silent -c 800 -rl 500 -t /home/jorge/nuclei-templates/ -json resultado.json
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err, "Error al escanear", subdomain+":"+port)
		log.Fatal(err)
	}
	stringout := string(out)
	var vulns []string = strings.Split(stringout, "\n")
	vulnstring := ParserVuln(vulns)
	database.UpdateVuln(database.ConnectDB(), subdomain, port, vulnstring)

}
