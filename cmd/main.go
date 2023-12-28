package main

import (
	"fmt"
	"github.com/aadi-1024/quickshare/pkg/auth"
	"github.com/aadi-1024/quickshare/pkg/config"
	"github.com/aadi-1024/quickshare/pkg/handlers"
	"log"
	"net/http"
	"os"
)

var app *config.Config

func main() {
	mux := NewRouter()
	pass := auth.CodeGen()
	app = &config.Config{}

	if len(os.Args) < 2 {
		log.Fatalln("Provide filename")
	}
	app.Filename = os.Args[1]
	app.InProd = false
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
