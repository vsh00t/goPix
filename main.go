package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

//var domain = "sygma-ec.com"
var domain string
var subdomain string
var port string

func nslookup(host string) {
	a, _ := net.LookupIP(host)
	for _, v := range a {
		ip := v.String()
		var d []string = strings.Split(ip, " ")
		fmt.Println(d)
	}

}

func init_database() {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS program (id INTEGER PRIMARY KEY, domain TEXT, activo INTEGER)")
	if err != nil {
		panic(err)
	}
	statement.Exec()
	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS targets (domain INTEGER PRIMARY KEY, subdomain TEXT, ip TEXT, port TEXT, vuln TEXT)")
	if err != nil {
		panic(err)
	}

	statement.Exec()
	statement, err = database.Prepare("INSERT INTO program (domain, activo) VALUES (?, ?)")
	if err != nil {
		panic(err)
	}
	statement.Exec("example.ec", 1)

}

func getdomain() {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()
	rows, err := database.Query("SELECT domain FROM program WHERE activo = 1 LIMIT 1")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		rows.Scan(&domain)
	}
}

func insertIps(subdomain string, ip string, port string) {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	statement, err := database.Prepare("INSERT INTO targets (subdomain, ip, port, activo) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	statement.Exec(subdomain, ip, port, 1)
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
	init_database()
	getdomain()
	if domain == "" {
		panic("No hay dominio activo")
	}

	fmt.Println(domain)
	cmd := exec.Command("subfinder", "-d", domain, "-silent", "-oJ")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	stringout := string(out)
	var subdomains []string = strings.Split(stringout, "\n")

	for _, subdomain := range subdomains {
		if subdomain != "" {
			var data map[string]interface{}
			err := json.Unmarshal([]byte(subdomain), &data)
			if err != nil {
				panic(err)
			}

			host := data["host"].(string)
			fmt.Println(host)
			//ip = nslookup(host)
			a, _ := net.LookupIP(host)
			for _, v := range a {
				ip := v.String()
				var d []string = strings.Split(ip, " ")

				//fmt.Println(d)

				for _, v := range d {
					//fmt.Println(v)

					if strings.Contains(v, ".") {
						//fmt.Println("Yes")
						cmd := exec.Command("naabu", "-host", ip, "-exclude-cdn", "-silent", "-json")
						out, err := cmd.CombinedOutput()
						if err != nil {
							log.Fatal(err)
						}
						stringout := string(out)
						var ports []string = strings.Split(stringout, "\n")
						//fmt.Println(ports)
						for _, port := range ports {
							if port != "" {
								var data map[string]interface{}
								err := json.Unmarshal([]byte(port), &data)
								if err != nil {
									panic(err)
								}
								//fmt.Println(data)
								openport := data["port"].(float64)
								s := fmt.Sprintf("%v", openport)
								//fmt.Println(s)
								//fmt.Println(reflect.TypeOf(s))
								insertIps(host, v, s)
							}

						}

					}

				}

			}

		}

	}
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

} // end main
