package ports

import (
	"net/http"

	"leal.co/internal/pkg/unleash"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var feature bool = unleash.IsEnabled()
	if feature {
		w.Write([]byte("hello world"))
	} else {
		w.Write([]byte("El feature no está activo"))
	}

}
