package demo

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
	"tf-cluster/internal/service/auth"
	"tf-cluster/internal/test"
)

var (
	svc *auth.HelloService
)

func init() {
	test.SetupServer()
	svc = auth.NewHelloService()
}

func Test_HelloService(t *testing.T) {
	// one
	gtest.C(t, func(t *gtest.T) {
		whereCondition := map[string]interface{}{
			"name": "tom",
		}
		_, err := svc.One(whereCondition)
		t.Assert(err, nil)
	})
	// list
	gtest.C(t, func(t *gtest.T) {
		whereCondition := map[string]interface{}{
			"name": "",
		}
		_, err := svc.List(whereCondition)
		t.Assert(err, nil)
	})
}
