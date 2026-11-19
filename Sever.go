package main

import (
	"encoding/json"
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
	STATS   chan string
)

func init() {
	PORT = 31337
	BASEURL = fmt.Sprintf("http://localhost:%d", PORT)
}

func main() {
	// Records statistics as soon as the server is started
	STATS = make(chan string)
	defer close(STATS)
	go recordStatistics(STATS)

	http.HandleFunc("/add/index", Shortener)
	http.HandleFunc("/log/view/", ViewLogFunc)
	http.HandleFunc("/r/", Redirector)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%d", PORT), nil))
}

func recordStatistics(ids <-chan string) {
	for id := range ids {
		url.resgistClick(id)
		fmt.Printf("%s - Click successfully registered for %s.\n", time.Now().Format("2006-01-02 15:04:05"), id)
	}
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

func replyWithJSON(w http.ResponseWriter, response string) {
	replyWith(w, http.StatusOK, Headers{
		"Content-Type": "application/json",
	})
	fmt.Fprintf(w, response)
}

func ExtractrUrl(r *http.Request) string {
	url := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(url)
	return string(url)
}

func Shortener(w http.ResponseWriter, r *http.Request) {

	// 405 - Only POST's Method Allowed
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
	replyWith(w, status, Headers{
		"Location": shortUrl,
		// Following a RESTful API standard
		"Link": fmt.Sprintf("<%s/log/view/%s> ; rel=\"view\"", BASEURL, url.Id),
	})
}

func Redirector(w http.ResponseWriter, r *http.Request) {

	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	if url := url.Search(id); url != nil {
		http.Redirect(w, r, url.Dest, http.StatusMovedPermanently)
		STATS <- id
	} else {
		http.NotFound(w, r)
	}
}

func ViewLogFunc(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	if url := url.Search(id); url != nil {
		json, err := json.Marshal(url.Stats())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		replyWithJSON(w, string(json))
	} else {
		http.NotFound(w, r)
	}
}
