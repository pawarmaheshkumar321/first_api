package main

import (
	"crypto/tls"
	"encoding/json"
	"first_api/internal/api/middlewares"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func main() {
	port := ":3000"

	// create tls mux server
	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	fmt.Println("Servers is running on port :", port)

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teacherHandler)

	mux.HandleFunc("/students/", studentHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// create custom server
	server := &http.Server{
		Addr:      port,
		Handler:   middlewares.SecurityHeaders(mux),
		TLSConfig: tlsConfig,
	}

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server :", err)
	}

	/*err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server :", err)
	}*/

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Root Route")
	_, err := w.Write([]byte("Hello Root Route"))
	if err != nil {
		log.Fatalln("Error writing :", err)
		return
	}
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Students Route")
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Execs Route")
}

func teacherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Teachers Route")
	fmt.Println("Method :", r.Method)

	switch r.Method {

	case http.MethodGet:
		w.Write([]byte("Hello Get Method in Teachers Route"))
		parsePathParameters(w, r)
		parseQueryParameters(w, r)

	case http.MethodPost:
		w.Write([]byte("Hello Post Method in Teachers Route"))
		parseFormElement(w, r)
		parseRawBodyElement(w, r)

	case http.MethodPut:
		w.Write([]byte("Hello Put Method in Teachers Route"))

	case http.MethodPatch:
		w.Write([]byte("Hello Patch Method in Teachers Route"))

	case http.MethodDelete:
		w.Write([]byte("Hello Delete Method in Teachers Route"))

	default:
		w.Write([]byte("Hello Unknown Method in Teachers Route"))
	}
}

func parseFormElement(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error Parsing Form :", http.StatusBadRequest)
		log.Fatalln("Error Parsing Form :", err)
	}

	fmt.Println("Form  :", r.Form)
	fmt.Println("Form  :", r.Form.Get("name"))

	response := make(map[string]interface{})
	for key, value := range r.Form {
		response[key] = value
	}

	fmt.Println("Processed response map  :", response)
	fmt.Println("Processed response map  :", response)

}

func parseRawBodyElement(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	defer r.Body.Close()

	fmt.Println("RawBody  :", body)
	fmt.Println("RawBody  :", string(body))

	// unmarshling json data with struct

	var userInstance user
	err = json.Unmarshal(body, &userInstance)
	if err != nil {
		fmt.Println("Error unMarshling Json :", err)
		return
	}

	fmt.Println("userInstance  :", userInstance)
	fmt.Println("userInstance Name :", userInstance.Name)
}

func parsePathParameters(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path  :", r.URL.Path)

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	userID := strings.TrimSuffix(path, "/")

	fmt.Println("Path2  :", path)
	fmt.Println("userID  :", userID)
}

func parseQueryParameters(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	fmt.Println("Query Params name :", queryParams.Get("key"))

}
