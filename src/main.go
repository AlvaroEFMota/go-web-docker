package main

import (
	"log"
	"net/http"
	"github.com/AlvaroEFMota/go-web-docker/src/handles"
)
	
func main(){
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("src/static"))))
	http.HandleFunc("/", handles.Index)
	http.HandleFunc("/Produto/", handles.ProdutoInicio)
	http.HandleFunc("/Produto/Create", handles.ProdutoCreateForm)
	http.HandleFunc("/Produto/Create/Process", handles.ProdutoCreateProcess)
	http.HandleFunc("/Produto/Show", handles.ProdutoShow)
	log.Fatal(http.ListenAndServe(":8080", nil))
}