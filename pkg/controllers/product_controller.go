package controllers

import (
	"net/http"
	"stickerfy/app/models"
	"stickerfy/app/services"

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
}

// NewProductController creates a new ProductController
func NewProductController(productService services.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

// GetAll returns all products
func (pc *productController) GetAll(c *fiber.Ctx) error {
	products, err := pc.productService.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"products": nil,
			"error":    true,
			"msg":      "there where no products found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"products": products,
		"error":    false,
		"msg":      nil,
	})
}

// Get returns a product by id
func (pc *productController) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"product": nil,
			"error":   true,
			"msg":     "invalid id",
		})
	}

	product, err := pc.productService.GetByID(id)
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
func (pc *productController) Post(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"msg":     "invalid product",
			"product": nil,
		})
	}
	product.ID = uuid.New()
	err := pc.productService.Post(product)
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

// Update updates a product by id
func (pc *productController) Update(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"msg":     "invalid product",
			"product": nil,
		})
	}

	err := pc.productService.Update(product)
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

// Delete deletes a product by id
func (pc *productController) Delete(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"msg":     "invalid product",
			"product": nil,
		})
	}
	err := pc.productService.Delete(product)
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
