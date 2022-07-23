package cluster

import (
	"github.com/jinzhu/gorm"
	"tf-cluster/internal/model"
)

// gorm文档: https://www.tizi365.com/archives/22.html

type Config struct {
	*model.Model
	Name         string `gorm:"column:name" json:"name"`
	ClusterId    string `json:"cluster_id"`
	Config       string `json:"config"`
	Server       string `json:"server"`
	Introduction string `json:"introduction"`
}

//func (this *Config) TableName() string{
//	// 静态表名(强制覆盖)
//	tableName := "config"
//	return fmt.Sprintf("%s%s",g.Config().GetString("database.auth.prefix"),tableName)
//}

func (this *Config) ListConfig(where map[string]interface{}, page, size int) ([]*Config, error) {
	var configs []*Config
	session := this.CommonWhere(where)
	err := session.Where(where).Offset(this.GetOffset(page, size)).Limit(size).Find(&configs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return configs, nil
		}
		return configs, err
	}
	return configs, nil
}

func (this *Config) CountConfig(where map[string]interface{}) (int, error) {
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

func (this *Config) OneConfig(where map[string]interface{}) (Config, error) {
	var config Config
	err := model.Db.Where(where).First(&config).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return config, nil
		}
		return config, err
	}
	return config, nil
}

func (this *Config) DeleteConfig(where map[string]interface{}) error {
	var config Config
	err := model.Db.Where(where).Delete(&config).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}

// ps: 当用结构体更新的时候，当结构体的值是""或者0，false等，就什么也不会更新。
func (this *Config) UpdateConfig(config Config, upMap map[string]interface{}) error {
	// 更新具体值
	err := model.Db.Model(&config).Update(upMap).Error
	// 更新模型
	if err != nil {
		return err
	}
	return nil
}

func (this *Config) CreateConfig(config Config) error {
	err := model.Db.Create(&config).Error
	if err != nil {
		return err
	}
	return nil
}

func (this *Config) AfterFind() error {
	return nil
}

func (this *Config) AfterUpdate(orm *gorm.DB) error {
	// hook save update复用一个即可
	return nil
}

func (this *Config) BeforeUpdate() error {
	return nil
}
