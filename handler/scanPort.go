//by vsh00t
package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
				//logging
				log.Default()
			}
		}
	}
}

func IsWeb(db *sql.DB, domain string, subdomain string, port string) {
	protocol := []string{"http://", "https://"}
	var state []string

	for _, protocolo := range protocol {
		resp, err := http.Get(protocolo + subdomain + ":" + port)
		if err != nil {
			state = append(state, "0")
			break
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			state = append(state, "200")
		} else {
			state = append(state, "0")
		}
	}
	if len(state) == 2 {
		if state[0] == "200" && state[1] == "200" {
			database.UpdateHttp(database.ConnectDB(), subdomain, port, 1)
		} else if state[0] == "200" && state[1] == "0" {
			database.UpdateHttp(database.ConnectDB(), subdomain, port, 2)
		} else if state[0] == "0" && state[1] == "200" {
			database.UpdateHttp(database.ConnectDB(), subdomain, port, 3)
		} else {
			database.UpdateHttp(database.ConnectDB(), subdomain, port, 0)
		}
	}
}

