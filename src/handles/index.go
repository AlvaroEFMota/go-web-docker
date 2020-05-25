package handles

import (
	"net/http"
	"html/template"
	"log"
)

//var tpl *template.Template
var count int

//Index mostra a pagina inicial
func Index(w http.ResponseWriter, r *http.Request){
	count++
	log.Printf("%d, %s\n",count,r.RemoteAddr)
	d := struct{
		Log int
		Addr string
	}{
		Log: count,
		Addr: r.RemoteAddr,
	}
	tpl, err := template.ParseFiles("src/templates/index.gohtml")
	if err != nil {
		log.Fatal("não abriu o arquivo")
	}
	tpl.Execute(w,d)
}