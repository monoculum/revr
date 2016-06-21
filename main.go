// Copyright 2016 Monoculum. All rights reserved.
// Use of this source code is governed by a MIT license that can be found
// in the LICENSE file.

package revr

import (
	"regexp"
	"strings"
)

// URL indicates the path and params of a url
type URL struct {
	path   string
	params []string
}

// URLs a map of urls by its name registered
type URLs map[string]URL

// URLStore a store of urls
type URLStore struct {
	store URLs
}

// New returns a URLStore
func New() *URLStore {
	return &URLStore{store: make(URLs)}
}

var rex = regexp.MustCompile(`:[\w]+`)

// MustAdd adds a route and it panic if error
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

// MustReverse get a reversed url and it panic if error
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
