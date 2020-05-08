package src

import "net/http"

var mux = &http.ServeMux{}

func init() {
	mux.HandleFunc("/parking/haikang/ipr/plate", plate)
}

// Mux api mux
func Mux() *http.ServeMux {
	return mux
}
