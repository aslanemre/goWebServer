package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./www"))

	http.Handle("/", fileServer)

	fmt.Printf("Web sunucu başlatıldı. https://localhost:443\n")

	if hata := http.ListenAndServeTLS(":443", "server_tls.crt", "server_tls.key", nil); hata != nil {

		log.Fatal(hata)

	}

}
