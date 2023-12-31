package main

import (
	"fmt"
	"github.com/aadi-1024/quickshare/pkg/auth"
	"github.com/aadi-1024/quickshare/pkg/config"
	"github.com/aadi-1024/quickshare/pkg/handlers"
	"hash/crc64"
	"log"
	"net/http"
	"os"
)

var app *config.Config

func main() {
	pass := auth.CodeGen()
	app = &config.Config{}

	if len(os.Args) < 2 {
		log.Fatalln("Provide filename")
	}
	app.Filename = os.Args[1]
	app.InProd = false

	hash := fmt.Sprintf("%d", crc64.Checksum([]byte(app.Filename), crc64.MakeTable(crc64.ISO)))
	app.Hash = hash

	mux := NewRouter(hash)
	handlers.InitRepo(app)

	fmt.Printf("Use passcode %s\n", pass)
	srv := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
