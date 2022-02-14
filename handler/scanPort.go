package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"main/database"
	"os/exec"
	"strings"
)

func scanPort(db *sql.DB, domain string, subdomain string, host string, ip string) {
	cmd := exec.Command("naabu", "-host", ip, "-exclude-cdn", "-silent", "-json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	stringout := string(out)
	var ports []string = strings.Split(stringout, "\n")
	for _, port := range ports {
		if port != "" {
			var data map[string]interface{}
			err := json.Unmarshal([]byte(port), &data)
			if err != nil {
				panic(err)
			}
			openport := data["port"].(float64)
			port := fmt.Sprintf("%v", openport)
			//fmt.Println(host, ip, openport)
			existe := database.Exists(database.ConnectDB(), host, ip, port)
			//fmt.Println(existe)
			if !existe {
				database.InsertHosts(database.ConnectDB(), domain, host, ip, port)
			} else {
				fmt.Println("ya existe")
			}
		}
	}
}
