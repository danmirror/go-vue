package main

import(
	"fmt"
	"net/http"
	"html/template"
	"path"
)
func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		// costume delim untuk membuat parse ke html 
		var pathfile = path.Join("templates","index.html")
		var tpl,  err= template.New("index").Delims("[[", "]]").ParseFiles(pathfile)
		// fmt.Println(tpl.DefinedTemplates()) 
		
		if err != nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		var data = map[string]interface{}{
			"Title" : "back-end from Golang",
			"body" :"ini body ",
			"message" : "ini go",
		}
	  	tpl.ExecuteTemplate(w, "index.html", data)
  })
  http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	fmt.Printf("running server")

	http.ListenAndServe(":8080",nil)
}