package main

import (
	"net/http"
	"fmt"
)


func main() {

	fmt.Printf("Começando!")
		

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Teste sexta-feira dia 12/8/2022- 15:22</h1>"))

	})


	fmt.Printf("Inciando!")

	http.ListenAndServe(":8080", nil)

	fmt.Printf("Fim!")

}