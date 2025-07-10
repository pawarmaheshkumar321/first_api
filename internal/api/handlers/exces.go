package handlers

import "net/http"

func ExecsHandler(w http.ResponseWriter, r *http.Request) {
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
