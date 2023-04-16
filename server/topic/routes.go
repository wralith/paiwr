package topic

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Routes struct {
	repo Repo
}

func NewRoutes(repo Repo) *Routes {
	return &Routes{repo: repo}
}

type CreateInput struct {
	Capacity int      `json:"capacity"`
	Category Category `json:"category"`
	Title    string   `json:"title"`
}

// Create
//
//	@ID			Topic-Create
//	@Summary	Create Topic
//	@Tags		topics
//	@Security	BearerAuth
//	@Accept		json
//	@Produce	json
//	@Param		options	body	CreateInput	true	"New Topic Information"
//	@Success	201
//	@Failure	400
//	@Failure	500
//	@Router		/topics [post]
func (r *Routes) Create(c *fiber.Ctx) error {
	var input CreateInput
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Unable to parse request"})
	}
	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	owner, err := uuid.Parse(user["sub"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Invalid UUID"})
	}

	topic := CreateTopic(CreateTopicOpts{
		Owner:    owner,
		Capacity: input.Capacity,
		Category: input.Category,
		Title:    input.Title,
	})

	err = r.repo.Save(context.Background(), &topic)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	return c.SendStatus(fiber.StatusCreated)
}

// FindByID
//
//	@ID			Topic-Find
//	@Summary	Find Topic By ID
//	@Tags		topics
//	@Produce	json
//	@Param		id	path		string	true	"Topic ID"
//	@Success	200	{object}	Topic
//	@Failure	400
//	@Failure	404
//	@Failure	500
//	@Router		/topics/{id} [get]
func (r *Routes) FindByID(c *fiber.Ctx) error {
	input := c.Params("id")
	id, err := uuid.Parse(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Invalid UUID"})
	}

	topic, err := r.repo.FindByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]string{"message": fmt.Sprintf("topic with id:`%s` not found", input)})
	}

	return c.Status(fiber.StatusOK).JSON(topic)
}

// FindByOwnerID
//
//	@ID			Topic-Find-Owned
//	@Summary	Find Topics By Owner ID
//	@Tags		topics
//	@Produce	json
//	@Param		id	path		string	true	"Owner ID"
//	@Success	200	{object}	[]Topic
//	@Failure	400
//	@Failure	500
//	@Router		/topics/owner/{id} [get]
func (r *Routes) FindByOwnerID(c *fiber.Ctx) error {
	input := c.Params("id")
	id, err := uuid.Parse(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Invalid UUID"})
	}

	topic, err := r.repo.FindByOwner(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	return c.Status(fiber.StatusOK).JSON(topic)
}

// FindInvolved
//
//	@ID			Topic-Find-Paired
//	@Summary	Find Paired Topics By User ID
//	@Tags		topics
//	@Produce	json
//	@Param		id	path		string	true	"Involved ID"
//	@Success	200	{object}	[]Topic
//	@Failure	400
//	@Failure	500
//	@Router		/topics/pair/{id} [get]
func (r *Routes) FindInvolved(c *fiber.Ctx) error {
	input := c.Params("id")
	id, err := uuid.Parse(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Invalid UUID"})
	}

	topic, err := r.repo.FindInvolved(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	return c.Status(fiber.StatusOK).JSON(topic)
}

// Delete
//
//	@ID			Topic-Delete
//	@Summary	Delete Topic
//	@Tags		topics
//	@Security	BearerAuth
//	@Produce	json
//	@Param		id	path	string	true	"Topic ID"
//	@Success	200
//	@Failure	400
//	@Failure	404
//	@Failure	500
//	@Router		/topics/{id} [delete]
func (r *Routes) Delete(c *fiber.Ctx) error {
	input := c.Params("id")
	id, err := uuid.Parse(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Invalid UUID"})
	}

	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	owner, err := uuid.Parse(user["sub"].(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": "Invalid UUID"})
	}

	topic, err := r.repo.FindByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]string{"message": fmt.Sprintf("topic with id:`%s` not found", input)})
	}

	if topic.Owner != owner {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	err = r.repo.Delete(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": "Unknown error"})
	}

	return c.SendStatus(fiber.StatusOK)
}
