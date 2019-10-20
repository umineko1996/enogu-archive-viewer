package no6

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
)

func Listen() error {
	log.Println("server listen localhost:6060")
	http.ListenAndServe("localhost:6060", http.HandlerFunc(handler))
	return nil
}

var routePath = filepath.Clean("./react/resource")

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
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

	// pathがRestAPIを指定している場合はapiの処理を行う
	if api := getGetMethodAPI(targetPath); api != nil {
		api.Do(w, r)
		return
	}

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

	func() {
		content := ""
		switch filepath.Ext(fpath) {

		case ".css":
			content = "text/css"
		case ".js":
			content = "js"
		default:
			return
		}
		w.Header().Set("Content-Type", content)
	}()

	w.WriteHeader(http.StatusOK)
	if _, err = io.Copy(w, f); err != nil {
		log.Println(err)
	}
}

type RestAPI interface {
	Do(w http.ResponseWriter, r *http.Request)
}

type RestAPIFunc func(w http.ResponseWriter, r *http.Request)

func (rf RestAPIFunc) Do(w http.ResponseWriter, r *http.Request) {
	rf(w, r)
	return
}

func getGetMethodAPI(path string) RestAPI {
	switch path {
	case "/search":
		return RestAPIFunc(searchAPI)
	default:
		return nil
	}
}

func searchAPI(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	log.Println(query)

	words := query.Get("w")
	if words == "" {
		http.Error(w, `please "w" query`, http.StatusBadRequest)
		return
	}

	f, err := os.Open(ArchivesListFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	f.Seek(int64(len(utf8BOM)), os.SEEK_CUR) // 先頭のBOMスキップ
	var archivesInfo []*archiveInfo

	// MEMO 一行づつ読んでく形にした方がいいかも
	if err := gocsv.Unmarshal(f, &archivesInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response = struct {
		Videos []*archiveInfo
		Len    int
	}{
		Videos: make([]*archiveInfo, 0, 20),
		Len:    0,
	}

	for _, archive := range archivesInfo {
		// TODO しっかり治す
		if strings.Contains(archive.Title, words) {
			response.Videos = append(response.Videos, archive) // サイズ合わせてる
			if response.Len++; response.Len >= 20 {
				break
			}
		}
	}
	body, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(string(body))
	w.Write(body)
	//io.WriteString(w, `{"videos" : [{"url": "test1"}, {"url": "test2"}]}`)
}
