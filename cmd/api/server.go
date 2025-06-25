package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3000"

	fmt.Println("Servers is running on port :", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Hello Root Route")
		_, err := w.Write([]byte("Hello Root Route"))
		if err != nil {
			log.Fatalln("Error writing :", err)
			return
		}
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello Teachers Route")
		fmt.Println("Method :", r.Method)

		switch r.Method {

		case http.MethodGet:
			w.Write([]byte("Hello Get Method in Teachers Route"))

		case http.MethodPost:
			w.Write([]byte("Hello Post Method in Teachers Route"))

		case http.MethodPut:
			w.Write([]byte("Hello Put Method in Teachers Route"))

		case http.MethodPatch:
			w.Write([]byte("Hello Patch Method in Teachers Route"))

		case http.MethodDelete:
			w.Write([]byte("Hello Delete Method in Teachers Route"))

		default:
			w.Write([]byte("Hello Unknown Method in Teachers Route"))
		}
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello Students Route")
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello Execs Route")
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server :", err)
	}

}
