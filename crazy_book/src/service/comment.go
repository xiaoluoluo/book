package service

import "crazy_book/src/models"

/**获取评论**/
func GetComment(userId uint64, questionId uint64) (bool, []CommentResp) {
	commentList := new(models.Comment).GetComment(questionId)
	userIdList := make([]uint64, 0, len(commentList))
	for _, c := range commentList {
		userIdList = append(userIdList, c.UserId)
	}
	commentUserList := new(models.User).GetUserList(userIdList)
	CommentRespList := make([]CommentResp, 0, len(userIdList))
	hasMe := false
	for _, c := range commentList {
		for _, u := range commentUserList {
			if u.UserId != c.UserId {
				continue
			}
			if u.UserId == userId {
				hasMe = true
			}
			resp := CommentResp{
				Comment: c,
				User:    u,
			}
			CommentRespList = append(CommentRespList, resp)
			break
		}
	}
	return hasMe, CommentRespList
}

/**获取问题的点赞数量**/
func GetQuestionLikeNum(userId uint64, questionId uint64) (bool, uint32) {
	likes := new(models.Liked).GetLiked(questionId)
	hasMe := false
	for _, like := range likes {
		if like.UserId == userId {
			hasMe = true
		}
	}
	return hasMe, uint32(len(likes))
}

func GetQuestionList(userId uint64, questions []models.Question) []QuestionResp {
	userIds := make([]uint64, 0, len(questions))
	for _, q := range questions {
		userIds = append(userIds, q.UserId)
	}
	userList := new(models.User).GetUserList(userIds)
	QuestionRespList := make([]QuestionResp, 0, len(userIds))
	for _, q := range questions {
		for _, u := range userList {
			if u.UserId != q.UserId {
				continue
			}
			meComment, commentList := GetComment(userId, q.QuestionId)
			meLike, likedNum := GetQuestionLikeNum(userId, q.QuestionId)
			resp := QuestionResp{
				Question:  q,
				User:      u,
				Comment:   commentList,
				LikedNum:  likedNum,
				MeComment: meComment,
				MeLiked:   meLike,
			}
			QuestionRespList = append(QuestionRespList, resp)
			break
		}
	}
	return QuestionRespList
}
