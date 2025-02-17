package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplateTest(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+html, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var html *template.Template
	var err error

	//check if we already hv a template in our map
	_, inMap := tc[t]
	if !inMap {
		//need to create the template
		log.Println("creating template and adding it to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//we have the template inside the cache
		log.Println("using cached template")
	}

	html = tc[t] //get the template from the cache by its name

	err = html.Execute(w, nil)
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), //t is the paramater that is received as the call to this function
		"./templates/base.layout.html",
	}

	//parse the template
	html, err := template.ParseFiles(templates...) //take all the entries in the templates slice and pass them as individual strings

	if err != nil {
		return err
	}

	//add template to the cache (map)
	tc[t] = html

	return nil
}
