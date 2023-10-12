package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"stickerfy/app/models"
	"stickerfy/app/services"
	"stickerfy/pkg/platform/cache"
	"stickerfy/pkg/utils"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ProductController is an interface for a product controller
type ProductController interface {
	GetAll(*fiber.Ctx) error
	GetByID(*fiber.Ctx) error
	Post(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
	Update(*fiber.Ctx) error
}

// productController is a implementation of ProductController
type productController struct {
	productService services.ProductService
	productsCache  cache.Cache
	context        context.Context
}

// NewProductController creates a new ProductController
func NewProductController(ctx context.Context, productService services.ProductService, productCache cache.Cache) ProductController {
	return &productController{
		productService: productService,
		productsCache:  productCache,
		context:        ctx,
	}
}

// GetAll returns all products
//
// @Summary		get all exists products
// @Description	Get all exists products.
// @ID				get-all-products
// @Tags			products
// @Accept			json
// @Produce		json
// @Success		200	{object}	utils.HTTPProducts
// @Failure		400	{object}	utils.HTTPError
// @Failure		500	{object}	utils.HTTPError
// @Router			/products [get]
func (pc *productController) GetAll(c *fiber.Ctx) error {
	cachedProducts, err := pc.productsCache.Get(pc.context, "products")
	if err == redis.Nil {
		products, err := pc.productService.GetAll(pc.context)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "there was an error getting the products",
			})
		}
		if products == nil {
			return c.Status(http.StatusNotFound).JSON(utils.HTTPError{
				Code:    http.StatusNotFound,
				Message: "there where no products found",
			})
		}
		encodedProducts, err := json.Marshal(products)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "there was an error marshaling the products",
			})
		}
		err = pc.productsCache.Set(pc.context, "products", encodedProducts, time.Minute*5)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "there was an error setting the products in cache",
			})
		}
		return c.Status(http.StatusOK).JSON(utils.HTTPProducts{
			Products: products,
		})
	}

	if err != nil {
		fmt.Printf("there was an error getting the products from cache: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "there was an error getting the products from cache",
		})
	}

	var products []models.Product
	err = json.Unmarshal([]byte(cachedProducts), &products)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "there was an error unmarshaling the products",
		})
	}

	return c.Status(http.StatusOK).JSON(utils.HTTPProducts{
		Products: products,
	})
}

// Get returns a product by id
//
//	@Summary		get product by given ID
//	@Description	Get product by given ID.
//	@ID				get-product-by-id
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Product ID" format(uuid) extensions(x-example=4a7cdb5c-bd2a-47f6-9d4a-3531b110d26d)
//	@Success		200	{object}	models.Product
//	@Failure		400	{object}	utils.HTTPError
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/products/{id} [get]
func (pc *productController) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid product id",
		})
	}

	product, err := pc.productService.GetByID(pc.context, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "there was an error getting the product",
		})
	}

	return c.Status(http.StatusOK).JSON(product)
}

// New creates a new product
//
//	@Summary		create a new product
//	@Description	Create a new product.
//	@ID				create-product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		models.Product	true	"Product"
//	@Success		201		{object}	models.Product
//	@Failure		400		{object}	utils.HTTPError
//	@Failure		500		{object}	utils.HTTPError
//	@Router			/products [post]
func (pc *productController) Post(c *fiber.Ctx) error {
	var product models.Product
	product.ID = uuid.New()
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid product",
		})
	}
	err := pc.productService.Post(pc.context, product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "there was an error creating the product",
		})
	}

	return c.Status(http.StatusCreated).JSON(product)
}

// Update updates a product
//
//	@Summary		update product
//	@Description	Update product.
//	@ID				update-product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		models.Product	true	"Product"
//	@Success		200	{object}	models.Product
//	@Failure		400	{object}	utils.HTTPError
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/products [put]
func (pc *productController) Update(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid product",
		})
	}

	err := pc.productService.Update(pc.context, product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "there was an error updating the product",
		})
	}

	return c.Status(http.StatusOK).JSON(product)
}

// Delete deletes a product
//
//	@Summary		delete product
//	@Description	Delete product.
//	@ID				delete-product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		models.Product	true	"Product"
//	@Success		200	{object}	models.Product
//	@Failure		400	{object}	utils.HTTPError
//	@Failure		500	{object}	utils.HTTPError
//	@Router			/products [delete]
func (pc *productController) Delete(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(utils.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid product",
		})
	}
	err := pc.productService.Delete(pc.context, product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "there was an error deleting the product",
		})
	}

	return c.Status(http.StatusOK).JSON(product)
}
