package cluster

import (
	"github.com/jinzhu/gorm"
	"tf-cluster/internal/model"
)

// gorm文档: https://www.tizi365.com/archives/22.html

type Node struct {
	*model.Model
	Ip            string `gorm:"column:ip" json:"ip"`
	Labels        string `json:"labels"`
	Taints        string `json:"taints"`
	Unschedulable int    `json:"Unschedulable"`
}

//func (this *Node) TableName() string{
//	// 静态表名(强制覆盖)
//	tableName := "Node"
//	return fmt.Sprintf("%s%s",g.Node().GetString("database.auth.prefix"),tableName)
//}

func (this *Node) ListNode(where map[string]interface{}, page, size int) ([]*Node, error) {
	var nodes []*Node
	session := this.CommonWhere(where)
	err := session.Where(where).Offset(this.GetOffset(page, size)).Limit(size).Find(&nodes).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nodes, nil
		}
		return nodes, err
	}
	return nodes, nil
}

func (this *Node) CountNode(where map[string]interface{}) (int, error) {
	var count int
	session := this.CommonWhere(where)
	err := session.Model(this).Where(where).Count(&count).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return count, nil
		}
		return count, err
	}
	return count, nil
}

func (this *Node) OneNode(where map[string]interface{}) (Node, error) {
	var node Node
	err := model.Db.Where(where).First(&node).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return node, nil
		}
		return node, err
	}
	return node, nil
}

func (this *Node) DeleteNode(where map[string]interface{}) error {
	var node Node
	err := model.Db.Where(where).Delete(&node).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}

// ps: 当用结构体更新的时候，当结构体的值是""或者0，false等，就什么也不会更新。
func (this *Node) UpdateNode(node Node, upMap map[string]interface{}) error {
	// 更新具体值
	err := model.Db.Model(&node).Update(upMap).Error
	// 更新模型
	if err != nil {
		return err
	}
	return nil
}

func (this *Node) CreateNode(node Node) error {
	err := model.Db.Create(&node).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *Node) AfterFind() error {
	return nil
}

func (this *Node) AfterUpdate(orm *gorm.DB) error {
	// hook save update复用一个即可
	return nil
}

func (this *Node) BeforeUpdate() error {
	return nil
}
