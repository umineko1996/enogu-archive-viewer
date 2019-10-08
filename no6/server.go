package no6

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Listen() error {
	log.Println("server listen localhost:6060")
	http.ListenAndServe("localhost:6060", http.HandlerFunc(handler))
	return nil
}

var routePath = filepath.Clean("./react/resource")

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMethodHandler(w, r)
	default:
		http.Error(w, "not supported method.", http.StatusBadRequest)
	}
}

func getMethodHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	targetPath := r.URL.Path
	if targetPath == "" || targetPath == "/" {
		targetPath = "index.html"
	}
	fpath := filepath.Join(routePath, targetPath)
	log.Println(fpath)
	if !strings.HasPrefix(fpath, routePath) {
		http.Error(w, "permission denied.", http.StatusForbidden)
		return
	}

	f, err := os.Open(fpath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = io.Copy(w, f); err != nil {
		log.Println(err)
	}
}
