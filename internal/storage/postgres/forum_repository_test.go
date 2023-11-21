package postgres

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/doutivity/research-forum-go/internal/domain"
	"github.com/doutivity/research-forum-go/internal/storage/postgres/dbs"
	"github.com/doutivity/research-forum-go/schema"

	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

const (
	dataSourceName = "postgresql://user:secretpassword@postgres1:5432/forum-db?sslmode=disable&timezone=UTC"
)

var (
	forumRepository *ForumRepository
	connection      *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	connection, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	err = schema.MigrateUp(connection)
	if err != nil {
		log.Fatalf("Failed to migrate up: %v", err)
	}

	repository, err := NewRepository(connection)
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	forumRepository = NewForumRepository(repository)

	code := m.Run()

	err = schema.MigrateDown(connection)
	if err != nil {
		log.Fatalf("Failed to migrate down: %v", err)
	}

	repository.Close()
	connection.Close()

	os.Exit(code)
}

func TestForumRepositoryTopics(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

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

	TopicByID, err := forumRepository.TopicByID(context.Background(), 1)
	require.NoError(t, err)
	require.Equal(t, expectedTopic, TopicByID)
}

func TestForumRepositoryComments(t *testing.T) {
	// add comment 1
	expectedComment1 := &domain.Comment{
		ID:              1,
		ParentCommentID: nil,
		Content:         "Great topic",
		CreatedAt:       time.Now().Truncate(time.Second).UTC(),
		Author: &domain.CommentAuthor{
			ID:       1,
			Username: "Admin",
		},
	}

	id1, err := forumRepository.CommentCreate(context.Background(), &domain.CommentCreate{
		ParentCommentID: expectedComment1.ParentCommentID,
		Content:         expectedComment1.Content,
		TopicID:         1,
	}, expectedComment1.CreatedAt, 1)
	require.NoError(t, err)
	require.Equal(t, expectedComment1.ID, id1)

	// add comment 2
	parentComentID := int64(1)
	expectedComment2 := &domain.Comment{
		ID:              2,
		ParentCommentID: &parentComentID,
		Content:         "no doubt",
		CreatedAt:       time.Now().Truncate(time.Second).UTC(),
		Author: &domain.CommentAuthor{
			ID:       1,
			Username: "Admin",
		},
	}

	id2, err := forumRepository.CommentCreate(context.Background(), &domain.CommentCreate{
		ParentCommentID: expectedComment2.ParentCommentID,
		Content:         expectedComment2.Content,
		TopicID:         1,
	}, expectedComment2.CreatedAt, 1)
	require.NoError(t, err)
	require.Equal(t, expectedComment2.ID, id2)

	// get comments by topic
	comments, err := forumRepository.CommentsByTopic(context.Background(), 1, 30, 0)
	require.NoError(t, err)
	require.Equal(t, []*domain.Comment{expectedComment1, expectedComment2}, comments)

	// update comment 2
	expectedComment2.Content = "NO DOUBT"
	err = forumRepository.CommentUpdate(context.Background(), &domain.CommentUpdate{
		ID:      expectedComment2.ID,
		Content: expectedComment2.Content,
	}, time.Now().Truncate(time.Second).UTC(), 1)
	require.NoError(t, err)

	// get comments by topic
	commentsUpdated, err := forumRepository.CommentsByTopic(context.Background(), 1, 30, 0)
	require.NoError(t, err)
	require.Equal(t, []*domain.Comment{expectedComment1, expectedComment2}, commentsUpdated)
}

func TestForumRepositoryLikes(t *testing.T) {
	// like comment 1
	like1time := time.Now().Truncate(time.Second).UTC()
	like1, err := forumRepository.LikeNew(context.Background(), &domain.LikeCreate{
		CommentID: 1,
		LikeAuthor: &domain.LikeAuthor{
			ID:       1,
			Username: "Admin",
		},
	}, like1time)
	require.NoError(t, err)
	require.Equal(t, like1, dbs.LikesNewRow{CommentID: 1, CreatedBy: 1})

	// like comment 2
	like2time := time.Now().Truncate(time.Second).UTC()
	like2, err := forumRepository.LikeNew(context.Background(), &domain.LikeCreate{
		CommentID: 2,
		LikeAuthor: &domain.LikeAuthor{
			ID:       1,
			Username: "Admin",
		},
	}, like2time)
	require.NoError(t, err)
	require.Equal(t, dbs.LikesNewRow{CommentID: 2, CreatedBy: 1}, like2)

	// unlike comment 1
	like1Updated, err := forumRepository.LikeNew(context.Background(), &domain.LikeCreate{
		CommentID: 1,
		LikeAuthor: &domain.LikeAuthor{
			ID:       1,
			Username: "Admin",
		},
	}, time.Now().Truncate(time.Second).UTC())
	require.NoError(t, err)
	require.Equal(t, dbs.LikesNewRow{CommentID: 1, CreatedBy: 1}, like1Updated)

	// get active likes for comments
	comments, err := forumRepository.CommentsByTopic(context.Background(), 1, 30, 0)
	require.NoError(t, err)

	var likesForComment1Expected []*domain.Like
	require.Equal(t, likesForComment1Expected, comments[0].Likes)

	likesForComment2Expected := []*domain.Like{
		{
			CreatedAt: like2time,
			LikeAuthor: &domain.LikeAuthor{
				ID:       1,
				Username: "Admin",
			},
		},
	}
	require.Equal(t, likesForComment2Expected, comments[1].Likes)
}
