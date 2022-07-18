package auth

import (
	"tf-cluster/internal/model/auth"
	"tf-cluster/library/log"
)

type HelloService struct {
	helloModel *auth.Hello
}

func NewHelloService() (s *HelloService) {
	h := &HelloService{}
	h.helloModel = &auth.Hello{}
	return h
}

func (this *HelloService) List(where map[string]interface{}) ([]*auth.Hello, error) {
	list, err := this.helloModel.ListHello(where)
	if err != nil {
		log.Logger.Errorf("HelloService ListError: %v", err)
		return nil, err
	}
	return list, nil
}

func (this *HelloService) One(where map[string]interface{}) (*auth.Hello, error) {
	one, err := this.helloModel.OneHello(where)
	if err != nil {
		log.Logger.Errorf("HelloService OneError: %v", err)
		return nil, err
	}
	return &one, nil
}
