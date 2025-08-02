package products

import (
	"net/http"
	"order-api/pkg/request"
	"order-api/pkg/response"
	"strconv"

	"gorm.io/gorm"
)

// Структура хендлера
type ProductsHandler struct {
	Repository *ProductsRepository
}

// Создание хендлера и инит роутов
func NewProductsHandler(router *http.ServeMux, repo *ProductsRepository) {
	handler := &ProductsHandler{
		Repository: repo,
	}
	router.HandleFunc("POST /product", handler.Create)
	router.HandleFunc("PATCH /product/{id}", handler.Patch)
	router.HandleFunc("GET /product/{id}", handler.GetByID)
	router.HandleFunc("DELETE /product/{id}", handler.Delete)
	router.HandleFunc("GET /products", handler.GetAll)

}

// Создание продукта
func (h *ProductsHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := request.HandleBody[ProductCreateRequest](&w, r)
	if err != nil {
		return
	}
	product := NewProduct(body.Name, body.Description, body.Images)
	productCreated, err := h.Repository.Create(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response.Json(w, productCreated, http.StatusCreated)

}

// Изменение продукта
func (h *ProductsHandler) Patch(w http.ResponseWriter, r *http.Request) {
	body, err := request.HandleBody[ProductCreateRequest](&w, r)
	if err != nil {
		return
	}
	idString := r.PathValue("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	productCreated, err := h.Repository.Update(&Product{
		Model:       gorm.Model{ID: uint(id)},
		Name:        body.Name,
		Description: body.Description,
		Images:      body.Images,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response.Json(w, productCreated, http.StatusCreated)
}

// Получение продукта по ID
func (h *ProductsHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := h.Repository.GetById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	response.Json(w, product, http.StatusOK)

}

// Удаление продукта
func (h *ProductsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = h.Repository.GetById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = h.Repository.Delete(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.Json(w, nil, http.StatusOK)

}

// Получение продуктов
func (h *ProductsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.Repository.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	response.Json(w, products, http.StatusOK)

}
