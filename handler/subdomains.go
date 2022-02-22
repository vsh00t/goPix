//by vsh00t
package handler

import (
	"log"
	"os/exec"
	"strings"

	"github.com/vsh00t/goPix/ui"
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
	count := len(subdomains)
	ui.NumSubdomains(count - 1)
	ui.InicioPorts()
	ParserSubdomain(domain, subdomains)

}
