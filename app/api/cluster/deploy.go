package cluster

import (
	"github.com/gin-gonic/gin"
	"tf-cluster/library/code"
	"tf-cluster/library/utils"
)

type RequestDeployList struct {
	Name string `json:"name" form:"name" valid:"name      @required#name不能为空"`
	Ip   string `json:"ip" form:"ip" valid:"ip      @required#ip不能为空"`
}

func DeployListApi(c *gin.Context) {
	res := map[string]interface{}{
		"result": "",
	}

	utils.Response(c, code.ErrSuccess, res)
}

type RequestDeployInfo struct {
	Name string `json:"name" form:"name" valid:"name      @required#name不能为空"`
	Ip   string `json:"ip" form:"ip" valid:"ip      @required#ip不能为空"`
}

func DeployInfoApi(c *gin.Context) {
	res := map[string]interface{}{
		"result": "",
	}

	utils.Response(c, code.ErrSuccess, res)
}
