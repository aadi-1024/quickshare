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
	"time"
)

var app *config.Config

func main() {
	pass := auth.CodeGen()
	app = &config.Config{}

	debug := os.Getenv("DEBUG")
	port := os.Getenv("PORT")

	if len(os.Args) < 2 {
		log.Fatalln("Provide filename")
	}
	app.Filename = os.Args[1]
	if debug == "1" || debug == "TRUE" {
		app.Debug = true
	} else {
		app.Debug = false
	}

	if port == "" {
		port = "8080"
	}
	//unix time is added so unique hash is generated every time for same file as well
	hash := fmt.Sprintf("%d", crc64.Checksum([]byte(app.Filename+fmt.Sprintf("%d", time.Now().Unix())), crc64.MakeTable(crc64.ISO)))
	app.Hash = hash

	mux := NewRouter(hash)
	handlers.InitRepo(app)

	fmt.Printf("Use passcode %s\n", pass)
	srv := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", port),
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
