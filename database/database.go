//by vsh00t
package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

const BDDFile = "data.db"

var domain string

//function to check error OK

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//funtion connect to database OK

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", BDDFile)
	db.SetMaxOpenConns(1)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//funtionto Init DB OK
func InitDB(db *sql.DB) {
	// Create table
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS program (id INTEGER PRIMARY KEY, domain TEXT, activo INTEGER)")
	CheckErr(err)
	stmt.Exec()
	stmt, err = db.Prepare("CREATE TABLE IF NOT EXISTS targets (id INTEGER PRIMARY KEY, domain TEXT, subdomain TEXT, ip TEXT, port TEXT, vuln TEXT, spider TEXT, http INT, activo INTEGER)")
	CheckErr(err)
	stmt.Exec()
	defer db.Close()
}

//function to get Domain OK
func GetDomain(db *sql.DB) string {
	// Select data
	rows, err := db.Query("SELECT domain FROM program WHERE activo = 1 LIMIT 1")
	CheckErr(err)
	for rows.Next() {
		rows.Scan(&domain)
	}
	defer db.Close()
	return domain
}

//function if exists OK
func Exists(db *sql.DB, subdomain string, ip string, port string) bool {
	// Select data
	var domain string
	rows, err := db.Query("SELECT domain FROM targets WHERE subdomain = ? AND ip = ? AND port = ?", subdomain, ip, port)
	CheckErr(err)
	for rows.Next() {
		rows.Scan(&domain)
	}
	defer db.Close()
	if domain == "" {
		return false
	} else {
		return true
	}
}

//function if active OK
func IsActiveScan(db *sql.DB, domain string) (hosts []string) {
	// Select data
	rows, err := db.Query("SELECT domain, subdomain, port FROM targets WHERE activo = ?", 1)
	CheckErr(err)
	for rows.Next() {
		var subdomain string
		var port string
		rows.Scan(&domain, &subdomain, &port)
		//fmt.Println(domain, subdomain, port)
		//print selected data
		hosts = append(hosts, domain+":"+subdomain+":"+port)
	}
	defer db.Close()
	return hosts
}

//function to insert hosts and ip
func InsertHosts(db *sql.DB, domain string, subdomain string, ip string, port string) {
	// Insert data
	stmt, err := db.Prepare("INSERT INTO targets (domain, subdomain, ip, port, vuln, activo) values(?,?,?,?,?,?)")
	CheckErr(err)
	stmt.Exec(domain, subdomain, ip, port, "None", 1)
	CheckErr(err)
	defer db.Close()
}

//function to update data
func UpdateVuln(db *sql.DB, subdomain string, port string, vuln string) {
	// Update data
	//fmt.Println("Insertando en " + subdomain + ":" + port + "La vulnerabilidad:" + vuln)
	stmt, err := db.Prepare("UPDATE targets SET vuln = ?, activo = ? WHERE subdomain = ? AND port = ?")
	CheckErr(err)
	stmt.Exec(vuln, 0, subdomain, port)
	defer db.Close()
}

func InsertDomain(db *sql.DB, domain string) {
	stmt, err := db.Prepare("INSERT INTO program (domain, activo) values(?,?)")
	CheckErr(err)
	stmt.Exec(domain, 1)
	defer db.Close()
}

func CuentaSubdomains(db *sql.DB, domain string) (count int) {
	rows, err := db.Query("SELECT COUNT(DISTINCT subdomain) FROM targets where domain = ?", domain)
	CheckErr(err)
	for rows.Next() {
		rows.Scan(&count)
	}
	defer db.Close()
	return count
}

func CuentaAllSubdomains(db *sql.DB, domain string) (count int) {
	rows, err := db.Query("SELECT COUNT(*) FROM targets where domain = ?", domain)
	CheckErr(err)
	for rows.Next() {
		rows.Scan(&count)
	}
	defer db.Close()
	return count
}

func CuentaProcSubdomains(db *sql.DB, domain string) (count int) {
	rows, err := db.Query("SELECT COUNT(*) FROM targets where domain = ? AND activo = ?", domain, 1)
	CheckErr(err)
	for rows.Next() {
		rows.Scan(&count)
	}
	defer db.Close()
	return count
}

//function to update http 0 no http or https 1 https 2 http 3 http y https
func UpdateHttp(db *sql.DB, subdomain string, port string, http int) {
	// Update data
	//fmt.Println("Insertando en " + subdomain + ":" + port + "La vulnerabilidad:" + vuln)
	stmt, err := db.Prepare("UPDATE targets SET http = ? WHERE subdomain = ? AND port = ?")
	CheckErr(err)
	stmt.Exec(http, subdomain, port)
	defer db.Close()
}
