// HTTP(Hyper text transfer protocol) // oddiy bir http yaratamiz.

package main

import (
	"io"
	"net/http"
)

// hello() funksiyasini yaratdik  va uning parametrlari
// res va req. va ular net/http paketidan ResponseWriter va Request
// qullanishyapti
func hello(res http.ResponseWriter, req *http.Request) {
	// res o'zgarivchisi orqali Header() va Set() funksiyasini qullanib
	// oddiy html sahifasini yaratamiz deb aytiyapmiz
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString( // io paketidagi WriteString funksiyasini qullanib
		res, // html kod yoziyabmiz
		`<DOCTYPE html>
    <html>
      <head>
        <title>Salom Gopherbek</title>
      </head>
      <body>
        Salom Gopherbek
      </body>
    </html>`,
		// backtic orqali oddiy bir html sahifasi yaratdik
	)
}

func main() {
	http.HandleFunc("/hello", hello) // HandleFunc funksiyasi orqali
	// hello rautiri yaratiyapmiz,
	http.ListenAndServe(":9000", nil) // va localhost:9000 portida eshitiyapmiz
	// localhost:9000/hello yurgizsangiz Salom Gopherbek habarini ko'rasiz
}
