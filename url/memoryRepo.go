package url

type memoryRepo struct {
	urls   map[string]*Url
	clicks map[string]int
}

func NewMemoryRepo() *memoryRepo {
	return &memoryRepo{
		make(map[string]*Url),
		make(map[string]int),
	}
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

func (r *memoryRepo) RegistClick(id string) {
	r.clicks[id] += 1
}

func (r *memoryRepo) Salv(url Url) error {
	r.urls[url.Id] = &url
	return nil
}
