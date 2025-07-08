package main

import (
	"crypto/tls"
	"encoding/json"
	mw "first_api/internal/api/middlewares"
	"first_api/internal/api/router"
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

	fmt.Println("Servers is running on port :", port)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	/*rl := mw.NewRateLimiter(5, time.Minute)

	hppOptions := mw.HPPOptions{
		CheckQuery:              true,
		CheckBody:               true,
		CheckBodyForContentType: "application/x-www-form-urlencoded",
		Whitelist:               []string{"sortBy", "sortOrder", "name", "age", "class"},
	}*/

	//secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	router := router.Router()
	secureMux := mw.SecurityHeaders(router)
	// create custom server
	server := &http.Server{
		Addr: port,
		//Handler:   nil,
		//Handler:   mux,
		//Handler:   middlewares.SecurityHeaders(mux),
		//Handler:   middlewares.Cors(mux),
		//Handler:   mw.Hpp(hppOptions)(rl.Middleware(mw.Compression(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Cors(mux)))))),
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server :", err)
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
