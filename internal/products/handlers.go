package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sebasL200/go-api-ecommerce.git/internal/json"
)

type handler struct {
	service Service
}

// CONSTRUCTOR
func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {

	products,err := h.service.ListProducts(r.Context())
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


  
    json.Write(w, http.StatusOK, products)
}

func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	
	idStr := chi.URLParam(r, "id") 

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("Error al convertir ID:", err)
		http.Error(w, "ID de producto inválido", http.StatusBadRequest)
		return
	}

	// 3. Llamar al servicio
	product, err := h.service.GetProduct(r.Context(), id)
	if err != nil {
		log.Println("Error al obtener producto:", err)
		// Aquí asumimos que si falla es porque no existe
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	json.Write(w, http.StatusOK, product)
}