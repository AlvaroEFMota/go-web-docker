package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"github.com/AlvaroEFMota/go-web-docker/src/database"
)

//ProdutoCreateForm handle para o formulário de criação de produto
func ProdutoCreateForm(w http.ResponseWriter, r *http.Request){
	
	tpl, err := template.ParseFiles("src/templates/produtoCreate.gohtml")
	if err != nil {
		log.Fatal("não foi possível abrir o arquivo [produtoCreate.gohtml]")
	}
	tpl.Execute(w, nil)
}

//ProdutoInicio mostra os produtos (index dos produtos)
func ProdutoInicio(w http.ResponseWriter, r *http.Request){
	/*if r.Method != "GET" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}*/
	var productWebPage ProductWebPage
	db := database.GetConexao()

	rows, err := db.Query("SELECT id, nome, preco, descricao FROM Produto")
	if err != nil {
		log.Fatal("[database] não foi possível fazer uma consulta")
	}

	lista := []Produto{}
	for rows.Next() {
		var pro Produto
		err := rows.Scan(&pro.ID, &pro.Nome, &pro.Preco, &pro.Descricao)
		if err != nil {
			log.Fatal("[database] erro ao ler os dados da consulta ao banco de dados")
		}
		lista = append(lista,pro)
	}
	productWebPage.Produtos = lista;
	productWebPage.ConstData.Title = "Alvaro 222"
	productWebPage.ConstData.PathToFiles = "/static"

	
	//lista = append(lista, Produto{ID: 1, Nome: "Teclado", Preco: 153.7, Descricao: "Teclado mecânico e blablabla"})
	//lista = append(lista, Produto{ID: 2, Nome: "Mouse", Preco: 47.0, Descricao: "Mouse blablabla"})
	tpl, err := template.ParseFiles("src/templates/produto.gohtml")
	if err != nil {
		log.Fatal("não foi possível abrir o arquivo [produto.gohtml]")
	}
	tpl.Execute(w, productWebPage)
}

//ProdutoCreateProcess É o processo que usa o metodo POST para criar um produto
func ProdutoCreateProcess(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
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

	db := database.GetConexao()
	db.Exec("INSERT INTO Produto (nome, preco, descricao)VALUES(?,?,?)",pro.Nome, pro.Preco, pro.Descricao)
	
	//http.Redirect(w,r,"/Produto/Show",http.StatusSeeOther)
	
	tpl, err := template.ParseFiles("src/templates/produtoCreated.gohtml")
	if err != nil {
		log.Fatal("não foi possível abrir o arquivo [produtoCreated.gohtml]")
	}
	tpl.Execute(w, pro)
}
//ProdutoShow mostra um produto especifico
func ProdutoShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	db := database.GetConexao()
	prodID := r.FormValue("id")
	var produto Produto

	row := db.QueryRow("SELECT id, nome, preco, descricao FROM Produto WHERE id = ?", prodID)	
	err := row.Scan(&produto.ID, &produto.Nome, &produto.Preco, &produto.Descricao)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}

	tpl, err := template.ParseFiles("src/templates/produtoShow.gohtml")
	if err != nil {
		log.Fatal("não foi possível abrir o arquivo [produtoShow.gohtml]")
	}
	//tpl.Execute(w, produto)
	tpl.Execute(w, produto)
	
}

//ProdutoDeleteProcess deleta um produto com base em seu id
func ProdutoDeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	
	
	db := database.GetConexao()
	prodID := r.FormValue("id")

	_, err := db.Exec("DELETE FROM Produto WHERE id = ?", prodID)
	if err != nil {
		log.Println("Erro ao deletar da tabela Produto onde ID =", prodID)
	}

	/*tpl, err := template.ParseFiles("src/templates/produtoDeleted.gohtml")
	if err != nil {
		log.Fatal("não foi possível abrir o arquivo [produtoDeleted.gohtml]")
	}
	tpl.Execute(w, nil)
	*/
	http.Redirect(w,r,"/Produto", http.StatusPermanentRedirect)
}

//ProdutoUpdate redireciona para um formulário para a edição de um produto
func ProdutoUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r," /", http.StatusSeeOther)
		return
	}
	var produto Produto
	prodID := r.FormValue("id")
	db := database.GetConexao()
	produto.ID, _ = strconv.Atoi(prodID)


	row := db.QueryRow("SELECT nome, preco, descricao FROM Produto WHERE id = ?", prodID)
	err := row.Scan(&produto.Nome, &produto.Preco, &produto.Descricao)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
	
	tpl, err := template.ParseFiles("src/templates/produtoUpdate.gohtml")
	if err != nil {
		log.Println("ERRO ao tentar abrir o aquivo [produtoUpdate.gohtml]")
	}

	tpl.Execute(w, produto)
}

//ProdutoUpdateProcess é o processo de atualização de um produto
func ProdutoUpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	db := database.GetConexao()
	var produto Produto
	produto.ID, _ = strconv.Atoi(r.FormValue("id_produto"))
	produto.Nome = r.FormValue("nome_produto")
	f64, _ := strconv.ParseFloat(r.FormValue("preco_produto"), 32)
	produto.Preco = float32(f64)
	produto.Descricao = r.FormValue("descricao_produto")
	log.Println(produto)

	_, err := db.Exec("UPDATE Produto SET nome = ?, preco = ?, descricao = ? WHERE id = ?",
																	produto.Nome,
																	produto.Preco,
																	produto.Descricao,
																	produto.ID)
	if err != nil {
		log.Println("Erro ao deletar da tabela Produto onde ID =", produto.ID)
	}

	tpl, err := template.ParseFiles("src/templates/produtoUpdated.gohtml")
	if err != nil {
		log.Println("ERRO ao abrir o arquivo [produtoUpdated.gohtml]")
	}

	tpl.Execute(w, produto)
}