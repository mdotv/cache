package cache

type Cache map[string]any

func New() Cache {
	return Cache{}
}

func (c Cache) Set(key string, value any) {
	c[key] = value
}

func (c Cache) Get(key string) any {
	return c[key]
}

func (c Cache) Delete(key string) {
	delete(c, key)
}
