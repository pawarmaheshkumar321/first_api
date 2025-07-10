package router

import (
	"first_api/internal/api/handlers"
	"net/http"
)

func Router() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)

	mux.HandleFunc("/teachers/", handlers.TeacherHandler)

	mux.HandleFunc("/students/", handlers.StudentHandler)

	mux.HandleFunc("/execs/", handlers.ExecsHandler)

	return mux
}
