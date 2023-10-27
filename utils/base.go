package utils

import "net/http"

// HelloWorld api Handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
