package topic

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var dummyUUID = uuid.New()
var dummyOpts = CreateTopicOpts{
	Category: Software,
	Title:    "Testing",
	Capacity: 4,
	Owner:    dummyUUID,
}

func TestCreateTopic(t *testing.T) {
	topic := CreateTopic(dummyOpts)
	assert.Equal(t, dummyUUID, topic.Parties[0])
	assert.WithinDuration(t, time.Now(), topic.CreatedAt, time.Millisecond)
}

func TestAddAndRemoveParties(t *testing.T) {
	topic := CreateTopic(dummyOpts)
	p1 := uuid.New()
	p2 := uuid.New()
	p3 := uuid.New()
	_ = topic.AddParties(p1, p2, p3)
	err := topic.AddParties(p1)
	assert.Error(t, err)
	assert.Contains(t, topic.Parties, p1)
	assert.Contains(t, topic.Parties, p2)
	topic.RemoveParties(p1, p2)
	assert.NotContains(t, topic.Parties, p1)
	assert.NotContains(t, topic.Parties, p2)
	assert.Contains(t, topic.Parties, p3)
}

func TestUpdateCapacity(t *testing.T) {
	topic := CreateTopic(dummyOpts)
	topic.UpdateCapacity(4)
	assert.Equal(t, 4, topic.Capacity)
	assert.WithinDuration(t, time.Now(), topic.UpdatedAt, time.Millisecond)
}

func TestFinish(t *testing.T) {
	topic := CreateTopic(dummyOpts)
	assert.False(t, topic.IsFinished())
	topic.Finish()
	assert.WithinDuration(t, time.Now(), topic.FinishedAt, time.Millisecond)
	assert.True(t, topic.IsFinished())
}
