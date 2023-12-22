package domain

import "time"

type TopicCreate struct {
	Title   string
	Content string
}

type TopicAuthor struct {
	ID       int64
	Username string
}

type Topic struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
	Author    *TopicAuthor
}

type TopicsWithUnreadCommentsNumber struct {
	Topic                *Topic
	UnreadCommentsNumber int64
}

type TopicWithUnreadComment struct {
	Topic   *Topic
	Comment *Comment
}
