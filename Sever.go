package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Null-Br-Group/TurnPoint/url"
)

type Headers map[string]string

type Url struct {
	Id      string
	Created time.Time
	Dest    string
}

var (
	PORT    int
	BASEURL string
)

func init() {
	PORT = 31337
	BASEURL = fmt.Sprintf("http://localhost:%d", PORT)
}

func main() {
	http.HandleFunc("/api/encurtar", Shortener)
	http.HandleFunc("/r/", Redirector)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%d", PORT), nil))
}

func Shortener(w http.ResponseWriter, r *http.Request) {

	// 405 - Only POST
	if r.Method != "POST" {
		replyWith(w, http.StatusMethodNotAllowed, Headers{
			"Allow": "POST",
		})
		return
	}

	url, new, err := url.SearchCreateNewUurl(ExtractrUrl(r))

	if err != nil {
		replyWith(w, http.StatusBadRequest, nil)
		return
	}

	var status int
	if new {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}

	shortUrl := fmt.Sprintf("%s/r/%s", BASEURL, url.Id)
	replyWith(w, status, Headers{"Location": shortUrl})
}

func replyWith(
	w http.ResponseWriter,
	status int,
	headers Headers,
) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
}

func ExtractrUrl(r *http.Request) string {
	url := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(url)
	return string(url)
}

func Redirector(w http.ResponseWriter, r *http.Request) {

	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	if url := url.Search(id); url != nil {
		http.Redirect(w, r, url.Dest, http.StatusMovedPermanently)
	} else {
		http.NotFound(w, r)
	}
}
