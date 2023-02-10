package controllers

import (
	"encoding/json"
	"net/http"
	"stickerfy/app/models"
	"stickerfy/app/services"
)

// ProductController is an interface for a product controller
type ProductController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

// productController is a implementation of ProductController
type productController struct {
	productService services.ProductService
}

// NewProductController creates a new ProductController
func NewProductController(productService services.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

// GetAll returns all products
func (pc *productController) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := pc.productService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Get returns a product by id
func (pc *productController) GetByID(w http.ResponseWriter, r *http.Request) {
	product, err := pc.productService.GetByID(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(product); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// New creates a new product
func (pc *productController) Post(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := pc.productService.Post(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Update updates a product by id
func (pc *productController) Update(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := pc.productService.Update(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete deletes a product by id
func (pc *productController) Delete(w http.ResponseWriter, r *http.Request) {
	err := pc.productService.Delete(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
