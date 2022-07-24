package cluster

import (
	"tf-cluster/internal/model/cluster"
	"tf-cluster/library/log"
)

type NodeService struct {
	nodeModel *cluster.Node
}

func NewNodeService() (s *NodeService) {
	h := &NodeService{}
	h.nodeModel = &cluster.Node{}
	return h
}

func (this *NodeService) List(where map[string]interface{}, page, size int) ([]*cluster.Node, error) {
	list, err := this.nodeModel.ListNode(where, page, size)
	if err != nil {
		log.Logger.Errorf("NodeService ListError: %v", err)
		return nil, err
	}
	return list, nil
}

func (this *NodeService) Count(where map[string]interface{}) (int, error) {
	count, err := this.nodeModel.CountNode(where)
	if err != nil {
		log.Logger.Errorf("NodeService CountError: %v", err)
		return 0, err
	}
	return count, nil
}

func (this *NodeService) One(where map[string]interface{}) (*cluster.Node, error) {
	one, err := this.nodeModel.OneNode(where)
	if err != nil {
		log.Logger.Errorf("NodeService OneError: %v", err)
		return nil, err
	}
	return &one, nil
}
