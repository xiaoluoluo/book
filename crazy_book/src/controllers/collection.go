package controllers

import (
	"crazy_book/src/models"
	"crazy_book/src/service"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

// 用户收藏错题
func (this *MainController) AddCollection() {
	req := struct {
		UserId     uint64 `json:"user_id"`
		QuestionId uint64 `json:"question_id"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	collectionCtr := new(models.Collection)
	qCollections := collectionCtr.GetQuestionCollection(req.UserId, req.QuestionId)
	if len(qCollections) > 0 {
		this.Ctx.WriteString(BuildSuccessResponse("已经收藏过了"))
		return
	}
	// 用户id  问题id
	insertId, err := collectionCtr.AddCollection(req.UserId, req.QuestionId)
	if err != nil {
		logs.Error("AddCollection.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	respon := struct {
		CollectionId uint64 `json:"collection_id"`
	}{}
	respon.CollectionId = uint64(insertId)
	jsonRespon, _ := json.Marshal(respon)
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonRespon)))
}

// 获取收藏的问题列表
func (this *MainController) GetCollectionQuestionList() {
	userId, err := this.GetUint64("user_id")
	if err != nil {
		logs.Error("GetUserLabel err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	collection := new(models.Collection).GetCollection(userId)
	questionIds := make([]uint64, 0, len(collection))
	for _, c := range collection {
		questionIds = append(questionIds, c.QuestionId)
	}
	questions := new(models.Question).GetQuestionListByIds(questionIds)
	questionRespList := service.GetQuestionList(userId, questions)
	jsonQuestionRespList, err := json.Marshal(questionRespList)
	if err != nil {
		logs.Error("GetQuestionList.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonQuestionRespList)))
	return
}

//取消收藏
func (this *MainController) CancelCollection() {
	req := struct {
		UserId     uint64 `json:"user_id"`
		QuestionId uint64 `json:"question_id"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	err = new(models.Collection).CancelCollection(req.UserId, req.QuestionId)
	if err != nil {
		logs.Error("CancelCollection.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	respon := struct {
		Ok bool `json:"ok"`
	}{}
	respon.Ok = true
	jsonRespon, _ := json.Marshal(respon)
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonRespon)))
}

//用户给题目点赞
func (this *MainController) AddLiked() {
	req := struct {
		UserId     uint64 `json:"user_id"`
		QuestionId uint64 `json:"question_id"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	likedCtr := new(models.Liked)
	likeds := likedCtr.GetQuestionLiked(req.UserId, req.QuestionId)
	if len(likeds) > 0 {
		this.Ctx.WriteString(BuildSuccessResponse("已经点过赞了"))
		return
	}
	insertId, err := new(models.Liked).AddLiked(req.UserId, req.QuestionId)
	if err != nil {
		logs.Error("AddLiked.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	respon := struct {
		LikeId uint64 `json:"like_id"`
	}{}
	respon.LikeId = uint64(insertId)
	jsonRespon, _ := json.Marshal(respon)
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonRespon)))
}

//取消点赞
func (this *MainController) CancelLiked() {
	req := struct {
		UserId     uint64 `json:"user_id"`
		QuestionId uint64 `json:"question_id"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	err = new(models.Liked).CancelLiked(req.UserId, req.QuestionId)
	if err != nil {
		logs.Error("CancelLiked.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	respon := struct {
		Ok bool `json:"ok"`
	}{}
	respon.Ok = true
	jsonRespon, _ := json.Marshal(respon)
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonRespon)))
}

//增加知识点
func (this *MainController) AddLabel() {
	req := struct {
		UserId      uint64 `json:"user_id"`
		SubjectCode uint32 `json:"subject_code"`
		Label       string `json:"label"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	labelCtr := new(models.Label)
	userLabels := labelCtr.GetUserSubjectLabel(req.UserId, req.SubjectCode)
	if len(userLabels) >= 10 {
		this.Ctx.WriteString(BuildErrResponse("知识点不可以超过十个"))
		return
	}
	insertId, err := labelCtr.AddUserLabel(req.UserId, req.SubjectCode, req.Label)
	if err != nil {
		logs.Error("AddLabel.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	respon := struct {
		Label uint64 `json:"label_id"`
	}{}
	respon.Label = uint64(insertId)
	jsonRespon, _ := json.Marshal(respon)
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonRespon)))
}

//删除知识点
func (this *MainController) DeleteLabel() {
	req := struct {
		UserId  uint64 `json:"user_id"`
		LabelId uint64 `json:"label_id"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	err = new(models.Label).DeletedUserLabel(req.UserId, req.LabelId)
	if err != nil {
		logs.Error("DeletedUserLabel.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	respon := struct {
		Ok bool `json:"ok"`
	}{}
	respon.Ok = true
	jsonRespon, _ := json.Marshal(respon)
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonRespon)))
}

//获取用户的所有的知识点标签
func (this *MainController) GetUserLabel() {
	userId, err := this.GetUint64("user_id")
	if err != nil {
		logs.Error("GetUserLabel err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	userLabel := new(models.Label).GetUserLabel(userId)
	if len(userLabel) <= 0 {
		//TODO 如果没有那么就需要根据年级初始化一些知识标签
	}
	labelRespMap := make(map[uint32][]models.Label, 0)
	for _, ul := range userLabel {
		respList,has:= labelRespMap[ul.SubjectCode]
		if !has {
			respList = make([]models.Label,0,1)
		}
		respList =append(respList,ul)
		labelRespMap[ul.SubjectCode] = respList
	}
	jsonUsers, err := json.Marshal(labelRespMap)
	if err != nil {
		logs.Error("GetUserLabel.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonUsers)))
}
