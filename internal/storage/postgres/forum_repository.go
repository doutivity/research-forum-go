package postgres

import (
	"context"
	"database/sql"
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

func (r *ForumRepository) LikesByCommentIDs(
	ctx context.Context,
	commentIDs []int64,
) ([]*domain.Like, error) {
	rows, err := r.db.Queries().LikesByCommentIDs(ctx, commentIDs)
	if err != nil {
		return nil, err
	}
	likes := make([]*domain.Like, len(rows))
	for i, row := range rows {
		likes[i] = &domain.Like{
			CommentID: row.CommentID,
			CreatedAt: row.CreatedAt,
			LikeAuthor: &domain.LikeAuthor{
				ID:       row.UserID,
				Username: row.Username,
			},
		}
	}
	return likes, nil
}

func (r *ForumRepository) LikeUpdate(
	ctx context.Context,
	likeCreate *domain.LikeUpdate,
	active bool,
	createdAt time.Time,
) error {
	return r.db.Queries().LikesUpsert(ctx, dbs.LikesUpsertParams{
		CommentID: likeCreate.CommentID,
		Active:    active,
		CreatedAt: createdAt,
		CreatedBy: likeCreate.LikeAuthor.ID,
		UpdatedAt: createdAt,
		UpdatedBy: likeCreate.LikeAuthor.ID,
	})
}

func (r *ForumRepository) CommentUpdate(
	ctx context.Context,
	comment *domain.CommentUpdate,
	updatedAt time.Time,
	updatedBy int64,
) error {
	return r.db.Queries().CommentUpdate(ctx, dbs.CommentUpdateParams{
		Content:   comment.Content,
		UpdatedAt: updatedAt,
		UpdatedBy: updatedBy,
		CommentID: comment.ID,
	})
}

func (r *ForumRepository) CommentsByTopic(
	ctx context.Context,
	topicId int64,
	limit int64,
	offset int64,
) ([]*domain.Comment, error) {
	rows, err := r.db.Queries().CommentsByTopic(ctx, dbs.CommentsByTopicParams{
		TopicID: topicId,
		Offset:  offset,
		Limit:   limit,
	})
	if err != nil {
		return nil, err
	}

	comments := make([]*domain.Comment, len(rows))
	for i, row := range rows {
		var parentCommentID *int64
		if row.ParentCommentID.Valid {
			parentCommentID = &row.ParentCommentID.Int64
		} else {
			parentCommentID = nil
		}
		comments[i] = &domain.Comment{
			ID:              row.CommentID,
			ParentCommentID: parentCommentID,
			Content:         row.Content,
			CreatedAt:       row.CreatedAt,
			Author: &domain.CommentAuthor{
				ID:       row.CreatedBy,
				Username: row.AuthorUsername,
			},
		}
	}
	return comments, nil
}

func (r *ForumRepository) CommentCreate(
	ctx context.Context,
	comment *domain.CommentCreate,
	createdAt time.Time,
	createdBy int64,
) (int64, error) {
	var parentCommentID sql.NullInt64
	if comment.ParentCommentID != nil {
		parentCommentID = sql.NullInt64{Int64: *comment.ParentCommentID, Valid: true}
	} else {
		parentCommentID = sql.NullInt64{Valid: false}
	}
	id, err := r.db.Queries().CommentsNew(ctx, dbs.CommentsNewParams{
		ParentCommentID: parentCommentID,
		Content:         comment.Content,
		TopicID:         comment.TopicID,
		CreatedAt:       createdAt,
		CreatedBy:       createdBy,
		UpdatedAt:       createdAt,
		UpdatedBy:       createdBy,
	})

	err = r.db.Queries().TopicLastUpdateNew(ctx, dbs.TopicLastUpdateNewParams{
		TopicID:       comment.TopicID,
		LastUpdatedAt: createdAt,
	})
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r *ForumRepository) TopicByID(
	ctx context.Context,
	topicID int64,
) (*domain.Topic, error) {
	row, err := r.db.Queries().TopicsGetByID(ctx, int64(topicID))
	if err != nil {
		return nil, err
	}
	return &domain.Topic{
		ID:        row.TopicID,
		Title:     row.Title,
		Content:   row.Content,
		CreatedAt: row.CreatedAt,
		Author: &domain.TopicAuthor{
			ID:       row.CreatedBy,
			Username: row.AuthorUsername,
		},
	}, nil
}

func (r *ForumRepository) TopicCreate(
	ctx context.Context,
	topic *domain.TopicCreate,
	createdAt time.Time,
	createdBy int64,
) (int64, error) {
	id, err := r.db.Queries().TopicsNew(ctx, dbs.TopicsNewParams{
		Title:     topic.Title,
		Content:   topic.Content,
		CreatedAt: createdAt,
		CreatedBy: createdBy,
		UpdatedAt: createdAt,
		UpdatedBy: createdBy,
	})
	if err != nil {
		return id, err
	}

	err = r.db.Queries().TopicLastUpdateNew(ctx, dbs.TopicLastUpdateNewParams{
		TopicID:       id,
		LastUpdatedAt: createdAt,
	})
	if err != nil {
		return id, err
	}
	return id, nil
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
