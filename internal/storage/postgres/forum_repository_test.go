package postgres

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/doutivity/research-forum-go/internal/domain"
	"github.com/doutivity/research-forum-go/schema"

	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

const (
	dataSourceName = "postgresql://user:secretpassword@postgres1:5432/forum-db?sslmode=disable&timezone=UTC"
)

func TestForumRepositoryTopics(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	connection, err := sql.Open("postgres", dataSourceName)
	require.NoError(t, err)
	defer connection.Close()

	require.NoError(t, schema.MigrateUp(connection))
	defer func() {
		require.NoError(t, schema.MigrateDown(connection))
	}()

	repository, err := NewRepository(connection)
	require.NoError(t, err)
	defer repository.Close()

	forumRepository := NewForumRepository(repository)

	var (
		expectedTopic = &domain.Topic{
			ID:        1,
			Title:     "DOU Forum on PostgreSQL",
			Content:   "https://github.com/doutivity/research-forum-go",
			CreatedAt: time.Now().Truncate(time.Second).UTC(),
			Author: &domain.TopicAuthor{
				ID:       1,
				Username: "Admin",
			},
		}
	)

	id, err := forumRepository.TopicCreate(context.Background(), &domain.TopicCreate{
		Title:   expectedTopic.Title,
		Content: expectedTopic.Content,
	}, expectedTopic.CreatedAt, 1)
	require.NoError(t, err)
	require.Equal(t, expectedTopic.ID, id)

	topics, err := forumRepository.Topics(context.Background(), 30, 0)
	require.NoError(t, err)
	require.Equal(t, []*domain.Topic{expectedTopic}, topics)

	// @TODO get topic by id
}

func TestForumRepositoryComments(t *testing.T) {
	// @TODO
	// add comment 1
	// add comment 2
	// get comments by topic

	// update comment 2
	// get comments by topic
}

func TestForumRepositoryLikes(t *testing.T) {
	// @TODO
	// like comment 1
	// like comment 2
	// unlike comment 1
	// get active likes for comments
}
