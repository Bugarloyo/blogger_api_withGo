package controllers

import (
	"blogger-api/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)


type BlogController struct {
	service *services.BlogService
}

func NewBlogController(service *services.BlogService) * BlogController {
	return &BlogController{service: service}
}

type BlogRequest struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Body string `json:"body"`
}

func (ctrl *BlogController) CreateBlog(c *fiber.Ctx) error {
	var req BlogRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	blog, err := ctrl.service.CreateBlog(req.Title, req.Author, req.Body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(blog)
}

func (ctrl *BlogController) GetAll(c *fiber.Ctx) error {
	blogs, err := ctrl.service.GetAllBlogs()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(blogs)
}

func (ctrl *BlogController) GetById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog, err := ctrl.service.GetBlogById(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Blog tidak ditemukan"})
	}
	return c.JSON(blog)
}

func (ctrl *BlogController) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req BlogRequest
	if err := c.BodyParser(&req); err !=nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	blog, err := ctrl.service.UpdateBlog(uint(id), req.Title, req.Author, req.Body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(blog)
}

func (ctrl *BlogController) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := ctrl.service.DeleteBlog(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Blog berhasil dihapus"})
}
