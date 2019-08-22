package service

import "crazy_book/src/models"

type QuestionResp struct {
	Question models.Question
	User     models.User
	Comment  []CommentResp
	LikedNum uint32
}

type CommentResp struct {
	Comment models.Comment
	User    models.User
}

type LabelResp struct {
	UserId      uint64
	SubjectCode uint32
	Labels      []models.Label
}
