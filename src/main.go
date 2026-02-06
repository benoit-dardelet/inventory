package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed www
var content embed.FS

func main() {
	log.Println("listening on :80")
	
	// On passe 'content' (nos fichiers) à la fonction router
	log.Fatal(http.ListenAndServe(":80", router(content)))	
}