package main

/*
import (
	"encoding/json"
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

	fmt.Println("Servers is running on port :", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Hello Root Route")
		_, err := w.Write([]byte("Hello Root Route"))
		if err != nil {
			log.Fatalln("Error writing :", err)
			return
		}
	})

	http.HandleFunc("/teachers/", teacherHandler)

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
*/
