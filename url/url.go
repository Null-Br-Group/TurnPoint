package url

import (
	"math/rand"
	"net/url"
	"time"
)

type Url struct {
	Id      string
	Created time.Time
	Dest    string
}

type Repo interface {
	IdExist(id string) bool
	FindId(id string) *Url
	FindUrl(url string) *Url
	Salve(url Url) error
	RegistClick(id string)
}

const (
	size    = 5
	symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+="
)

var repo Repo

func SettingUpRepo(r Repo) {
	repo = r
}

func init() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(100)
}

func SearchCreateNewUurl(dest string) (
	u *string,
	new bool,
	err error,
) {
	if u = repo.FindUrl(dest); u != nil {
		return u, false, nil
	}

	if _, err = url.ParseRequestURI(dest); err != nil {
		return nil, false, err
	}

	url := Url{generateID(), time.Now(), dest}
	repo.Salve(url)
	return &url, true, nil
}

func generateID() string {
	newID := func() string {
		id := make([]byte, size, size)
		for i := range id {
			id[i] = symbols[rand.Intn(len(symbols))]
		}
		return string(id)
	}

	for {
		if id := newID(); !repo.IdExist(id) {
			return id
		}
	}
}

func Search(id string) *Url {
	return repo.FindId(id)
}
