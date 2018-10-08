package service

import (
	"time"

	"github.com/taichi-hagiwara/ezrpc"
)

func ChatService() ezrpc.Service {
	return &chatService{}
}

type chatService struct {
}

func (s *chatService) Init(r *ezrpc.ServiceRegistry) error {
	r.Register("post", &PostArgs{}, &PostResult{})
	return nil
}

type Post struct {
	Time time.Time
	Name string
	Text string
}

type PostArgs struct {
	Text string
}

type PostResult struct {
	Log []*Post
}
