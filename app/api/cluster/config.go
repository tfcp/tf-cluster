package cluster

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"tf-cluster/internal/service/cluster"
	"tf-cluster/library/code"
	"tf-cluster/library/utils"
)

type RequestConfigList struct {
	Name     string `json:"name" form:"name"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

func ConfigListApi(c *gin.Context) {
	var reqConfigList RequestConfigList
	c.Bind(&reqConfigList)
	if err := gvalid.CheckStruct(c, reqConfigList, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := cluster.NewConfigService()
	whereCondition := map[string]interface{}{}
	if reqConfigList.Name != "" {
		whereCondition["name"] = reqConfigList.Name
	}
	ListInfo, err := hs.List(whereCondition, reqConfigList.Page, reqConfigList.PageSize)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": ListInfo,
	}
	utils.Response(c, code.ErrSuccess, res)
}

type RequestConfigCount struct {
	Name string `json:"name" form:"name"`
}

func ConfigCountApi(c *gin.Context) {
	var reqConfigCount RequestConfigCount
	c.Bind(&reqConfigCount)
	if err := gvalid.CheckStruct(c, reqConfigCount, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := cluster.NewConfigService()
	whereCondition := map[string]interface{}{}
	if reqConfigCount.Name != "" {
		whereCondition["name"] = reqConfigCount.Name
	}
	ListInfo, err := hs.Count(whereCondition)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": ListInfo,
	}
	utils.Response(c, code.ErrSuccess, res)
}

type RequestConfigSave struct {
	Id           int    `json:"id" form:"id"`
	Name         string `json:"name" form:"name" valid:"name      @required#name不能为空"`
	ClusterId    string `json:"cluster_id" form:"cluster_id" valid:"cluster_id      @required#cluster_id不能为空"`
	Config       string `json:"config" form:"config" valid:"config      @required#config不能为空"`
	Introduction string `json:"introduction" form:"introduction"`
}

func ConfigSaveApi(c *gin.Context) {
	var reqConfigSave RequestConfigSave
	c.Bind(&reqConfigSave)
	if err := gvalid.CheckStruct(c, reqConfigSave, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := cluster.NewConfigService()
	if hs.ConfigExists(reqConfigSave.Id, reqConfigSave.Name) {
		utils.Response(c, code.ErrExist, nil)
		return
	}
	err := hs.Save(reqConfigSave.Id,
		reqConfigSave.Name,
		reqConfigSave.ClusterId,
		reqConfigSave.Config,
		reqConfigSave.Introduction,
	)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	utils.Response(c, code.ErrSuccess, map[string]interface{}{})
}

type RequestConfigInfo struct {
	Id int `json:"id" form:"id" valid:"id      @required#id不能为空"`
}

func ConfigInfoApi(c *gin.Context) {
	var reqConfigInfo RequestConfigInfo
	c.Bind(&reqConfigInfo)
	if err := gvalid.CheckStruct(c, reqConfigInfo, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := cluster.NewConfigService()
	whereCondition := map[string]interface{}{}
	whereCondition["id"] = reqConfigInfo.Id
	configDetail, err := hs.One(whereCondition)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": configDetail,
	}
	utils.Response(c, code.ErrSuccess, res)
}
