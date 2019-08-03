package controllers

import (
	"crazy_book/src/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
)

type MainController struct {
	beego.Controller
}

//获取OpendId
func (this *MainController) GetWxOpenId() {
	req := struct {
		AppId  string `json:"app_id"`
		Secret string `json:"secret"`
		Code   string `json:"code"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + req.AppId + "&secret=" + req.Secret +
		"&js_code=" + req.Code + "&grant_type=authorization_code"
	resp, err := http.Get(url)
	if err != nil {
		logs.Error("GetWxOpenId is err:", err.Error(), req.AppId, req.Secret, req.Code)
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	this.Ctx.WriteString(BuildSuccessResponse(string(body)))
}

// 注册
func (this *MainController) Register() {
	req := struct {
		UserWid     string `json:"user_wid"`
		UserName    string `json:"user_name"`
		UserHeadPic string `json:"user_head_pic"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	users := new(models.User).Login(req.UserWid)
	if len(users) > 0 {
		logs.Error("Register.Login:", req.UserWid, req.UserName, req.UserHeadPic)
		this.Ctx.WriteString(BuildErrResponse("已经注册过了"))
		return
	}
	insertId, err := new(models.User).Register(req.UserWid, req.UserName, req.UserHeadPic)
	if err != nil {
		logs.Error("Register err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库出现错误"))
		return
	}
	respon := struct {
		UserId uint64 `json:"user_id"`
	}{}
	respon.UserId = uint64(insertId)
	jsonRespon, _ := json.Marshal(respon)
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonRespon)))
	return
}

// 登录
func (this *MainController) Login() {
	userWid := this.GetString("wid")
	users := new(models.User).Login(userWid)
	if len(users) <= 0 {
		logs.Error("Login no user:")
		this.Ctx.WriteString(BuildErrResponse("数据库没有数据"))
		return
	}
	jsonUsers, err := json.Marshal(users[0])
	if err != nil {
		logs.Error("Login.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonUsers)))
}

// 修改用户的年级
func (this *MainController) UpdateGrade() {
	req := struct {
		UserId    uint64 `json:"user_id"`
		UserGrade uint32 `json:"user_grade"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	err = new(models.User).UpdateUserGrade(req.UserId, req.UserGrade)
	if err != nil {
		logs.Error("UpdateUserGrade err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库出现错误"))
		return
	}
	this.Ctx.WriteString(BuildSuccessResponse("ok"))
}

// 增加我的错题
func (this *MainController) AddMyQuestion() {
	req := struct {
		UserId        uint64 `json:"user_id"`
		QuestionTitle string `json:"question_title"`
		AnswerPic     string `json:"answer_pic"`
		SubjectCode   uint32 `json:"subject_code"`
		TrueTitle     string `json:"true_title"`
		TruePic       string `json:"true_pic"`
		FalseTitle    string `json:"false_title"`
		FalsePic      string `json:"false_pic"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	users := new(models.User).GetUserById(req.UserId)
	if len(users) <= 0 {
		logs.Error("GetUserById userId is  err:", req.UserId)
		this.Ctx.WriteString(BuildErrResponse("用户id有误"))
		return
	}
	userGrade := users[0].UserGrade
	insertId, err := new(models.Question).AddMyQuestion(req.UserId, userGrade, req.QuestionTitle, req.AnswerPic, req.SubjectCode, req.TrueTitle, req.TruePic, req.FalseTitle, req.FalsePic)
	if err != nil {
		logs.Error("AddMyQuestion err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库出现错误"))
		return
	}
	respon := struct {
		QuestionId uint64 `json:"question_id"`
	}{}
	respon.QuestionId = uint64(insertId)
	jsonRespon, _ := json.Marshal(respon)
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonRespon)))
}

// 更新题目信息
func (this *MainController) UpdateQuestion() {
	req := struct {
		QuestionId    uint64 `json:"question_id"`
		UserId        uint64 `json:"user_id"`
		QuestionTitle string `json:"question_title"`
		AnswerPic     string `json:"answer_pic"`
		SubjectCode   uint32 `json:"subject_code"`
		TrueTitle     string `json:"true_title"`
		TruePic       string `json:"true_pic"`
		FalseTitle    string `json:"false_title"`
		FalsePic      string `json:"false_pic"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	err = new(models.Question).UpdateQuestion(req.QuestionId, req.UserId, req.QuestionTitle, req.AnswerPic, req.SubjectCode, req.TrueTitle, req.TruePic, req.FalseTitle, req.FalsePic)
	if err != nil {
		logs.Error("UpdateQuestion err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	this.Ctx.WriteString(BuildSuccessResponse("ok"))
}

// 获取我的所有错题
func (this *MainController) GetMyAllQuestion() {
	userId, err := this.GetUint64("user_id")
	if err != nil {
		logs.Error("GetMyAllQuestion err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	page, err := this.GetInt("page")
	if err != nil {
		logs.Error("GetMyAllQuestion err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	subjectCode, err := this.GetUint32("subject_code", 0)
	if err != nil {
		logs.Error("GetMyAllQuestion err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	users := new(models.User).GetUserById(userId)
	if len(users) <= 0 {
		logs.Error("GetUserById user is empty userId:", userId)
		this.Ctx.WriteString(BuildErrResponse("用户不存在，请联系管理员"))
		return
	}
	grade := users[0].UserGrade
	var questions []models.Question
	if subjectCode <= 0 {
		questions = new(models.Question).GetMyAllQuestion(userId, grade, 10, page)
	} else {
		questions = new(models.Question).GetMyQuestionBySubject(userId, grade, subjectCode, 10, page)
	}
	jsonQuestions, err := json.Marshal(questions)
	if err != nil {
		logs.Error("GetMyAllQuestion.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonQuestions)))
}

// 根据题目id 获取题目信息
func (this *MainController) GetQuestionById() {
	questionId, err := this.GetUint64("question_id")
	if err != nil {
		logs.Error("GetQuestionById err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	questions := new(models.Question).GetQuestionById(questionId)
	jsonQuestions, err := json.Marshal(questions)
	if err != nil {
		logs.Error("GetQuestionById.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonQuestions)))
}

//广场中的所有错题
func (this *MainController) GetQuestionList() {
	page, err := this.GetInt("page", 0)
	if err != nil {
		logs.Error("GetQuestionList err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	userId, err := this.GetUint64("user_id")
	if err != nil {
		logs.Error("GetQuestionList err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	subjectCode, err := this.GetUint32("subject_code", 0)
	if err != nil {
		logs.Error("GetQuestionList err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	var questions []models.Question
	if subjectCode <= 0 {
		questions = new(models.Question).GetQuestionList(10, page)
	} else {
		users := new(models.User).GetUserById(userId)
		if len(users) <= 0 {
			logs.Error("GetUserById user is empty userId:", userId)
			this.Ctx.WriteString(BuildErrResponse("用户不存在，请联系管理员"))
			return
		}
		grade := users[0].UserGrade
		questions = new(models.Question).GetQuestionByGradeAndSubject(grade, subjectCode, 10, page)
	}

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
			resp := QuestionResp{
				Question: q,
				User:     u,
			}
			QuestionRespList = append(QuestionRespList, resp)
			break
		}
	}
	jsonQuestionRespList, err := json.Marshal(QuestionRespList)
	if err != nil {
		logs.Error("GetQuestionList.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonQuestionRespList)))
}

// 删除我的题目
func (this *MainController) DeletedMyQuestion() {
	req := struct {
		QuestionId uint64 `json:"question_id"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	// 删除我的题目
	err = new(models.Question).DeletedMyQuestion(req.QuestionId)
	if err != nil {
		logs.Error("DeletedMyQuestion.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	this.Ctx.WriteString(BuildSuccessResponse("ok"))
}

// 增加问题的评论
func (this *MainController) AddQuestionComment() {
	req := struct {
		UserId       uint64 `json:"user_id"`
		QuestionId   uint64 `json:"question_id"`
		CommentIntro string `json:"comment_intro"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	// 用户id  问题id  评论
	insertId, err := new(models.Comment).AddComment(req.UserId, req.QuestionId, req.CommentIntro)
	if err != nil {
		logs.Error("AddComment.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	respon := struct {
		CommentId uint64 `json:"comment_id"`
	}{}
	respon.CommentId = uint64(insertId)
	jsonRespon, _ := json.Marshal(respon)
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonRespon)))
}

// 获取问题的评论
func (this *MainController) GetQuestionComment() {
	// 问题id
	questionId, err := this.GetUint64("question_id")
	if err != nil {
		logs.Error("GetQuestionComment err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
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
	jsonCommentRespList, err := json.Marshal(CommentRespList)
	if err != nil {
		logs.Error("CommentRespList.Marshal err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库报错"))
		return
	}
	this.Ctx.WriteString(BuildSuccessResponse(string(jsonCommentRespList)))
}
