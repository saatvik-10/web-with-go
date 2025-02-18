package render

import (
	"bytes"
//	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, html string) {
	//create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal("Error parsing template", err)
	}

	//get req. template from cache
	t, ok := tc[html]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer) //in-memory writer

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	//render the template
//	parsedTemplate, _ := template.ParseFiles("./templates/"+html, "./templates/base.layout.html")
//	err := parsedTemplate.Execute(w, nil)
//	if err != nil {
//		fmt.Println("error parsing template: ", err)
//		return
//	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	//get all of the files with *.page.html
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	//range through all *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}


//var tc = make(map[string]*template.Template)

//func RenderTemplateTest(w http.ResponseWriter, t string) {
//	var html *template.Template
//	var err error
//
//	//check if we already hv a template in our map
//	_, inMap := tc[t]
//	if !inMap {
//		//need to create the template
//		log.Println("creating template and adding it to cache")
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		//we have the template inside the cache
//		log.Println("using cached template")
//	}
//
//	html = tc[t] //get the template from the cache by its name
//
//	err = html.Execute(w, nil)
//}
//
//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t), //t is the paramater that is received as the call to this function
//		"./templates/base.layout.html",
//	}
//
//	//parse the template
//	html, err := template.ParseFiles(templates...) //take all the entries in the templates slice and pass them as individual strings
//
//	if err != nil {
//		return err
//	}
//
//	//add template to the cache (map)
//	tc[t] = html
//
//	return nil
//}


