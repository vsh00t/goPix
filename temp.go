package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"os/exec"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

// select domain from program where activo = 1

var subdomain string
var port string
var domain string

func getDom() {

	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	rows, err := database.Query("SELECT subdomain,port FROM targets WHERE activo = 1 LIMIT 1")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		rows.Scan(&subdomain, &port)
	}
}

func setInactive(i string, j string) {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	fmt.Println(i, j)
	statement, err := database.Prepare("UPDATE targets SET activo = 0 WHERE subdomain = ? AND port = ?")
	if err != nil {
		panic(err)
	}
	statement.Exec(i, j)

}

func updateVulns(i string, j string, k string) {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	statement, err := database.Prepare("UPDATE targets SET vuln = ? WHERE subdomain = ? AND port = ?")
	if err != nil {
		panic(err)
	}
	statement.Exec(i, j, k)

}

func scanVulns() {
	cmd := exec.Command("nuclei", "-u", subdomain+":"+port, "--silent", "-c", "800", "-rl", "500", "-t", "/home/jorge/nuclei-templates/", "-json", "resultado.json") //nuclei --silent -c 800 -rl 500 -t /home/jorge/nuclei-templates/ -json resultado.json
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	stringout := string(out)
	var vulns []string = strings.Split(stringout, "\n")
	for _, vuln := range vulns {
		if vuln != "" {
			fmt.Println(vuln)
			updateVulns(vuln, subdomain, port)

		}
	}
}

func main() {
	getDom()
	fmt.Printf("%s:%s\n", subdomain, port)

	//for results query sql
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	rows, err := database.Query("SELECT subdomain,port FROM targets WHERE activo = 1")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		rows.Scan(&subdomain, &port)
		fmt.Printf("%s:%s\n", subdomain, port)
		scanVulns()
		setInactive(subdomain, port)
	}

}
