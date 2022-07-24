package cluster

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gvalid"
	"tf-cluster/internal/service/cluster"
	"tf-cluster/library/code"
	"tf-cluster/library/utils"
)

type RequestNodeList struct {
	Ip        string `json:"ip" form:"ip"`
	ClusterId string `json:"cluster_id" form:"cluster_id"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
}

func NodeListApi(c *gin.Context) {
	var reqNodeList RequestNodeList
	c.Bind(&reqNodeList)
	if err := gvalid.CheckStruct(c, reqNodeList, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := cluster.NewNodeService()
	whereCondition := map[string]interface{}{}
	if reqNodeList.Ip != "" {
		whereCondition["ip"] = reqNodeList.Ip
	}
	if reqNodeList.ClusterId != "" {
		whereCondition["cluster_id"] = reqNodeList.ClusterId
	}
	ListInfo, err := hs.List(whereCondition, reqNodeList.Page, reqNodeList.PageSize)
	if err != nil {
		utils.Response(c, code.ErrSystem, err.Error())
		return
	}
	res := map[string]interface{}{
		"result": ListInfo,
	}
	utils.Response(c, code.ErrSuccess, res)
}

type RequestNodeCount struct {
	Ip        string `json:"ip" form:"ip"`
	ClusterId string `json:"cluster_id" form:"cluster_id"`
}

func NodeCountApi(c *gin.Context) {
	var reqNodeCount RequestNodeCount
	c.Bind(&reqNodeCount)
	if err := gvalid.CheckStruct(c, reqNodeCount, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := cluster.NewNodeService()
	whereCondition := map[string]interface{}{}
	if reqNodeCount.Ip != "" {
		whereCondition["ip"] = reqNodeCount.Ip
	}
	if reqNodeCount.ClusterId != "" {
		whereCondition["cluster_id"] = reqNodeCount.ClusterId
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

type RequestNodeInfo struct {
	Id int `json:"id" form:"id" valid:"id      @required#id不能为空"`
}

func NodeInfoApi(c *gin.Context) {
	var reqNodeInfo RequestNodeInfo
	c.Bind(&reqNodeInfo)
	if err := gvalid.CheckStruct(c, reqNodeInfo, nil); err != nil {
		utils.Response(c, code.ErrSystem, err.FirstString())
		return
	}
	hs := cluster.NewNodeService()
	whereCondition := map[string]interface{}{}
	whereCondition["id"] = reqNodeInfo.Id
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
