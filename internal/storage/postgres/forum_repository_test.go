package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
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
	// add topic 1
	var (
		expectedTopic1 = &domain.Topic{
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

	fmt.Println(time.Now().Truncate(time.Second).UTC())

	id, err := forumRepository.TopicCreate(context.Background(), &domain.TopicCreate{
		Title:   expectedTopic1.Title,
		Content: expectedTopic1.Content,
	}, expectedTopic1.CreatedAt, 1)
	require.NoError(t, err)
	require.Equal(t, expectedTopic1.ID, id)

	topics, err := forumRepository.Topics(context.Background(), 30, 0)
	require.NoError(t, err)
	require.Equal(t, []*domain.Topic{expectedTopic1}, topics)

	TopicByID, err := forumRepository.TopicByID(context.Background(), 1)
	require.NoError(t, err)
	require.Equal(t, expectedTopic1, TopicByID)

	// add topic 2
	var (
		expectedTopic2 = &domain.Topic{
			ID:        2,
			Title:     "DOU Forum on PostgreSQL",
			Content:   "https://github.com/doutivity/research-forum-go",
			CreatedAt: time.Now().Truncate(time.Second).UTC().Add(time.Second),
			Author: &domain.TopicAuthor{
				ID:       1,
				Username: "Admin",
			},
		}
	)

	id, err = forumRepository.TopicCreate(context.Background(), &domain.TopicCreate{
		Title:   expectedTopic2.Title,
		Content: expectedTopic2.Content,
	}, expectedTopic2.CreatedAt, 1)
	require.NoError(t, err)
	require.Equal(t, expectedTopic2.ID, id)

	TopicByID, err = forumRepository.TopicByID(context.Background(), 2)
	require.NoError(t, err)
	require.Equal(t, expectedTopic2, TopicByID)

	// get topics
	topics, err = forumRepository.Topics(context.Background(), 30, 0)
	require.NoError(t, err)
	require.Equal(t, []*domain.Topic{expectedTopic2, expectedTopic1}, topics)
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

func TestForumRepositoryTopicsOrder(t *testing.T) {
	// add comment to topic 1
	expectedComment1 := &domain.Comment{
		ID:              3,
		ParentCommentID: nil,
		Content:         "Great topic",
		CreatedAt:       time.Now().Truncate(time.Second).UTC().Add(time.Second * 2),
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

	// get topics
	topics, err := forumRepository.Topics(context.Background(), 30, 0)
	require.NoError(t, err)
	require.Equal(t, int64(1), topics[0].ID)
	require.Equal(t, int64(2), topics[1].ID)
}

func TestForumRepositoryLikes(t *testing.T) {
	// like comment 1
	like1time := time.Now().Truncate(time.Second).UTC()
	err := forumRepository.LikeUpdate(context.Background(), &domain.LikeUpdate{
		CommentID: 1,
		LikeAuthor: &domain.LikeAuthor{
			ID:       1,
			Username: "Admin",
		},
	}, true, like1time)
	require.NoError(t, err)

	// like comment 2
	like2time := time.Now().Truncate(time.Second).UTC()
	err = forumRepository.LikeUpdate(context.Background(), &domain.LikeUpdate{
		CommentID: 2,
		LikeAuthor: &domain.LikeAuthor{
			ID:       1,
			Username: "Admin",
		},
	}, true, like2time)
	require.NoError(t, err)

	// unlike comment 1
	err = forumRepository.LikeUpdate(context.Background(), &domain.LikeUpdate{
		CommentID: 1,
		LikeAuthor: &domain.LikeAuthor{
			ID:       1,
			Username: "Admin",
		},
	}, false, time.Now().Truncate(time.Second).UTC())
	require.NoError(t, err)

	// get active likes for comments
	likesForComments, err := forumRepository.LikesByCommentIDs(context.Background(), []int64{1, 2})
	require.NoError(t, err)
	require.Equal(t, []*domain.Like{
		{
			CommentID: 2,
			CreatedAt: like2time,
			LikeAuthor: &domain.LikeAuthor{
				ID:       1,
				Username: "Admin",
			},
		}}, likesForComments)
}

func TestForumRepositoryTopicsWithUnreadCommentsNumber(t *testing.T) {
	// get topics
	topics, err := forumRepository.TopicsWithUnreadCommentsNumber(context.Background(), 1, 30, 0)
	require.NoError(t, err)
	require.Equal(t, int64(3), topics[0].UnreadCommentsNumber)
	require.Equal(t, int64(0), topics[1].UnreadCommentsNumber)

	// read comment 1 and comment 2 from topic 1
	err = forumRepository.LastReadCommentsCreate(context.Background(), domain.ReadComment{
		TopicId:   1,
		CommentId: 2,
	}, 1)
	require.NoError(t, err)

	// get topics
	topics, err = forumRepository.TopicsWithUnreadCommentsNumber(context.Background(), 1, 30, 0)
	require.NoError(t, err)
	require.Equal(t, int64(1), topics[0].UnreadCommentsNumber)
	require.Equal(t, int64(0), topics[1].UnreadCommentsNumber)

	// last read comment after reading comment 1 and comment 2 from topic 1
	topic, err := forumRepository.TopicByIDWithLastReadComment(context.Background(), 1, 1)
	require.NoError(t, err)
	require.Equal(t, int64(2), topic.Comment.ID)

	// last read comment = null from topic 1
	topic, err = forumRepository.TopicByIDWithLastReadComment(context.Background(), 2, 1)
	require.NoError(t, err)
	require.Nil(t, topic.Comment)
}
