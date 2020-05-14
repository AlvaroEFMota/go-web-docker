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
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main(){

	http.Handle("/templates/build/", http.StripPrefix("/templates/build/",
		http.FileServer(http.Dir("templates/build"))))

	http.HandleFunc("/", index)
	//http.HandleFunc("/process", process)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func index(w http.ResponseWriter, r *http.Request){
	count++
	fmt.Printf("%d, %s\n",count,r.RemoteAddr)

	d := struct{
		Log int
		Addr string
	}{
		Log: count,
		Addr: r.RemoteAddr,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", d)
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