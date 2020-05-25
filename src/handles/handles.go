package handles

import (
	"net/http"
)

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