package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

// On ajoute 'content' en paramètre
func router(content embed.FS) http.Handler {
	mux := http.NewServeMux()

	// --- ROUTES API ---
	mux.HandleFunc("GET /health", HealthHandle)
	mux.HandleFunc("GET /cpu", CPUHandler)
	mux.HandleFunc("GET /ps", PSHandler)
	mux.HandleFunc("GET /ps/{user}", PSUserHandler)
	mux.HandleFunc("GET /net", NetHandler)
	mux.HandleFunc("GET /net/{card}", NetNameHandler)
	mux.HandleFunc("GET /mem", MemHandler)
	mux.HandleFunc("GET /disk", DiskHandler)
	mux.HandleFunc("GET /load", LoadHandler)
	mux.HandleFunc("GET /ps/kill/{pid}", KillProcessHandler)

	// --- FICHIERS STATIQUES (SITE WEB) ---
	
	// 1. On récupère le sous-dossier "www" dans l'embed
	wwwFS, err := fs.Sub(content, "www")
	if err != nil {
		log.Fatal("Erreur critique : dossier 'www' introuvable !", err)
	}

	// 2. On sert ce dossier à la racine "/"
	// Si on demande /procs.html, Go cherchera dans content/www/procs.html
	mux.Handle("/", http.FileServer(http.FS(wwwFS)))

	return mux
}