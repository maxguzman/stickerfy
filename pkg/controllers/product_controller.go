package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"stickerfy/app/models"
	"stickerfy/app/services"
	"stickerfy/pkg/platform/cache"
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
// @Description Get all exists products.
// @Summary get all exists products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func (pc *productController) GetAll(c *fiber.Ctx) error {
	cachedProducts, err := pc.productsCache.Get(pc.context, "products")
	if err == redis.Nil {
		products, err := pc.productService.GetAll(pc.context)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"products": nil,
				"error":    true,
				"msg":      "there where no products found",
			})
		}
		encodedProducts, err := json.Marshal(products)
		if err != nil {
			return err
		}
		err = pc.productsCache.Set(pc.context, "products", encodedProducts, time.Minute*30)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"products": nil,
				"error":    true,
				"msg":      "there was an error caching the products",
			})
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"products": products,
			"error":    false,
			"msg":      nil,
		})
	}

	if err != nil {
		fmt.Printf("there was an error getting the products from cache: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"products": nil,
			"error":    true,
			"msg":      "there was an error getting the products from cache",
		})
	}

	var products []models.Product
	err = json.Unmarshal([]byte(cachedProducts), &products)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"products": nil,
			"error":    true,
			"msg":      "there was an error unmarshaling the products",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"products": products,
		"error":    false,
		"msg":      nil,
	})
}

// Get returns a product by id
// @Description Get product by given ID.
// @Summary get product by given ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Router /product/{id} [get]
func (pc *productController) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"product": nil,
			"error":   true,
			"msg":     "invalid id",
		})
	}

	product, err := pc.productService.GetByID(pc.context, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"product": nil,
			"error":   true,
			"msg":     "there was an error getting the product",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"product": product,
		"error":   false,
		"msg":     nil,
	})
}

// New creates a new product
// @Description Create a new product.
// @Summary create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param image_path body string true "Image Path"
// @Param title body string true "Title"
// @Param description body string true "Description"
// @Param price body number true "Price"
// @Success 200 {object} models.Product
// @Router /product [post]
func (pc *productController) Post(c *fiber.Ctx) error {
	var product models.Product
	product.ID = uuid.New()
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"msg":     "invalid product",
			"product": nil,
		})
	}
	err := pc.productService.Post(pc.context, product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"product": nil,
			"error":   true,
			"msg":     "there was an error creating the product",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"product": product,
		"error":   false,
		"msg":     nil,
	})
}

// Update updates a product
// @Description Update product.
// @Summary update product
// @Tags Product
// @Accept json
// @Produce json
// @Param id body string true "Product ID"
// @Param image_path body string true "Image Path"
// @Param title body string true "Title"
// @Param description body string true "Description"
// @Param price body number true "Price"
// @Success 202 {string} status "ok"
// @Router /product [put]
func (pc *productController) Update(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"msg":     "invalid product",
			"product": nil,
		})
	}

	err := pc.productService.Update(pc.context, product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"product": nil,
			"error":   true,
			"msg":     "there was an error updating the product",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"product": product,
		"error":   false,
		"msg":     nil,
	})
}

// Delete deletes a product
// @Description Delete product.
// @Summary delete product
// @Tags Product
// @Accept json
// @Produce json
// @Param id body string true "Product ID"
// @Param image_path body string true "Image Path"
// @Param title body string true "Title"
// @Param description body string true "Description"
// @Param price body number true "Price"
// @Success 204 {string} status "ok"
// @Router /product [delete]
func (pc *productController) Delete(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"msg":     "invalid product",
			"product": nil,
		})
	}
	err := pc.productService.Delete(pc.context, product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"product": nil,
			"error":   true,
			"msg":     "there was an error deleting the product",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"product": product,
		"error":   false,
		"msg":     nil,
	})
}
