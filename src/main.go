package main

import (
	"github.com/AlvaroEFMota/go-web-docker/handles"
	"net/http"
	"html/template"
	"fmt"
	"log"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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
	var err error
	db, err = sql.Open("mysql", "root:toor@tcp(localhost:3306)/go_web_docker")
	if err != nil {
		panic(err)
	}
	
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println("Connected to database!")

		tpl = template.Must(template.ParseGlob("src/templates/*.gohtml"))
}
	
func main(){
	
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("src/static"))))
	http.HandleFunc("/", handles.index)
	http.HandleFunc("/Produto/Show", produtoShow)
	http.HandleFunc("/Produto/Create", produtoCreateForm)
	http.HandleFunc("/Produto/Create/Process", produtoCreateProcess)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
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
}*/

func produtoCreateForm(w http.ResponseWriter, r *http.Request){
	
	tpl.ExecuteTemplate(w, "produtoCreate.gohtml", nil)
}

func produtoCreateProcess(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	var pro Produto
	pro.Nome = r.FormValue("nome_produto")
	precoTmp := r.FormValue("preco_produto")
	pro.Descricao = r.FormValue("descricao_produto")
	//Validaçãp
	if(pro.Nome == "" || precoTmp == "" || pro.Descricao == "") {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	f64, err := strconv.ParseFloat(precoTmp,32)//Ainda continua 64
	if err != nil {
		http.Error(w,http.StatusText(406)+"Por favor, informe um valor válido para o preço", http.StatusNotAcceptable)
	}
	pro.Preco = float32(f64)
	tpl.ExecuteTemplate(w, "produtoCreate.gohtml", nil)
}

func produtoShow(w http.ResponseWriter, r *http.Request){
	/*if r.Method != "GET" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}*/
	lista := []Produto{}
	lista = append(lista, Produto{ID: 1, Nome: "Teclado", Preco: 153.7, Descricao: "Teclado mecânico e blablabla"})
	lista = append(lista, Produto{ID: 2, Nome: "Mouse", Preco: 47.0, Descricao: "Mouse blablabla"})
	tpl.ExecuteTemplate(w, "produtoShow.gohtml", lista)
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
	tpl.ExecuteTemplate(w, "mensagem.gohtml", dict)
	//fmt.Println(name,age)
}