package web

import "net/http"

func Router() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(assetFS())))
}
