package controllers

import (
	"net/http"

	"github.com/PrinceNarteh/go-users/models"
	"github.com/PrinceNarteh/go-users/repository"
	"github.com/PrinceNarteh/go-users/responses"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	user, err := repository.GetUserById(userId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:     "error",
			StatusCode: http.StatusBadRequest,
			Data:       models.User{},
		})
	}
	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: "success", StatusCode: http.StatusOK, Data: user})
}
