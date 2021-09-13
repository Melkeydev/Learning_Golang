package helpers

import (
  "io/ioutil"
  "example.com/Melkey/structures"
)

func LoadPage(title string) (*structures.Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename) 

  if err != nil {
    return nil, err
  }

  return &structures.Page{Title: title, Body:body}, nil
}
