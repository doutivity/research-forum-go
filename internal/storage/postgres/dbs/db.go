// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package dbs

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.commentUpdateStmt, err = db.PrepareContext(ctx, commentUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query CommentUpdate: %w", err)
	}
	if q.commentsByIDStmt, err = db.PrepareContext(ctx, commentsByID); err != nil {
		return nil, fmt.Errorf("error preparing query CommentsByID: %w", err)
	}
	if q.commentsByTopicStmt, err = db.PrepareContext(ctx, commentsByTopic); err != nil {
		return nil, fmt.Errorf("error preparing query CommentsByTopic: %w", err)
	}
	if q.commentsNewStmt, err = db.PrepareContext(ctx, commentsNew); err != nil {
		return nil, fmt.Errorf("error preparing query CommentsNew: %w", err)
	}
	if q.lastReadCommentsNewStmt, err = db.PrepareContext(ctx, lastReadCommentsNew); err != nil {
		return nil, fmt.Errorf("error preparing query LastReadCommentsNew: %w", err)
	}
	if q.likesByCommentIDsStmt, err = db.PrepareContext(ctx, likesByCommentIDs); err != nil {
		return nil, fmt.Errorf("error preparing query LikesByCommentIDs: %w", err)
	}
	if q.likesUpsertStmt, err = db.PrepareContext(ctx, likesUpsert); err != nil {
		return nil, fmt.Errorf("error preparing query LikesUpsert: %w", err)
	}
	if q.topicLastUpdateNewStmt, err = db.PrepareContext(ctx, topicLastUpdateNew); err != nil {
		return nil, fmt.Errorf("error preparing query TopicLastUpdateNew: %w", err)
	}
	if q.topicsStmt, err = db.PrepareContext(ctx, topics); err != nil {
		return nil, fmt.Errorf("error preparing query Topics: %w", err)
	}
	if q.topicsCommentStatsCountIncStmt, err = db.PrepareContext(ctx, topicsCommentStatsCountInc); err != nil {
		return nil, fmt.Errorf("error preparing query TopicsCommentStatsCountInc: %w", err)
	}
	if q.topicsCommentStatsNewStmt, err = db.PrepareContext(ctx, topicsCommentStatsNew); err != nil {
		return nil, fmt.Errorf("error preparing query TopicsCommentStatsNew: %w", err)
	}
	if q.topicsGetByIDStmt, err = db.PrepareContext(ctx, topicsGetByID); err != nil {
		return nil, fmt.Errorf("error preparing query TopicsGetByID: %w", err)
	}
	if q.topicsGetByIDWithLastReadCommentStmt, err = db.PrepareContext(ctx, topicsGetByIDWithLastReadComment); err != nil {
		return nil, fmt.Errorf("error preparing query TopicsGetByIDWithLastReadComment: %w", err)
	}
	if q.topicsNewStmt, err = db.PrepareContext(ctx, topicsNew); err != nil {
		return nil, fmt.Errorf("error preparing query TopicsNew: %w", err)
	}
	if q.topicsWithUnreadCommentsNumberStmt, err = db.PrepareContext(ctx, topicsWithUnreadCommentsNumber); err != nil {
		return nil, fmt.Errorf("error preparing query TopicsWithUnreadCommentsNumber: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.commentUpdateStmt != nil {
		if cerr := q.commentUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing commentUpdateStmt: %w", cerr)
		}
	}
	if q.commentsByIDStmt != nil {
		if cerr := q.commentsByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing commentsByIDStmt: %w", cerr)
		}
	}
	if q.commentsByTopicStmt != nil {
		if cerr := q.commentsByTopicStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing commentsByTopicStmt: %w", cerr)
		}
	}
	if q.commentsNewStmt != nil {
		if cerr := q.commentsNewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing commentsNewStmt: %w", cerr)
		}
	}
	if q.lastReadCommentsNewStmt != nil {
		if cerr := q.lastReadCommentsNewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing lastReadCommentsNewStmt: %w", cerr)
		}
	}
	if q.likesByCommentIDsStmt != nil {
		if cerr := q.likesByCommentIDsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing likesByCommentIDsStmt: %w", cerr)
		}
	}
	if q.likesUpsertStmt != nil {
		if cerr := q.likesUpsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing likesUpsertStmt: %w", cerr)
		}
	}
	if q.topicLastUpdateNewStmt != nil {
		if cerr := q.topicLastUpdateNewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing topicLastUpdateNewStmt: %w", cerr)
		}
	}
	if q.topicsStmt != nil {
		if cerr := q.topicsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing topicsStmt: %w", cerr)
		}
	}
	if q.topicsCommentStatsCountIncStmt != nil {
		if cerr := q.topicsCommentStatsCountIncStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing topicsCommentStatsCountIncStmt: %w", cerr)
		}
	}
	if q.topicsCommentStatsNewStmt != nil {
		if cerr := q.topicsCommentStatsNewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing topicsCommentStatsNewStmt: %w", cerr)
		}
	}
	if q.topicsGetByIDStmt != nil {
		if cerr := q.topicsGetByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing topicsGetByIDStmt: %w", cerr)
		}
	}
	if q.topicsGetByIDWithLastReadCommentStmt != nil {
		if cerr := q.topicsGetByIDWithLastReadCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing topicsGetByIDWithLastReadCommentStmt: %w", cerr)
		}
	}
	if q.topicsNewStmt != nil {
		if cerr := q.topicsNewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing topicsNewStmt: %w", cerr)
		}
	}
	if q.topicsWithUnreadCommentsNumberStmt != nil {
		if cerr := q.topicsWithUnreadCommentsNumberStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing topicsWithUnreadCommentsNumberStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                   DBTX
	tx                                   *sql.Tx
	commentUpdateStmt                    *sql.Stmt
	commentsByIDStmt                     *sql.Stmt
	commentsByTopicStmt                  *sql.Stmt
	commentsNewStmt                      *sql.Stmt
	lastReadCommentsNewStmt              *sql.Stmt
	likesByCommentIDsStmt                *sql.Stmt
	likesUpsertStmt                      *sql.Stmt
	topicLastUpdateNewStmt               *sql.Stmt
	topicsStmt                           *sql.Stmt
	topicsCommentStatsCountIncStmt       *sql.Stmt
	topicsCommentStatsNewStmt            *sql.Stmt
	topicsGetByIDStmt                    *sql.Stmt
	topicsGetByIDWithLastReadCommentStmt *sql.Stmt
	topicsNewStmt                        *sql.Stmt
	topicsWithUnreadCommentsNumberStmt   *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                   tx,
		tx:                                   tx,
		commentUpdateStmt:                    q.commentUpdateStmt,
		commentsByIDStmt:                     q.commentsByIDStmt,
		commentsByTopicStmt:                  q.commentsByTopicStmt,
		commentsNewStmt:                      q.commentsNewStmt,
		lastReadCommentsNewStmt:              q.lastReadCommentsNewStmt,
		likesByCommentIDsStmt:                q.likesByCommentIDsStmt,
		likesUpsertStmt:                      q.likesUpsertStmt,
		topicLastUpdateNewStmt:               q.topicLastUpdateNewStmt,
		topicsStmt:                           q.topicsStmt,
		topicsCommentStatsCountIncStmt:       q.topicsCommentStatsCountIncStmt,
		topicsCommentStatsNewStmt:            q.topicsCommentStatsNewStmt,
		topicsGetByIDStmt:                    q.topicsGetByIDStmt,
		topicsGetByIDWithLastReadCommentStmt: q.topicsGetByIDWithLastReadCommentStmt,
		topicsNewStmt:                        q.topicsNewStmt,
		topicsWithUnreadCommentsNumberStmt:   q.topicsWithUnreadCommentsNumberStmt,
	}
}
