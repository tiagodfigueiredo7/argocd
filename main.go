package main

import (
	"net/http"
	"fmt"
)


func main() {

	fmt.Printf("Começando!")
		

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Teste segunda-feira dia 10/8/2022- 21:50</h1>"))

	})


	fmt.Printf("Inciando!")

	http.ListenAndServe(":8080", nil)

	fmt.Printf("Fim!")

}