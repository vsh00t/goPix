//by vsh00t
package handler

import (
	"database/sql"
	"fmt"
	"main/database"
	"main/ui"
	"os"
	"os/exec"
)

var nucleiTemplates string = "/home/jorge/nuclei-templates/"

func ScanVulns(db *sql.DB, subdomain string, port string, domain string) {
	fmt.Println("")
	fmt.Println("Escaneando", subdomain+":"+port)
	fmt.Println("")
	cmd := "echo " + subdomain + "| httpx --silent -rl 500 -ports " + port + " | nuclei --silent -rl 500 -c 800 -nts -t " + nucleiTemplates + " -o " + subdomain + "_" + port + ".txt"
	ui.ProgBar()
	exec.Command("/bin/bash", "-c", cmd).CombinedOutput()
	//open file and show content
	file := subdomain + "_" + port + ".txt"
	dat, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	database.UpdateVuln(database.ConnectDB(), subdomain, port, string(dat))

}
