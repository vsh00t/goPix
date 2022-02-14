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

var ip string
var domain string
var subdomain string
var port string

//var stringout string
var host string
var nucleiTemplates string = "/Users/moyapj/nuclei-templates/"

const BDDFile = "data.db"

//function to check error
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func nslookup(host string) {
	a, _ := net.LookupIP(host)
	for _, v := range a {
		ip := v.String()
		var d []string = strings.Split(ip, " ")
		fmt.Println(d)
	}

}

//funtion connect to database
func connect() *sql.DB {
	// Open the database
	db, err := sql.Open("sqlite3", BDDFile)
	checkErr(err)
	return db
}

//function to get Domain
func getdomain(db *sql.DB) {
	// Select data
	rows, err := db.Query("SELECT domain FROM program WHERE activo = 1 LIMIT 1")
	checkErr(err)
	for rows.Next() {
		rows.Scan(&domain)
	}
}

//function to create table
func init_database(db *sql.DB) {
	// Create table
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS program (id INTEGER PRIMARY KEY, domain TEXT, activo INTEGER)")
	checkErr(err)
	stmt.Exec()
	stmt, err = db.Prepare("CREATE TABLE IF NOT EXISTS targets (id INTEGER PRIMARY KEY, domain TEXT, subdomain TEXT, ip TEXT, port TEXT, vuln TEXT, activo INTEGER)")
	checkErr(err)
	stmt.Exec()
}

//function to insert ip address
func insertIps(db *sql.DB, domain, subdomain string, ip string, port string) {
	// Insert data
	stmt, err := db.Prepare("SELECT id FROM targets WHERE subdomain = ? AND port = ? LIMIT 1")
	checkErr(err)
	rows, err := stmt.Query(subdomain, port)
	checkErr(err)
	defer rows.Close()
	if rows.Next() {
		fmt.Println("Ya existe")
	} else {
		stmt, err := db.Prepare("INSERT INTO targets (domain, subdomain, ip, port, activo, vuln) VALUES (?, ?, ?, ?, ?, ?)")
		checkErr(err)
		res, err := stmt.Exec(domain, subdomain, ip, port, 1, "None")
		checkErr(err)
		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println(id)
	}
}

//function to update vulnerability
func updateVulns(db *sql.DB, i string, j string, k string) {
	// Update data
	stmt, err := db.Prepare("UPDATE targets SET vuln = ?, activo = ? WHERE subdomain = ? AND port = ?")
	checkErr(err)
	if i != "" {
		res, err := stmt.Exec(i, 0, j, k)
		checkErr(err)
		affected, err := res.RowsAffected()
		checkErr(err)
		fmt.Println(affected, err, "No se pudo actualizar la vulnerabilidad")
	} else {
		res, err := stmt.Exec("None", 0, j, k)
		checkErr(err)
		affected, err := res.RowsAffected()
		checkErr(err)
		fmt.Println(affected, err, "No se pudo actualizar")
	}
}

func checkVuln(db *sql.DB, subdomain string, port string) {
	// Select data
	rows, err := db.Query("SELECT  FROM targets WHERE subdomain = ? AND port = ? LIMIT 1")
	checkErr(err)
	for rows.Next() {
		rows.Scan(&subdomain)
	}
}

func scanVulns(db *sql.DB, subdomain string, port string) {
	cmd := exec.Command("nuclei", "-u", subdomain+":"+port, "--silent", "-c", "800", "-rl", "500", "-t", nucleiTemplates, "-json", "resultado.json") //nuclei --silent -c 800 -rl 500 -t /home/jorge/nuclei-templates/ -json resultado.json
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err, "Error al escanear", subdomain+":"+port)
		log.Fatal(err)
	}
	stringout := string(out)
	var vulns []string = strings.Split(stringout, "\n")
	for _, vuln := range vulns {
		fmt.Println(vuln)
		updateVulns(db, vuln, subdomain, port)
	}
	fmt.Sprintln("Escaneo terminado en", subdomain+":"+port)
}

//function to insert data
func insertData(db *sql.DB, domain string, activo int) {
	// Insert data
	stmt, err := db.Prepare("INSERT INTO program (domain, activo) values(?,?)")
	checkErr(err)
	res, err := stmt.Exec(domain, activo)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

//function to update data
func updateData(db *sql.DB, domain string, activo int) {
	// Update data
	stmt, err := db.Prepare("UPDATE program SET domain = ?, activo = ? WHERE id = 1")
	checkErr(err)
	stmt.Exec(domain, activo)
}

//function to select data
func selectData(db *sql.DB) {
	// Select data
	rows, err := db.Query("SELECT id, domain, activo FROM program")
	checkErr(err)
	for rows.Next() {
		var id int
		var domain string
		var activo int
		err = rows.Scan(&id, &domain, &activo)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(domain)
		fmt.Println(activo)
	}
}

func algo() {

	//connect to database
	db := connect()
	getdomain(db)
	//create table
	//createTable(db)
	//insert data
	insertData(db, "google.com", 1)
	//select data
	selectData(db)
	//update data
	updateData(db, "google.com", 1)
	//select data
	selectData(db)
}

func listSubdomains(db *sql.DB, domain string) {
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
			a, _ := net.LookupIP(host)
			for _, v := range a {
				ip := v.String()
				var d []string = strings.Split(ip, " ")

				for _, v := range d {

					if strings.Contains(v, ".") {
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
								s := fmt.Sprintf("%v", openport)
								insertIps(db, domain, host, v, s)
							}

						}

					}

				}

			}

		}

	}
}

//function to select data
func checkVulns(db *sql.DB, domain string) {
	var i = []string{}
	// Select data
	rows, err := db.Query("SELECT subdomain,port,ip FROM targets WHERE activo = 1")
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var subdomain string
		var port string
		var ip string
		err = rows.Scan(&subdomain, &port, &ip)
		checkErr(err)
		i = append(i, subdomain+":"+port+":"+ip)
	}
	fmt.Println(i)
	defer rows.Close()
	//for loop length i
	for _, host := range i {
		checkErr(err)
		host := strings.Split(host, ":")
		fmt.Println("Escaneando", host[0]+":"+host[1])
		cmd := exec.Command("nuclei", "-u", host[0]+":"+host[1], "--silent", "-c", "800", "-rl", "500", "-t", nucleiTemplates, "-json", "resultado.json") //nuclei --silent -c 800 -rl 500 -t /home/jorge/nuclei-templates/ -json resultado.json
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err, "Error al escanear", host[0]+":"+host[1])
			log.Fatal(err)
		}
		stringout := string(out)
		fmt.Println(stringout)
		fmt.Println("resultado:", stringout, "activo:", 0, "subdomain", host[0], "port", host[1], "ip", host[2])
		stmt, err := db.Prepare("UPDATE targets SET vuln = ?, activo = ? WHERE ip = ? AND port = ?")
		checkErr(err)
		res, err := stmt.Exec(stringout, 0, host[2], host[1])
		checkErr(err)
		affected, err := res.RowsAffected()
		checkErr(err)
		fmt.Println(affected, err, "No se pudo actualizar la vulnerabilidad")
		//subdomain := strings.Split(host, ":")

	}
}

func main() {
	db, err := sql.Open("sqlite3", BDDFile)
	db.SetMaxOpenConns(1)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	init_database(db)
	getdomain(db)
	if domain == "" {
		panic("No hay dominio activo cree uno en la tabla program")
	}
	listSubdomains(db, domain)
	checkVulns(db, domain)

} // end main
