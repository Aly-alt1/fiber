package handler

import (
	"fiber3-app/backend/domain"
	"fiber3-app/backend/repository"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type MouseHandler struct {
	repository *repository.MouseRepository
}

func NewMouseHandler() *MouseHandler {
	return &MouseHandler{repository: repository.NewMouseRepository()}
}

func (h *MouseHandler) Create(c fiber.Ctx) error {
	var mouse domain.Mouse

	if err := c.Bind().Body(&mouse); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	mouse.ID = uuid.NewString()
	if err := h.repository.Create(&mouse); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(strings.Split(err.Error(), "\n")[2])
	}

	return c.SendString(mouse.ID)
}

func (h *MouseHandler) GetAll(c fiber.Ctx) error {
	mice := h.repository.GetAll()

	return c.JSON(mice)
}

func (h *MouseHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")

	data := h.repository.GetByID(id)
	return c.JSON(data)
}

func (h *MouseHandler) Update(c fiber.Ctx) error {
	var mouse domain.Mouse

	if err := c.Bind().Body(&mouse); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := h.repository.Update(&mouse); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(strings.Split(err.Error(), "\n")[2])
	}

	return c.SendString("success")
}
