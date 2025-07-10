package handlers

import "net/http"

func StudentHandler(w http.ResponseWriter, r *http.Request) {
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
