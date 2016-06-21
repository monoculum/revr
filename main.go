package revr

import (
	"regexp"
	"strings"
)

type URL struct {
	path   string
	params []string
}

type URLs map[string]URL

type URLStore struct {
	store URLs
}

func New() *URLStore {
	return &URLStore{store: make(URLs)}
}

var rex = regexp.MustCompile(`:[\w]+`)

func (u *URLStore) MustAdd(name, path string) string {
	m := URL{path: path}
	s := rex.FindAll([]byte(path), -1)
	for i := range s {
		m.params = append(m.params, string(s[i]))
	}
	if _, ok := u.store[name]; ok {
		panic("The url already exists.")
	}
	u.store[name] = m
	return path
}

func (u *URLStore) MustReverse(name string, params ...string) string {
	url, ok := u.store[name]
	if !ok {
		panic("The url not exists...")
	}
	if len(params) != len(url.params) {
		panic("The length of params argument is different to url's params")
	}
	ur := url.path
	for i, v := range url.params {
		ur = strings.Replace(ur, v, params[i], 1)
	}
	return ur
}
