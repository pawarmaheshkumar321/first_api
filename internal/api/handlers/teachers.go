package handlers

import (
	"encoding/json"
	"first_api/internal/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var teachers = make(map[int]models.Teacher)
var mutex = &sync.Mutex{}
var nextID = 1

func TeacherHandler(w http.ResponseWriter, r *http.Request) {
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

		teachersList := make([]models.Teacher, 0, len(teachers))

		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
				teachersList = append(teachersList, teacher)
			}
		}

		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   []models.Teacher `json:"data"`
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

	var newTeachers []models.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	addedTeachers := make([]models.Teacher, len(newTeachers))

	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextID
		teachers[nextID] = newTeacher
		addedTeachers[i] = newTeacher
		nextID++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error Encoding JSON response", http.StatusPartialContent)
		return
	}

}

//-----POST Teacher Handlers----//

//------------------Teacher Handlers------------------------------//

func init() {
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Math",
	}
	nextID++
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextID++
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Kola",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextID++
}
