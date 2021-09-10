package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Struct for the wiki page
type Page struct {
	Title string
	Body []byte
}



func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
// is have a follow up function of some sorts that implements a if err != nil:

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	// Read the file if it exists
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}


	return &Page{Title: title, Body:body}, nil
}


func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)

	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}



func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8084", nil))

	//p1 := &Page{Title: "TestPage", Body: []byte("This is a simple page")}
	//p1.save()

	//p2, err := loadPage("TestPage")
	//if err != nil {
		//log.Fatal(err)
	//}


}
