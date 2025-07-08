package main

import (
	"crypto/tls"
	"encoding/json"
	mw "first_api/internal/api/middlewares"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

type Teacher struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Class     string `json:"class"`
	Subject   string `json:"subject"`
}

var teachers = make(map[int]Teacher)
var mutex = &sync.Mutex{}
var nextID = 1

func init() {
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Math",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Kola",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextID++
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

	/*rl := mw.NewRateLimiter(5, time.Minute)

	hppOptions := mw.HPPOptions{
		CheckQuery:              true,
		CheckBody:               true,
		CheckBodyForContentType: "application/x-www-form-urlencoded",
		Whitelist:               []string{"sortBy", "sortOrder", "name", "age", "class"},
	}*/

	//secureMux := applyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	secureMux := mw.SecurityHeaders(mux)
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

// middleware is a function that hhtp.Handler with additional Functanality
type Middleware func(http.Handler) http.Handler

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
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
	switch r.Method {

	case http.MethodGet:
		w.Write([]byte("Hello Get Method in Student Route"))

	case http.MethodPost:
		w.Write([]byte("Hello Post Method in Student Route"))

	case http.MethodPut:
		w.Write([]byte("Hello Put Method in Student Route"))

	case http.MethodPatch:
		w.Write([]byte("Hello Patch Method in Student Route"))

	case http.MethodDelete:
		w.Write([]byte("Hello Delete Method in Student Route"))

	default:
		w.Write([]byte("Hello Unknown Method in Student Route"))
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Write([]byte("Hello Get Method in Execs Route"))

	case http.MethodPost:
		w.Write([]byte("Hello Post Method in Execs Route"))
	case http.MethodPut:
		w.Write([]byte("Hello Put Method in Execs Route"))

	case http.MethodPatch:
		w.Write([]byte("Hello Patch Method in Execs Route"))

	case http.MethodDelete:
		w.Write([]byte("Hello Delete Method in Execs Route"))

	default:
		w.Write([]byte("Hello Unknown Method in Execs Route"))
	}
}

func teacherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Teachers Route")
	fmt.Println("Method :", r.Method)

	switch r.Method {

	case http.MethodGet:
		getTeacherHandler(w, r)
		/*w.Write([]byte("Hello Get Method in Teachers Route"))
		parsePathParameters(w, r)
		parseQueryParameters(w, r)*/

	case http.MethodPost:
		addTeacherHandler(w, r)
		/*w.Write([]byte("Hello Post Method in Teachers Route"))
		parseFormElement(w, r)
		parseRawBodyElement(w, r)*/

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

// ------------------Teacher Handlers------------------------------//
// -----Get Teacher Handlers----//
func getTeacherHandler(w http.ResponseWriter, r *http.Request) {

	// for extracting ID from url or path for filtering data
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	fmt.Println("ID : ", idStr)

	if idStr == "" {

		//https://localhost:3000/teachers/?first_name=Jane
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teachersList := make([]Teacher, 0, len(teachers))

		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
				teachersList = append(teachersList, teacher)
			}
		}

		response := struct {
			Status string    `json:"status"`
			Count  int       `json:"count"`
			Data   []Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teachers),
			Data:   teachersList,
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Error Encoding JSON", http.StatusPartialContent)
			return
		}
	} else {

		//https://localhost:3000/teachers/2
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Error 78:", err)
		}

		teacher, exists := teachers[id]
		if !exists {
			http.Error(w, "Teacher Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(teacher)
		if err != nil {
			http.Error(w, "Error Encoding JSON", http.StatusPartialContent)
			return
		}
	}

}

//-----Get Teacher Handlers----//

// -----POST Teacher Handlers----//
func addTeacherHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var newTeachers []Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	addedTeachers := make([]Teacher, len(newTeachers))

	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextID
		teachers[nextID] = newTeacher
		addedTeachers[i] = newTeacher
		nextID++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := struct {
		Status string    `json:"status"`
		Count  int       `json:"count"`
		Data   []Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}

	json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error Encoding JSON response", http.StatusPartialContent)
		return
	}

}

//-----POST Teacher Handlers----//

//------------------Teacher Handlers------------------------------//
