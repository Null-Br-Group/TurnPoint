package url

type memoryRepo struct {
	urls map[string]*Url
}

func NewMemoryRepo() *memoryRepo {
	return &memoryRepo{make(map[string]*Url)}
}

func (r *memoryRepo) IdExist(id string) bool {
	_, existe := r.urls[id]
	return existe
}

func (r *memoryRepo) FindId(id string) *Url {
	return r.urls[id]
}

func (r *memoryRepo) FindUrl(url string) *Url {
	for _, u := range r.urls {
		if u.Dest == url {
			return u
		}
	}

	return nil
}

func (r *memoryRepo) Salv(url Url) error {
	r.urls[url.Id] = &url
	return nil
}
