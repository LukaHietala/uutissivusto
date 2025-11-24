## Uutissivusto - koulu projekti

Golla ja Gin- frameworkilla toteutettu ylen tapainen uutissivusto. 

- `main.go`, Reitit ja palvelimen käynnistys
- `internal` sisäinen logiikka, https://github.com/golang-standards/project-layout/blob/master/internal/README.md
- `internal/auth` autentikointi middleware ja kirjautuminen
- `internal/database` tietokanta yhteydet ja queryt
- `templates` HTML/tmpl mallit ja staattiset tiedostot 

## Pyörittäiminen

- `go mod tidy` lataa riippuvuudet
- `go run main.go` käynnistää palvelimen paikallisesti porttiin 8080, mysql yhteydet pitää hoitaa `internal/database/connect.go` tiedostossa

<img width="1920" height="1052" alt="Screenshot From 2025-11-24 19-43-55" src="https://github.com/user-attachments/assets/2ed91b76-4e08-4deb-9973-a84f07df6ff9" />
