package topic

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Category string

const (
	Software       Category = "software"
	SocialSciences Category = "social_sciences"
	Other          Category = "other"
)

// Topic represents a topic to study and ID's of users involved
type Topic struct {
	ID       uuid.UUID `json:"id"`
	Category Category  `json:"category"`
	Title    string    `json:"title"`

	// Max expected parties
	Capacity int `json:"capacity"`

	// ID of the owner User
	Owner uuid.UUID `json:"owner"`

	// IDs of invloved Users
	Parties []uuid.UUID `json:"parties"`

	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	FinishedAt time.Time `json:"finished_at"`
}

// TODO: Validations?

// CreateTopicOpts are options struct to create a new Topic
type CreateTopicOpts struct {
	Category Category  `json:"category"`
	Title    string    `json:"title"`
	Capacity int       `json:"capacity"`
	Owner    uuid.UUID `json:"owner"`
}

// CreateTopic creates a new Topic with new ID
func CreateTopic(opts CreateTopicOpts) Topic {
	now := time.Now()
	return Topic{
		ID:        uuid.New(),
		Category:  opts.Category,
		Title:     opts.Title,
		Capacity:  opts.Capacity,
		Owner:     opts.Owner,
		Parties:   []uuid.UUID{opts.Owner},
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (t *Topic) AddParties(ids ...uuid.UUID) error {
	emptySpots := t.Capacity - len(t.Parties)
	if len(ids) > emptySpots {
		return fmt.Errorf("error while adding %d parties to topic with %d empty spots", len(ids), emptySpots)
	}
	for _, id := range ids {
		if !slices.Contains(t.Parties, id) {
			t.Parties = append(t.Parties, id)
		}
	}
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Topic) RemoveParties(ids ...uuid.UUID) {
	for _, id := range ids {
		if i := slices.Index(t.Parties, id); i != -1 {
			t.Parties = slices.Delete(t.Parties, i, i+1)
		}
	}
	t.UpdatedAt = time.Now()
}

func (t *Topic) UpdateCapacity(newCapactiy int) {
	t.Capacity = newCapactiy
	t.UpdatedAt = time.Now()
}

func (t *Topic) Finish() {
	now := time.Now()
	if t.FinishedAt.IsZero() {
		t.FinishedAt = now
		t.UpdatedAt = now
	}
}

func (t *Topic) IsFinished() bool {
	return !t.FinishedAt.IsZero()
}
