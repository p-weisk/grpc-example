# Build

protobuf compiler und go compiler/sdk werden benötigt
* .proto datei im Ordner API mit protoc und Zielplattform Go kompilieren

## server

* im ordner server den befehl in build.sh ausführen (ggf. vorher `go get -v` ausführen, um dependencies zu laden)

## client
* im Ordner client den befehl `go build && go install` ausführen

# Run

## 1. Server
* in den ordner server wechseln
* docker-compose up ausführen
*
## 2. Client
* in den ordner wechseln, in den go installiete programme legt (`$GOPATH/bin`)
* dort das programm `client` ausführen
