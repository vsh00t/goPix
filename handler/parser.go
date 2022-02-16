package handler

import (
	"encoding/json"
	"fmt"
	"main/database"
	"net"
	"strings"
)

func ParserSubdomain(domain string, subdomain []string) {
	for _, subdomain := range subdomains {
		if subdomain != "" {
			var data map[string]interface{}
			err := json.Unmarshal([]byte(subdomain), &data)
			if err != nil {
				panic(err)
			}
			host := data["host"].(string)
			a, _ := net.LookupIP(host)
			for _, ip := range a {
				ip := ip.String()
				var d []string = strings.Split(ip, " ")

				for _, ip := range d {

					if strings.Contains(ip, ".") {
						scanPort(database.ConnectDB(), domain, subdomain, host, ip)
						//fmt.Println("subdominio: ", host, "Direccion IP:", v)

					}

				}

			}

		}

	}

}

func ParserVuln(vulns []string) (vulnstring string) {
	for _, vuln := range vulns {
		if vuln != "" {
			//string json to map
			var data map[string]interface{}
			err := json.Unmarshal([]byte(vuln), &data)
			if err != nil {
				panic(err)
			}
			//get data
			//fmt.Println(data["extracted-results"])
			//fmt.Println(data["info"])
			r := data["extracted-results"]
			result := fmt.Sprintf("%v", r)
			i := data["info"]
			info := fmt.Sprintf("%v", i)
			var vulnsall []string
			vulnsall = append(vulnsall, result, info, "|")
			vulnstring = fmt.Sprintf("%v", vulnsall)
			//fmt.Println(vulnstring)

		}

	}
	vulnstring = strings.Replace(vulnstring, "[", "", -1)
	vulnstring = strings.Replace(vulnstring, "]", "", -1)
	vulnstring = strings.Replace(vulnstring, "reference:<nil>", "", -1)
	return vulnstring
}
