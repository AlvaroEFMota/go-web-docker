package main

import (
	"html/template"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

var db *sql.DB
var tpl *template.Template

func init()  {
	/*var err error
	db, err = sql.Open("mysql", "root:toor@tcp(localhost:3306)/go_web_docker")
	if err != nil {
		panic(err)
	}
	
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println("Connected to database!")
	*/
		//tpl = template.Must(template.ParseGlob("src/templates/*.gohtml"))
}
	
func main(){
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("src/static"))))
	http.HandleFunc("/", handles.Index)
	http.HandleFunc("/Produto/Show", handles.ProdutoShow)
	http.HandleFunc("/Produto/Create", handles.ProdutoCreateForm)
	http.HandleFunc("/Produto/Create/Process", handles.ProdutoCreateProcess)
	log.Fatal(http.ListenAndServe(":8081", nil))
}