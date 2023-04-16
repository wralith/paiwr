package user

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// TODO: Update Email, Password, Bio etc.
// TODO: Fix repetitions, especially on error handling

type Routes struct {
	repo      Repo
	jwtSecret string
}

func NewRoutes(repo Repo, jwtSecret string) *Routes {
	return &Routes{repo: repo, jwtSecret: jwtSecret}
}

// Register
//
//	@ID			User-Register
//	@Summary	Register New User
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		options	body	CreateUserOpts	true	"New User Info"
//	@Success	201
//	@Failure	400
//	@Failure	500
//	@Router		/users/register [post]
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

type LoginResult struct {
	ID       uuid.UUID `json:"id"`
	Token    string    `json:"token"`
	Username string    `json:"username"`
	Exp      int64     `json:"exp"`
}

// Login
//
//	@ID			User-Login
//	@Summary	Login With Credentials
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		credentials	body	LoginInput	true	"User Credentials"
//	@Success	200
//	@Failure	400
//	@Failure	500
//	@Router		/users/login [post]
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

	user, err := r.repo.FindByUsername(context.Background(), input.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	exp := time.Now().Add(time.Hour).Unix()
	claims := jwt.MapClaims{"name": user.Username, "sub": user.ID, "exp": exp}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	return c.Status(fiber.StatusOK).JSON(LoginResult{ID: user.ID, Token: token, Username: user.Username, Exp: exp})
}

// FindByID
//
//	@ID			User-FindByID
//	@Summary	Find User By ID
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"User ID"
//	@Success	200	{object}	User
//	@Failure	400
//	@Failure	500
//	@Router		/users/{id} [get]
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
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

// UpdatePassword
//
//	@ID			User-Update-Password
//	@Summary	Upate Password
//	@Tags		users
//	@Security	BearerAuth
//	@Accept		json
//	@Produce	json
//	@Param		credentials	body		UpdatePasswordInput	true	"User Credentials and New Password"
//	@Success	200			{object}	User
//	@Failure	400
//	@Failure	500
//	@Router		/users/update-password [patch]
func (r *Routes) UpdatePassword(c *fiber.Ctx) error {
	var input UpdatePasswordInput
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Unable to parse request"})
	}

	userClaims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	owner, err := uuid.Parse(userClaims["sub"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Invalid UUID"})
	}

	user, err := r.repo.FindByID(context.Background(), owner)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	// FIXME: WHY Query WHY?
	err = r.repo.VerifyPassword(context.Background(), user.Username, input.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Username and password does not match"})
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
