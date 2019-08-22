package service

import "crazy_book/src/models"

/**获取评论**/
func GetComment(questionId uint64) []CommentResp {
	commentList := new(models.Comment).GetComment(questionId)
	userIdList := make([]uint64, 0, len(commentList))
	for _, c := range commentList {
		userIdList = append(userIdList, c.UserId)
	}
	commentUserList := new(models.User).GetUserList(userIdList)
	CommentRespList := make([]CommentResp, 0, len(userIdList))
	for _, c := range commentList {
		for _, u := range commentUserList {
			if u.UserId != c.UserId {
				continue
			}
			resp := CommentResp{
				Comment: c,
				User:    u,
			}
			CommentRespList = append(CommentRespList, resp)
			break
		}
	}
	return CommentRespList
}

/**获取问题的点赞数量**/
func GetQuestionLikeNum(questionId uint64) uint32 {
	likes := new(models.Liked).GetLiked(questionId)
	return uint32(len(likes))
}

func GetQuestionList(questions []models.Question) []QuestionResp {
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
			commentList := GetComment(q.QuestionId)
			likedNum := GetQuestionLikeNum(q.QuestionId)
			resp := QuestionResp{
				Question: q,
				User:     u,
				Comment:  commentList,
				LikedNum: likedNum,
			}
			QuestionRespList = append(QuestionRespList, resp)
			break
		}
	}
	return QuestionRespList
}
