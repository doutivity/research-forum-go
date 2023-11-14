package postgres

import (
	"context"
	"time"

	"github.com/doutivity/research-forum-go/internal/domain"
	"github.com/doutivity/research-forum-go/internal/storage/postgres/dbs"
)

type ForumRepository struct {
	db *Repository
}

func NewForumRepository(db *Repository) *ForumRepository {
	return &ForumRepository{db: db}
}

func (r *ForumRepository) TopicCreate(
	ctx context.Context,
	topic *domain.TopicCreate,
	createdAt time.Time,
	createdBy int64,
) (int64, error) {
	return r.db.Queries().TopicsNew(ctx, dbs.TopicsNewParams{
		Title:     topic.Title,
		Content:   topic.Content,
		CreatedAt: createdAt,
		CreatedBy: createdBy,
		UpdatedAt: createdAt,
		UpdatedBy: createdBy,
	})
}

func (r *ForumRepository) Topics(
	ctx context.Context,
	limit int64,
	offset int64,
) ([]*domain.Topic, error) {
	rows, err := r.db.Queries().Topics(ctx, dbs.TopicsParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	topics := make([]*domain.Topic, len(rows))
	for i, row := range rows {
		topics[i] = &domain.Topic{
			ID:        row.TopicID,
			Title:     row.Title,
			Content:   row.Content,
			CreatedAt: row.CreatedAt,
			Author: &domain.TopicAuthor{
				ID:       row.CreatedBy,
				Username: row.AuthorUsername,
			},
		}
	}
	return topics, nil
}
