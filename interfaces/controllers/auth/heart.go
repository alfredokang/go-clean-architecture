package auth

import (
	"net/http"
)

func Heart(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		return
	}
}
