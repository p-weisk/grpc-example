# Build

protobuf compiler und go compiler/sdk sowie Docker werden benötigt
* .proto datei im Ordner API mit protoc und Zielplattform Go kompilieren
```
go get -u github.com/golang/protobuf/protoc-gen-go
cd ${GOPATH}/src/github.com/golang/protobuf/protoc-gen-go && go build && go install
export $PATH="$PATH:$GOPATH/bin"

protoc -I api/ \
    -I ${GOPATH}/src \
    --go_out=plugins=grpc:api \
    api/api.proto
```

## server

* im ordner server den befehl in build.sh ausführen (ggf. vorher `go get -v` ausführen, um dependencies zu laden)
* `sudo docker-compose build` ausführen
## client
* im Ordner client/js `npm install` ausführen

# Run

## 1. Server
* in den ordner server wechseln
* `sudo docker-compose up` ausführen
*
## 2. Client
* in den ordner client/js wechseln
* dort ausführen:
    ```
    node findInvoiceById.js <arg>
    node getVolumeOfSales.js <arg>
    node createInvoice.js <args>
    ```
