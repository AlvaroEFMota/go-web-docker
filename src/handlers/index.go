package handlers

import (
	"net/http"
	"html/template"
	"log"
)

var count int
var addrList []struct{
	Access int
	Addr string
}

//Index mostra a pagina inicial
func Index(w http.ResponseWriter, r *http.Request){
	count++
	//log.Printf("%d, %s\n",count,r.RemoteAddr)
	entrou := false
	for i, elem := range addrList {
		if r.RemoteAddr == elem.Addr {
			entrou = true
			addrList[i].Access++
		}
	}
	log.Println(entrou)
	if entrou == false {
		addrList = append(addrList, struct{
			Access int
			Addr string
		}{
			Addr: r.RemoteAddr,
			Access: 0,
		})
	}
	d := struct{
		Log int
		Addr string
	}{
		Log: count,
		Addr: r.RemoteAddr,
	}
	log.Println(addrList)
	tpl, err := template.ParseFiles("src/templates/index.gohtml")
	if err != nil {
		log.Fatal("não abriu o arquivo")
	}
	tpl.Execute(w,d)
}