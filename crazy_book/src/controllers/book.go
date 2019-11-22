package controllers

import (
	"crazy_book/config"
	"crazy_book/src/models"
	"crazy_book/src/service"
	"encoding/json"
	"fmt"
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
		Code string `json:"code"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", config.AppId, config.Secret, req.Code)
	resp, err := http.Get(url)
	if err != nil {
		logs.Error("GetWxOpenId is err:", err.Error(), config.AppId, config.Secret, req.Code)
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
	userId := uint64(0)
	if len(users) > 0 {
		userId = users[0].UserId
	} else {
		insertId, err := new(models.User).Register(req.UserWid, req.UserName, req.UserHeadPic)
		if err != nil {
			logs.Error("Register err:", err.Error())
			this.Ctx.WriteString(BuildErrResponse("数据库出现错误"))
			return
		}
		userId = uint64(insertId)
	}
	respon := struct {
		UserId uint64 `json:"user_id"`
	}{}
	respon.UserId = uint64(userId)
	this.Ctx.WriteString(BuildSuccessResponse(respon))
	return
}

// 登录
func (this *MainController) Login() {
	userWid := this.GetString("user_wid")
	users := new(models.User).Login(userWid)
	if len(users) <= 0 {
		logs.Error("Login no user wid:%s", userWid)
		this.Ctx.WriteString(BuildErrResponse("数据库没有数据"))
		return
	}
	this.Ctx.WriteString(BuildSuccessResponse(users[0]))
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
		QuestionPic   string `json:"question_pic"`
		SubjectCode   uint32 `json:"subject_code"`
		TruePic1      string `json:"true_pic1"`
		TruePic2      string `json:"true_pic2"`
		Point         string `json:"point"`
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
	insertId, err := new(models.Question).AddMyQuestion(req.UserId, userGrade, req.QuestionTitle, req.QuestionPic, req.SubjectCode, req.TruePic1, req.TruePic2, req.Point)
	if err != nil {
		logs.Error("AddMyQuestion err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("数据库出现错误"))
		return
	}
	respon := struct {
		QuestionId uint64 `json:"question_id"`
	}{}
	respon.QuestionId = uint64(insertId)
	this.Ctx.WriteString(BuildSuccessResponse(respon))
}

// 更新题目信息
func (this *MainController) UpdateQuestion() {
	req := struct {
		QuestionId    uint64 `json:"question_id"`
		UserId        uint64 `json:"user_id"`
		QuestionTitle string `json:"question_title"`
		QuestionPic   string `json:"question_pic"`
		SubjectCode   uint32 `json:"subject_code"`
		TruePic1      string `json:"true_pic1"`
		TruePic2      string `json:"true_pic2"`
		Point         string `json:"point"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if err != nil {
		logs.Error("json.Unmarshal is err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("请求参数错误"))
		return
	}
	err = new(models.Question).UpdateQuestion(req.QuestionId, req.UserId, req.QuestionTitle, req.QuestionPic, req.SubjectCode, req.TruePic1, req.TruePic2, req.Point)
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
	if questions == nil {
		questions = make([]models.Question, 0, 1)
	}
	this.Ctx.WriteString(BuildSuccessResponse(questions))
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
	this.Ctx.WriteString(BuildSuccessResponse(questions))
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
	questionRespList := service.GetQuestionList(userId, questions)
	this.Ctx.WriteString(BuildSuccessResponse(questionRespList))
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
	this.Ctx.WriteString(BuildSuccessResponse(respon))
}

// 获取问题的评论
func (this *MainController) GetQuestionComment() {
	// 问题id
	questionId, err := this.GetUint64("question_id")
	if err != nil {
		logs.Error("GetQuestionComment err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("参数缺少question_id"))
		return
	}
	userId, err := this.GetUint64("user_id")
	if err != nil {
		logs.Error("GetQuestionComment err:", err.Error())
		this.Ctx.WriteString(BuildErrResponse("参数缺少user_id"))
		return
	}
	_, commentRespList := service.GetComment(userId, questionId)
	this.Ctx.WriteString(BuildSuccessResponse(commentRespList))
}
