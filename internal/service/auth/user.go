package auth

import (
	"tf-cluster/internal/model/auth"
	"tf-cluster/library/log"
	"tf-cluster/library/utils"
)

type UserService struct {
	userModel *auth.User
}

func NewUserService() (s *UserService) {
	h := &UserService{}
	h.userModel = &auth.User{}
	return h
}

func (this *UserService) List(where map[string]interface{}, page, size int) ([]*auth.User, error) {
	list, err := this.userModel.ListUser(where, page, size)
	if err != nil {
		log.Logger.Errorf("UserService ListError: %v", err)
		return nil, err
	}
	return list, nil
}

func (this *UserService) Count(where map[string]interface{}) (int, error) {
	count, err := this.userModel.CountUser(where)
	if err != nil {
		log.Logger.Errorf("UserService CountError: %v", err)
		return 0, err
	}
	return count, nil
}

func (this *UserService) One(where map[string]interface{}) (*auth.User, error) {
	one, err := this.userModel.OneUser(where)
	if err != nil {
		log.Logger.Errorf("UserService OneError: %v", err)
		return nil, err
	}
	return &one, nil
}

func (this *UserService) Info(token string) (map[string]interface{}, error) {
	info, err := utils.ParseToken(token)
	if err != nil {
		log.Logger.Errorf("UserService InfoError: %v", err)
		return nil, err
	}
	resInfo := map[string]interface{}{
		"avatar":       info.Avatar,
		"introduction": info.Introduction,
		"name":         info.Name,
		"roles":        info.Roles,
	}
	return resInfo, nil
}

func (this *UserService) Delete(id int) error {
	whereDelete := map[string]interface{}{
		"id": id,
	}
	err := this.userModel.DeleteUser(whereDelete)
	if err != nil {
		log.Logger.Errorf("UserService DeleteError: %v", err)
		return err
	}
	return nil
}

func (this *UserService) ChangeStatus(id, status int) error {
	whereOne := map[string]interface{}{
		"id": id,
	}
	user, err := this.userModel.OneUser(whereOne)
	if err != nil {
		log.Logger.Errorf("UserService ChangeStatus OneError: %v", err)
		return err
	}
	if err := this.userModel.UpdateUser(user, map[string]interface{}{"status": status}); err != nil {
		log.Logger.Errorf("UserService ChangeStatus UpdateUserError: %v", err)
		return err
	}
	return nil
}

func (this *UserService) Save(id, status, role int, name, avatar, introduction string) error {
	newUser := auth.User{
		Name:         name,
		Status:       status,
		Role:         role,
		Pwd:          "",
		Avatar:       avatar,
		Introduction: introduction,
	}
	if id > 0 {
		// update
		whereUp := map[string]interface{}{
			"id": id,
		}
		userUp, err := this.userModel.OneUser(whereUp)
		if err != nil {
			log.Logger.Errorf("UserService Save OneUserError: %v", err)
			return err
		}
		upMap := map[string]interface{}{}
		if status > 0 {
			upMap["status"] = status
		}
		if role > 0 {
			upMap["role"] = role
		}
		if name != "" {
			upMap["name"] = name
		}
		if avatar != "" {
			upMap["avatar"] = avatar
		}
		if introduction != "" {
			upMap["introduction"] = introduction
		}
		if err := this.userModel.UpdateUser(userUp, upMap); err != nil {
			log.Logger.Errorf("UserService Save UpdateUserError: %v", err)
			return err
		}
		return nil
	}
	if err := this.userModel.CreateUser(newUser); err != nil {
		log.Logger.Errorf("UserService Save CreateUserError: %v", err)
		return nil
	}
	return nil
}
