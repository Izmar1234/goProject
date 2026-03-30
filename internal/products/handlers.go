package products

import (
	"net/http"

	"github.com/Izmar1234/goProject/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {

	products := []string{"Hello", "World"}
	json.Write(w, http.StatusOK, products)

}
