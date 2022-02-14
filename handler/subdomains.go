package handler

import (
	"log"
	"os/exec"
	"strings"
)

var subdomains []string

func ListSubdomains(domain string) {
	cmd := exec.Command("subfinder", "-d", domain, "-silent", "-oJ")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	stringout := string(out)
	subdomains = strings.Split(stringout, "\n")
	ParserSubdomain(domain, subdomains)

}
