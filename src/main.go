package main

import (
	"net/http"
	"html/template"
	"fmt"
	"log"
)

var count int = 0
var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main(){
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fs)
	//http.HandleFunc("/", index)
	//http.HandleFunc("/process", process)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request){
	fmt.Printf("%d, %s\n",count,r.RemoteAddr)
	/*d := struct{
		Logo string
	}{
		Logo: "logo.png",
	}*/
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func process(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	count ++
	name := r.FormValue("name")
	age := r.FormValue("age")
	fmt.Printf("%d, %s, %s, %s\n",count,r.RemoteAddr, name, age)
	dict := struct {
		Name string
		Age string
	}{
		Name: name,
		Age: age,
	}
	tpl.ExecuteTemplate(w, "process.gohtml", dict)
	//fmt.Println(name,age)
}