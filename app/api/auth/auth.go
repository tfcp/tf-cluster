package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gvalid"
	"golang.org/x/crypto/bcrypt"
	"tf-cluster/internal/enum"
	"tf-cluster/internal/service/auth"
	"tf-cluster/library/code"
	"tf-cluster/library/utils"
)

type RequestLogin struct {
	Name string `json:"username" form:"username" valid:"username      @required#username不能为空"`
	Pwd  string `json:"password" form:"password" valid:"password      @required#password不能为空"`
}

func LoginApi(c *gin.Context) {
	var reqLogin RequestLogin
	c.Bind(&reqLogin)
	if err := gvalid.CheckStruct(c, reqLogin, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	us := auth.UserService{}
	whereMap := map[string]interface{}{
		"username": reqLogin.Name,
	}
	um, err := us.One(whereMap)
	if err != nil {
		utils.Response(c, code.ErrAccount, err.Error())
		return
	}
	// 用户禁用
	if um.Status == enum.UserUnable {
		utils.Response(c, code.ErrUnable, err.Error())
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(um.Pwd), []byte(reqLogin.Pwd)) //加密处理
	if err != nil {
		utils.Response(c, code.ErrPwd, err.Error())
		return
	}
	token, err := utils.GenerateToken(um.Name, um.Avatar, um.Introduction, um.Role)
	if err != nil {
		utils.Response(c, code.ErrAccount, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, map[string]interface{}{
		"token": token,
	})
}

type RequestInfo struct {
	Token string `json:"token" form:"token" valid:"token      @required#token不能为空"`
}

func InfoApi(c *gin.Context) {
	var reqInfo RequestInfo
	c.Bind(&reqInfo)
	if err := gvalid.CheckStruct(c, reqInfo, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	//token := c.Query("token")
	hs := auth.NewUserService()
	oneInfo, err := hs.Info(reqInfo.Token)
	//oneInfo, err := hs.Info(token)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, oneInfo)
}

func TestApi(c *gin.Context) {
	testInfo := g.Config().GetString("test", "3")
	utils.Response(c, code.ErrSuccess, testInfo)

}
