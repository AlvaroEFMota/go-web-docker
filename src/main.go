package main

import (
	"log"
	"net/http"
	"github.com/AlvaroEFMota/go-web-docker/src/handles"
)
//Produto ,estrutura de produtos
type Produto struct {
	ID int32
	Nome string
	Preco float32
	Descricao string
}

var count int = 0

	
func main(){
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("src/static"))))
	http.HandleFunc("/", handles.Index)
	http.HandleFunc("/Produto/Show", handles.ProdutoShow)
	http.HandleFunc("/Produto/Create", handles.ProdutoCreateForm)
	http.HandleFunc("/Produto/Create/Process", handles.ProdutoCreateProcess)
	log.Fatal(http.ListenAndServe(":8081", nil))
}