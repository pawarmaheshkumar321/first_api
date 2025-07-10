package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Root Route")
	_, err := w.Write([]byte("Hello Root Route"))
	if err != nil {
		log.Fatalln("Error writing :", err)
		return
	}
}
