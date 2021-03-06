package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

// SignUpHandler
// @Summary 注册账号接口
// @Description 注册账号
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param object query models.ParamSignUp true "账号密码信息"
// @Success 200 {object} ResponseData
// @Router /api/v1/signup [post]
func SignUpHandler(c *gin.Context){
	//参数获取
	p := new( models.ParamSignUp )
	if err := c.ShouldBindJSON(p); err != nil{
		zap.L().Error("SignUp with invalid param",zap.Error(err) )
		//判断error是不是valid类型,如果是才需要翻译
		errs,ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c,CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParam,removeTopStruct( errs.Translate(trans) ))
		return
	}
	//业务处理
	if err := logic.SignUp(p); err!=nil{

		ResponseError(c,CodeServerBusy)
	}
	//返回响应
	ResponseSuccess(c,nil)
}

func LoginHandler(c *gin.Context) {
	//获取请求参数

	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil{
		zap.L().Error("Login with invalid param",zap.Error(err) )
		//判断error是不是valid类型,如果是才需要翻译
		errs,ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c,CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParam,removeTopStruct( errs.Translate(trans) ))
		return
	}

	token, err := logic.Login(p)
	if err!=nil {
		zap.L().Error("Logic.Login failed",zap.Error(err))
		ResponseErrorWithMsg(c,CodeInvalidPassword,"用户名或密码错误")
		return
	}
//	var temp map[string]interface{}
//	temp["token"] = token

	ResponseSuccess(c,gin.H{ "token":token } )
	return
}