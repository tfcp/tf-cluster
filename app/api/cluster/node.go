package cluster

import (
	"github.com/gin-gonic/gin"
	"tf-cluster/library/code"
	"tf-cluster/library/utils"
)

type RequestNodeList struct {
	Name string `json:"name" form:"name" valid:"name      @required#name不能为空"`
	Ip   string `json:"ip" form:"ip" valid:"ip      @required#ip不能为空"`
}

func NodeListApi(c *gin.Context) {
	res := map[string]interface{}{
		"result": "",
	}
	utils.Response(c, code.ErrSuccess, res)
}

type RequestNodeInfo struct {
	Name string `json:"name" form:"name" valid:"name      @required#name不能为空"`
	Ip   string `json:"ip" form:"ip" valid:"ip      @required#ip不能为空"`
}

func NodeInfoApi(c *gin.Context) {
	res := map[string]interface{}{
		"result": "",
	}
	utils.Response(c, code.ErrSuccess, res)
}

type RequestDrainList struct {
	Name string `json:"name" form:"name" valid:"name      @required#name不能为空"`
	Ip   string `json:"ip" form:"ip" valid:"ip      @required#ip不能为空"`
}

func DrainListApi(c *gin.Context) {
	res := map[string]interface{}{
		"result": "",
	}
	utils.Response(c, code.ErrSuccess, res)
}

type RequestDrainCount struct {
	Name string `json:"name" form:"name" valid:"name      @required#name不能为空"`
	Ip   string `json:"ip" form:"ip" valid:"ip      @required#ip不能为空"`
}

func DrainCountApi(c *gin.Context) {
	res := map[string]interface{}{
		"result": "",
	}
	utils.Response(c, code.ErrSuccess, res)
}
