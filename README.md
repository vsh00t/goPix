# goPix
Usa esta herramienta para realizar escaneos automatizados de descubrimiento de superficie de ataque en tu organización. Solo deberás ingresar el dominio. 

## Prerequisitos

- go (https://go.dev/doc/install)
- nuclei (go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest)
- httpx (go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest)
- naabu (sudo apt install -y libpcap-dev; go install -v github.com/projectdiscovery/naabu/v2/cmd/naabu@latest)
- subfinder (go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest)

## Uso
1. git clone https://github.com/vsh00t/goPix.git
2. cd goPix
3. go build -o gopix
4. ./gopix
5. Ingrese el dominio a escanear. 
6. Todos los resultados se almacenan en el archivo data.db (sqlite file)

## ToDo
goPix actualmente tiene una funcionalidad básica, se irán agregando más. 

1. Docker container
2. Reporte automatizado. 
3. Escaneo web
4. Escaneo cloud
