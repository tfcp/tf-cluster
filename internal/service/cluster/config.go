package cluster

import (
	"tf-cluster/internal/model/cluster"
	"tf-cluster/library/log"
)

type ConfigService struct {
	configModel *cluster.Config
}

func NewConfigService() (s *ConfigService) {
	h := &ConfigService{}
	h.configModel = &cluster.Config{}
	return h
}

func (this *ConfigService) List(where map[string]interface{}, page, size int) ([]*cluster.Config, error) {
	list, err := this.configModel.ListConfig(where, page, size)
	if err != nil {
		log.Logger.Errorf("ConfigService ListError: %v", err)
		return nil, err
	}
	return list, nil
}

func (this *ConfigService) Count(where map[string]interface{}) (int, error) {
	count, err := this.configModel.CountConfig(where)
	if err != nil {
		log.Logger.Errorf("ConfigService CountError: %v", err)
		return 0, err
	}
	return count, nil
}

func (this *ConfigService) One(where map[string]interface{}) (*cluster.Config, error) {
	one, err := this.configModel.OneConfig(where)
	if err != nil {
		log.Logger.Errorf("ConfigService OneError: %v", err)
		return nil, err
	}
	return &one, nil
}

func (this *ConfigService) Save(id int, name, clusterId, config, introduction string) error {
	var err error
	newConfig := cluster.Config{
		Name:         name,
		ClusterId:    clusterId,
		Config:       config,
		Introduction: introduction,
	}
	if id > 0 {
		// update
		whereUp := map[string]interface{}{
			"id": id,
		}
		userUp, err := this.configModel.OneConfig(whereUp)
		if err != nil {
			log.Logger.Errorf("ConfigService Save OneUserError: %v", err)
			return err
		}
		upMap := map[string]interface{}{}
		if name != "" {
			upMap["name"] = name
		}
		if clusterId != "" {
			upMap["cluster_id"] = clusterId
		}
		if config != "" {
			upMap["config"] = config
		}
		if introduction != "" {
			upMap["introduction"] = introduction
		}
		if err = this.configModel.UpdateConfig(userUp, upMap); err != nil {
			log.Logger.Errorf("ConfigService Save UpdateUserError: %v", err)
			return err
		}
		return nil
	}
	if err = this.configModel.CreateConfig(newConfig); err != nil {
		log.Logger.Errorf("ConfigService Save CreateUserError: %v", err)
		return nil
	}
	return nil
}

func (this *ConfigService) ConfigExists(id int, name string) bool {
	if id > 0 {
		return false
	}
	whereUser := map[string]interface{}{
		"name": name,
	}
	user, err := this.configModel.OneConfig(whereUser)
	if err != nil {
		log.Logger.Errorf("ConfigExists OneConfigError: %v", err)
		return true
	}
	if user.ID != 0 {
		return true
	}
	return false
}
