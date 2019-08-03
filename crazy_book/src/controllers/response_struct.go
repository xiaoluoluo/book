package controllers

import "crazy_book/src/models"

type QuestionResp struct {
	Question models.Question
	User     models.User
}

type CommentResp struct {
	Comment models.Comment
	User    models.User
}