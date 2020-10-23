package main

import (
	"log"
	"net/http"
	"github.com/AlvaroEFMota/go-web-docker/src/handlers"
)
	
func main(){
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("src/static"))))
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/Produto/", handlers.ProdutoInicio)
	http.HandleFunc("/Produto/Create", handlers.ProdutoCreateForm)
	http.HandleFunc("/Produto/Create/Process", handlers.ProdutoCreateProcess)
	http.HandleFunc("/Produto/Show", handlers.ProdutoShow)
	http.HandleFunc("/Produto/Delete/Process", handlers.ProdutoDeleteProcess)
	log.Fatal(http.ListenAndServe(":8080", nil))
}