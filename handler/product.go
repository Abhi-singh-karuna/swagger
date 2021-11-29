package handler

import (
	"path/database"
	"path/model"

	"github.com/gofiber/fiber/v2"
)

//
func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "ya this api is running  good !! ",
	})
}

// for getting all products
func GetAllProducts(c *fiber.Ctx) error {
	db := database.DB
	var products []model.Product
	db.Find(&products)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All products",
		"data":    products,
	})
}

// for getting single products by ID
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var product model.Product
	db.Find(&product, id)
	if product.Title == "" {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "No product found with ID",
			"data":    nil,
		})

	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product found",
		"data":    product,
	})
}

// creating a new products in database
func CreateProduct(c *fiber.Ctx) error {
	db := database.DB
	product := new(model.Product)
	if err := c.BodyParser(product); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create product",
			"data":    err,
		})
	}
	db.Create(&product)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created product",
		"data":    product,
	})
}

// DeleteProduct by the products id
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var product model.Product
	db.First(&product, id)
	if product.Title == "" {
		return c.JSON(fiber.Map{"status": "error",
			"message": "No product found with ID",
			"data":    nil,
		})

	}
	db.Delete(&product)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product successfully deleted",
		"data":    nil,
	})
}
