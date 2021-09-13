package structures

import (
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) Save() error {
  filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
