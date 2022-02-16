package ui

import (
	"fmt"

	"github.com/fatih/color"
)

func Colorize() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"
	colorMagenta := "\033[35m"
	colorBold := "\033[1m"
	colorUnderline := "\033[4m"
	colorReset = "\033[0m"
	color.Set(color.FgRed, color.Bold)
	fmt.Println(colorRed, ".", colorReset, colorGreen, ":", colorReset, colorYellow, ".", colorReset, colorBlue, ":", colorReset, colorPurple, ".", colorReset, colorCyan, ":", colorReset, colorWhite, ".", colorReset, colorMagenta, ":", colorReset, colorBold, ".", colorReset, colorUnderline, ":", colorReset, colorReset, colorReset)
	fmt.Println(colorBlue, "	▄▄            ")
	fmt.Println(colorBlue, "  ▄▄█▀▀▀█▄█         ▀███▀▀▀██▄  ██            ")
	fmt.Println(colorBlue, "▄██▀     ▀█           ██   ▀██▄               ")
	fmt.Println(colorBlue, "██▀       ▀  ▄██▀██▄  ██   ▄██▀███ ▀██▀   ▀██▀")
	fmt.Println(colorBlue, "█▓          ██▀   ▀██ ███████   ██   ▀██ ▄█▀  ")
	fmt.Println(colorBlue, "█▓▄    ▀██████     ██ ██        ▓█     ███    ")
	fmt.Println(colorBlue, "▀▓█▄     ██ ██     ▓█ █▓        ▓█     ▓▓██   ")
	fmt.Println(colorBlue, "▓▓▓    ▀▓█▓▓▓█     ▓▓ █▓        ▓▓     ▓▓█    ")
	fmt.Println(colorBlue, "▀▒▓▓     ▓▓ ▓▓▓   ▓▓▓ ▓▓        ▓▓   ▓▓▀ ▓▓▓  ")
	fmt.Println(colorBlue, "  ▒▒▒ ▒ ▒▒   ▒ ▒ ▒ ▒▒▓▒▓▒     ▒ ▒ ▒ ▒▒    ▒▓▒ ")
	fmt.Println(colorBlue, "")
	fmt.Println(colorBlue, "by: vSh00t\n")
	fmt.Println(colorBlue, "Twitter: @Jim0ya")
}

func Inicio(domain string) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorBlue := "\033[34m"
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\n")
	fmt.Println(colorRed, "[+]", colorReset, colorBlue, "Iniciando enumeración de subdominios de:", domain)
	fmt.Printf("\n")

}

func InicioPorts() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorBlue := "\033[34m"
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\n")
	fmt.Println(colorRed, "[+]", colorReset, colorBlue, "Iniciando el escaneo de puertos de todos los subdominios identificados")
	fmt.Printf("\n")

}

func InicioScaneo() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorBlue := "\033[34m"
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\n")
	fmt.Println(colorRed, "[+]", colorReset, colorBlue, "Iniciando el escaneo de Vulnerabilidades")
	fmt.Printf("\n")
}

func NumSubdomains(numSubdomains int) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\n")
	fmt.Println(colorRed, "[-]", colorReset, colorGreen, "Se han identificado", colorReset, colorRed, numSubdomains, colorReset, colorGreen, " subdominios")
	fmt.Printf("\n")
}

func NumAllSubdomains(numAllSubdomains int) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	color.Set(color.FgRed, color.Bold)
	fmt.Printf("\n")
	fmt.Println(colorRed, "[-]", colorReset, colorGreen, "Se han identificado", colorReset, colorRed, numAllSubdomains, colorReset, colorGreen, " servicios activos")
	fmt.Printf("\n")
}
