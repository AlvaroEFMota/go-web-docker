package handlers

//Produto define os atibutos de um produto
type Produto struct {
	ID int
	Nome string
	Preco float32
	Descricao string
}

//Constant define the constants about a web page
type Constant struct {
	Title string
	Bootstrap string
	PathToFiles string
}

//ProductWebPage stores the data about a web page
type ProductWebPage struct {
	ConstData Constant
	Produtos []Produto
}