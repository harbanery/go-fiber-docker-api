package controllers

import (
	"fmt"
	"go-fiber-docker-api/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {

	products := models.GetAllProducts()
	if len(products) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "no content",
			"statusCode": 202,
			"message":    "Product is empty.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       products,
	})
}

func GetDetailProduct(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	product := models.GetProductbyID(id)
	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Product not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       product,
	})
}

func AddProduct(c *fiber.Ctx) error {

	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})
	}

	if err := models.CreateProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    "Failed to create product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"message":    "Product created successfully",
	})
}

func EditProduct(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	existProduct := models.GetProductbyID(id)
	if existProduct.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Product not found",
		})
	}

	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid request body",
		})
	}

	if err := models.UpdateProduct(id, &product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    fmt.Sprintf("Failed to update product with ID %d", id),
		})
	} else {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":     "success",
			"statusCode": 200,
			"message":    fmt.Sprintf("Product with ID %d updated successfully", id),
		})
	}
}

func DeleteProduct(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":     "bad request",
			"statusCode": 400,
			"message":    "Invalid ID format",
		})
	}

	product := models.GetProductbyID(id)
	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":     "not found",
			"statusCode": 404,
			"message":    "Product not found",
		})
	}

	if err := models.DeleteProduct(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":     "server error",
			"statusCode": 500,
			"message":    fmt.Sprintf("Failed to delete product with ID %d", id),
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":     "success",
			"statusCode": 200,
			"message":    fmt.Sprintf("Product with ID %d deleted successfully", id),
		})
	}
}
