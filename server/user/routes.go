package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// TODO: Update Email, Password, Bio etc.
// TODO: Fix repetitions, especially on error handling

type Routes struct {
	repo Repo
}

func NewRoutes(repo Repo) *Routes {
	return &Routes{repo: repo}
}

func (r *Routes) Register(c *fiber.Ctx) error {
	var input CreateUserOpts
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Unable to parse request"})
	}

	user, err := CreateUser(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	err = r.repo.Save(context.Background(), &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	return c.SendStatus(fiber.StatusCreated)
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TODO: Send Token
func (r *Routes) Login(c *fiber.Ctx) error {
	var input LoginInput
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Unable to parse request"})
	}

	err = r.repo.VerifyPassword(context.Background(), input.Username, input.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Username and password does not match"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (r *Routes) FindByID(c *fiber.Ctx) error {
	input := c.Params("id")
	id, err := uuid.Parse(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Invalid UUID"})
	}

	user, err := r.repo.FindByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	return c.Status(fiber.StatusOK).JSON(&user)
}

type UpdatePasswordInput struct {
	Username    string `json:"username"` // TODO: From token?
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

func (r *Routes) UpdatePassword(c *fiber.Ctx) error {
	var input UpdatePasswordInput
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Unable to parse request"})
	}

	err = r.repo.VerifyPassword(context.Background(), input.Username, input.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Username and password does not match"})
	}

	user, err := r.repo.FindByUsername(context.Background(), input.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	err = user.UpdatePassword(input.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	err = r.repo.Update(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	return c.SendStatus(fiber.StatusOK)
}
