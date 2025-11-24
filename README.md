## Uutissivusto - koulu projekti

Golla ja Gin- frameworkilla toteutettu ylen tapainen uutissivusto. 

- `main.go`, Reitit ja palvelimen käynnistys
- `internal` sisäinen logiikka, https://github.com/golang-standards/project-layout/blob/master/internal/README.md
- `internal/auth` autentikointi middleware ja kirjautuminen
- `internal/database` tietokanta yhteydet ja queryt
- `templates` HTML/tmpl mallit ja staattiset tiedostot 

## Pyörittäiminen

`go mod tidy` lataa riippuvuudet
`go run main.go` käynnistää palvelimen paikallisesti porttiin 8080, mysql yhteydet pitää hoitaa `internal/database/connect.go` tiedostossa