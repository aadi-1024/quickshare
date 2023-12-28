package main

import (
	"fmt"
	"github.com/aadi-1024/quickshare/pkg/auth"
	"net/http"
)

func main() {
	mux := NewRouter()
	pass := auth.CodeGen()
	fmt.Printf("Use passcode %s\n", pass)
	srv := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
